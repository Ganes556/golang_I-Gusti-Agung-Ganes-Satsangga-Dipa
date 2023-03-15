
-- CREATE DATABASE alta_online_shop
CREATE DATABASE alta_online_shop;

-- Implementasi schema Olshop yang telah dikerjakan

-- 1-to-many: user dengan alamat.
CREATE TABLE user_addresses (
  address_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  address VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE users (
  user_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  user_address_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  gender CHAR(1) NOT NULL,
  date_of_birth DATETIME,
  status ENUM('active', 'inactive', 'blocked' ,'deleted'),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_address_id) REFERENCES user_addresses(address_id)
);

CREATE TABLE product_descriptions (
  product_description_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  description TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE product_types (
  product_type_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  name VARCHAR(50),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE product_operators (
  product_operator_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  name VARCHAR(50),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE products (
  product_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  product_description_id INT NOT NULL,
  product_type_id INT NOT NULL,
  product_operator_id INT NOT NULL,
  name VARCHAR(100) NOT NULL,
  price FLOAT NOT NULL,
  stock INT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (product_description_id) REFERENCES product_descriptions(product_description_id),
  FOREIGN KEY (product_type_id) REFERENCES product_types(product_type_id),
  FOREIGN KEY (product_operator_id) REFERENCES product_operators(product_operator_id)
);

-- 1-to-1: payment method description.
CREATE TABLE payment_descriptions (
  payment_description_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  description TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE payment_methods (
  payment_method_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  payment_description_id INT NOT NULL,
  name VARCHAR(50),
  status ENUM('active', 'inactive'),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (payment_description_id) REFERENCES payment_descriptions(payment_description_id)
);

-- many-to-many: user dengan payment method menjadi user_payment_method_detail.
CREATE TABLE user_payment_method_details (
  user_payment_method_detail_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  user_id INT NOT NULL,
  payment_method_id INT NOT NULL,
  card_number VARCHAR(50),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(user_id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_methods(payment_method_id)
)

CREATE TABLE transactions (
  transaction_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  payment_method_id INT NOT NULL,
  user_id INT NOT NULL,
  total_qty INT NOT NULL,
  total_price FLOAT NOT NULL,
  status ENUM('pending', 'paid', 'canceled', 'delivered'),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(user_id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_methods(payment_method_id)
);

CREATE TABLE transaction_details (
  transaction_detail_id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  transaction_id INT NOT NULL,
  product_id INT NOT NULL,
  qty INT NOT NULL,
  price FLOAT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id),
  FOREIGN KEY (product_id) REFERENCES products(product_id)
);

-- Create tabel kurir dengan field id, name, created_at, updated_at.

CREATE TABLE kurir (
  id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  name VARCHAR(50),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

-- Tambahkan ongkos_dasar column di tabel kurir

ALTER TABLE kurir ADD COLUMN ongkos_dasar FLOAT;

-- Rename tabel kurir menjadi shipping.

ALTER TABLE kurir RENAME TO shipping;

-- Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.

DROP TABLE shipping;



