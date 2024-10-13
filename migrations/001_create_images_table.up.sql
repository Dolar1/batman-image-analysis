-- migrations/001_create_images_table.up.sql

CREATE TABLE image (
    id SERIAL PRIMARY KEY,
    image_id VARCHAR(255) UNIQUE NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    original_filename VARCHAR(255) NOT NULL,
    upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    width INT,
    height INT,
    file_size BIGINT,
    file_type VARCHAR(50)
);