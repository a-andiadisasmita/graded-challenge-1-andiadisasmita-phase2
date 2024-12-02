CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO customers (name, email, phone) VALUES
('John Doe', 'john.doe@example.com', '1234567890'),
('Jane Smith', 'jane.smith@example.com', '0987654321'),
('Alice Johnson', 'alice.johnson@example.com', '1112223333'),
('Bob Brown', 'bob.brown@example.com', '4445556666'),
('Charlie White', 'charlie.white@example.com', '7778889999'),
('Diana Black', 'diana.black@example.com', '3334445555'),
('Ella Green', 'ella.green@example.com', '6667778888'),
('Frank Red', 'frank.red@example.com', '9990001111'),
('Grace Yellow', 'grace.yellow@example.com', '2223334444'),
('Hank Purple', 'hank.purple@example.com', '5556667777');
