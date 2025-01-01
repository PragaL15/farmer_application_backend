CREATE OR REPLACE PROCEDURE get_all_users()
LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY 
    SELECT * FROM user_table;
END;
$$;
