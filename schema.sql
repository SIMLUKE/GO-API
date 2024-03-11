CREATE TABLE IF NOT EXISTS users
(
  Id text PRIMARY KEY,
  Name text,
  Email text,
  Password text
);

CREATE TABLE IF NOT EXISTS articles
(
    Id     text PRIMARY KEY,
    Name   text,
    Size   text,
    Price  float,
    State  text,
    Brand  text,
    Color  text,
    UserId text,
    FOREIGN KEY (UserId) REFERENCES users(Id)
);

-- default users
INSERT INTO users (Id, Name, Email, Password)
VALUES
    ('f449d522-875a-41a1-8e09-620a40598acb', 'Root', 'john.doe@example.com', '79ba8af6932beeba43b6748cefed436566f636249828f102780b57133e322d2e'), --password = very_secure
    ('56be9407-0170-4dd9-a5ef-570b94d883a7', 'Alice Smith', 'alice.smith@example.com', 'c6ba91b90d922e159893f46c387e5dc1b3dc5c101a5a4522f03b987177a24a91'); --password = password456

-- Insert random products
INSERT INTO articles (Id, Name, Size, Price, State, Brand, Color, UserId)
VALUES
    ('101', 'T-Shirt', 'M', 15.99, 'New', 'Nike', 'Blue', '56be9407-0170-4dd9-a5ef-570b94d883a7'),
    ('102', 'Jeans', 'L', 39.99, 'Recent', 'Levis', 'Black', '56be9407-0170-4dd9-a5ef-570b94d883a7'),
    ('103', 'Running Shoes', '39', 79.99, 'New', 'Adidas', 'White', '56be9407-0170-4dd9-a5ef-570b94d883a7'),
    ('104', 'Dress', 'S', 49.99, 'New', 'H&M', 'Red', '56be9407-0170-4dd9-a5ef-570b94d883a7'),
    ('105', 'Backpack', 'M', 29.99, 'Old', 'North Face', 'Green', '56be9407-0170-4dd9-a5ef-570b94d883a7');
