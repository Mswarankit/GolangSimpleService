## URL Shortener

To run this url-shortener

```
go run main.go url_shortener.go
```

__Shorten URL__

##### Request
```
curl --location 'http://localhost:8080/shorten' \
--header 'Content-Type: application/json' \
--data '{
    "url": "https://www.google.com/search?q=golang+tutorial"
}
'
```

##### Response

```
{"short_url":"nRiaSct_y1g="}
```


##### Request

```
curl --location --request GET 'http://localhost:8080/nRiaSct_y1g=' 
```



__Checking Metrics__

##### Request
```
curl --location 'http://localhost:8080/metrics' \
--data ''
```

##### Response
```
{"top_domains":[{"domain":"google.com","count":3}]}
```