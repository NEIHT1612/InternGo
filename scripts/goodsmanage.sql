CREATE DATABASE manage_goods;

CREATE SCHEMA goods_management;

CREATE TABLE goods_management.categories (
    category_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    category_name VARCHAR(100) NOT NULL
);

CREATE TABLE goods_management.suppliers (
    supplier_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    supplier_name VARCHAR(150) NOT NULL,
    address VARCHAR(200),
    phone VARCHAR(20)
);

CREATE TABLE goods_management.products (
    product_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_name VARCHAR(150) NOT NULL,
    category_id UUID REFERENCES goods_management.categories(category_id) ON DELETE SET NULL,
    supplier_id UUID REFERENCES goods_management.suppliers(supplier_id) ON DELETE SET NULL,
    unit VARCHAR(50),
    unit_price NUMERIC(15,2) CHECK (unit_price >= 0)
);

CREATE TABLE goods_management.customers (
    customer_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(150) NOT NULL,
    address VARCHAR(200),
    phone VARCHAR(20),
    username VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL 
);

CREATE TABLE goods_management.invoices (
    invoice_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    customer_id UUID REFERENCES goods_management.customers(customer_id) ON DELETE SET NULL
);

CREATE TABLE goods_management.invoice_details (
    invoice_id UUID REFERENCES goods_management.invoices(invoice_id) ON DELETE CASCADE,
    product_id UUID REFERENCES goods_management.products(product_id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    unit_price NUMERIC(12,2) NOT NULL,
    PRIMARY KEY (invoice_id, product_id)
);

INSERT INTO goods_management.categories (category_name)
VALUES
('Electronics'),
('Home Appliances'),
('Groceries'),
('Stationery');

INSERT INTO goods_management.suppliers (supplier_name, address, phone)
VALUES
('TechWorld Co.', '123 High Street, New York', '123456789'),
('FreshMart Ltd.', '45 Market Road, Los Angeles', '987654321'),
('OfficePlus', '22 Business Park, Chicago', '555666777');

INSERT INTO goods_management.products (product_name, category_id, supplier_id, unit, unit_price)
VALUES
('Laptop Lenovo ThinkPad', 
    (SELECT category_id FROM goods_management.categories WHERE category_name = 'Electronics'),
    (SELECT supplier_id FROM goods_management.suppliers WHERE supplier_name = 'TechWorld Co.'),
    'pcs', 1200.00),

('Air Conditioner LG 1.5HP', 
    (SELECT category_id FROM goods_management.categories WHERE category_name = 'Home Appliances'),
    (SELECT supplier_id FROM goods_management.suppliers WHERE supplier_name = 'TechWorld Co.'),
    'pcs', 750.00),

('Organic Rice 5kg', 
    (SELECT category_id FROM goods_management.categories WHERE category_name = 'Groceries'),
    (SELECT supplier_id FROM goods_management.suppliers WHERE supplier_name = 'FreshMart Ltd.'),
    'bag', 8.50),

('Ballpoint Pen', 
    (SELECT category_id FROM goods_management.categories WHERE category_name = 'Stationery'),
    (SELECT supplier_id FROM goods_management.suppliers WHERE supplier_name = 'OfficePlus'),
    'box', 2.50);

INSERT INTO goods_management.customers (full_name, address, phone, username, password_hash)
VALUES
('Alice Johnson', '12 Park Lane, New York', '111222333', 'john123', '123'),
('Bob Smith', '88 Sunset Blvd, Los Angeles', '444555666', 'bob123', '123'),
('Charlie Nguyen', '100 Lake View, Chicago', '777888999', 'char123', '123');

INSERT INTO goods_management.invoices (customer_id)
VALUES
((SELECT customer_id FROM goods_management.customers WHERE username = 'john123')),
((SELECT customer_id FROM goods_management.customers WHERE username = 'bob123')),
((SELECT customer_id FROM goods_management.customers WHERE username = 'char123'));

INSERT INTO goods_management.invoice_details (invoice_id, product_id, quantity, unit_price)
VALUES
((SELECT invoice_id FROM goods_management.invoices LIMIT 1 OFFSET 0),
 (SELECT product_id FROM goods_management.products WHERE product_name = 'Laptop Lenovo ThinkPad'),
 1, 1200.00),

((SELECT invoice_id FROM goods_management.invoices LIMIT 1 OFFSET 1),
 (SELECT product_id FROM goods_management.products WHERE product_name = 'Air Conditioner LG 1.5HP'),
 1, 750.00),
((SELECT invoice_id FROM goods_management.invoices LIMIT 1 OFFSET 1),
 (SELECT product_id FROM goods_management.products WHERE product_name = 'Organic Rice 5kg'),
 2, 8.50),

((SELECT invoice_id FROM goods_management.invoices LIMIT 1 OFFSET 2),
 (SELECT product_id FROM goods_management.products WHERE product_name = 'Ballpoint Pen'),
 5, 2.50);
