
/*
    Remove triggers
*/
DROP TRIGGER IF EXISTS user_status_updated_at on user_status;
DROP TRIGGER IF EXISTS note_updated_at on note;

/*
    Remove tables
*/
DROP TABLE IF EXISTS note;
DROP TABLE IF EXISTS app_user;
DROP TABLE IF EXISTS user_status;

/*
    Remove functions
*/
DROP FUNCTION IF EXISTS update_updated_at_column();
