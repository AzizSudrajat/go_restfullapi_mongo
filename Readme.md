# Golang RESTfull API
RESTfull API build on :
- Go Language
- MVC Pattern
- Lib Go : [mux](github.com/gorilla/mux), [monggo-driver](go.mongodb.org/mongo-driver/mongo), [godotenv](github.com/joho/godotenv)
- Mongo DB Community
## Installation
Golang required [Link](https://golang.org/doc/install) go1.15.* to run
Mongo DB Community required [Link](https://www.mongodb.com/try/download/community)
```sh
- Download or Clone from Github -
cd go_restfullapi_mongo
touch .env
echo "PORT=:8000" >> .env
echo "CONNECTION_STRING=mongodb://{yourhost}/{yourcollection}" >> .env
```
Running Code
```sh
go build main.go
go run ./
```
## API Endpoint
| METHOD | ENDPOINT | Description |
| ------ | ------ | ------ |
| GET | http://{yourhost}:8000/api/books | GET ALL Book |
| GET | http://{yourhost}:8000/api/books/{id} | GET ONE Book  |
| POST | http://{yourhost}:8000/api/books | CREATE Book |
| PUT | http://{yourhost}:8000/api/books/{id} | UPDATE Book |
| DELETE | http://{yourhost}:8000/api/books/{id} | DELETE Book |

## Postman Collection
File on Directory demo/postman

## License
MIT
