# My App

Repository for a SPA web app

## Desired Architecture Overview
![Overview](images/overview.jpg)

## Instructions to start dev environemnt
docker-compose up --build

Here are the routes (you can find them in internal/api/routes.go)
1. `GET     http://0.0.0.0:5300/health`
2. `GET     http://0.0.0.0:5300/api/v1/thing`
3. `POST    http://0.0.0.0:5300/api/v1/thing`
4. `GET     http://0.0.0.0:5300/api/v1/thing/{id}`
5. `PUT     http://0.0.0.0:5300/api/v1/thing/{id}`
6. `DELETE  http://0.0.0.0:5300/api/v1/thing/{id}`

Now open Postman and send request to any of the routes with the following body as an example:

`POST` -> `http://0.0.0.0:5300/api/v1/thing`

`{
     "name": "aaa Thing",
     "features": ["aaaa", "b", "c"]
 }`


