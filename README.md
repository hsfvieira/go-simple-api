# Simple API with Golang

### Routes:

- (GET) `/users/`: List all users
- (GET) `/users/:id`: Show user by id
- (POST) `/users/`: Register new user. Example: `{"nome": "Teste", "idade": 11}`
- (DELETE) `/users/:id`: Remove user by id
- (PUT) `/users/:id`: Update user by id. Example: `{"nome": "Teste", "idade": 11}`