## TODO list

### Auth
1. Add cache to store token after it's been verified and minimize amount of future requests.
2. Add store and reset OAuth cookies (on login and logout).
3. 	SetCustomClaims to token
- test auth after SetFirebaseUID, does token (access or id?) contains firebaseUID
- Add key to (session) cache if SetFirebaseUID doesn't work

4. How to set or create user with roles or privileges in OIDC (OAuth) flow ? Set them to token claims.
- Sign in flow first. Set roles and privileges (permissions)
- ??? Create user with given roles and privileges. 
- Optional: send email with confirmation link (with expiration time).

### Tests

#### Unit tests
1. add service layer test that is mostly copy of repo tests.
2. add service facade layer test that is partially copy of repo tests.