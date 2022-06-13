## Run
### Go

To run the first task:
```
(cd task1 && cat ../input.json | go run ./ currency country city)
```

To run the second task:
```
(cd task2 && go run ./)
```

### Docker

```
docker build -t flaconi-challenge ./
```
then
```
docker run -p=8080:8080 flaconi-challenge
```

## Decisions and assumptions

Since this is a small project, we're using hardcoded values instead of environment variables.


There was no defintion on where/ how to run the API so we took the liberty to split into two directories.


Regarding the API, we implemented post with the nesting level as query parameter with a comma separated input, the input.json is inside the body.


For example: 
```
http://localhost:8080/nesting/json?nesting_levels=currency,country,city
```


Other assumptions and decisions are commented through the code.

