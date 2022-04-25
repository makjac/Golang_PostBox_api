package service

type LoginService interface {
	Login(username string, passwd string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPasswd   string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "qwe",
		authorizedPasswd:   "123",
	}
}

func (service *loginService) Login(username string, passwd string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPasswd == passwd
}
