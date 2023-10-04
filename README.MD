## Test Automation training server.

Provides simple REST API server with base login, logout and retrieve data if valid access token provided.
Could be used for API test automation training courses.

Contains example data, which could be retrived using GET /api/events request.

Implementation is simplified so shouldn't be used as example to build REST API.

Check OPEN API spec in /api folder.

## Watch Open API specification: 

- open https://editor.swagger.io
- Copy paste to it /api/openapi.yaml

## Run using pre-built docker image

- docker pull akirraint/api-automation-demo:latest
- docker run --rm -p 8081:8081 akirraint/api-automation-demo:latest

## Run using docker-compose.yml with the app and SwaggerUI:

- docker-compose up (optional -d flag) in project root folder

  App: http://localhost/8081/auth/login (/api/events, etc)
  SwaggerUI: http://localhost/8087
