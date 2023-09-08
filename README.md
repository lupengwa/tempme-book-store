# Start up docker postgres db
- use the README.md in dat package to create db and insert data

# Test APIs
- import bookstore api.postman_collection to postman
- to create order, users need to provide bookid and quantity in to body

# About Test
- I only unit tests in order_test.go as an example

# API function
## Check service liveness
Get path: "/liveness"

## Add new user
Post path: "/users" 
Body like: 
{
"email":"test1@gmail.com"
}

## Get all books
Get path: "/books"

## Create orders
Post path "/order"
Header required, example User:user@test.com
Body like:
{
    "items":[
        {
        "bookId" : "8249bf7c-d99d-46fd-88ac-103f4549e228",
        "quantity": 2
        }
    ]
}

## Find orders
Get path "/order"
Header required, example User:user@test.com

