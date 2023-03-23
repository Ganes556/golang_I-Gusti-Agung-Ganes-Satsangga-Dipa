# Redis (Key-Value Database)

- ## CREATE data users

  ```
    HSET users username1 budi password1 bobby123 username2 ratni password2 helboy123
  ```

- ## Get semua data users

  ```
    HGETALL users
  ```

- ## Get data user1

  ```
    HSCAN users 0 MATCH *1*
  ```

- ## Update data users

  ```
    HSET users username1 budiman password1 budi1235 username2 ratna password2 helboy124
  ```

- ## Delete data users

  ```
    DEL users
  ```

# Neo4j (Graph Database)

- ## Insert data users

  ```
    CREATE (:users {username: "budi", password: "bobby123"}),
       (:users {username: "ratni", password: "helboy123"});
  ```

- ## GET semua data users

  ```
    MATCH (users) RETURN *;
  ```

- ## Get data user1

  ```
    MATCH (users {username: "budi"}) RETURN *;
  ```

- ## Update data users

  ```
    MATCH (u:users {username: "budi"}) SET u.username = "budiman" RETURN u;
  ```

- ## Delete data users
  ```
    MATCH (users) DELETE users;
  ```

# Casandra (Column Database)

- ## Create keyspace

  ```
    CREATE KEYSPACE IF NOT EXISTS simple_keyspace WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

  ```

- ## Use keyspace

  ```
    USE simple_keyspace;
  ```

- ## Create table users
  ```
    CREATE TABLE users (
      id UUID PRIMARY KEY,
      username text,
      password text
    );
  ```
- ## Insert data users

  ```
    BEGIN BATCH
      INSERT INTO users (id, username, password) VALUES (uuid(), 'budi', 'bobby123');
      INSERT INTO users (id, username, password) VALUES (uuid(), 'ratni', 'helboy123');
    APPLY BATCH;
  ```

- ## Get semua data users

  ```
    SELECT * FROM users;
  ```

- ## Update data

  ```
    UPDATE users SET username = 'ratnawati' WHERE id = 522f720f-c50b-48ac-9717-6809881a3dbf;
  ```

- ## Delete data users

  ```
    DELETE FROM users WHERE id = 522f720f-c50b-48ac-9717-6809881a3dbf;
  ```
