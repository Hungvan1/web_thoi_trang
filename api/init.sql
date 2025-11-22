CREATE DATABASE clothing_shop USE clothing_shop
CREATE TABLE
    users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100),
        phone CHAR(10),
        address VARCHAR(225),
        username VARCHAR(20),
        password VARCHAR(20),
        role VARCHAR(20)
    );

CREATE TABLE
    category (
        id INT AUTO_INCREMENT PRIMARY KEY,
        category_name VARCHAR(100)
    );

CREATE TABLE
    products (
        id INT AUTO_INCREMENT PRIMARY KEY,
        product_name VARCHAR(255),
        price DECIMAL(10, 2),
        number INT,
        detail VARCHAR(300),
        status VARCHAR(50),
        size VARCHAR(5),
        gender VARCHAR(20),
        color VARCHAR(50),
        category_id INT,
        user_id INT,
        image VARCHAR(255),
        FOREIGN KEY (category_id) REFERENCES category (id),
        FOREIGN KEY (user_id) REFERENCES users (id)
    );

CREATE TABLE
    orders (
        id INT AUTO_INCREMENT PRIMARY KEY,
        order_date DATETIME,
        ship_address VARCHAR(225),
        user_id INT,
        total_amount DECIMAL(10, 2),
        FOREIGN KEY (user_id) REFERENCES users (id)
    );

CREATE TABLE
    order_items (
        id INT AUTO_INCREMENT PRIMARY KEY,
        unit_price DECIMAL(10, 2),
        quantity INT,
        order_id INT,
        product_id INT,
        FOREIGN KEY (order_id) REFERENCES orders (id),
        FOREIGN KEY (product_id) REFERENCES products (id)
    );

CREATE TABLE
    review (
        review_id INT AUTO_INCREMENT PRIMARY KEY,
        text VARCHAR(300),
        user_id INT,
        product_id INT,
        order_id INT,
        FOREIGN KEY (user_id) REFERENCES users (id),
        FOREIGN KEY (product_id) REFERENCES products (id),
        FOREIGN KEY (order_id) REFERENCES orders (id)
    );