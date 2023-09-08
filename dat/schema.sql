CREATE SCHEMA bookstore;
CREATE TABLE bookstore.users (
   user_id UUID PRIMARY KEY,
   email VARCHAR(255) UNIQUE
);


CREATE TABLE bookstore.books (
      book_id UUID PRIMARY KEY,
      name VARCHAR UNIQUE,
      price NUMERIC (12,2)
);

CREATE TABLE bookstore.orders (
    order_id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    FOREIGN KEY (USER_ID) references bookstore.users(USER_ID)
);

CREATE TABLE bookstore.order_items (
    order_id UUID NOT NULL,
    book_id  UUID NOT NULL,
    quantity SMALLINT NOT NULL,
    FOREIGN KEY (BOOK_ID) references bookstore.books(BOOK_ID),
    FOREIGN KEY (ORDER_ID) references bookstore.orders(ORDER_ID)
);
