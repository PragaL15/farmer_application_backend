CREATE OR REPLACE FUNCTION manage_user_bank_details(
    operation_type VARCHAR, 
    p_id INTEGER DEFAULT NULL,
    p_user_id INTEGER DEFAULT NULL,
    p_card_number CHAR(16) DEFAULT NULL,
    p_upi_id VARCHAR(255) DEFAULT NULL,
    p_ifsc_code CHAR(11) DEFAULT NULL,
    p_account_number VARCHAR(20) DEFAULT NULL,
    p_account_holder_name VARCHAR(255) DEFAULT NULL,
    p_bank_name VARCHAR(255) DEFAULT NULL,
    p_branch_name VARCHAR(255) DEFAULT NULL,
    p_status BOOLEAN DEFAULT NULL
)
RETURNS VOID AS $$
BEGIN
    -- Handle INSERT operation
    IF operation_type = 'INSERT' THEN
        INSERT INTO user_bank_details (
            user_id, card_number, upi_id, ifsc_code, account_number, account_holder_name, bank_name, branch_name, date_of_creation, status
        ) VALUES (
            p_user_id, p_card_number, p_upi_id, p_ifsc_code, p_account_number, p_account_holder_name, p_bank_name, p_branch_name, CURRENT_TIMESTAMP, p_status
        );

    -- Handle UPDATE operation
    ELSIF operation_type = 'UPDATE' THEN
        UPDATE user_bank_details
        SET
            user_id = COALESCE(p_user_id, user_id),
            card_number = COALESCE(p_card_number, card_number),
            upi_id = COALESCE(p_upi_id, upi_id),
            ifsc_code = COALESCE(p_ifsc_code, ifsc_code),
            account_number = COALESCE(p_account_number, account_number),
            account_holder_name = COALESCE(p_account_holder_name, account_holder_name),
            bank_name = COALESCE(p_bank_name, bank_name),
            branch_name = COALESCE(p_branch_name, branch_name),
            status = COALESCE(p_status, status)
        WHERE id = p_id;

    -- Handle DELETE operation
    ELSIF operation_type = 'DELETE' THEN
        DELETE FROM user_bank_details WHERE id = p_id;

    -- Handle invalid operation types
    ELSE
        RAISE EXCEPTION 'Invalid operation type: %', operation_type;
    END IF;
END;
$$ LANGUAGE plpgsql;
