CREATE TABLE user_roles(
    id INT PRIMARY KEY NOT NULL,
    role_name VARCHAR(100) ,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);


CREATE TABLE users (
	id INT PRIMARY KEY NOT NULL, 
	username VARCHAR(225) NOT NULL,
    phone_number VARCHAR(100) NOT NULL,
	password VARCHAR(225) NOT NULL,
    email VARCHAR(225) NOT NULL,
    role_id int ,
    is_verified BOOLEAN,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_role FOREIGN KEY (role_id) REFERENCES user_roles(id)
);

CREATE TABLE user_details(
	id INT PRIMARY KEY NOT NULL,
    user_id INT,
	balance float,
    credential_proof VARCHAR(225),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
	CONSTRAINT fk_user_details FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE user_top_ups (
    id INT PRIMARY KEY NOT NULL,
    user_id INT ,
    payment_method_id INT,
    created_at TIMESTAMP ,
    updated_at TIMESTAMP ,
    CONSTRAINT fk_user_top_up FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE payment_methods (
    id INT PRIMARY KEY NOT NULL,
    payment_method VARCHAR(225) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE courts(
    id int PRIMARY KEY NOT NULL,
    court_name VARCHAR(100) NOT NULL,
    description VARCHAR(225),
    court_price float,
    is_available BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP 
);

CREATE TABLE vouchers (
    id INT PRIMARY KEY NOT NULL,
    voucher_code VARCHAR(225) NOT NULL,
    is_available BOOLEAN ,
    discount float,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE transaction_status(
    id INT PRIMARY KEY NOT NULL,
    transaction_status VARCHAR(225) NOT NULL,
    created_at TIMESTAMP , 
    updated_at TIMESTAMP
);

CREATE TABLE bookings(
    id INT PRIMARY KEY NOT NULL, 
    booking_number VARCHAR(225),
    user_id INT NOT NULL,
    court_id INT NOT NULL,
    payment_method_id INT NOT NULL,
    voucher_id INT,
    total_transaction float,
    transaction_status_id INT ,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_voucher FOREIGN KEY (voucher_id) REFERENCES vouchers(id),
    CONSTRAINT fk_trx_stat FOREIGN KEY (transaction_status_id) REFERENCES transaction_status(id),
    CONSTRAINT fk_user_booking FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_court_booking FOREIGN KEY (court_id) REFERENCES courts(id)
);


CREATE TABLE booking_details(
    id_detail INT PRIMARY KEY,
    booking_id INT ,
    date_book TIMESTAMP NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    CONSTRAINT fk_booking_detail FOREIGN KEY (booking_id) REFERENCES bookings(id)
);

-- dirun sendiri
ALTER TABLE bookings
ADD CONSTRAINT fk_payment_book FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id);

ALTER TABLE user_top_ups
ADD CONSTRAINT fk_payment_topup FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id);
