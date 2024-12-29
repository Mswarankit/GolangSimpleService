## User Service (SET,GET, LIST)


To run this project
```
go run/cmd/main.go -port 8080
```


To test the APIs
# create user
```
curl -X POST -u admin:pa$$worD184 http://localhost:8080/users -d 
'{"id":"1","name":"Virat Kohli","signupTime":17354681722377}'

curl -u admin:pa$$worD184 http://localhost:8080/users/1

curl  -u admin:pa$$worD184 http://localhost:8080/users

```
