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
    
    INSERT INTO product (name, brand, category, quantity, price) VALUES ('Needle Rice', 'Sigala', 'Rice', 3, 1.40);
    INSERT INTO product (name, brand, category, quantity, price) VALUES ('Milk', 'Mimosa', 'Dairy', 2, 2.00);

    RAISE NOTICE 'Sucessfully processed init.sql';
END $$;
