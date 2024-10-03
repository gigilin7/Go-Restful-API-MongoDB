# Go-Restful-API-MongoDB
Build a mini project using Golang, Gin and MongoDB.
This mini project can do CRUD for management users.

## Database
[MongoDB Atlas](https://www.mongodb.com/products/platform/atlas-database) ( the Cloud-Native Document Database as a Service. )
  + Region: AWS Hong Kong (ap-east-1)
  + version: 7.0.12
 
## Structure
```
.
├─config
├─delivery
│  ├─http
│  └─response
├─domain
│  ├─models
│  ├─repository       
│  └─usecase
├─repository
└─usecase
```
Explain the individual functions and functions of the above folders:
+ `config`: The setting of database.
+ `delivery`:
  + `http`: A UserController using the Gin framework is defined to register a set of user-related routes, including APIs for creating, obtaining single or all users, updating and deleting users.
  + `response`: Indicates the format of the API response, including status code, message content and returned data.
+ `domain`:
  + `models`: The data object.
  + `repository`: Interface of CRUD operations.
  + `usecase`: Connect the business process flow with the data flow that is processed in the repository section.
+ `repository`: Functions for every usecase in folder domain.
+ `usecase`: Functions for every repository in folder domain.

## Requirement
Install godotenv package and load .env file
```
go get github.com/joho/godotenv
```
Create a .env file in the project directory and write the MongoDB password into it:
```
MONGODB_PASSWORD=your_password_here
```

## Implement
```
 go run main.go
```
CRUD operation instructions:
```
curl http://localhost:9090/v1/user/getall
```

```
curl http://localhost:9090/v1/user/get/:name
```

```
curl -X POST http://localhost:9090/v1/user/create \
-H "Content-Type: application/json" \
-d '{
  "name": "Amy", 
  "age": 23,
  "address": {
    "state": "Taiwan", 
    "city": "Taipei", 
    "pincode": 90002
  }
}'
```

```
curl -X PATCH http://localhost:9090/v1/user/update \
-H "Content-Type: application/json" \
-d '{
  "name": "Amy", 
  "age": 22
}'
```

```
curl -X DELETE http://localhost:9090/v1/user/delete/:name
```

## Result
<img src="https://github.com/gigilin7/Go-Restful-API-MongoDB/blob/main/picture/mongodb1.jpg" height=200>

<img src="https://github.com/gigilin7/Go-Restful-API-MongoDB/blob/main/picture/mongodb2.jpg" height=400><img src="https://github.com/gigilin7/Go-Restful-API-MongoDB/blob/main/picture/mongodb3.jpg" height=400>

<img src="https://github.com/gigilin7/Go-Restful-API-MongoDB/blob/main/picture/mongodb-amy.jpg" height=400>

[Reference for learning](https://medium.com/@hoseahutahuruk/build-rest-api-golang-gin-mongodb-b6ac5713440f)




