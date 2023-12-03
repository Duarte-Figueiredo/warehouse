DO $$ 
BEGIN
    RAISE NOTICE 'Starting init.sql...';
    
    CREATE TABLE product (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100)
    );

    RAISE NOTICE 'Table product created';
    
    INSERT INTO product (name) VALUES ('Cadeira');
    INSERT INTO product (name) VALUES ('Mesa');

    RAISE NOTICE 'Sucessfully processed init.sql';
END $$;
