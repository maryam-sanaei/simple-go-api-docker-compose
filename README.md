# simple-go-api-docker-compose

Example api shows how to use Docker compose with Go application and mysql. The repository contains a simple api written in Golang using gin and mysql that contains two endpoints to display persons list and add a new person.


Clone the repository and type the following command to start the app

```bash
$ docker-compose up
```
Get persons list
```bash
$ curl http://localhost:8080

```

Add new person
```bash
$ curl http://localhost:8080/addperson \
--include     --header "Content-Type: application/json" \
--request "POST" --data '{"firstname": "myfirstname","lastname": "mylastname"}'

```
