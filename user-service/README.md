## User Service (SET,GET, LIST)


To run this project
```
go run/cmd/main.go -port 8080
```


To test the APIs

__Create Users__
##### Request
```
curl --location 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic YWRtaW46cGEkJHdvckQxODQ=' \
--data '{"id":"1","name":"Virat Kohli","signupTime":17354381702377}'
```
##### Response
```
{
    "id": "1",
    "name": "Virat Kohli",
    "signupTime": 17354381702377
}
```

__Get Single User__

##### Request
```
curl --location 'http://localhost:8080/users/1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic YWRtaW46cGEkJHdvckQxODQ=' \
--data ''
```

##### Response
```
{
    "id": "1",
    "name": "Virat Kohli",
    "signupTime": 1704532200000
}
```


__Get Users Lists__

##### Request
```
curl --location 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic YWRtaW46cGEkJHdvckQxODQ=' \
--data ''
```

##### Response
```
[
    {
        "id": "1",
        "name": "Virat Kohli",
        "signupTime": 1704532200000
    },
    {
        "id": "2",
        "name": "MS Dhoni",
        "signupTime": 1704532200000
    }
]
```