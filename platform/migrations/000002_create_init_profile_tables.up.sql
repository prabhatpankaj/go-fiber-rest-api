-- Create profile table
CREATE TABLE profile (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) UNIQUE,
    full_name VARCHAR(100),
    email VARCHAR(100),
    role VARCHAR (255) NOT NULL,
    verified_status INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL
);