CREATE TABLE cats (
    name TEXT,
    shop_id INT
);
CREATE TABLE shops (
    id INT,
    shop_name TEXT
);
INSERT INTO cats (name,shop_id) VALUES ('Vlas',1),('Nemo',2);

INSERT INTO shops (id,shop_name) VALUES (1,'Four_Paws'),(2,'Mr.Zoo');

SELECT name,shop_name FROM cats JOIN shops ON cats.shop_id=shops.id;

SELECT name,shop_name FROM cats,shops WHERE cats.shop_id=shops.id;

INSERT INTO cats (name,shop_id)VALUES('Fenua',3);

SELECT name,shop_name FROM cats LEFT JOIN shops ON cats.shop_id=shops.id;

SELECT name FROM cats LEFT JOIN shops ON cats.shop_id=shops.id WHERE shop_name IS NULL;

BEGIN;
INSERT INTO cats (name,shop_id)VALUES('Barsik',4);
INSERT INTO shops (id,shop_name) VALUES (4,'Zoo_back');
COMMIT;

SELECT COUNT(*) FROM cats;

SELECT COUNT(*) FROM cats LEFT JOIN shops ON cats.shop_id=shops.id WHERE shop_name IS NULL;

SELECT COUNT(*) FROM cats JOIN shops ON cats.shop_id=shops.id;

INSERT INTO cats (name,shop_id) VALUES ('Gosha',1);

SELECT shop_name,COUNT(*) AS count_of_cats FROM shops JOIN cats ON shop_id=id GROUP BY shop_name;

SELECT shop_name,COUNT(1) FROM shops JOIN cats ON shop_id=id GROUP BY shop_name;

ALTER TABLE cats ADD COLUMN type TEXT;

UPDATE cats SET type='home' WHERE name='Vlas' OR name='Nemo' OR name='Fenua';

UPDATE cats SET type='street' WHERE name IN('Barsik','Gosha');

SELECT shop_name,name,type FROM cats JOIN shops ON shop_id=id;

SELECT type ,COUNT(*) FROM cats GROUP BY type;

SELECT type ,COUNT(*) FROM cats JOIN shops ON shop_id=id GROUP BY type;

ALTER TABLE cats ADD COLUMN price INT;

UPDATE cats SET price=100;

SELECT SUM(price) FROM cats;

SELECT SUM(price),type,COUNT(*) FROM cats GROUP BY type;

SELECT SUM(price),type,COUNT(*) FROM cats JOIN shops ON shop_id=id  GROUP BY type;

UPDATE cats SET price = 200 WHERE name = 'Barsik';

UPDATE cats SET price = 50 WHERE name = 'Fenua';

SELECT name,price FROM cats ORDER BY price ASC LIMIT 1;

SELECT name,price FROM cats ORDER BY price DESC LIMIT 1;

SELECT name,price,shop_name FROM cats JOIN shops ON shop_id=id ORDER BY price DESC LIMIT 1 ;

SELECT c.name FROM cats c;
















