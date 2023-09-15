-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Asia/Kolkata";

-- Create user table
CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    username VARCHAR (50) NOT NULL,
    password VARCHAR (1000) NOT NULL,
    active_status INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL
);