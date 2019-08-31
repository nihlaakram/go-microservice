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
export TEST_DB_USER="<test_db_user>"
export TEST_DB_NAME="<test_db_name>"
export TEST_DB_HOST="<test_db_hostname>"
export TEST_DB_PORT="<test_db_port>"
export TEST_DB_PASS="<test_db_password>"

// If you want to build the executable
export DB_USER="<db_user>"
export DB_NAME="<db_name>"
export DB_HOST="<db_hostname>"
export DB_PORT="<db_port>"
export DB_PASS="<db_password>"

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
