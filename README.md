# Mini Project Golang CAP 2022 By Celerates

## Component Mini Project
- CRUD Inventory
- Login
- Register
- JWT Token
- Hexagonal architecture

### Struktur project Hexagonal architecture inventory 
```
Mini Project    
│
└───app
│   │app.go   
│   
└───auth
|   │auth.go
|   
└───domain
|   |inventoryRepositoruDB.go
|   |inventory.go
|   |users.go
|
└───dto
|   |inventoryResponse.go
|   |pagination.go
|   |userResponse.go
|
└───errs
|   |errors.go
|
└───handler
|   |inventoryHandler.go
|   |usersHandler.go
| 
└───logger
|   |logger.go
|
└───service
|   |inventoryService.go
|   |usersService.go
|
|env
|go.mod
|go.sum
|main.go
|README.md
```

## Postgre SQL
Create SQL used in this project

### Table Users
```sql
CREATE TABLE "users" (
  "user_id" serial primary key not null,
  "email" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "role" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

INSERT INTO "users" (email, password, role, created_at)
VALUES
	('celerate12@gmail.com','abc123','admin',2020-08-09 10:27:22.000),
	('test13@gmail.com','abc123','user',2020-08-09 10:27:22.000);
```

### Table Inventories
```sql
CREATE TABLE "inventories" (
  "item_id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(100) NOT NULL,
  "category_item" varchar(100) NOT NULL,
  "stock" int NOT NULL,
  "price" int NOT NULL,
  "status" smallint NOT NULL DEFAULT '1'
);

INSERT INTO "inventories" (name, category_item, stock, price, status)
VALUES
	 ('Asus Zenbook','Laptop',100,9000,1),
   ('Acer Swift','Laptop',100,7000,1),
	 ('Lenovo IdeaPad','Laptop',100,6000,1);
   
```

## Rest Api Inventory
### Get Inventory
```
localhost:9000/inventory?page=1&limit=3
```
### Example Rest Api Get Inventory
```
localhost:9000/inventory?page=1&limit=3
```
