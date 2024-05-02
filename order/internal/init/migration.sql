CREATE TABLE products
(
    product_id VARCHAR(100) UNIQUE ,
    title VARCHAR(100) NOT NULL ,
    description VARCHAR(200) NOT NULL ,
    price BIGINT NOT NULL,
    PRIMARY KEY(product_id)
);

CREATE TABLE cart
(
    product_id VARCHAR(100) NOT NULL ,
    quantity INT NOT NULL DEFAULT 1,
    CONSTRAINT fk_products_shopping_cart FOREIGN KEY (product_id) REFERENCES products(product_id)
);

CREATE TABLE orders
(
    order_id VARCHAR(100) NOT NULL ,
    airwaybill VARCHAR(100),
    invoice VARCHAR(100),
    status varchar(50) NOT NULL ,
    total_price BIGINT NOT NULL
);

CREATE TABLE order_prodcuts
(
  order_id VARCHAR(100) NOT NULL,
  product_id VARCHAR(100) NOT NULL,
  quantity INT NOT NULL DEFAULT 1,
  price INT NOT NULL, 
  CONSTRAINT fk_order_mp FOREIGN KEY (order_id) REFERENCES orders(order_id),
  CONSTRAINT fk_products_mp FOREIGN KEY (product_id) REFERENCES products(product_id)
)
