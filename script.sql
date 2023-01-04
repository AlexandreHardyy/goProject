CREATE TABLE IF NOT EXISTS Product (
   Id serial PRIMARY KEY,
   name VARCHAR(255) NOT NULL UNIQUE,
   price REAL NOT NULL,
   createdAt TIMESTAMP NOT NULL,
   updatedAt TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS Payment (
   Id serial PRIMARY KEY,
   productId INTEGER NOT NULL REFERENCES Product(id),
   pricePaid REAL NOT NULL,
   createdAt TIMESTAMP NOT NULL,
   updatedAt TIMESTAMP NOT NULL
);

-------- FIXTURES --------

INSERT INTO product(name, price, createdAt, updatedAt)
VALUES
    ('Lampe connectée', 30, now(), now()),
    ('Iphone 14 Pro Max', 1300, now(), now()),
    ('Batterie externe', 45.50, now(), now()),
    ('TV Oled 4k', 1670.99, now(), now()),
    ('MacBook Pro 13 M1', 1500, now(), now()),
    ('Chiffonette Apple (pas trop cher)', 25, now(), now());
   
    