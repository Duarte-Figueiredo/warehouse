DO $$ 
BEGIN
    RAISE NOTICE 'Starting init.sql...';
    
    CREATE TABLE product (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        brand VARCHAR(100),
        category VARCHAR(100),
        quantity INTEGER,
        price FLOAT
    );

    RAISE NOTICE 'Table product created';

    CREATE TABLE request (
        id SERIAL PRIMARY KEY,
        categoriesProducts VARCHAR(100),
        maxPrices VARCHAR(100),
        quantities VARCHAR(100)
    );
    
    RAISE NOTICE 'Table request created';

    INSERT INTO product (name, brand, category, quantity, price) VALUES ('Needle Rice', 'Sigala', 'Rice', 3, 1.40);
    INSERT INTO product (name, brand, category, quantity, price) VALUES ('Black Rice', 'Sigala', 'Rice', 3, 1.40);
    INSERT INTO product (name, brand, category, quantity, price) VALUES ('Basmati Rice', 'Sigala', 'Rice', 3, 5.00);
    INSERT INTO product (name, brand, category, quantity, price) VALUES ('Fat Milk', 'Mimosa', 'Milk', 2, 2.00);
    INSERT INTO product (name, brand, category, quantity, price) VALUES ('Fat Milk', 'Vigor', 'Milk', 2, 2.50);
    INSERT INTO product (name, brand, category, quantity, price) VALUES ('White Eggs', 'Ovifatima', 'Eggs', 2, 2.00);
    INSERT INTO product (name, brand, category, quantity, price) VALUES ('Brown Eggs', 'Mirtania', 'Eggs', 0, 1.00);
    INSERT INTO product (name, brand, category, quantity, price) VALUES ('Goat', 'Filadelfia', 'Cheese', 1, 1.00);
    INSERT INTO product (name, brand, category, quantity, price) VALUES ('Blue', 'Filadelfia', 'Cheese', 0, 4.00);



    INSERT INTO request (categoriesProducts, maxPrices, quantities) VALUES ('Rice, Milk, Cheese', '2, 1, 2', '2, 2, 2');

    RAISE NOTICE 'Sucessfully processed init.sql';
END $$;
