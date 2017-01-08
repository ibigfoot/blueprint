
/*
	Create tables
*/

CREATE TABLE user_status (
    id SERIAL PRIMARY KEY,
    
    status VARCHAR(25) NOT NULL,

    created_at TIMESTAMP NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP NULL DEFAULT current_timestamp,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);


CREATE TABLE app_user (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password CHAR(60) NOT NULL,

    status_id INTEGER NOT NULL DEFAULT 1 REFERENCES user_status(id) ON DELETE CASCADE ON UPDATE CASCADE,

    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

INSERT INTO user_status (id, status, created_at, updated_at, deleted_at) VALUES
(1, 'active',   CURRENT_TIMESTAMP,  NULL,  NULL),
(2, 'inactive', CURRENT_TIMESTAMP,  NULL,  NULL);

CREATE TABLE note (
    id SERIAL PRIMARY KEY,
    
    name TEXT NOT NULL,
    
    user_id INTEGER NOT NULL REFERENCES app_user(id),
    
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

COMMIT;

/*
    FUNCTIONS 
*/

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now(); 
   RETURN NEW;
END;
$$ language 'plpgsql';

/*
    TRIGGERS  
    used for setting the update_at timestamp fields on database execute
*/

CREATE TRIGGER user_status_updated_at BEFORE UPDATE
    ON user_status FOR EACH ROW EXECUTE PROCEDURE 
    update_updated_at_column();

CREATE TRIGGER note_updated_at BEFORE UPDATE
    ON note FOR EACH ROW EXECUTE PROCEDURE 
    update_updated_at_column();