    -- Enable pgvector extension
    CREATE EXTENSION IF NOT EXISTS vector;

    -- Create basic schema for ecommerce
    CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        description TEXT,
        price DECIMAL(10, 2) NOT NULL,
        category VARCHAR(100),
        embedding vector(1536),
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS interactions (
        id SERIAL PRIMARY KEY,
        user_id INTEGER REFERENCES users(id),
        product_id INTEGER REFERENCES products(id),
        action VARCHAR(50) NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS orders (
        id SERIAL PRIMARY KEY,
        user_id INTEGER REFERENCES users(id),
        total_amount DECIMAL(10, 2) NOT NULL,
        status VARCHAR(50) DEFAULT 'PENDING',
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS order_items (
        id SERIAL PRIMARY KEY,
        order_id INTEGER REFERENCES orders(id),
        product_id INTEGER REFERENCES products(id),
        quantity INTEGER NOT NULL,
        unit_price DECIMAL(10, 2) NOT NULL
    );

    -- Seed some mock data
    INSERT INTO products (name, description, price, category) VALUES
    ('Wireless Noise-Canceling Headphones', 'Premium over-ear headphones with active noise cancellation and 30-hour battery life.', 299.99, 'Electronics'),
    ('Ergonomic Office Chair', 'Adjustable lumbar support, breathable mesh back, and comfortable foam seat cushion.', 199.50, 'Furniture'),
    ('Smart Fitness Watch', 'Tracks heart rate, sleep, steps, and includes built-in GPS for outdoor runs.', 149.00, 'Wearables'),
    ('Organic Arabica Coffee Beans', 'Medium roast, 100% organic beans with hints of chocolate and caramel. 1 lb bag.', 18.99, 'Food & Beverage');

    INSERT INTO users (email, name) VALUES
    ('alice@test.com', 'Alice Smith'),
    ('bob@test.com', 'Bob Johnson');

    INSERT INTO orders (user_id, total_amount, status) VALUES
    (1, 499.49, 'COMPLETED'),
    (2, 149.00, 'PENDING');

    INSERT INTO order_items (order_id, product_id, quantity, unit_price) VALUES
    (1, 1, 1, 299.99),
    (1, 2, 1, 199.50),
    (2, 3, 1, 149.00);
