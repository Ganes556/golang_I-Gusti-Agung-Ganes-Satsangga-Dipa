-- Gabungkan data transaksi dari user id 1 dan user id 2.

SELECT * FROM transactions WHERE user_id = 1 UNION SELECT * FROM transactions WHERE user_id = 2;

-- Tampilkan jumlah harga transaksi user id 1.

SELECT SUM(total_price) FROM transactions WHERE user_id = 1;

-- Tampilkan total transaksi dengan product type 2.

SELECT COUNT(t.transaction_id) FROM transactions t JOIN transaction_details td ON t.transaction_id = td.transaction_id JOIN products p ON p.product_id = td.product_id WHERE p.product_type_id = 2;

-- Tampilkan semua field table product dan field name table product type yang saling berhubungan.

SELECT p.*, pt.name FROM products p JOIN product_types pt ON p.product_type_id = pt.product_type_id;

-- Tampilkan semua field table transaction, field name table product dan field name table user.

SELECT t.*, p.name AS product_name, u.name AS username FROM transactions t JOIN transaction_details td ON t.transaction_id = td.transaction_id JOIN products p ON td.product_id = p.product_id JOIN users u ON t.user_id = u.user_id;

-- Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.

DELIMITER $$
CREATE FUNCTION del_transactions(transaction_id INT) 
RETURNS INT DETERMINISTIC
BEGIN
  DELETE FROM transaction_details td WHERE td.transaction_id = transaction_id;
  DELETE FROM transactions t WHERE t.transaction_id = transaction_id;
  RETURN 1;
END$$
DELIMITER ;



-- Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.

DELIMITER $$
CREATE FUNCTION update_total_qty(transaction_id INT)
RETURNS INT DETERMINISTIC
BEGIN 
  DECLARE total_qty_now INT;
  UPDATE transactions t 
  SET total_qty = (
    SELECT COALESCE(SUM(qty),0) FROM transaction_details td
    WHERE td.transaction_id = transaction_id
    )
  WHERE t.transaction_id = transaction_id;
 SELECT total_qty INTO total_qty_now FROM transactions t WHERE t.transaction_id = transaction_id;
RETURN total_qty_now;
END$$
DELIMITER ;

-- Fungsi tambahan biar lengkap
DELIMITER $$
CREATE FUNCTION update_total_price(transaction_id INT)
RETURNS INT DETERMINISTIC
BEGIN 
  DECLARE total_price_now INT;
  UPDATE transactions t 
  SET total_price = (
    SELECT COALESCE(SUM(td.price),0) FROM transaction_details td
    JOIN products p ON p.product_id = td.product_id
    WHERE td.transaction_id = transaction_id
    )
  WHERE t.transaction_id = transaction_id;
  SELECT total_price INTO total_price_now FROM transactions t WHERE t.transaction_id = transaction_id;
RETURN total_price_now;
END$$
DELIMITER ;

-- Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.

SELECT * FROM products WHERE product_id NOT IN (SELECT product_id FROM transaction_details);
