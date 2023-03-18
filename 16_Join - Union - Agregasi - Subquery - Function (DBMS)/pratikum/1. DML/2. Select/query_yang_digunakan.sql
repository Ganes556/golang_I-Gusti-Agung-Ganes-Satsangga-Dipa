-- Tampilkan nama users dengan gender Laki-laki / M

SELECT name FROM users WHERE gender = 'M';

-- Tampilkan products dengan id = 3

SELECT * FROM products WHERE product_id = 3;

-- Tampilkan data users yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’

SELECT * FROM users WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY) AND name LIKE '%a%';

-- Hitung jumlah users dengan status gender Perempuan

SELECT COUNT(gender) FROM users WHERE gender = 'F';

-- Tampilkan data users dengan urutan sesuai nama abjad

SELECT * FROM users ORDER BY name;

-- Tampilkan 5 data pada tabel products

SELECT * FROM products LIMIT 5;


