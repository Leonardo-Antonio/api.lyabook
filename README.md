# API RESTful for Mor - LyaBook

### Generate certificates
> can be created inside a folder ```/certificate```
- private
```shell
openssl genrsa -out app.rsa 1024
```
- public
```shell
openssl rsa -in app.rsa -pubout > app.rsa.pub
```

## Documentation
### Roles
- Admin
- Manager
- Client

---
##### SignUp - Crear cuenta
```json
- http://localhost:8080/api/v1/users/sign-up/dni
{
    "password": "cmcx100pre",
    "rol": "Admin",
    "dni": "71062235" 
}
- http://localhost:8080/api/v1/users/sign-up/email
{
    "name": "leonardo antONio",
    "last_name": "noLASco LEyva",
    "password": "cmcx100pre",
    "rol": "Client",
    "email": "example@example.com" 
}
```
---