# Golang restful API with GIN+GORM
This repository is a part of my learning process. Decided to use golang because my laptop terribly slow when opening Spring Boot project. Using mysql as database (of course you can change this when configuring it using gorm)

## What's in this API?
Here some documentation:  

### /api/auth/login
body: email, password  
method: POST  
example request:  
```
{
	"email":"yudhanewbie@gmail.com",
	"password":"yudhanewbie"
}
```
example response (OK):  
```
{
  "status": true,
  "message": "OK!",
  "error": null,
  "data": {
    "id": 8,
    "fullname": "Prieyudha Akadita S",
    "email": "yudhanewbie@gmail.com",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiOCIsImlzX2FkbWluIjpmYWxzZSwiZXhwIjoxNjA3MTQyMjAzLCJpYXQiOjE2MDY4ODMwMDMsImlzcyI6InlkaG53YiJ9.f0t1Ga5LYgkshnlrYYY7ByfWRd9wI8jfz86MPc8rARo"
  }
}
```
example response (error):  
```
{
  "status": false,
  "message": "Cannot authenticate! Check again your credentials",
  "error": [
    "Invalid credentials"
  ],
  "data": null
}
```

### /api/auth/
body: email, password  
method: POST  
example request:  
```
{
	"email":"test7@gmail.com",
	"password": "yudhanewbie",
	"fullname": "Prieyudha Akadita S"
}
```
example response (ok):  
```
{
  "status": true,
  "message": "OK!",
  "error": null,
  "data": {
    "id": 17,
    "fullname": "Prieyudha Akadita S",
    "email": "test7@gmail.com",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiMCIsImlzX2FkbWluIjpmYWxzZSwiZXhwIjoxNjA3MDkyNDc1LCJpYXQiOjE2MDY4MzMyNzUsImlzcyI6InlkaG53YiJ9.DM5b69kCZF_4HuTk28dQgn3ea7-HU269q4jBDKT2LJA"
  }
}

```
example response (error: duplicate email):
```
{
  "status": false,
  "message": "Failed to process your data",
  "error": [
    "Duplicate email"
  ],
  "data": null
}
```
