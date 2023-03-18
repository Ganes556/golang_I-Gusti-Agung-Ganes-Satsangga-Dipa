
-- Delete dulu tabel transaction_details dengan product id yang sama 
-- Karena tabel transaction_details menyimpan FK yang merupakan PK dari tabel products
DELETE FROM transaction_details WHERE product_id = 1;

-- Delete data pada tabel products dengan id = 1
DELETE FROM products WHERE product_id = 1;



-- Delete dulu tabel transaction_details dengan product_type_id yang sama
-- Karena tabel transaction_details menyimpan FK yang merupakan PK dari tabel products
DELETE FROM transaction_details t WHERE t.product_id IN (SELECT product_id FROM products WHERE product_type_id = 1);

-- Baru Delete pada pada tabel products dengan product type id 1
DELETE FROM products WHERE product_type_id = 1;