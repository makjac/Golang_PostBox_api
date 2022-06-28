/*
	cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 1 * time.Minute,
		Credentials: false,
		ValidateHeaders: false,
	}
*/
package cors

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	AllowOriginKey      string = "Access-Control-Allow-Origin"
	AllowCredentialsKey        = "Access-Control-Allow-Credentials"
	AllowHeadersKey            = "Access-Control-Allow-Headers"
	AllowMethodsKey            = "Access-Control-Allow-Methods"
	MaxAgeKey                  = "Access-Control-Max-Age"

	OriginKey         = "Origin"
	RequestMethodKey  = "Access-Control-Request-Method"
	RequestHeadersKey = "Access-Control-Request-Headers"
	ExposeHeadersKey  = "Access-Control-Expose-Headers"
)

const (
	optionsMethod = "OPTIONS"
)

type Config struct {
	ValidateHeaders bool

	Origins string
	origins []string

	RequestHeaders string
	requestHeaders []string

	ExposedHeaders string

	Methods string
	methods []string

	MaxAge time.Duration
	maxAge string

	Credentials bool
	credentials string
}

func (config *Config) prepare() {
	if config.Origins == "*" && config.Credentials == true {
		panic("Do not use Origins = \"*\" and Credentials = true together.")
	}

	config.origins = strings.Split(config.Origins, ", ")
	config.methods = strings.Split(config.Methods, ", ")
	config.requestHeaders = strings.Split(config.RequestHeaders, ", ")
	config.maxAge = fmt.Sprintf("%.f", config.MaxAge.Seconds())

	config.credentials = fmt.Sprintf("%t", config.Credentials)

	for idx, header := range config.requestHeaders {
		config.requestHeaders[idx] = strings.ToLower(header)
	}
}

func Middleware(config Config) gin.HandlerFunc {
	forceOriginMatch := false

	if config.Origins == "" {
		panic("You must set at least a single valid origin. If you don't want CORS, to apply, simply remove the middleware.")
	}

	if config.Origins == "*" && config.Credentials == false {
		forceOriginMatch = true
	}

	config.prepare()

	return func(context *gin.Context) {
		currentOrigin := context.Request.Header.Get(OriginKey)
		context.Writer.Header().Add("Vary", OriginKey)

		if currentOrigin == "" {
			return
		}

		originMatch := false
		if !forceOriginMatch {
			originMatch = matchOrigin(currentOrigin, config)
		}

		if forceOriginMatch || originMatch {
			valid := false
			preflight := false

			if context.Request.Method == optionsMethod {
				requestMethod := context.Request.Header.Get(RequestMethodKey)
				if requestMethod != "" {
					preflight = true
					valid = handlePreflight(context, config, requestMethod)
				}
			}

			if !preflight {
				valid = handleRequest(context, config)
			}

			if valid {

				if config.Credentials {
					context.Writer.Header().Set(AllowCredentialsKey, config.credentials)
					context.Writer.Header().Set(AllowOriginKey, currentOrigin)
				} else if forceOriginMatch {
					context.Writer.Header().Set(AllowOriginKey, "*")
				} else {
					context.Writer.Header().Set(AllowOriginKey, currentOrigin)
				}

				if preflight {
					context.AbortWithStatus(200)
				}
				return
			}
		}

		context.Abort()
	}
}

func handlePreflight(context *gin.Context, config Config, requestMethod string) bool {
	if ok := validateRequestMethod(requestMethod, config); ok == false {
		return false
	}

	if ok := validateRequestHeaders(context.Request.Header.Get(RequestHeadersKey), config); ok == true {
		context.Writer.Header().Set(AllowMethodsKey, config.Methods)
		context.Writer.Header().Set(AllowHeadersKey, config.RequestHeaders)

		if config.maxAge != "0" {
			context.Writer.Header().Set(MaxAgeKey, config.maxAge)
		}

		return true
	}

	return false
}

func handleRequest(context *gin.Context, config Config) bool {
	if config.ExposedHeaders != "" {
		context.Writer.Header().Set(ExposeHeadersKey, config.ExposedHeaders)
	}

	return true
}

func matchOrigin(origin string, config Config) bool {
	for _, value := range config.origins {
		if value == origin {
			return true
		}
	}
	return false
}

func validateRequestMethod(requestMethod string, config Config) bool {
	if !config.ValidateHeaders {
		return true
	}

	if requestMethod != "" {
		for _, value := range config.methods {
			if value == requestMethod {
				return true
			}
		}
	}

	return false
}

func validateRequestHeaders(requestHeaders string, config Config) bool {
	if !config.ValidateHeaders {
		return true
	}

	headers := strings.Split(requestHeaders, ",")

	for _, header := range headers {
		match := false
		header = strings.ToLower(strings.Trim(header, " \t\r\n"))

		for _, value := range config.requestHeaders {
			if value == header {
				match = true
				break
			}
		}

		if !match {
			return false
		}
	}

	return true
}
