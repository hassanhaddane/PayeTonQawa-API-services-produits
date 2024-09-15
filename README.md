# PayeTonQawa-API-service-clients
**MSPR 4**

Utilisation de go (golang). Utilisation du framework Gin pour la partie REST API.


## Utilisation

### Build and run

1. docker build -t service-products .
2. docker run -p 8092:8080 service-products
3. open a terminal an use the curl commands bellow

### Requetes CRUD

#### Create product

```sh
curl -X POST http://localhost:8092/products \
-H "Content-Type: application/json" \
-d '{"id":"1", "username":"john_doe", "email":"john@example.com"}'
```


#### Read product

Get a list of all products

```sh
curl -X GET http://localhost:8092/products
```


#### Update product

using product with id 1 as an example:

```sh
curl -X PUT http://localhost:8092/products/1 -H "Content-Type: application/json" -d '{
  "id": "1",
  "username": "<updated_username>",
  "email": "<updated_email>"
}'
```

#### Delete product

```sh
curl -X DELETE http://localhost:8092/products/1
```


