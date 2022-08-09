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
   ('Lenovo IdeaPad','Laptop',100,6000,1),
   ('Corsair ssd 225Gb','Ram',100,7000,1),
   ('vGen ssd 225Gb','Ram',100,5000,1),
   ('Kingston ssd 225Gb','Ram',100,4000,1);
   
```

## Rest Api Inventory
### GET All Inventory
```
localhost:9000/inventory?page=1&limit=4
```
### Example Rest Api Get All Inventory
```
{
    "code": 200,
    "message": "Success get all data inventory",
    "data": {
        "limit": 3,
        "page": 1,
        "total_rows": 6,
        "total_pages": 1,
        "rows": [
            {
                "itemID": 4,
                "name": "Corsair ssd 225Gb",
                "categoryItem": "Ram",
                "stock": 100,
                "price": 7000,
                "status": "1"
            },
            {
                "itemID": 5,
                "name": "vGen ssd 225Gb",
                "categoryItem": "Ram",
                "stock": 100,
                "price": 5000,
                "status": "1"
            },
            {
                "itemID": 6,
                "name": "Kingston ssd 225Gb",
                "categoryItem": "Ram",
                "stock": 100,
                "price": 4000,
                "status": "1"
            }
        ]
    }
}
```

### GET By itemID Inventory
```
localhost:9000/inventory/1
```
### Example Rest Api Get By itemID Inventory
```
{
    "code": 200,
    "message": "Success get data inventory by item id",
    "data": "1"
}{
    "itemID": 1,
    "name": "Asus Zenbook",
    "categoryItem": "Laptop",
    "stock": 100,
    "price": 9000,
    "status": "1"
}
```

### PUT Inventory
```
localhost:9000/inventory/1
```
### Example Rest Api Get By itemID Inventory
```
{
    "code": 200,
    "message": "Success get data inventory by item id",
    "data": "1"
}{
    "itemID": 1,
    "name": "Asus Zenbook",
    "categoryItem": "Laptop",
    "stock": 110,
    "price": 9000,
    "status": "1"
}
```
