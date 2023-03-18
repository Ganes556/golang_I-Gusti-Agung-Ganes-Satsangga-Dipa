-- Insert 5 operators pada table operators
INSERT INTO product_operators(name, created_at, updated_at) VALUES
('Kingz Notebook', NOW(),NOW()),
('Supplier Herbal Timur Tengah', NOW(), NOW()),
('Studio Ponsel', NOW(), NOW()),
('Lasallefood', NOW(), NOW()),
('Greenara', NOW(), NOW());

-- Insert 3 product type
INSERT INTO product_types(name, created_at, updated_at) VALUES
('handphone', NOW(), NOW()),
('electronic', NOW(), NOW()),
('food & drink', NOW(), NOW());

-- Insert product description pada setiap product
INSERT INTO product_descriptions(description, created_at, updated_at) VALUES
('Produk ini merupakan handphone dari brand Apple dengan kapasitas penyimpanan 256GB, 128GB, dan 64GB. Handphone ini hanya dapat digunakan melalui koneksi WiFi dan tersedia dalam warna merah tengah malam.', NOW(), NOW()),
('Produk ini merupakan handphone dari brand Xiaomi dengan kapasitas penyimpanan 128GB dan RAM 6GB. Handphone ini menggunakan prosesor MediaTek Helio G95 dan tersedia dalam warna abu-abu.', NOW(), NOW()),
('Produk ini merupakan laptop dari brand Asus dengan spesifikasi prosesor Ryzen 5 5600H dan RAM 8GB. Laptop ini juga dilengkapi dengan storage 512GB SSD dan kartu grafis Radeon Vega.', NOW(), NOW()),
('Produk ini merupakan laptop dari brand Infinix dengan spesifikasi prosesor Intel i3 1005G1 dan RAM 4GB. Laptop ini dilengkapi dengan storage 256GB SSD dan layar 14 inci beresolusi FHD.', NOW(), NOW()),
('Produk ini merupakan laptop dari brand Acer dengan spesifikasi prosesor Ryzen 5 5500U dan RAM 8GB. Laptop ini juga dilengkapi dengan storage 256GB SSD dan layar 15,6 inci beresolusi FHD.', NOW(), NOW()),
('Produk ini merupakan sarden berbumbu cabai dari brand Delmonte dalam kemasan 155 gram dan terdiri dari 3 sachet.', NOW(), NOW()),
('Produk ini merupakan mayonnaise dari brand Maestro dalam kemasan pouch dengan berat 180 gram dan terdiri dari 6 sachet.', NOW(), NOW()),
('Produk ini merupakan sirup rasa cocopandan dari brand Marjan dalam kemasan 460 ml dan terdiri dari 6 botol.', NOW(), NOW());

-- Insert 2 product dengan product type id = 1, dan operators id = 3
-- type-id 1 = electronic, operator-id 3 = Studio Ponsel
INSERT INTO products(product_description_id, product_type_id, product_operator_id, name, price, stock, created_at, updated_at) VALUES
(1, 1, 3, 'iPhone SE 3 2022 256GB 128GB 64GB - WIFI ONLY, 128GB BLACK', 6840000.00, 100, NOW(), NOW()),
(2, 1, 3, 'Xiaomi Poco M5s 6/128GB - GREY', 2640000.00, 100, NOW(), NOW());

-- Insert 3 product dengan product type id = 2, dan operators id = 1
-- type-id 2 = handphone, operator-id 1 = Kingz Notebook
INSERT INTO products(product_description_id, product_type_id, product_operator_id, name, price, stock, created_at, updated_at) VALUES
(3, 2, 1, 'Asus Vivobook 14X M1403QA Ryzen 5 5600H 8GB 512SS - Silver, 8GB/512SSD', 8219999.00, 5, NOW(), NOW()),
(4, 2, 1, 'Infinix INbook X2 i3 1005G1 4GB 256GB SSD - Grey', 4519000.00, 50, NOW(), NOW()),
(5, 2, 1, 'Acer Aspire 5 A515 Ryzen 5 5500U 8GB 256SSD - Silver, 8GB/256SSD', 7139000.00, 100, NOW(), NOW());

-- Insert 3 product dengan product type id = 3, dan operators id = 4
-- type-id 3 = food & drink, operator-id 4 = Lasallefood
INSERT INTO products(product_description_id, product_type_id, product_operator_id, name, price, stock, created_at, updated_at) VALUES
(6, 3, 4, 'Delmonte Sarden Chilli 155gr - isi 3', 23485.00, 100, NOW(), NOW()),
(7, 3, 4, 'Maestro Mayonnaise Pouch 180 g - isi 6', 52000.00, 100, NOW(), NOW()),
(8, 3, 4, 'Marjan Sirup Cocopandan 460 ml - isi 6', 126900.00, 50, NOW(), NOW());

-- Insert 3 payment description
INSERT INTO payment_descriptions(description, created_at, updated_at) VALUES
('Payment via bank transfer to the account number listed on the invoice such as Bank BCA, BRI, etc.', NOW(), NOW()),
('Payment via credit card registered such as Visa , Mastercard, etc.', NOW(), NOW()),
('Payment via electronic wallet such as OVO, DANA, etc.', NOW(), NOW());

-- Insert 3 payment methods
INSERT INTO payment_methods(payment_description_id, name, created_at, updated_at) VALUES
(1, 'Bank Transfer',NOW(), NOW()),
(2, 'Credit Card', NOW(), NOW()),
(3, 'Electronic Wallet', NOW(), NOW());


