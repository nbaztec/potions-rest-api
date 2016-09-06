# About

Test potions is a simple REST test service that exposes the following end points:

 
## Authorization

The API expects an `Authorization` header containing the base64 encoded key. The key is set
at server level and defaults to **mixOfRandomPotions**


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
$ ./test-potions
```

# Test

```
$ go test ./...
$ go test ./... -cover
```