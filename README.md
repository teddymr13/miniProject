# Mini Project Golang CAP 2022 By Celerates

## Component Mini Project
- CRUD Inventory
- Login
- Register
- JWT Token
- Hexagonal architecture

### Structure File Project Hexagonal Architecture Inventory 
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

### Database Name
```sql
miniproject
```

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
	('test1@gmail.com','abc123','admin',2022-08-09 14:07:12.497),
	('testUser@gmail.com','abc123','user',2022-08-09 14:07:12.497);
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

### POST Inventory (CREATE)
```
localhost:9000/inventory
```

### Example Body POST Inventory
```
{
    "name": "Asus ROG",
    "categoryItem": "Laptop",
    "stock": 10,
    "price": 10000,
    "status": "1"
}
```

### EXAMPLE Rest Api Create POST Inventory 
```
{
    "code": 200,
    "message": "Success create data inventory",
    "data": {
        "itemID": 6,
        "name": "Asus ROG",
        "categoryItem": "Laptop",
        "stock": 10,
        "price": 10000,
        "status": "1"
    }
}
```

### GET All Inventory with Pagination (READ) 
```
localhost:9000/inventory?page=0&limit=4
```
### Example Rest Api Get All Inventory with Pagination
```
{
    "code": 200,
    "message": "Success get all data inventory",
    "data": {
        "limit": 4,
        "page": 0,
        "total_rows": 7,
        "total_pages": 1,
        "rows": [
            {
                "itemID": 1,
                "name": "Asus Zenbook",
                "categoryItem": "Laptop",
                "stock": 100,
                "price": 9000,
                "status": "1"
            },
            {
                "itemID": 2,
                "name": "Acer Swift",
                "categoryItem": "Laptop",
                "stock": 100,
                "price": 7000,
                "status": "1"
            },
            {
                "itemID": 3,
                "name": "Lenovo IdeaPad",
                "categoryItem": "Laptop",
                "stock": 100,
                "price": 6000,
                "status": "1"
            },
            {
                "itemID": 4,
                "name": "Corsair ssd 225Gb",
                "categoryItem": "Ram",
                "stock": 100,
                "price": 7000,
                "status": "1"
            }
        ]
    }
}
```

### GET By itemID Inventory (READ BY ID)
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

### PUT Inventory (UPDATE)
```
localhost:9000/inventory/1
```
### Example Body PUT Inventory
```
{
    "name": "Asus Zenbook",
    "categoryItem": "Laptop",
    "stock": 110,
    "price": 9000,
    "status": "1"
}
```

### Example Rest Api PUT By itemID Inventory
```
{
    "code": 200,
    "message": "Success get data inventory by item id",
    "data": "7"
}{
    "itemID": 7,
    "name": "Asus Zenbook",
    "categoryItem": "Laptop",
    "stock": 110,
    "price": 9000,
    "status": "1"
}
```

### DELETE Inventory
```
localhost:9000/inventory/6
```
### Example Rest Api Get By itemID Inventory
```
{
    "code": 200,
    "message": "Success delete data inventory by item id 6",
    "data": null
}
```

### Register
```
localhost:9000/register
```
### Example Body Register
```
{
    "email": "test1@gmail.com",
    "password": "abc123",
    "role": "admin"
}
```

### Example Rest Api Register
```
{
    "code": 200,
    "message": "Register Success",
    "data": {
        "id": 1,
        "email": "test1@gmail.com",
        "role": "admin",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo3fQ._aC0pIGssTV5v-H6BQnd9oC-znivHJNBfwwt7cerHDo"
    }
}
```

### Login
```
localhost:9000/login
```
### Example Body PUT Inventory
```
{
   "email": "test1@gmail.com",
   "password": "abc123"
}
```

### Example Rest Api Login
```
{
    "code": 200,
    "message": "Login Success",
    "data": {
        "id": 1,
        "email": "test1@gmail.com",
        "role": "admin",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo3fQ._aC0pIGssTV5v-H6BQnd9oC-znivHJNBfwwt7cerHDo"
    }
}
```

