# API RESTful for Mor - LyaBook
### Kill ports
```shell
lsof -i:8080
# out console
main    286077  leo    9u  IPv6 1398322      0t0  TCP *:http-alt (LISTEN)
main    286077  leo   11u  IPv6 1395667      0t0  TCP localhost:http-alt->localhost:40034 (CLOSE_WAIT)
# end out

kill -9 286077
```
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
##### SignUp - Crear cuenta [POST]
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

##### Verify - Verificar cuenta [POST]
```json
- http://localhost:8080/api/v1/users/verify
{
    "verification_code": "VQ25IO8X",
    "email": "leo2001.nl08@gmail.com"
}
```

---
##### LogIn - Ingresar [POST]
```json
- http://localhost:8080/api/v1/users/log-in/dni
{
    "password": "cmcx100pre",
    "dni": "71062235" 
}
- http://localhost:8080/api/v1/users/log-in/email
{
    "password": "cmcx100pre",
    "email": "example@example.com" 
}
```
---
##### Update - Actualizar usuario [PUT]
```json
- http://localhost:8080/api/v1/users/:id 
{
    "name": "Alexandra",
    "last_name": "Navarro",
    "password": "cmcx100pre",
}
```
---

## Category

#### Create [POST]
```json
- http://localhost:8080/api/v1/category
{
    "name": "terror nocturno"
}
```
#### Update [PUT]
```json
- http://localhost:8080/api/v1/category/:id
{
    "name": "terror nocturno"
}
```
---


## Book

#### Create [POST]
```json
// crea libro digital 
- http://localhost:8080/api/v1/books/d
{
    "name": "Misterio, una pasión",
    "author": "Aldo Miashiro",
    "editorial": "Macro",
    "price_current": 150.55,
    "description": "Misterio, una historia basada en hechos reales de la creación de la trinchera norte",
    "type": {
        "digital": {
            "format": "asds",
            "details": ["dsdsdsd"]
        }
    },
    "categories": ["61101d22b88c55b02dbc5f2c"],
    "src": "https://imge.com"
}



// crea libro fisico 
- http://localhost:8080/api/v1/books/f
{
    "name": "Misterio, una pasión",
    "author": "Aldo Miashiro",
    "editorial": "Macro",
    "price_current": 150.55,
    "description": "Misterio, una historia basada en hechos reales de la creación de la trinchera norte",
    "type": {
        "fisico": {
            "details": ["ssdsa"],
            "log": "151525411",
            "lat": "-45854514515",
            "stock": 4
        }
    },
    "categories": ["61101d22b88c55b02dbc5f2c"],
    "src": "https://imge.com"
}



// crea libro fisico y digital
- http://localhost:8080/api/v1/books/df
{
    "name": "Misterio, una pasión",
    "author": "Aldo Miashiro",
    "editorial": "Macro",
    "price_current": 150.55,
    "description": "Misterio, una historia basada en hechos reales de la creación de la trinchera norte",
    "type": {
        "digital": {
            "format": "asds",
            "details": ["dsdsdsd"]
        },
        "fisico": {
            "details": ["ssdsa"],
            "log": "151525411",
            "lat": "-45854514515",
            "stock": 4
        }
    },
    "categories": ["61101d22b88c55b02dbc5f2c"],
    "src": "https://imge.com"
}
```