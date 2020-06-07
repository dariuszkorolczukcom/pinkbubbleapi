CREATE TABLE users ( id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, email VARCHAR (100) NOT NULL UNIQUE, password VARCHAR (255) NOT NULL, role INT UNSIGNED, first_name varchar (255), last_name varchar (255), created_at DATETIME, updated_at DATETIME, deleted_at DATETIME );

CREATE TABLE categories ( id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, name VARCHAR (100), active INT, position INT, description VARCHAR (255), created_at DATETIME, updated_at DATETIME, deleted_at DATETIME );

CREATE TABLE products ( id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, name VARCHAR (100) NOT NULL, short_description VARCHAR (255), description TEXT, active INT, position INT, navcolor VARCHAR (20) NOT NULL, price INT UNSIGNED NOT NULL, category_id INT UNSIGNED NOT NULL, FOREIGN KEY (category_id) REFERENCES categories (id), created_at DATETIME, updated_at DATETIME, deleted_at DATETIME );

CREATE TABLE images ( id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, name VARCHAR (255), product_id INT UNSIGNED NOT NULL, active INT, position INT, created_at DATETIME, updated_at DATETIME, deleted_at DATETIME, FOREIGN KEY (product_id) REFERENCES products (id) );

CREATE TABLE orders ( id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, user_id INT UNSIGNED NOT NULL, status INT NOT NULL, price INT UNSIGNED NOT NULL, FOREIGN KEY (user_id) REFERENCES `users` (`id`), created_at DATETIME, updated_at DATETIME, deleted_at DATETIME );

CREATE TABLE order_items ( id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, order_id INT UNSIGNED NOT NULL, product_id INT UNSIGNED NOT NULL, quantity INT NOT NULL, price INT UNSIGNED NOT NULL, FOREIGN KEY (order_id) REFERENCES orders(id), FOREIGN KEY (product_id) REFERENCES products (id), created_at DATETIME, updated_at DATETIME, deleted_at DATETIME );