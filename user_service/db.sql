-- Table: users

CREATE TYPE valid_roles AS ENUM ('admin', 'recipient', 'donor');

CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role valid_roles NOT NULL,
    deposit DECIMAL(10, 2) DEFAULT 0 CHECK(deposit>=0)
);

-- Table: UserDetail
CREATE TABLE IF NOT EXISTS user_details (
    user_detail_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id) UNIQUE NOT NULL,
    fname VARCHAR(255),
    lname VARCHAR(255),
    address TEXT,
    age INT DEFAULT 0 CHECK(age>=0),
    phone_number VARCHAR(20),
    profile_picture_url varchar(255)
);


--data seed 

-- Seed data for users table
INSERT INTO users (username, email, password, role, deposit) VALUES
('admin1', 'admin1@example.com', 'password1', 'admin', 100.00),
('recipient1', 'recipient1@example.com', 'password2', 'recipient', 50.00),
('donor1', 'donor1@example.com', 'password3', 'donor', 200.00),
('recipient2', 'recipient2@example.com', 'password4', 'recipient', 30.00),
('donor2', 'donor2@example.com', 'password5', 'donor', 150.00);

-- Seed data for user_details table
INSERT INTO user_details (user_id, fname, lname, address, age, phone_number, profile_picture_url) VALUES
(1, 'Admin', 'One', '123 Admin Street', 35, '123-456-7890', 'http://example.com/admin1.jpg'),
(2, 'Recipient', 'One', '456 Recipient Avenue', 28, '123-456-7891', 'http://example.com/recipient1.jpg'),
(3, 'Donor', 'One', '789 Donor Road', 40, '123-456-7892', 'http://example.com/donor1.jpg'),
(4, 'Recipient', 'Two', '101 Recipient Blvd', 22, '123-456-7893', 'http://example.com/recipient2.jpg'),
(5, 'Donor', 'Two', '202 Donor Lane', 33, '123-456-7894', 'http://example.com/donor2.jpg');