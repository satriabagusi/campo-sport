CREATE TABLE user_roles(
    id INT PRIMARY KEY,
    role_name VARCHAR(100) ,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);


CREATE TABLE users (
	id INT PRIMARY KEY, 
	username VARCHAR(225) NOT NULL,
    phone_number VARCHAR(100) NOT NULL,
	password VARCHAR(225) NOT NULL,
    email VARCHAR(225) NOT NULL,
    role_id int default 3,
    is_verified BOOLEAN default false,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_user_role FOREIGN KEY (role_id) REFERENCES user_roles(id)
);

CREATE TABLE user_details(
	id INT PRIMARY KEY,
    user_id INT,
	balance float,
    credential_proof VARCHAR(225),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
	CONSTRAINT fk_user_details FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE user_top_ups (
    id INT PRIMARY KEY,
    user_id INT ,
    payment_method_id INT,
    created_at TIMESTAMP ,
    updated_at TIMESTAMP ,
    CONSTRAINT fk_user_top_up FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE payment_methods (
    id INT PRIMARY KEY,
    payment_method VARCHAR(225),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE courts(
    id int PRIMARY KEY,
    court_name VARCHAR(100) NOT NULL,
    description VARCHAR(225),
    court_price float NOT NULL,
    is_available BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP 
);

CREATE TABLE vouchers (
    id INT PRIMARY KEY ,
    voucher_code VARCHAR(225) NOT NULL,
    is_available BOOLEAN ,
    discount float,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);



CREATE TABLE transaction_status(
    id INT PRIMARY KEY,
    transaction_status VARCHAR(225),
    created_at TIMESTAMP , 
    updated_at TIMESTAMP
);

CREATE TABLE bookings(
    id INT PRIMARY KEY, 
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


ALTER TABLE user_top_ups
ADD COLUMN 
order_number VARCHAR(255) ;

ALTER TABLE user_top_ups
ADD COLUMN 
amount float;

ALTER TABLE user_top_ups
ADD COLUMN 
transaction_status_id int;

ALTER TABLE user_top_ups
ADD CONSTRAINT fk_status_top_up
FOREIGN KEY (transaction_status_id)
REFERENCES transaction_status (id);


-- Create the sequence
CREATE SEQUENCE court_id_seq;
-- Associate the sequence with the id column in the users table
ALTER TABLE courts ALTER COLUMN id SET DEFAULT nextval('court_id_seq'::regclass);


-- Create the sequence
CREATE SEQUENCE users_id_seq;
-- Associate the sequence with the id column in the users table
ALTER TABLE users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);

-- Create the sequence
CREATE SEQUENCE user_detail_id_seq;
-- Associate the sequence with the id column in the users table
ALTER TABLE user_details ALTER COLUMN id SET DEFAULT nextval('user_detail_id_seq'::regclass);

-- Create the sequence
CREATE SEQUENCE voucher_id_seq;
-- Associate the sequence with the id column in the users table
ALTER TABLE vouchers ALTER COLUMN id SET DEFAULT nextval('voucher_id_seq'::regclass);

ALTER TABLE vouchers
ALTER COLUMN created_at SET DEFAULT current_timestamp;

ALTER TABLE vouchers
ALTER COLUMN updated_at SET DEFAULT current_timestamp;

ALTER TABLE vouchers
ALTER COLUMN created_at SET DEFAULT current_timestamp;

ALTER TABLE vouchers
ALTER COLUMN updated_at SET DEFAULT current_timestamp;

ALTER TABLE user_details
ALTER COLUMN created_at SET DEFAULT current_timestamp;

ALTER TABLE user_details
ALTER COLUMN updated_at SET DEFAULT current_timestamp;

ALTER TABLE courts
ALTER COLUMN created_at SET DEFAULT current_timestamp;

ALTER TABLE courts
ALTER COLUMN updated_at SET DEFAULT current_timestamp;


ALTER TABLE vouchers
ALTER COLUMN created_at SET DEFAULT current_timestamp;

ALTER TABLE vouchers
ALTER COLUMN updated_at SET DEFAULT current_timestamp;


-- Create the sequence
CREATE SEQUENCE topup_id_seq;
ALTER TABLE user_top_ups ALTER COLUMN id SET DEFAULT nextval('topup_id_seq'::regclass);


ALTER TABLE user_details
ALTER COLUMN created_at SET DEFAULT current_timestamp;

ALTER TABLE user_details
ALTER COLUMN updated_at SET DEFAULT current_timestamp;
