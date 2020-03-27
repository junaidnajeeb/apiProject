
## Setup go path in VS code
settings.json file
```json
{
  "window.zoomLevel": 1,
  "editor.renderWhitespace": "all",
  "files.trimTrailingWhitespace": true,
  "editor.renderControlCharacters": false,
  "breadcrumbs.enabled": true,
  "editor.minimap.enabled": true,
  "editor.insertSpaces": false,
  "editor.detectIndentation": false,
  "diffEditor.ignoreTrimWhitespace": true,
  "go.formatTool": "goimports",
  "go.useLanguageServer": true,
  "go.gopath": "/Users/junaid/Documents/gitRepos/go-workspace",
}
```

## Setup go path in .bash_profile
```bash
# GOLANG SETTINGS
export GOPATH=$HOME/Documents/gitRepos/go-workspace # don't forget to change your path correctly!
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
export GOBIN=$GOPATH/bin

```
## Under workspace->src directory
```bash
	go install apiProject
```
## Run from src folder
```bash
	./apiProject
```

## Dependency installed
```
	go get -u github.com/gorilla/mux
	go get -u github.com/stamblerre/gocode
	go get -u github.com/ramya-rao-a/go-outline
	go get -u github.com/sqs/goreturns
	go get -u github.com/spf13/viper
	go get -u github.com/go-sql-driver/mysql
	go get -u github.com/jinzhu/gorm
	go get -u github.com/sirupsen/logrus
```

# API details


**GET home**
----
  Home/Index call

* **URL**

  /

* **Method:**

  `GET`

*  **URL Params**

   **Required:**

   None

   **Optional:**

   None

* **Data Params**

  None

* **Success Response:**

  ```json
	{
		"home": {
			"version": "1.0",
			"message": "Welcome home!"
		},
		"message": "Home page",
		"status": true
	}
  ```

**Fetch one Account**
----
  call to get account details by id

* **URL**

  /accounts/id

* **Method:**

  `GET`

*  **URL Params**

   **Required:**

   `id=integer`

   **Optional:**

   None

* **Data Params**

  None

* **Success Response:**

  `200 OK`
  ```json
  {
    "account": {
        "id": 1,
        "name": "account name 2",
        "description": "Some Description of the account",
        "status": "active",
        "created_at": "2020-03-26T15:17:08-07:00",
        "updated_at": "2020-03-26T15:17:08-07:00",
        "deleted_at": null
    },
    "message": "Success",
    "status": true
  }
  ```
* **Error Response:**

  ```json
  {
    "message": "Account not found",
    "status": false
  }
  ```

**Create Account**
----
  call to create a new account

* **URL**

  /accounts

* **Method:**

  `POST`

*  **URL Params**

   **Required:**

   None

   **Optional:**

   None

* **Data Params**

  ```json
  {
	"Name":"account name 4",
	"Description":"Description of the account",
	"Status" : "active"
  }
  ```

* **Success Response:**

  `201 Created`
  ```json
  {
    "account": {
        "id": 17,
        "name": "account name 4",
        "description": "Some Description of the account",
        "status": "inactive",
        "created_at": "2020-03-26T23:34:55.160182-07:00",
        "updated_at": "2020-03-26T23:34:55.160182-07:00",
        "deleted_at": null
    },
    "message": "Account has been created",
    "status": true
  }
  ```
* **Error Response:**

  ```json
  {
    "message": "Failed to create account: Error 1062: Duplicate entry 'account name 4' for key 'name'",
    "status": false
  }
  ```

#### POST create/update User
```javascript
curl -v --location --request POST 'http://localhost:8080/users' \
> --header 'Content-Type: application/json' \
> --data-raw '{
>     "id": 100,
>     "Name": "Some Name updated",
>     "Description": "Some Description"
> }'
> POST /users HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
> Content-Type: application/json
> Content-Length: 89
>
< HTTP/1.1 201 Created
< Content-Type: application/json
< Date: Mon, 23 Mar 2020 19:55:55 GMT
< Content-Length: 38
<
{"id":100,"name":"Some Name updated"}
```

#### GET All Users
```javascript
curl -v --location --request GET 'http://localhost:8080/users/'
> GET /users/ HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: /users
< Date: Mon, 23 Mar 2020 19:57:33 GMT
< Content-Length: 41
<
> GET /users HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Mon, 23 Mar 2020 19:57:33 GMT
< Content-Length: 46
<
{"100":{"id":100,"name":"Some Name updated"}}
```

#### GET One User
```javascript
curl -v --location --request GET 'http://localhost:8080/users/100'
> GET /users/100 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Mon, 23 Mar 2020 19:58:59 GMT
< Content-Length: 38
<
{"id":100,"name":"Some Name updated"}
```

#### DELETE One User
```javascript
curl -v --location --request DELETE 'http://localhost:8080/users/100' \
> --header 'Content-Type: application/json'
> DELETE /users/100 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
> Content-Type: application/json
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Mon, 23 Mar 2020 20:00:48 GMT
< Content-Length: 10
<
"Deleted"
```

##### User if not exist
```javascript
curl -v --location --request DELETE 'http://localhost:8080/users/100' \
> --header 'Content-Type: application/json'
> DELETE /users/100 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
> Content-Type: application/json
>
< HTTP/1.1 404 Not Found
< Content-Type: application/json
< Date: Mon, 23 Mar 2020 20:00:52 GMT
< Content-Length: 58
<
{"ErrorMessage":"Requested user not found","ErrorCode":0}
```

### Visual Studio Code launch.json
```json
{
	// Use IntelliSense to learn about possible attributes.
	// Hover to view descriptions of existing attributes.
	// For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
	"version": "0.2.0",
	"configurations": [
		{
			"name": "Launch",
			"type": "go",
			"request": "launch",
			"mode": "auto",
			"program": "${fileDirname}",
			"env": {
				"API_PORT" : "8089"
			},
			"args": []
		}
	]
}
```