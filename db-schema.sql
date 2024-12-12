-- Users Table
CREATE TABLE users (
    user_id VARCHAR(10) PRIMARY KEY COMMENT 'Unique 10-character identifier for each user (student ID)',
    user_name VARCHAR(255) NOT NULL COMMENT 'Full name of the user',
    user_tel VARCHAR(15) COMMENT 'Contact phone number of the user',
    user_password VARCHAR(255) NOT NULL COMMENT 'Hashed password for user authentication',
    user_ban_status TINYINT(1) DEFAULT 0 COMMENT 'Ban status: 1 = banned, 0 = not banned'
);

-- Staffs Table
CREATE TABLE staffs (
    staff_id VARCHAR(10) PRIMARY KEY COMMENT 'Unique 10-character identifier for each staff',
    staff_name VARCHAR(255) NOT NULL COMMENT 'Full name of the staff member',
    staff_password VARCHAR(255) NOT NULL COMMENT 'Hashed password for staff authentication'
);

-- Admin Table
CREATE TABLE admin (
    admin_id VARCHAR(10) PRIMARY KEY COMMENT 'Unique 10-character identifier for each admin',
    admin_name VARCHAR(255) NOT NULL COMMENT 'Full name of the admin',
    admin_password VARCHAR(255) NOT NULL COMMENT 'Hashed password for admin authentication'
);

-- Items Table
CREATE TABLE items (
    item_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each item',
    item_name VARCHAR(255) NOT NULL COMMENT 'Name of the item',
    item_current_amount INT NOT NULL COMMENT 'Current stock amount of the item',
    item_image VARCHAR(255) COMMENT 'URL to the image of the item',
    item_borrow_count INT DEFAULT 0 COMMENT 'Total number of times this item has been borrowed'
);

-- Orders Table
CREATE TABLE orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each order',
    user_id VARCHAR(10) NOT NULL COMMENT 'Identifier of the user who placed the order',
    user_org VARCHAR(255) COMMENT 'Organization name of the borrower',
    borrow_datetime DATETIME NOT NULL COMMENT 'Datetime when the items are borrowed',
    return_datetime DATETIME NOT NULL COMMENT 'Scheduled return datetime for the items'
);

-- Borrows Table
CREATE TABLE borrows (
    borrow_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each borrow entry',
    order_id INT NOT NULL COMMENT 'Identifier of the related order',
    item_id INT NOT NULL COMMENT 'Identifier of the borrowed item',
    borrow_amount INT NOT NULL COMMENT 'Quantity of the item borrowed',
    pickup_datetime DATETIME DEFAULT NULL COMMENT 'Actual datetime of item pickup'
);

-- Returns Table
CREATE TABLE returns (
    return_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each return entry',
    order_id INT NOT NULL COMMENT 'Identifier of the related order',
    item_id INT NOT NULL COMMENT 'Identifier of the returned item',
    return_amount INT NOT NULL COMMENT 'Quantity of the item returned',
    dropoff_datetime DATETIME DEFAULT NULL COMMENT 'Actual datetime of item return'
);

-- Losts Table
CREATE TABLE losts (
    lost_id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'Unique identifier for each lost entry',
    order_id INT NOT NULL COMMENT 'Identifier of the related order',
    item_id INT NOT NULL COMMENT 'Identifier of the lost item',
    lost_amount INT NOT NULL COMMENT 'Quantity of the item reported lost'
);
