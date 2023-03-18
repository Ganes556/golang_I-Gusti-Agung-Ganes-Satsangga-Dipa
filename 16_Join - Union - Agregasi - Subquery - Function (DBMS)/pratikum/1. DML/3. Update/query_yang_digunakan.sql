-- Ubah data products id 1 dengan nama ‘product dummy’

UPDATE products SET name = 'product dummy' WHERE product_id = 1;

-- Update qty = 3 pada transaction details dengan product id = 1
-- NOTE : tidak terlalu terlihat karena sudah terlanjur semua qtynya 3. Tapi secara query seharusnya benar

UPDATE transaction_details SET qty = 3 WHERE product_id = 1;


