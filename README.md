# Letter Counter Web Service
Simple web service for lettercounter

## Getting Started 
1. run ``go get github.com/IndBrony/lettercounterwebservice`` to pull the package
2. then ``go run github.com/IndBrony/lettercounterwebservice`` to run the service 
3. you're good to go

## Endpoint
 - url: {host}:8080/
 - method: POST
 - Content-Type: application/x-www-form-urlencoded
 - request.body : text
 - response: {vocals, consonants}  ->  json format
