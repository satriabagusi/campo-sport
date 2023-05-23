-- insert user roles

INSERT INTO user_roles (id, role_name, created_at, updated_at)
VALUES 
(1, 'Admin', current_timestamp, current_timestamp),
(2, 'Member', current_timestamp, current_timestamp),
(3, 'Non-member', current_timestamp, current_timestamp);

--insert payment_methods

INSERT INTO payment_methods (id, payment_method, created_at, updated_at)
VALUES
    (1, 'Tunai', current_timestamp, current_timestamp),
    (2, 'Ballance', current_timestamp, current_timestamp),
    (3, 'Online', current_timestamp, current_timestamp);

-- insert courts
INSERT INTO courts (id, court_name, description, court_price, is_available, created_at, updated_at)
VALUES
    (1, 'Indoor court', 'Luas 550 meter persegi.', 200000, true, current_timestamp, current_timestamp),
    (2, 'Outdoor court', 'Luas 11.000 meter persegi', 250000, true, current_timestamp, current_timestamp),
    (3, 'Tennis court', 'Luas 260,8 meter persegi. ', 150000, false, current_timestamp, current_timestamp);

-- insert couchers
INSERT INTO vouchers (id, voucher_code, is_available, discount, created_at, updated_at)
VALUES
    (1, 'Voucher001', true, 10000, current_timestamp, current_timestamp),
    (2, 'Voucher002', true, 15000, current_timestamp, current_timestamp),
    (3, 'Voucher003', false, 20000, current_timestamp, current_timestamp);

-- insert transaction status
INSERT INTO transaction_status (id, transaction_status, created_at, updated_at)
VALUES
    (1, 'Pending', current_timestamp, current_timestamp),
    (2, 'Success', current_timestamp, current_timestamp),
    (3, 'Failed', current_timestamp, current_timestamp);

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


INSERT INTO users (id, username,phone_number, password,email,role_id,is_verified,   created_at, updated_at)
VALUES (1, 'admin','081231915158','admin','admin@gmail.com',1,true, current_timestamp, current_timestamp);
