CREATE TABLE IF NOT EXISTS users(
   id          BIGINT PRIMARY KEY,
   username VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   age INT NOT NULL,
   refreshtoken TEXT,
   refreshtoken_expiretime TIMESTAMP WITH TIME ZONE
);
