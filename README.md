### Go 

### task1

Given an input as JSON array of flat objects, the program will parse the JSON and return a nested JSON of objects of arrays with keys specified in command line arguments. The leaf values are are arrays of flat objects matching appropriate groups.

## Run
### Go

To run the first task:
```
(cd task1 && cat ../input.json | go run ./ currency country city)
```

### task2 

REST service of task1. The input is a JSON that is received in POST request, nesting paramaters are in the request params.

To run the second task:
```
(cd task2 && go run ./)
```

### Docker

```
docker build -t golang-nested-json ./
```
then
```
docker run -p=8080:8080 golang-nested-json
```

## Decisions and assumptions

Since this is a small project, we're using hardcoded values instead of environment variables.

Regarding the API, we implemented post with the nesting level as query parameter with a comma separated input, the input.json is inside the body.


For example: 
```
http://localhost:8080/nesting/json?nesting_levels=currency,country,city
```


Other assumptions and decisions are commented through the code.

