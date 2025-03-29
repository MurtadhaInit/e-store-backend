-- TODO: Create a new UTF-8 database.
CREATE DATABASE ecomm CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Switch to using that database.
USE ecomm;

CREATE TABLE customers (
  customer_id INT PRIMARY KEY AUTO_INCREMENT,
  first_name VARCHAR(40) NOT NULL,
  last_name VARCHAR(40) NOT NULL,
  email VARCHAR(40) NOT NULL,
  birth_day DATE NOT NULL,
  phone_number INT NOT NULL,
  address VARCHAR(40) NOT NULL,
  date_added TIMESTAMP NOT NULL
);

CREATE TABLE products (
  product_id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  description TEXT NOT NULL,
  price DECIMAL(6, 2) NOT NULL,
  stock_quantity INT NOT NULL,
  date_added TIMESTAMP NOT NULL
);

CREATE TABLE orders (
  order_id INT PRIMARY KEY AUTO_INCREMENT,
  customer_id INT FOREIGN KEY REFERENCES customers(customer_id) ON DELETE SET NULL,
  order_status VARCHAR(40) NOT NULL,
  date_placed TIMESTAMP NOT NULL,
  order_total INT NOT NULL
);

-- TODO: Add an index on the date_added column.
CREATE INDEX idx_orders_placed ON orders(date_placed);

CREATE TABLE order_items (
  order_id INT,
  product_id INT,
  quantity INT NOT NULL,
  date_added TIMESTAMP NOT NULL,
  PRIMARY KEY(order_id, product_id),
  FOREIGN KEY(order_id) REFERENCES orders(order_id) ON DELETE CASCADE,
  FOREIGN KEY(product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

CREATE TABLE carts (
  cart_id INT PRIMARY KEY,
  customer_id INT FOREIGN KEY REFERENCES customers(customer_id) ON DELETE CASCADE,
  cart_total INT NOT NULL,
  last_updated TIMESTAMP NOT NULL
);

CREATE TABLE cart_items (
  cart_id INT,
  product_id INT,
  quantity INT NOT NULL,
  date_added TIMESTAMP NOT NULL,
  PRIMARY KEY(cart_id, product_id),
  FOREIGN KEY(cart_id) REFERENCES carts(cart_id) ON DELETE CASCADE,
  FOREIGN KEY(product_id) REFERENCES products(product_id) ON DELETE CASCADE
);
