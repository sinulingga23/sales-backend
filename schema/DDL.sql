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