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

- docker pull akirraint/api-server:latest
- docker run --rm -p 8081:8081 akirraint/api-server:latest
