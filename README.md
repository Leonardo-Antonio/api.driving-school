# Api Rest - Driving School

## EndPoints
- Users
```json
```
### Commands for generate certificates
- private certificate:
    ```shell
    openssl genrsa -out app.rsa 1024
    ```
- public certificate:
    ```shell
    openssl rsa -in app.rsa -pubout > app.rsa.pub
    ```

### Deploy
```shell
$ sh deploy.sh
```

creo q debe haber un campo en la tb users como teacher-assigned
los envios de correos como notificaciones
valdiar q solo un cliente pueda comprar un pquete