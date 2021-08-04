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