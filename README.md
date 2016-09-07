# About

Test potions is a simple REST test service that exposes the following end points:

 
## Authorization

The service is protected by a secret key that is set at the server level.  
All API communication expects an `Authorization` header 
containing the **base64 encoded** secret key as the header value. 

The default secret key is: **mixOfRandomPotions**


## GET, POST /


Check API health.

**status** string  
**version** string


## GET /list

Retrieve list of potions.

**items** list of potions having the format:  
* **id** integer
* **name** string


## GET /item?id=XX

Retrieve detailed information about a single potion.

**item** list of potions having the format:  
* **id** integer
* **name** string
* **toxicity** float


## POST /mix

Mix 2 or more potions and return the final toxicity.

Request JSON body:

**ids** list of potion IDs. Eg: `{ "ids": [2, 5, 6] }` 

&nbsp;

Response JSON body:

**toxicity** Resulting toxicity achieved by mixing the requested potions


# Run

```
$ go build
$ ./potions-rest-api
```

# Test

```
$ go test ./...
$ go test ./... -cover
```


# CURL Examples

```
$ curl -H 'Authorization:XXXXX' /
{"status":"Ok","version":"1.0"}
  
$ curl -H 'Authorization:XXXXX' /list
{"items":[{"id":1,"string":"Acidic Tonic"},{"id":2,"string":"Brew of Exalted Cats"},{"id":....
  
$ curl -H 'Authorization:XXXXX' /item?id=1
{"item":{"id":1,"string":"Acidic Tonic","toxicity":10.5}}
  
$ curl -X POST -H 'Authorization:XXXX' /mix --data '{"ids": [1,2,3]}'
{"toxicity":55.9}
```