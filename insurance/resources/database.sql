-- Create the database
CREATE DATABASE insurance;
\c insurance

-- Drop and create `customers` table
DROP TABLE IF EXISTS clients;
CREATE TABLE clients (
  client_id SERIAL PRIMARY KEY,
  fname VARCHAR(100) NOT NULL,
  lname VARCHAR(100) NOT NULL,
  birthdate DATE NOT NULL, 
  id_card_1 VARCHAR(100) NOT NULL, 
  id_no_1 VARCHAR(100) NOT NULL,
  id_card_2 VARCHAR(100) NOT NULL, 
  id_no_2 VARCHAR(100) NOT NULL,
  birthplace VARCHAR(100) NOT NULL,
  contact_no VARCHAR(30) NOT NULL,
  status VARCHAR(10) NOT NULL DEFAULT TRUE,
  gender VARCHAR(20) NOT NULL
);

-- Insert data into `customers` table
INSERT INTO clients (client_id, fname, lname, birthdate, id_card_1, id_no_1, id_card_2, id_no_2, birthplace, contact_no, status, gender) VALUES
  (110001, 'John Michael', 'Echavez', '1990-09-16', 'National Id', '001-001-001', 'Driver License', '001-002-002', 'General Santos City', '09000000001', "TRUE", 'Male'),
  (110002, 'JM', 'Echavez', '1990-09-16', 'National Id', '002-002-002', 'Driver License', '002-002-002', 'Paranaque', '09000000002', "TRUE", 'Male');

