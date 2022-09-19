CREATE DATABASE `sales_backend`;

CREATE TABLE `category_product` (
  category_product_id VARCHAR(25) NOT NULL,
  category VARCHAR(150) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL,
  PRIMARY KEY(category_product_id)
);

CREATE TABLE `product` (
  product_id VARCHAR(25) NOT NULL,
  category_product_id VARCHAR(25) NULL,
  name VARCHAR(150) NOT NULL,
  unit VARCHAR(50) NOT NULL,
  price DECIMAL(19,2) NOT NULL,
  stock INT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL,
  PRIMARY KEY(product_id),
  CONSTRAINT `product_category_product_id_foreign` FOREIGN KEY(category_product_id)
  REFERENCES category_product(category_product_id)
  ON UPDATE CASCADE ON DELETE NO ACTION
);



CREATE TABLE `province` (
  province_id INT NOT NULL,
  province VARCHAR(150) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL,
  PRIMARY KEY(province_id)
);

CREATE TABLE `city` (
  city_id INT NOT NULL,
  province_id INT NULL,
  city VARCHAR(150) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL,
  PRIMARY KEY(city_id),
  CONSTRAINT `city_province_id_foreign` FOREIGN KEY(province_id)
  REFERENCES province(province_id)
  ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE `sub_district` (
  sub_district_id INT NOT NULL,
  city_id INT NULL,
  sub_district VARCHAR(150) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL,
  PRIMARY KEY(sub_district_id),
  CONSTRAINT `sub_district_city_id_foreign` FOREIGN KEY(city_id)
  REFERENCES city(city_id)
  ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE `address` (
  address_id INT NOT NULL,
  province_id INT NULL,
  city_id INT NULL,
  sub_district_id INT NULL,
  address VARCHAR(150) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL,
  PRIMARY KEY(address_id),
  CONSTRAINT `address_province_id_foreign` FOREIGN KEY(province_id)
  REFERENCES province(province_id)
  ON UPDATE CASCADE ON DELETE NO ACTION,

  CONSTRAINT `address_city_id_foreign` FOREIGN KEY(city_id)
  REFERENCES city(city_id)
  ON UPDATE CASCADE ON DELETE NO ACTION,

  CONSTRAINT `address_sub_district_id_foreign` FOREIGN KEY(sub_district_id)
  REFERENCES sub_district(sub_district_id)
  ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE `customer` (
  customer_id VARCHAR(25) NOT NULL,
  address_id INT NULL,
  first_name VARCHAR(75) NOT NULL,
  last_name VARCHAR(75) NULL,
  gender VARCHAR(3) NULL,
  address VARCHAR(150) NULL,
  email VARCHAR(150) NOT NULL UNIQUE,
  phone_number VARCHAR(25) NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL,
  PRIMARY KEY(customer_id),
  CONSTRAINT `customer_address_id` FOREIGN KEY(address_id)
  REFERENCES address(address_id)
  ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE `employee` (
  employee_id VARCHAR(25) NOT NULL,
  address_id INT NULL,
  first_name VARCHAR(75) NOT NULL,
  last_name VARCHAR(75) NULL,
  gender VARCHAR(3) NULL,
  address VARCHAR(150) NULL,
  email VARCHAR(150) NOT NULL UNIQUE,
  phone_number VARCHAR(25) NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL,
  PRIMARY KEY(employee_id),
  CONSTRAINT `employee_address_id` FOREIGN KEY(address_id)
  REFERENCES address(address_id)
  ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE `transaction` (
  transaction_id  VARCHAR(25) NOT NULL,
  customer_id VARCHAR(25) NULL,
  employee_id VARCHAR(25) NULL,
  date DATE NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL,
  PRIMARY KEY(transaction_id),
  CONSTRAINT `transaction_customer_id_foreign` FOREIGN KEY(customer_id)
  REFERENCES customer(customer_id)
  ON UPDATE CASCADE ON DELETE NO ACTION,

  CONSTRAINT `transaction_employee_id_foreign` FOREIGN KEY(employee_id)
  REFERENCES employee(employee_id)
  ON UPDATE CASCADE ON DELETE NO ACTION
);

CREATE TABLE `transaction_detail` (
  transaction_detail_id VARCHAR(25) NOT NULL,
  transaction_id VARCHAR(25) NULL,
  product_id VARCHAR(25) NULL,
  price DECIMAL(19,2) NOT NULL,
  quantity INT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NULL,
  CONSTRAINT `transaction_detail_transaction_id_foreign` FOREIGN KEY(transaction_id)
  REFERENCES transaction(transaction_id)
  ON UPDATE CASCADE ON DELETE NO ACTION,

  CONSTRAINT `transaction_detail_product_id_foreign` FOREIGN KEY(product_id)
  REFERENCES product(product_id)
  ON UPDATE CASCADE ON DELETE NO ACTION
);

LOCK TABLE address WRITE, province WRITE;
ALTER TABLE address
DROP FOREIGN KEY `address_province_id_foreign`;

LOCK TABLE city WRITE;
ALTER TABLE `city`
DROP FOREIGN KEY `city_province_id_foreign`;

LOCK TABLE province WRITE;
ALTER TABLE province
MODIFY `province_id` INT NOT NULL AUTO_INCREMENT;

LOCK TABLE city WRITE;
ALTER TABLE city ADD CONSTRAINT `city_province_id_foreign` FOREIGN KEY(province_id)
REFERENCES province(province_id)
ON UPDATE CASCADE ON DELETE NO ACTION;

LOCK TABLE address WRITE;
ALTER TABLE address
ADD CONSTRAINT `address_province_id_foreign` FOREIGN KEY(province_id)
REFERENCES province(province_id)
ON UPDATE CASCADE ON DELETE NO ACTION;

UNLOCK TABLES;

LOCK TABLE city WRITE, address WRITE;
ALTER TABLE address DROP FOREIGN KEY `address_city_id_foreign`;
ALTER TABLE sub_district DROP FOREIGN KEY `sub_district_city_id_foreign`;
ALTER TABLE city MODIFY `city_id` INT NOT NULL AUTO_INCREMENT;
ALTER TABLE `address` ADD CONSTRAINT `address_city_id_foreign` FOREIGN KEY(city_id) REFERENCES city(city_id) ON UPDATE CASCADE ON DELETE NO ACTION;
ALTER TABLE `sub_district` ADD CONSTRAINT `sub_district_city_id_foreign` FOREIGN KEY(city_id) REFERENCES city(city_id) ON UPDATE CASCADE ON DELETE NO ACTION;

LOCK TABLE sub_district WRITE;
LOCK TABLE address WRITE;
ALTER TABLE address DROP FOREIGN KEY `address_sub_district_id_foreign`;
ALTER TABLE sub_district MODIFY `sub_district_id` INT NOT NULL AUTO_INCREMENT;
ALTER TABLE address ADD CONSTRAINT `address_sub_district_id_foreign` FOREIGN KEY(sub_district_id) REFERENCES sub_district(sub_district_id) ON UPDATE CASCADE ON DELETE NO ACTION;
UNLOCK TABLES;

LOCK TABLE permissions WRITE;
ALTER TABLE permissions DROP FOREIGN KEY `permissions_role_id_foreign`;
UNLOCK TABLES;

LOCK TABLE users_roles WRITE;
ALTER TABLE users_roles DROP FOREIGN KEY `users_roles_role_id_foreign`;
UNLOCK TABLES;

LOCK TABLE roles WRITE;
ALTER TABLE roles MODIFY `role_id` INT NOT NULL AUTO_INCREMENT;
UNLOCK TABLES;

LOCK TABLE permissions WRITE;
ALTER TABLE permissions ADD CONSTRAINT `permissions_role_id_foreign` FOREIGN KEY(role_id) REFERENCES roles(role_id) ON UPDATE CASCADE ON DELETE CASCADE;
UNLOCK TABLES;

LOCK TABLE users_roles WRITE;
ALTER TABLE users_roles ADD CONSTRAINT `users_roles_role_id_foreign` FOREIGN KEY(role_id) REFERENCES roles(role_id) ON UPDATE CASCADE ON DELETE CASCADE;
UNLOCK TABLES;

LOCK TABLE users WRITE;
ALTER TABLE users DROP FOREIGN KEY `customer_address_id`;
UNLOCK TABLES;

LOCK TABLE address WRITE;
ALTER TABLE address MODIFY `address_id` INT NOT NULL AUTO_INCREMENT;
UNLOCK TABLES;

LOCK TABLE users WRITE;
ALTER TABLE users ADD CONSTRAINT `users_address_id_foreign` FOREIGN KEY(address_id) REFERENCES address(address_id) ON UPDATE CASCADE ON DELETE CASCADE;
UNLOCK TABLES;

ALTER TABLE permissions MODIFY permission_id INT NOT NULL AUTO_INCREMENT;
