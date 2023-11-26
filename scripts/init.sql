DO $$ 
BEGIN
    RAISE NOTICE 'Iniciando execução do script init.sql...';
    
    CREATE TABLE produtos (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100)
    );

    RAISE NOTICE 'Tabela produtos criada com sucesso.';
    
    INSERT INTO produtos (name) VALUES ('Cadeira');
    INSERT INTO produtos (name) VALUES ('Mesa');

    RAISE NOTICE 'Dados inseridos na tabela produtos.';
END $$;
