# Api Rest - Driving School


Para poder usar debe crear un file .env y crear los certificados para crear el firmar y verificar el token
```.env
MONGO_URI = mongodb://127.0.0.1:27017/
NAME_DATABASE = driving-school
PORT = 8080
EMAIL = leo2001.nl08@gmail.com
PASSWORD_EMAIL = gxyt rxyx vimx yshl
BASE_URI = /api/v1

# Api Reniec
API_RENIEC_RUC = https://dniruc.apisperu.com/api/v1/ruc
API_RENIEC_DNI = https://dniruc.apisperu.com/api/v1/dni
TOKEN_API_RENIC = eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbWFpbCI6ImxlbzIwMDEubmwwOEBnbWFpbC5jb20ifQ.AmjGab5gCfOAPEBsmK75qqUNFDmxsJS-eHHjxJayG1g
```
### Commands for generate certificates
estos certificados deben ser creados en la suguiente ruta src > authorizatin > keys
- private certificate:
    ```shell
    openssl genrsa -out app.rsa 1024
    ```
- public certificate:
    ```shell
    openssl rsa -in app.rsa -pubout > app.rsa.pub
    ```

## EndPoints
- Users
```json
http://localhost:8080/api/v1/users [GET] obtine todos los usuarios
http://localhost:8080/api/v1/users/sign-up/dni [POST] crear cuenta por dni
http://localhost:8080/api/v1/users/sign-up/email [POST] crear cuenta por email
http://localhost:8080/api/v1/users/log-in/dni [POST] 
http://localhost:8080/api/v1/users/log-in/email [POST] 
http://localhost:8080/api/v1/users/:id [delete]
http://localhost:8080/api/v1/users/ [PUT] actualizar usiarrio

> Roles [Admin, Instructor, Client]

Body (edit)
{
    "_id" : "60bf05b51fce3a50eca01e3c",
    "email" : "71062sa25@gmail.dcom",
	"names" : "Alexandras",
	"last_names" : "Navarros2s",
	"password" : "cmcx100pre",
	"rol": "admin"
}

Body (create client dni)
{
    "dni" : "<dni>",
	"password" : "cmcx100pre",
	"rol": "client"
}

Body (create client email)
{
    "email" : "71062sa25@gmail.dcom",
	"names" : "Alexandras",
	"last_names" : "Navarros2s",
	"password" : "cmcx100pre",
	"rol": "client"
}

GetAll User requiere un token de ADmin
Headers {
    Authorization: <token>
```



- Pacakges
```json
http://localhost:8080/api/v1/packages [POST]
http://localhost:8080/api/v1/packages [GET]
http://localhost:8080/api/v1/packages/:id [GET]
http://localhost:8080/api/v1/packages/name/:namePackage [GET]
http://localhost:8080/api/v1/packages/ [PUT]
http://localhost:8080/api/v1/packages/:id [DELETE]

> Roles [Admin, Instructor, Client]

Body (edit)
{
    "_id": "60c4e8bd9ff0139141647e07" ,
    "name": "packageUpe",
    "description": "description update",
    "content": ["holaU", "holaU", "sadsad"],
    "price": 505
}
Headers {
    Authorization: <token[admin]>

Body (create package)
{
    "name": "packageUpe",
    "description": "description",
    "content": ["sadsaa", "asd", "sadsad"],
    "price": 505
}
Headers {
    Authorization: <token[admin]>
```

- Admin - Instructor
```json
http://localhost:8080/api/v1/student-teacher/:turn [GET] buscar alumnos y profesores por turno
http://localhost:8080/api/v1/student-teacher/ [POST] asiganar a alumno a profesor
http://localhost:8080/api/v1/student-teacher/teacher/:id [GET] buscar alumnos de un profesor

> Turns [Morning, Afternoon, Night]

Body (assign)
{
    "id_client": "60c3140a1602681625f71548",
    "id_teacher": "60c315771b91ae631e610610"
}
Headers {
    Authorization: <token[admin]>
```

- Admin - Instructor
```json
http://localhost:8080/api/v1/sales comprar un apckete

> Turns [Morning, Afternoon, Night]
body
{
    "id_client": "<id>",
    "id_package": "<id>",
    "turn": "Morning"
}
Headers {
    Authorization: <token[cliente]>
```
## Run Bin Project

```
./main
```
## Dev and Deploy project 
```go
dev
go run src/main.go
prod
go build src/main.go
```


### Deploy [deprecated]
```shell
$ sh deploy.sh
```
