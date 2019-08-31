# go-microservice

This repo contains a simple microservice with the following API specification.

### Add article
- Method: `POST`
- Path: `/articles`
- Request Body:
```
{
    "title": "this is your title",
    "content": "Why you even read me?",
    "author": "Nihla",
}
```
- Response Header: `HTTP 201`
- Response Body:
```
{
    "status": 201,
    "message": "Success",
    "data": {
      "id": <article_id>
    }
}
```
or
- Response Header: `HTTP <HTTP_CODE>`
- Response Body:
```
{
    "status": <HTTP_CODE>,
    "message": <ERROR_DESCRIPTION>,
    "data": null
}
```

### Get an article by its id
- Method: `GET`
- Path: `/articles/<article_id>`
- Response Header: `HTTP 200`
- Response Body:
```
{
    "status": 200,
    "message": "Success",
    "data": [
      {
        "id": <article_id>,
        "title":<article_title>,
        "content":<article_content>,
        "author":<article_author>,
      }
    ]
}
```
or
- Response Header: `HTTP <HTTP_CODE>`
- Response Body:
```
{
    "status": <HTTP_CODE>,
    "message": <ERROR_DESCRIPTION>,
    "data": null
}
```

### Get all articles
- Method: `GET`
- Path: `/articles`
- Response Header: `HTTP 200`
- Response Body:
```
{
    "status": 200,
    "message": "Success",
    "data": [
      {
        "id": <article_id>,
        "title":<article_title>,
        "content":<article_content>,
        "author":<article_author>,
      },
      {
        "id": <article_id>,
        "title":<article_title>,
        "content":<article_content>,
        "author":<article_author>,
      }
    ]
}
```
or
- Response Header: `HTTP <HTTP_CODE>`
- Response Body:
```
{
    "status": <HTTP_CODE>,
    "message": <ERROR_DESCRIPTION>,
    "data": null
}
```

### Setting up

Note: Make sure you have Golang and MySQL installed.

(1). Clone this repository in to `src/github.com/nihlaakram` directory of your `$GOPATH`.

(2). Export the required environment variables.

```
// If you want to run tests
export TEST_DB_USER=""
export TEST_DB_NAME=""
export TEST_DB_HOST=""
export TEST_DB_PORT=""
export TEST_DB_PASS=""

// If you want to build the executable
export DB_USER=""
export DB_NAME=""
export DB_HOST=""
export DB_PORT=""
export DB_PASS=""

```
(3). Run `start.sh test` to run tests and `start.sh` to build and run the executable.

### Test Enviornment
* GOOS - Darwin
* GoLang Version 1.12.9
* MySQL 6.3

<br/>
<br/>
<br/>
<br/>
Coded with Love,
<br/> Nihla Akram
