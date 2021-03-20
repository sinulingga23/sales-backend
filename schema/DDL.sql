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

