USE vand_interview_crud;

CREATE TABLE USER (
  user_id INT PRIMARY KEY AUTO_INCREMENT,
  user_username VARCHAR(60),
  user_password VARCHAR(130),
  user_email VARCHAR(60),
  user_role VARCHAR(60)
);

CREATE TABLE STORE (
    store_id INT PRIMARY KEY AUTO_INCREMENT,
    store_name VARCHAR(60),
    store_description TEXT,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES USER(user_id)
);

CREATE TABLE PRODUCT (
    product_id INT PRIMARY KEY AUTO_INCREMENT,
    product_name varchar(60),
    product_price INT,
    product_variant varchar(60),
    store_id INT,
    FOREIGN KEY (store_id) REFERENCES STORE(store_id)
);