# CloudBread
Simple Microservices of cloud bread product using golang, gorilla/mux and postgresql databases, deployed on heroku
<br>
Incoming update, i'll create auth service to serve jwt web token
## Endpoint
| Method  | Endpoint | Description |
| ------------- | ------------- | ------------- |
| GET  | /breads/  | display all product data |
| GET by id  | /breads/{id}  | display specified product data |
| POST  | /breads/  | create product data |
| PUT  | /breads/{id}  | update specified product data |
| DELETE by id  | /breads/{id}  | delete specified product data |