-- Insert 5 user address pada tabel user_addresses
INSERT INTO user_addresses(address, created_at, updated_at) VALUES
('Jalan Raya 1', NOW(), NOW()),
('Jalan Raya 2', NOW(), NOW()),
('Jalan Raya 3', NOW(), NOW()),
('Jalan Raya 4', NOW(), NOW()),
('Jalan Raya 5', NOW(), NOW());

-- Insert 5 user pada tabel users
INSERT INTO users(user_address_id, name, gender, date_of_birth, created_at, updated_at) VALUES
(1, 'Ahmad Farhan', 'M', '1990-01-01', NOW(), NOW()),
(2, 'Dewi Putri', 'F', '1992-05-15', NOW(), NOW()),
(3, 'Rudi Hartono', 'M', '1985-07-23', NOW(), NOW()),
(4, 'Sari Indah', 'F', '1998-12-10', NOW(), NOW()),
(5, 'Irfan Maulana', 'M', '1980-06-05', NOW(), NOW());


-- Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
INSERT INTO transactions(user_id, payment_method_id, status, total_qty, total_price, created_at, updated_at) VALUES
(1, 1, 'paid', 3, 16320000.00, NOW(), NOW()),
(1, 2, 'unpaid', 3, 20958998.00, NOW(), NOW()),
(1, 3, 'paid', 3, 14301485.00, NOW(), NOW()),
(2, 1, 'paid', 3, 7242000.00, NOW(), NOW()),
(2, 2, 'unpaid', 3, 20958998.00, NOW(), NOW()),
(2, 3, 'paid', 3, 16320000.00, NOW(), NOW()),
(3, 1, 'paid', 3, 7242000.00, NOW(), NOW()),
(3, 2, 'unpaid', 3, 20958998.00, NOW(), NOW()),
(3, 3, 'paid', 3, 16320000.00, NOW(), NOW()),
(4, 1, 'paid', 3, 16320000.00, NOW(), NOW()),
(4, 2, 'unpaid', 3, 20958998.00, NOW(), NOW()),
(4, 3, 'paid', 3, 7242000.00, NOW(), NOW()),
(5, 1, 'paid', 3, 20958998.00, NOW(), NOW()),
(5, 2, 'unpaid', 3, 14301485.00, NOW(), NOW()),
(5, 3, 'paid', 3, 7242000.00, NOW(), NOW());

-- Insert 3 product di masing-masing transaksi

INSERT INTO transaction_details(transaction_id, product_id, qty, price, created_at, updated_at) VALUES
(1, 1, 2, 13680000.00, NOW(), NOW()),
(1, 2, 1, 2640000.00, NOW(), NOW()),
(2, 3, 2, 16439998.00, NOW(), NOW()),
(2, 4, 1, 4519000.00, NOW(), NOW()),
(3, 5, 2, 14278000.00, NOW(), NOW()),
(3, 6, 1, 23485.00, NOW(), NOW()),
(4, 7, 2, 104000.00, NOW(), NOW()),
(4, 5, 1, 7139000.00, NOW(), NOW()),
(5, 3, 2, 16439998.00, NOW(), NOW()),
(5, 4, 1, 4519000.00, NOW(), NOW()),
(6, 1, 2, 13680000.00, NOW(), NOW()),
(6, 2, 1, 2640000.00, NOW(), NOW()),
(7, 7, 2, 104000.00, NOW(), NOW()),
(7, 5, 1, 7139000.00, NOW(), NOW()),
(8, 3, 2, 16439998.00, NOW(), NOW()),
(8, 4, 1, 4519000.00, NOW(), NOW()),
(9, 1, 2, 13680000.00, NOW(), NOW()),
(9, 2, 1, 2640000.00, NOW(), NOW()),
(10, 1, 2, 13680000.00, NOW(), NOW()),
(10, 2, 1, 2640000.00, NOW(), NOW()),
(11, 3, 2, 16439998.00, NOW(), NOW()),
(11, 4, 1, 4519000.00, NOW(), NOW()),
(12, 7, 2, 104000.00, NOW(), NOW()),
(12, 5, 1, 7139000.00, NOW(), NOW()),
(13, 3, 2, 16439998.00, NOW(), NOW()),
(13, 4, 1, 4519000.00, NOW(), NOW()),
(14, 5, 2, 14278000.00, NOW(), NOW()),
(14, 6, 1, 23485.00, NOW(), NOW()),
(15, 7, 2, 104000.00, NOW(), NOW()),
(15, 5, 1, 7139000.00, NOW(), NOW());

-- Insert payment number untuk masing-masing transaksi

INSERT INTO user_payment_method_details(user_id, payment_method_id, payment_number, created_at, updated_at) VALUES
(1, 1, '1234567890', NOW(), NOW()),
(1, 2, '1234-5678-9012', NOW(), NOW()),
(1, 3, '081234567890', NOW(), NOW()),
(2, 1, '0987654321', NOW(), NOW()),
(2, 2, '0987-6543-2109', NOW(), NOW()),
(2, 3, '089876543210', NOW(), NOW()),
(3, 1, '1234567890', NOW(), NOW()),
(3, 2, '1234-5678-9012', NOW(), NOW()),
(3, 3, '081234567890', NOW(), NOW()),
(4, 1, '0987654321', NOW(), NOW()),
(4, 2, '0987-6543-2109', NOW(), NOW()),
(4, 3, '089876543210', NOW(), NOW()),
(5, 1, '1234567890', NOW(), NOW()),
(5, 2, '1234-5678-9012', NOW(), NOW()),
(5, 3, '081234567890', NOW(), NOW());
