DROP  TABLE IF EXISTS GeneralImport;
CREATE TABLE GeneralImport
(
invoiceNo VARCHAR(15),
invoiceDate TIMESTAMP,
quantity INT,
unitPrice DOUBLE PRECISION,
itemId VARCHAR(15),
description TEXT,
customerId BIGINT,
country TEXT
);

DROP  TABLE IF EXISTS Customers;
CREATE TABLE Customers
(
customerName TEXT
);

DROP  TABLE IF EXISTS General;
CREATE TABLE General
(
invoiceNo VARCHAR(15),
invoiceDate TIMESTAMP,
quantity INT,
unitPrice DOUBLE PRECISION,
itemId VARCHAR(15),
description TEXT,
customerId BIGINT,
country TEXT
);
