# **Ringkasan Materi Join, Union, Agregasi, Subquery, Function (DBMS)**

# **1. Join**

- Sebuah klausa untuk mengkombinasikan dari dua atau lebih tabel
- Jenis Join standar SQL :
  - Inner Join : Mengembalikan baris-baris dari dua tabel atau lebih yang memenuhi syarat
    - Contoh :
      ```sql
      SELECT * FROM table1
      INNER JOIN table2
      ON table1.column_name = table2.column_name;
      ```
  - Left Join : Mengembalikan semua baris dari tabel kiri yang dikenai kondisi ON dan hanya baris dari baris tabel di kanan yang memenuhi kondisi join
    - Contoh :
      ```sql
      SELECT * FROM table1
      LEFT JOIN table2
      ON table1.column_name = table2.column_name;
      ```
  - Right Join : Mengembalikan semua baris dari tabel kanan yang dikenai kondisi ON dan hanya baris dari baris tabel di kiri yang memenuhi kondisi join
    - Contoh :
      ```sql
      SELECT * FROM table1
      RIGHT JOIN table2
      ON table1.column_name = table2.column_name;
      ```

# **2. Union**

- Menggabungkan hasil dari dua atau lebih SELECT statement
- Syarat:
  - Jumlah kolom harus sama
  - Tipe data dari kolom harus sama (tapi bila memang bisa di konversi ke tipe data yang sama, maka bisa)
  - Urutan kolom harus sama
- Contoh :
  ```sql
  SELECT column_name FROM table1
  UNION
  SELECT column_name FROM table2;
  ```

# **3. Fungsi Agregasi**

- Fungsi di mana nilai dari beberapa baris dikelompokan bersama untuk membentuk nilai ringkasan tunggal
- Beberapa fungsi agregasi yang umum digunakan :
  - COUNT() : jumlah baris yang tidak kosong
  - SUM() : jumlah dari nilai dalam kolom
  - AVG() : rata-rata dari nilai dalam kolom
  - MAX() : nilai terbesar dalam kolom
  - MIN() : nilai terkecil dalam kolom
- Contoh penggunaan fungsi agregasi `COUNT`:
  ```sql
  SELECT COUNT(column_name) FROM table_name;
  ```
- HAVING : digunakan untuk menambahkan kondisi pada fungsi agregasi
  - Contoh penggunaan fungsi agregasi `SUM`:
  ```sql
  SELECT user_id FROM tweets GROUP BY user_id HAVING SUM(retweet_count) > 2;
  ```

# **4. Subquery**

- Subquery adalah query yang berada di dalam query lain
- Subquery dapat berada di dalam SELECT, FROM, WHERE, HAVING, dan JOIN
- Contoh:
  ```sql
  SELECT * FROM table1
  WHERE column_name IN (SELECT column_name FROM table2);
  ```

# **5. Trigger**

- Trigger adalah sebuah prosedur yang akan dijalankan secara otomatis ketika terjadi suatu event
- Jenis-jenis trigger :
  - BEFORE : trigger akan dijalankan sebelum event terjadi
  - AFTER : trigger akan dijalankan setelah event terjadi
- Contoh :
  ```sql
    DELIMITER $$
    CREATE TRIGGER after_del_transactions
    BEFORE DELETE ON transactions
    FOR EACH ROW
    BEGIN
      DELETE FROM transaction_details WHERE transaction_id = OLD.transaction_id;
    END$$
    DELIMITER ;
  ```

# **6. Function**

- Sekumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya
- Contoh:

  ```sql
    -- Membuat function
    -- deterministic (private) --> kepemilikan nya si user yang login saja,
    DELIMITER $$
    CREATE FUNCTION tweet_per_user
      (user_id_p int) returns int deterministic
      BEGIN
        DECLARE total_tweet int;
        SELECT count(*) into total_tweet from tweet
        WHERE user_id = user_id_p;
        return total_tweet;
      END$$
    DELIMITER ;
  ```

  ```sql
    -- Memanggil function
    SELECT tweet_per_user(1);
  ```
