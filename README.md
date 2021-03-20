# go-jsonapi-example

`make && make run`

## Create a new car:
`curl -X POST http://localhost:8080/v1/cars -d '{"data" : {"type" : "cars" , "attributes": {"brand" : "bmw", "model": "x5", "price": 999, "status": "OnTheWay"}}}'`

## List cars:
`curl -X GET http://localhost:8080/v1/cars`

## List paginated cars:
`curl -X GET 'http://localhost:8080/v1/cars?page\[offset\]=0&page\[limit\]=2'`
**OR**
`curl -X GET 'http://localhost:8080/v1/cars?page\[number\]=1&page\[size\]=2'`

## Update:
`curl -X PATCH http://localhost:8080/v1/cars/1 -d '{ "data" : {"type" : "cars", "id": "1", "attributes": {"model" : "x3"}}}'`

## Delete:
`curl -X DELETE http://localhost:8080/v1/cars/2`
