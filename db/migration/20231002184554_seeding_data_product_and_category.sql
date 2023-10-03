-- +goose Up
-- +goose StatementBegin
INSERT INTO product_categories (id, category_name)
VALUES (1,'sayur dan buah'),(2,'bumbu dapur'),(3,'daging'),(4,'perlengkapan rumah tangga'),(5,'makanan siap saji');

INSERT INTO products (product_name, category_id, description, qty, selling_price, promo_price)
VALUES ('jeruk mandarin 1kg',1,'jeruk mandarin manis asli',20,30000,25000),
       ('apel malang 0.5 kg',1,'apel malng budidaya asli dari malang',30,20000,17500),
       ('tomat 1kg',1,'tomat segar hasil budidaya',40,15000,14500),
       ('bawang merah 1 kg',2,'bawang merah dari jwa timur',70,29700,27200),
       ('bawang putih 1 kg',2,'bawang putih',20,30000,28700),
       ('jeruk nipis 0.5 kg',2,'jeruk nipis pelengkap makanan',10,12000,12000),
       ('ayam potong filet 1kg',3,'daging ayam potong filet segar',20,35000,32500),
       ('ayam potong paha 0.5 kg',3,'daging ayam potong filet segar',20,20000,18500),
       ('wagyu a3 250gram',3,'wagyu import australia',20,90000,88900),
       ('sunlight 500gram',4,'sabun cuci piring',10,55000,50000),
       ('lifeboy sabun cair 200gram',4,'sabun mandi',20,30000,25000),
       ('sunsilk shampo',4,'shampo',20,30000,29000),
       ('nasi goreng ayam',5,'makanan siap saji',20,12500,12500),
       ('nasi ayam geprek',5,'makanan siap saji',20,16500,15000),
       ('nasi katsu sambal bbq',5,'makanan siap saji',20,18000,17500);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
