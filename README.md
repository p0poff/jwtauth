# jwtauth
go jwt auth test

## API
### **POST** /login
**request body**
```
{
"login": "foo",
"pass": "pass"
}
```
**response**
```
{
  "code": 200,
  "result": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImJhciIsImV4cCI6MTU5NDA2ODAxMiwiaWF0IjoxNTk0MDY3MDEyLCJpc3MiOiJqd3QifQ.r_ks3JSDu5Q5zOuqa9zvNrSfNdSpIgEV4SPD-GrccWQ"
  }
}
```

### **POST** /check
**request body**
```
{
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImJhciIsImV4cCI6MTU5NDA2ODAxMiwiaWF0IjoxNTk0MDY3MDEyLCJpc3MiOiJqd3QifQ.r_ks3JSDu5Q5zOuqa9zvNrSfNdSpIgEV4SPD-GrccWQ"
}
```
**response**
```
{
  "code": 200,
  "result": {
    "username": "bar"
  }
}
```
