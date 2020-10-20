--BEFORE RUNNING 
--import names.csv to Customers (import encoding utf8, header yes, delimiter coma)
--import E-CommerceOrder - data to generalImport, get mini version in case you're having performance issues when running
--this script

DROP  TABLE IF EXISTS CustomerTemp;
CREATE TABLE CustomerTemp
(
customerName TEXT
);

INSERT INTO customerTemp(customerName)
SELECT
--*
customerName--, COUNT(customerName) kk--, customerId
FROM customers
GROUP BY customerName--, customerId
--ORDER BY kk DESC
LIMIT 18287;

--SELECT * FROM customerTemp

ALTER TABLE CustomerTemp ADD COLUMN customerId SERIAL PRIMARY KEY;
--ALTER TABLE CustomerTemp DROP COLUMN customerId

TRUNCATE TABLE Customers;
DROP TABLE Customers;

DROP  TABLE IF EXISTS Customers;
CREATE TABLE Customers
(
customerId BIGINT,
customerName TEXT,
country TEXT
);

INSERT INTO general
SELECT DISTINCT * FROM generalImport;

INSERT INTO Customers (customerId, customerName, country)
SELECT
t.customerId, t.customerName, g.country
FROM CustomerTemp t
INNER JOIN 
(SELECT customerId, country from general group by customerId, country)
g 
ON g.customerId = t.CustomerId;

--to test if there is any missing customer
--SELECT customerId from general where customerId not in (select c.customerId from (select distinct * from customers)c)

DROP TABLE IF EXISTS customerTemp;

ALTER TABLE general ADD COLUMN invoiceLineId SERIAL PRIMARY KEY;

DROP  TABLE IF EXISTS Orders;
CREATE TABLE Orders
(
orderId SERIAL,
invoiceNo VARCHAR(15),
customerId BIGINT
);

DROP  TABLE IF EXISTS Invoices;
CREATE TABLE Invoices
(
invoiceNo VARCHAR(15),
invoiceDate TIMESTAMP
);

DROP  TABLE IF EXISTS OrderDetails;
CREATE TABLE OrderDetails
(
orderDetailId SERIAL,
invoiceNo VARCHAR(15),
invoiceLineId BIGINT
);

DROP  TABLE IF EXISTS InvoiceLines;
CREATE TABLE InvoiceLines
(
invoiceLineID BIGINT,
quantity INT,
unitPrice DOUBLE PRECISION
);

DROP  TABLE IF EXISTS OrderDetailItems;
CREATE TABLE OrderDetailItems
(
orderDetailItem SERIAL,
invoiceLineID BIGINT,
itemId VARCHAR(15)
);

DROP  TABLE IF EXISTS Items;
CREATE TABLE Items
(
itemId VARCHAR(15),
description TEXT,
unitPrice DOUBLE PRECISION,
priceMultiplier DOUBLE PRECISION
);

INSERT INTO OrderDetails (invoiceNo, invoiceLineId)
SELECT
invoiceNo, invoiceLineId
FROM general;

--SELECT * FROM OrderDetails
--ALTER TABLE OrderDetails ADD COLUMN orderDetailsId SERIAL PRIMARY KEY;
--SELECT customerId, customername, country, count(customername) rp
--FROM Customers group by customerId, customername, country
--order by rp DESC 
--SELECT * FROM Customers

INSERT INTO Invoices (invoiceNo, invoiceDate) --25943
SELECT
invoiceNo, invoiceDate
FROM general
GROUP BY invoiceNo, invoiceDate;

--SELECT Invoicelineid from general where customerId is null -> 135037

INSERT INTO Orders (invoiceNo, customerId) --22317 without null customers, 25943 with
SELECT
i.invoiceNo, g.customerId
FROM Invoices i
INNER JOIN (SELECT customerId, invoiceNo FROM general group by
			customerId, invoiceNo
		   ) g on g.InvoiceNo = i.InvoiceNo;
--INNER JOIN customers c on c.customerId = g.customerId

--SELECT * FROM Orders

INSERT INTO InvoiceLines (invoiceLineId, unitPrice, quantity) --536641
SELECT
invoiceLineId, unitPrice, quantity
FROM general;

INSERT INTO OrderDetailItems (invoiceLineId, itemId)
SELECT
invoiceLineId, itemId
FROM general;

-------------------- item cleanup

DROP  TABLE IF EXISTS ItemsTemp;
CREATE TABLE ItemsTemp
(
itemId VARCHAR(15),
description TEXT,
unitPrice DOUBLE PRECISION,
priceMultiplier DOUBLE PRECISION,
invoiceDate TIMESTAMP,
invoiceId VARCHAR(15)
);

INSERT INTO itemsTemp  --531808
SELECT itemId, description, unitPrice, 1 priceMultiplier, invoiceDate, invoiceNo  --13356, 4688, 4667   531808
	FROM general GROUP BY itemId, description, unitPrice, invoiceNo, invoiceDate 
	ORDER BY itemId;
	
--SELECT itemId from general WHERE itemId NOT IN (SELECT itemId FROM itemsTemp);

--SELECT * FROM itemsTemp
	
--DELETE FROM itemsTemp a
--WHERE a.itemId <> (SELECT max(b.InvoiceId)
--                  FROM   itemsTemp b
--                  WHERE  a.itemId = b.itemId);
				  
DELETE FROM itemsTemp a USING (
      SELECT MAX(invoiceId) as invoiceId, itemId
        FROM itemsTemp 
        GROUP BY itemId HAVING COUNT(*) > 1
      ) b
      WHERE a.itemId = b.itemId 
      AND a.invoiceId <> b.invoiceId;
	  
INSERT INTO items
(SELECT DISTINCT itemId, description, unitPrice, priceMultiplier FROM itemsTemp);

--SELECT itemId from general where itemId NOT in (SELECT itemId from items)

DROP TABLE itemsTemp;
	  
DROP TABLE generalImport;