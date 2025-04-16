## TODO list

### Auth
1. SetCustomClaims to token (for example, SetFirebaseUID)
Test: does token contain custom claims ?
2. Sign-in flow. Set roles and privileges (permissions)

### Integration tests
Add integration tests using https://github.com/ory/dockertest

### Controller layer & Router

Add router and connect with controller (handlers) layer.

### App errors wrapper

Add errors wrapper. It allows to return error description by API handlers.

### Graceful shutdown

Add graceful shutdown for opened connections and http server.