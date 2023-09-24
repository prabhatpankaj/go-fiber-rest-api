-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Asia/Kolkata";

CREATE TABLE roles (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT
);

CREATE TABLE permissions (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT
);

-- Create the "role_permissions" junction table with UUID
CREATE TABLE role_permissions (
    role_id UUID NOT NULL REFERENCES roles(id),
    permission_id UUID NOT NULL REFERENCES permissions(id),
    PRIMARY KEY (role_id, permission_id)
);

-- Create the "users" table
CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    full_name VARCHAR(25) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    user_status INT NOT NULL
);

-- Create the "user_roles" junction table to allow users to have multiple roles
CREATE TABLE user_roles (
    user_id UUID NOT NULL REFERENCES users(id),
    role_id UUID NOT NULL REFERENCES roles(id),
    PRIMARY KEY (user_id, role_id)
);

-- Add indexes
CREATE INDEX active_users ON users (id) WHERE user_status = 1;


-- -- Insert data into the "roles" table
-- INSERT INTO roles (id, name, description) VALUES
--     ('a6cb87e2-5d6a-4a3d-8c35-62f2f8f17bf1', 'admin', 'Administrator role'),
--     ('4a0e95a4-0e07-4d6a-a78c-f66d82b8a5d9', 'user', 'User role');

-- -- Insert data into the "permissions" table
-- INSERT INTO permissions (id, name, description) VALUES
--     ('56cb41a9-32ed-45e7-9003-c31e463df216', 'read_user', 'Read access to user'),
--     ('fc95a271-2e0f-43ab-9ff5-5c1f0a824c42', 'edit_user', 'Edit access to user'),
--     ('ac2b875d-21f6-4c22-9e84-33a59077a3b3', 'delete_user', 'Delete access to user');

-- -- Insert data into the "users" table
-- INSERT INTO users (id, created_at, email, full_name, password_hash, user_status) VALUES
--     ('b6948329-228e-41c7-adc7-f6184abd7024', NOW(), 'prabhat@redcliffelabs.com', 'Prabhat Pankaj', '$2a$04$CQ/NSGUYdEqNEB/uEF3.meazCEUgxmHLniy3dGePzAvZU9PIcRBCC', 1),
--     ('dfcebd8f-aa79-4f23-8a7b-bf3707891aa1', NOW(), 'admin@example.com', 'admin admin', '$2a$04$CQ/NSGUYdEqNEB/uEF3.meazCEUgxmHLniy3dGePzAvZU9PIcRBCC', 1);

-- -- Insert data into the "role_permissions" junction table
-- INSERT INTO role_permissions (role_id, permission_id) VALUES
--     ('a6cb87e2-5d6a-4a3d-8c35-62f2f8f17bf1', '56cb41a9-32ed-45e7-9003-c31e463df216'), -- Admin role has read_user permission
--     ('a6cb87e2-5d6a-4a3d-8c35-62f2f8f17bf1', 'fc95a271-2e0f-43ab-9ff5-5c1f0a824c42'), -- Admin role has edit_user permission
--     ('a6cb87e2-5d6a-4a3d-8c35-62f2f8f17bf1', 'ac2b875d-21f6-4c22-9e84-33a59077a3b3'), -- Admin role has delete_user permission
--     ('4a0e95a4-0e07-4d6a-a78c-f66d82b8a5d9', '56cb41a9-32ed-45e7-9003-c31e463df216'); -- User role has read_user permission


-- -- Insert data into the "user_roles" junction table to associate users with roles
-- INSERT INTO user_roles (user_id, role_id) VALUES
--     ('dfcebd8f-aa79-4f23-8a7b-bf3707891aa1', 'a6cb87e2-5d6a-4a3d-8c35-62f2f8f17bf1') -- Admin has Admin role
--     ('b6948329-228e-41c7-adc7-f6184abd7024', 'a6cb87e2-5d6a-4a3d-8c35-62f2f8f17bf1'); -- Admin has User role
--     ('b6948329-228e-41c7-adc7-f6184abd7024', '4a0e95a4-0e07-4d6a-a78c-f66d82b8a5d9'); -- User has User role
