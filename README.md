
![Header](https://i.imgur.com/IXVWoeC.png)
# introduction

### Topic

This is a project for an imaginary courier company. PostBox Is a company that provides courier services using parcel machines. The idea is based on the Inpost company. 

The application was created just for customers. It is supposed to allow him:
* to set up/log in an account, 
* order a package, 
* view the list of incoming/shipped packages, 
* see the history of selected packages, 
* view/edit his profile.

## architekture

For this purpose, I created:
* database in Postgres
* api in Golang using gin-gonic and go-pg packages
* user interface using Dart language and Flutter framework.

The application server was hosted on Google Cloud Platform, sending emails was done using Mailjet platform

![architekture](https://i.imgur.com/pQgaNFm.png)
