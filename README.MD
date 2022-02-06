##Test Automation training server.

Provides simple REST API server with base login, logout and retrieve data if valid access token provided.
Could be used for API test automation training courses.

Contains example data, which could be retrived using GET /api/events request.

Implementation is simplified so shouldn't be used as example to build REST API.

Check OPEN API spec in /api folder.

##Run using pre-built docker image

- docker push akirraint/api-server:1.0
- docker run --rm -p 8081:8081 akirraint/api-server:1.0