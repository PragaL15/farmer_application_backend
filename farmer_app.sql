--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'WIN1252';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: acceptance_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.acceptance_status AS ENUM (
    'pending',
    'accepted',
    'rejected'
);


ALTER TYPE public.acceptance_status OWNER TO postgres;

--
-- Name: add_business(integer, character varying, character varying, integer, text, text, character varying, character varying, character varying); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.add_business(_b_typeid integer, _b_name character varying, _b_location character varying, _b_mandiid integer DEFAULT NULL::integer, _b_address text DEFAULT NULL::text, _b_comt text DEFAULT NULL::text, _b_email character varying DEFAULT NULL::character varying, _b_gstnum character varying DEFAULT NULL::character varying, _b_pan character varying DEFAULT NULL::character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO business_table (b_typeid, b_name, b_location, b_mandiid, b_address, b_comt, b_email, b_gstnum, b_pan)
    VALUES (_b_typeid, _b_name, _b_location, _b_mandiid, _b_address, _b_comt, _b_email, _b_gstnum, _b_pan);
END;
$$;


ALTER FUNCTION public.add_business(_b_typeid integer, _b_name character varying, _b_location character varying, _b_mandiid integer, _b_address text, _b_comt text, _b_email character varying, _b_gstnum character varying, _b_pan character varying) OWNER TO postgres;

--
-- Name: add_cash_payment(character varying); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.add_cash_payment(payment_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO cash_payment_list (payment_type) VALUES (payment_name);
END;
$$;


ALTER FUNCTION public.add_cash_payment(payment_name character varying) OWNER TO postgres;

--
-- Name: add_mode_of_payment(character varying); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.add_mode_of_payment(p_payment_mode character varying) RETURNS TABLE(id integer, payment_mode character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO mode_of_payments_list (payment_mode)
    VALUES (p_payment_mode)
    RETURNING mode_of_payments_list.id, mode_of_payments_list.payment_mode;
END;
$$;


ALTER FUNCTION public.add_mode_of_payment(p_payment_mode character varying) OWNER TO postgres;

--
-- Name: add_user_bank_details(integer, character, character varying, character, character varying, character varying, character varying, character varying, boolean); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.add_user_bank_details(p_user_id integer, p_card_number character, p_upi_id character varying, p_ifsc_code character, p_account_number character varying, p_account_holder_name character varying, p_bank_name character varying, p_branch_name character varying, p_status boolean DEFAULT true) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO user_bank_details (
        user_id,
        card_number,
        upi_id,
        ifsc_code,
        account_number,
        account_holder_name,
        bank_name,
        branch_name,
        status
    ) VALUES (
        p_user_id,
        p_card_number,
        p_upi_id,
        p_ifsc_code,
        p_account_number,
        p_account_holder_name,
        p_bank_name,
        p_branch_name,
        p_status
    );
END;
$$;


ALTER FUNCTION public.add_user_bank_details(p_user_id integer, p_card_number character, p_upi_id character varying, p_ifsc_code character, p_account_number character varying, p_account_holder_name character varying, p_bank_name character varying, p_branch_name character varying, p_status boolean) OWNER TO postgres;

--
-- Name: delete_category(integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_category(IN cat_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    DELETE FROM master_category_table WHERE category_id = cat_id;
END;
$$;


ALTER PROCEDURE public.delete_category(IN cat_id integer) OWNER TO postgres;

--
-- Name: delete_driver(integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_driver(IN p_driver_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    DELETE FROM master_driver_table WHERE driver_id = p_driver_id;
END;
$$;


ALTER PROCEDURE public.delete_driver(IN p_driver_id integer) OWNER TO postgres;

--
-- Name: delete_location(integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_location(IN p_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    DELETE FROM master_location WHERE id = p_id;
END;
$$;


ALTER PROCEDURE public.delete_location(IN p_id integer) OWNER TO postgres;

--
-- Name: delete_master_mandi(integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_master_mandi(IN p_mandi_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    DELETE FROM public.master_mandi_table
    WHERE mandi_id = p_mandi_id;
END;
$$;


ALTER PROCEDURE public.delete_master_mandi(IN p_mandi_id integer) OWNER TO postgres;

--
-- Name: delete_master_product(integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_master_product(IN p_product_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    DELETE FROM public.master_product
    WHERE product_id = p_product_id;
END;
$$;


ALTER PROCEDURE public.delete_master_product(IN p_product_id integer) OWNER TO postgres;

--
-- Name: delete_master_state(integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_master_state(IN p_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    DELETE FROM public.master_states
    WHERE id = p_id;
END;
$$;


ALTER PROCEDURE public.delete_master_state(IN p_id integer) OWNER TO postgres;

--
-- Name: delete_master_vehicle(integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_master_vehicle(IN p_vehicle_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    DELETE FROM public.master_vehicle_table
    WHERE vehicle_id = p_vehicle_id;
END;
$$;


ALTER PROCEDURE public.delete_master_vehicle(IN p_vehicle_id integer) OWNER TO postgres;

--
-- Name: delete_master_violation(integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_master_violation(IN p_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    DELETE FROM public.master_violation_table
    WHERE id = p_id;
END;
$$;


ALTER PROCEDURE public.delete_master_violation(IN p_id integer) OWNER TO postgres;

--
-- Name: delete_user(integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_user(IN p_user_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    DELETE FROM public.user_table
    WHERE user_id = p_user_id;
END;
$$;


ALTER PROCEDURE public.delete_user(IN p_user_id integer) OWNER TO postgres;

--
-- Name: generate_invoice_number(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.generate_invoice_number() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE
    new_invoice_number TEXT;
BEGIN
    -- Generate the invoice number with today's date and a sequential ID
    new_invoice_number := 'INV-' || TO_CHAR(CURRENT_DATE, 'YYYYMMDD') || '-' || LPAD(NEW.id::TEXT, 5, '0');

    -- Assign the generated number to invoice_number column
    NEW.invoice_number := new_invoice_number;
    
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.generate_invoice_number() OWNER TO postgres;

--
-- Name: get_all_businesses_func(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_all_businesses_func() RETURNS TABLE(bid integer, b_type_id integer, b_typename character varying, b_name character varying, b_location_id integer, location character varying, b_state_id integer, state_name character varying, b_mandi_id integer, b_address text, b_phone_num text, b_email character varying, b_gst_num character varying, b_pan_num character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        b.bid, 
        b.b_typeid AS b_type_id,   -- ? Ensured column alias matches return type
        bt.b_typename, 
        b.b_name, 
        b.b_location_id, 
        ml.location, 
        b.b_state_id, 
        ms.state AS state_name, 
        b.b_mandi_id, 
        b.b_address, 
        b.b_phone_num,  
        b.b_email,  
        b.b_gst_num,  
        b.b_pan_num  
    FROM public.business_table b
    LEFT JOIN public.business_type_table bt ON b.b_typeid = bt.b_typeid
    LEFT JOIN public.master_states ms ON b.b_state_id = ms.id  
    LEFT JOIN public.master_location ml ON b.b_location_id = ml.id;
END;
$$;


ALTER FUNCTION public.get_all_businesses_func() OWNER TO postgres;

--
-- Name: get_all_drivers(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_all_drivers() RETURNS TABLE(driver_id integer, driver_name character varying, driver_age integer, driver_license character varying, driver_number character varying, driver_address character varying, driver_status character varying, date_of_joining date, experience_years integer, license_expiry_date date, emergency_contact character varying, assigned_route_id integer, d_o_b date)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        master_driver_table.driver_id,
        master_driver_table.driver_name,
        master_driver_table.driver_age,
        master_driver_table.driver_license,
        master_driver_table.driver_number,
        master_driver_table.driver_address,
        master_driver_table.driver_status,
        master_driver_table.date_of_joining,
        master_driver_table.experience_years,
        master_driver_table.license_expiry_date,
        master_driver_table.emergency_contact,
        master_driver_table.assigned_route_id,
        master_driver_table.d_o_b
    FROM master_driver_table;
END;
$$;


ALTER FUNCTION public.get_all_drivers() OWNER TO postgres;

--
-- Name: get_all_users(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_all_users() RETURNS TABLE(user_id integer, user_type_id integer, user_type_name character varying, name character varying, mobile_num character varying, email character varying, address text, pincode character varying, location integer, location_name character varying, state integer, state_name character varying, status integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT
        u.user_id,
        u.user_type_id,
        ut.user_type AS user_type_name,
        u.name,
        u.mobile_num,
        u.email,
        u.address,
        u.pincode,
        u.location,
        ml.location AS location_name,  -- Corrected column name
        u.state,
        ms.state AS state_name,  -- Corrected column name
        u.status
    FROM user_table u
    LEFT JOIN master_location ml ON u.location = ml.id
    LEFT JOIN master_states ms ON u.state = ms.id
    LEFT JOIN user_type_table ut ON u.user_type_id = ut.user_type_id;
END;
$$;


ALTER FUNCTION public.get_all_users() OWNER TO postgres;

--
-- Name: get_business_by_id(integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_business_by_id(business_id integer) RETURNS TABLE(bid integer, b_typeid integer, b_typename text, b_name text, b_location_id integer, location text, b_state_id integer, state_name text, b_mandiid integer, b_address text, b_phone_num character varying, b_email character varying, b_gstnum character varying, b_pan character varying, driver_id integer, driver_name character varying, violation_id integer, violation_name text, entry_date date)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT
        b.bid, b.b_typeid, bt.b_typename, b.b_name,
        b.b_location_id, l.location, b.b_state_id, s.state_name,
        b.b_mandiid, b.b_address, b.b_phone_num, b.b_email,
        b.b_gstnum, b.b_pan,
        dv.driver_id, d.driver_name, dv.violation_id, v.violation_name, dv.entry_date
    FROM business_table b
    JOIN business_type_table bt ON b.b_typeid = bt.b_typeid  -- ? Fixed column name
    JOIN location_table l ON b.b_location_id = l.location_id
    JOIN state_table s ON b.b_state_id = s.state_id
    LEFT JOIN driver_violation_table dv ON b.bid = dv.driver_id  -- ? Ensure correct join
    LEFT JOIN master_driver_table d ON dv.driver_id = d.driver_id
    LEFT JOIN master_violation_table v ON dv.violation_id = v.id
    WHERE b.bid = business_id;
END;
$$;


ALTER FUNCTION public.get_business_by_id(business_id integer) OWNER TO postgres;

--
-- Name: get_business_types(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_business_types() RETURNS TABLE(b_typeid integer, b_typename character varying, remarks text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT * FROM business_type_table;
END;
$$;


ALTER FUNCTION public.get_business_types() OWNER TO postgres;

--
-- Name: get_cash_payment_list(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_cash_payment_list() RETURNS TABLE(id integer, payment_type character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT * FROM cash_payment_list;
END;
$$;


ALTER FUNCTION public.get_cash_payment_list() OWNER TO postgres;

--
-- Name: get_categories(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_categories() RETURNS TABLE(category_id integer, category_name character varying, super_cat_id integer, remarks character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        mct.category_id, 
        mct.category_name::VARCHAR(255), 
        mct.super_cat_id,   -- Keep null values as they are
        mct.remarks::VARCHAR(255)  
    FROM master_category_table AS mct;
END;
$$;


ALTER FUNCTION public.get_categories() OWNER TO postgres;

--
-- Name: get_daily_price_update(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_daily_price_update() RETURNS TABLE(product_id integer, product_name character varying, price character varying, status integer, unit_id integer, unit_name character varying, wholeseller_id integer, wholeseller_name character varying, b_mandi_id integer, mandi_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        dpu.product_id,
        mp.product_name,
        dpu.price,
        dpu.status,
        dpu.unit_id,
        ut.unit_name,
        dpu.wholeseller_id,
        bt.b_name AS wholeseller_name,
        bt.b_mandi_id,
        mmt.mandi_name
    FROM daily_price_update dpu
    JOIN master_product mp ON dpu.product_id = mp.product_id
    JOIN units_table ut ON dpu.unit_id = ut.id
    JOIN business_table bt ON dpu.wholeseller_id = bt.bid
    LEFT JOIN master_mandi_table mmt ON bt.b_mandi_id = mmt.mandi_id;
END;
$$;


ALTER FUNCTION public.get_daily_price_update() OWNER TO postgres;

--
-- Name: get_driver_violations(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_driver_violations() RETURNS TABLE(id integer, driver_id integer, driver_name character varying, violation_id integer, violation_name text, entry_date date)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        dv.id,
        dv.driver_id, 
        d.driver_name, 
        dv.violation_id, 
        v.violation_name, 
        dv.entry_date
    FROM driver_violation_table dv
    JOIN master_driver_table d ON dv.driver_id = d.driver_id
    JOIN master_violation_table v ON dv.violation_id = v.id;
END;
$$;


ALTER FUNCTION public.get_driver_violations() OWNER TO postgres;

--
-- Name: get_invoice_details(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_invoice_details() RETURNS TABLE(id integer, invoice_number character varying, order_id integer, total_amount numeric, discount_amount numeric, invoice_date timestamp without time zone, due_date date, pay_mode integer, pay_mode_name character varying, pay_type integer, pay_type_name character varying, final_amount numeric)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        i.id,
        i.invoice_number,
        i.order_id,
        i.total_amount,
        i.discount_amount,
        i.invoice_date,
        i.due_date,
        i.pay_mode,
        m.payment_mode AS pay_mode_name,
        i.pay_type,
        c.payment_type AS pay_type_name,
        i.final_amount
    FROM invoice_table i
    LEFT JOIN mode_of_payments_list m ON i.pay_mode = m.id
    LEFT JOIN cash_payment_list c ON i.pay_type = c.id;
END;
$$;


ALTER FUNCTION public.get_invoice_details() OWNER TO postgres;

--
-- Name: get_invoice_details_with_business_info(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_invoice_details_with_business_info() RETURNS TABLE(id integer, invoice_number text, order_id bigint, retailer_id bigint, retailer_name text, retailer_email text, retailer_phone text, retailer_address text, retailer_location_id integer, retailer_state_id integer, wholeseller_id bigint, wholeseller_name text, wholeseller_email text, wholeseller_phone text, wholeseller_address text, wholeseller_location_id integer, wholeseller_state_id integer, total_amount numeric, discount_amount numeric, invoice_date date, due_date date, pay_mode integer, pay_mode_name text, pay_type integer, pay_type_name text, final_amount numeric, order_item_id bigint, product_id bigint, product_name text, retailer_status text, wholeseller_state_name text, retailer_state_name text, wholeseller_location_name text, retailer_location_name text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT
        i.id, 
        i.invoice_number::TEXT,  
        i.order_id,
        b_retailer.bid AS retailer_id, 
        b_retailer.b_name::TEXT AS retailer_name,  -- ? Explicit cast
        b_retailer.b_email::TEXT AS retailer_email,
        b_retailer.b_phone_num::TEXT AS retailer_phone,
        b_retailer.b_address::TEXT AS retailer_address, 
        b_retailer.b_location_id,
        b_retailer.b_state_id,

        b_wholeseller.bid AS wholeseller_id, 
        b_wholeseller.b_name::TEXT AS wholeseller_name,
        b_wholeseller.b_email::TEXT AS wholeseller_email,
        b_wholeseller.b_phone_num::TEXT AS wholeseller_phone,
        b_wholeseller.b_address::TEXT AS wholeseller_address, 
        b_wholeseller.b_location_id,
        b_wholeseller.b_state_id,

        i.total_amount, 
        i.discount_amount, 
        i.invoice_date::DATE,  
        i.due_date::DATE,  
        m.id AS pay_mode, 
        m.payment_mode::TEXT AS pay_mode_name,
        c.id AS pay_type, 
        c.payment_type::TEXT AS pay_type_name,
        i.final_amount,
        oi.order_item_id,
        p.product_id, 
        p.product_name::TEXT,
        idt.retailer_status::TEXT,

        ms_wholeseller.state::TEXT AS wholeseller_state_name,
        ms_retailer.state::TEXT AS retailer_state_name,
        ml_wholeseller.location::TEXT AS wholeseller_location_name,
        ml_retailer.location::TEXT AS retailer_location_name

    FROM invoice_table i
    JOIN order_table o ON i.order_id = o.order_id
    JOIN business_table b_retailer ON o.retailer_id = b_retailer.bid
    JOIN business_table b_wholeseller ON o.wholeseller_id = b_wholeseller.bid
    JOIN mode_of_payments_list m ON i.pay_mode = m.id
    JOIN cash_payment_list c ON i.pay_type = c.id
    JOIN order_item_table oi ON o.order_id = oi.order_id
    JOIN master_product p ON oi.product_id = p.product_id
    JOIN invoice_details_table idt ON i.id = idt.invoice_id
    LEFT JOIN retailer_status rst ON idt.retailer_status = rst.id

    -- Join to get state names
    LEFT JOIN master_states ms_wholeseller ON b_wholeseller.b_state_id = ms_wholeseller.id
    LEFT JOIN master_states ms_retailer ON b_retailer.b_state_id = ms_retailer.id

    -- Join to get location names
    LEFT JOIN master_location ml_wholeseller ON b_wholeseller.b_location_id = ml_wholeseller.id
    LEFT JOIN master_location ml_retailer ON b_retailer.b_location_id = ml_retailer.id;
END;
$$;


ALTER FUNCTION public.get_invoice_details_with_business_info() OWNER TO postgres;

--
-- Name: get_invoice_details_with_order_info(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_invoice_details_with_order_info() RETURNS TABLE(id integer, invoice_number character varying, order_id bigint, retailer_id integer, wholeseller_id integer, total_amount numeric, discount_amount numeric, invoice_date timestamp without time zone, due_date date, pay_mode integer, pay_mode_name character varying, pay_type integer, pay_type_name character varying, final_amount numeric)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        i.id,
        i.invoice_number,
        i.order_id,
        o.retailer_id,
        o.wholeseller_id,
        i.total_amount,
        i.discount_amount,
        i.invoice_date,
        i.due_date,
        i.pay_mode,
        m.payment_mode AS pay_mode_name,
        i.pay_type,
        c.payment_type AS pay_type_name,
        i.final_amount
    FROM invoice_table i
    LEFT JOIN order_table o ON i.order_id = o.order_id
    LEFT JOIN mode_of_payments_list m ON i.pay_mode = m.id
    LEFT JOIN cash_payment_list c ON i.pay_type = c.id;
END;
$$;


ALTER FUNCTION public.get_invoice_details_with_order_info() OWNER TO postgres;

--
-- Name: get_master_categories(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_master_categories() RETURNS TABLE(category_id integer, category_name character varying, super_cat_id integer, created_at timestamp without time zone, updated_at timestamp without time zone, col1 text, col2 text, remarks text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        mc.category_id,
        mc.category_name,
        mc.super_cat_id,
        mc.created_at,
        mc.updated_at,
        mc.col1,
        mc.col2,
        mc.remarks
    FROM master_category_table mc;
END;
$$;


ALTER FUNCTION public.get_master_categories() OWNER TO postgres;

--
-- Name: get_master_locations(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_master_locations() RETURNS TABLE(id integer, location character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        ml.id, ml.location -- Use alias to avoid ambiguity
    FROM master_location ml;
END;
$$;


ALTER FUNCTION public.get_master_locations() OWNER TO postgres;

--
-- Name: get_master_mandis(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_master_mandis() RETURNS TABLE(mandi_id integer, mandi_location text, mandi_number text, mandi_incharge text, mandi_incharge_num text, mandi_pincode text, mandi_address text, remarks text, mandi_city integer, mandi_state integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT
        m.mandi_id,
        m.mandi_location::TEXT,     -- Casting to TEXT to avoid type mismatches
        m.mandi_number::TEXT,
        m.mandi_incharge::TEXT,
        m.mandi_incharge_num::TEXT,
        m.mandi_pincode::TEXT,
        m.mandi_address::TEXT,
        m.remarks::TEXT,
        m.mandi_city,
        m.mandi_state
    FROM master_mandi_table m;
END;
$$;


ALTER FUNCTION public.get_master_mandis() OWNER TO postgres;

--
-- Name: get_master_products(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_master_products() RETURNS TABLE(product_id integer, category_id integer, category_name character varying, product_name character varying, status integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        p.product_id, 
        p.category_id, 
        c.category_name,
        p.product_name, 
        p.status
    FROM master_product p
    LEFT JOIN master_category_table c ON p.category_id = c.category_id;
END;
$$;


ALTER FUNCTION public.get_master_products() OWNER TO postgres;

--
-- Name: get_master_states(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_master_states() RETURNS TABLE(id integer, state character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        ms.id, ms.state
    FROM master_states ms;
END;
$$;


ALTER FUNCTION public.get_master_states() OWNER TO postgres;

--
-- Name: get_master_vehicles(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_master_vehicles() RETURNS TABLE(vehicle_id integer, insurance_id integer, vehicle_name character varying, vehicle_manufacture_year character varying, vehicle_warranty character varying, vehicle_make integer, vehicle_make_name character varying, vehicle_model integer, vehicle_model_name character varying, vehicle_registration_no character varying, vehicle_engine_type integer, vehicle_engine_type_name character varying, vehicle_purchase_date date, vehicle_color character varying, vehicle_insurance_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        mv.vehicle_id,
        mv.insurance_id,
        mv.vehicle_name,
        mv.vehicle_manufacture_year,
        mv.vehicle_warranty,
        mv.vehicle_make,
        vm.make AS vehicle_make_name,
        mv.vehicle_model,
        vmod.model AS vehicle_model_name,
        mv.vehicle_registration_no,
        mv.vehicle_engine_type,
        vet.engine_type AS vehicle_engine_type_name,
        mv.vehicle_purchase_date,
        mv.vehicle_color,
        mv.vehicle_insurance_id
    FROM master_vehicle_table mv
    LEFT JOIN vehicle_make vm ON mv.vehicle_make = vm.id
    LEFT JOIN vehicle_model vmod ON mv.vehicle_model = vmod.id
    LEFT JOIN vehicle_engine_type vet ON mv.vehicle_engine_type = vet.id;
END;
$$;


ALTER FUNCTION public.get_master_vehicles() OWNER TO postgres;

--
-- Name: get_master_violations(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_master_violations() RETURNS TABLE(id integer, violation_name text, level_of_serious text, status integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        mv.id, mv.violation_name, mv.level_of_serious, mv.status
    FROM master_violation_table mv;
END;
$$;


ALTER FUNCTION public.get_master_violations() OWNER TO postgres;

--
-- Name: get_mode_of_payment_details(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_mode_of_payment_details() RETURNS TABLE(id integer, pay_mode integer, pay_type integer, pay_mode_name character varying, pay_type_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        mp.id, 
        mp.pay_mode, 
        mp.pay_type, 
        mpl.payment_mode AS pay_mode_name, 
        cpl.payment_type AS pay_type_name
    FROM mode_of_payment mp
    LEFT JOIN mode_of_payments_list mpl ON mp.pay_mode = mpl.id
    LEFT JOIN cash_payment_list cpl ON mp.pay_type = cpl.id;
END;
$$;


ALTER FUNCTION public.get_mode_of_payment_details() OWNER TO postgres;

--
-- Name: get_mode_of_payments(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_mode_of_payments() RETURNS TABLE(id integer, payment_mode character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT * FROM mode_of_payments_list;
END;
$$;


ALTER FUNCTION public.get_mode_of_payments() OWNER TO postgres;

--
-- Name: get_order_details(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_order_details() RETURNS TABLE(order_id bigint, order_item_id bigint, date_of_order timestamp without time zone, expected_delivery_date timestamp without time zone, actual_delivery_date timestamp without time zone, order_status text, retailer_id bigint, retailer_name text, wholeseller_id bigint, wholeseller_name text, location_name text, state_name text, pincode text, address text, total_order_amount numeric, product_id bigint, product_name text, quantity numeric, unit_id integer, unit_name text, amt_of_order_item numeric, order_item_status text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        o.order_id::BIGINT,  
        oi.order_item_id::BIGINT,  
        o.date_of_order,
        o.expected_delivery_date, -- ? Moved before order_status
        o.actual_delivery_date,   -- ? Moved before order_status
        os.order_status::TEXT,  
        o.retailer_id::BIGINT, 
        r.b_name::TEXT,  
        o.wholeseller_id::BIGINT, 
        w.b_name::TEXT,  
        l.location::TEXT,  
        s.state::TEXT,  
        o.pincode::TEXT,  
        o.address::TEXT,  
        o.total_order_amount,
        oi.product_id,
        mp.product_name::TEXT,  
        oi.quantity,
        oi.unit_id,
        u.unit_name::TEXT,  
        oi.amt_of_order_item,
        ois.order_status::TEXT  
    FROM order_table o
    JOIN order_item_table oi ON o.order_id = oi.order_id
    LEFT JOIN order_status_table os ON o.order_status = os.order_status_id  
    LEFT JOIN business_table r ON o.retailer_id = r.bid  
    LEFT JOIN business_table w ON o.wholeseller_id = w.bid  
    LEFT JOIN master_location l ON o.location_id = l.id  
    LEFT JOIN master_states s ON o.state_id = s.id  
    LEFT JOIN master_product mp ON oi.product_id = mp.product_id
    LEFT JOIN units_table u ON oi.unit_id = u.id  
    LEFT JOIN order_status_table ois ON oi.order_item_status = ois.order_status_id;
END;
$$;


ALTER FUNCTION public.get_order_details() OWNER TO postgres;

--
-- Name: get_order_history(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_order_history() RETURNS TABLE(order_id integer, date_of_order timestamp without time zone, order_status integer, expected_delivery_date timestamp without time zone, actual_delivery_date timestamp without time zone, retailer_id integer, retailer_name character varying, wholeseller_id integer, wholeseller_name character varying, location_id integer, location_name character varying, state_id integer, state_name character varying, pincode character varying, address text, delivery_completed_date timestamp without time zone, history_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        oh.order_id,
        oh.date_of_order,
        oh.order_status,
        oh.expected_delivery_date,
        oh.actual_delivery_date,
        oh.retailer_id,
        COALESCE(b1.b_name, 'Unknown') AS retailer_name,
        oh.wholeseller_id,
        COALESCE(b2.b_name, 'Unknown') AS wholeseller_name,
        oh.location_id,
        COALESCE(ml.location, 'Unknown') AS location_name,
        oh.state_id,
        COALESCE(ms.state, 'Unknown') AS state_name,
        oh.pincode,
        oh.address,
        oh.delivery_completed_date,
        oh.history_id
    FROM order_history_table oh
    LEFT JOIN business_table b1 ON oh.retailer_id = b1.bid
    LEFT JOIN business_table b2 ON oh.wholeseller_id = b2.bid
    LEFT JOIN master_location ml ON oh.location_id = ml.id
    LEFT JOIN master_states ms ON oh.state_id = ms.id;
END;
$$;


ALTER FUNCTION public.get_order_history() OWNER TO postgres;

--
-- Name: get_orders_with_items(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_orders_with_items() RETURNS TABLE(order_id integer, date_of_order timestamp without time zone, order_status integer, expected_delivery_date timestamp without time zone, actual_delivery_date timestamp without time zone, retailer_id integer, retailer_name character varying, wholeseller_id integer, wholeseller_name character varying, location_id integer, location_name character varying, state_id integer, state_name character varying, total_order_amount numeric, order_item_id bigint, product_id bigint, product_name character varying, quantity numeric, amt_of_order_item numeric)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        o.order_id,
        o.date_of_order,
        o.order_status,
        o.expected_delivery_date,
        o.actual_delivery_date,
        o.retailer_id,
        r.b_name AS retailer_name,
        o.wholeseller_id,
        w.b_name AS wholeseller_name,
        o.location_id,
        ml.location AS location_name,
        o.state_id,
        ms.state AS state_name,
        o.total_order_amount,
        oi.order_item_id,
        oi.product_id,
        mp.product_name,
        oi.quantity,
        oi.amt_of_order_item
    FROM order_table o
    LEFT JOIN business_table r ON o.retailer_id = r.bid
    LEFT JOIN business_table w ON o.wholeseller_id = w.bid
    LEFT JOIN master_location ml ON o.location_id = ml.id
    LEFT JOIN master_states ms ON o.state_id = ms.id
    LEFT JOIN order_item_table oi ON o.order_id = oi.order_id
    LEFT JOIN master_product mp ON oi.product_id = mp.product_id;
END;
$$;


ALTER FUNCTION public.get_orders_with_items() OWNER TO postgres;

--
-- Name: insert_business(integer, integer, integer, text, character varying, character varying, character varying, text, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.insert_business(IN p_b_typeid integer, IN p_b_location_id integer, IN p_b_state_id integer, IN p_b_address text, IN p_b_email character varying, IN p_b_gst_num character varying, IN p_b_pan_num character varying, IN p_b_phone_num text DEFAULT NULL::text, IN p_b_mandi_id integer DEFAULT NULL::integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO business_table (
        b_typeid, b_location_id, b_state_id, b_mandi_id,
        b_address, b_phone_num, b_email, b_gst_num, b_pan_num
    ) VALUES (
        p_b_typeid, p_b_location_id, p_b_state_id, p_b_mandi_id,
        p_b_address, p_b_phone_num, p_b_email, p_b_gst_num, p_b_pan_num
    );
END;
$$;


ALTER PROCEDURE public.insert_business(IN p_b_typeid integer, IN p_b_location_id integer, IN p_b_state_id integer, IN p_b_address text, IN p_b_email character varying, IN p_b_gst_num character varying, IN p_b_pan_num character varying, IN p_b_phone_num text, IN p_b_mandi_id integer) OWNER TO postgres;

--
-- Name: insert_business(integer, character varying, integer, integer, text, character varying, character varying, character varying, integer, text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_business(p_b_typeid integer, p_b_name character varying, p_b_location_id integer, p_b_state_id integer, p_b_address text, p_b_email character varying, p_b_gst_num character varying, p_b_pan_num character varying, p_b_mandi_id integer DEFAULT NULL::integer, p_b_phone_num text DEFAULT NULL::text) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO public.business_table (
        b_typeid, b_name, b_location_id, b_state_id,
        b_address, b_email, b_gst_num, b_pan_num,
        b_mandi_id, b_phone_num
    ) VALUES (
        p_b_typeid, p_b_name, p_b_location_id, p_b_state_id,
        p_b_address, p_b_email, p_b_gst_num, p_b_pan_num,
        NULLIF(p_b_mandi_id, 0), NULLIF(p_b_phone_num, '')
    );
END;
$$;


ALTER FUNCTION public.insert_business(p_b_typeid integer, p_b_name character varying, p_b_location_id integer, p_b_state_id integer, p_b_address text, p_b_email character varying, p_b_gst_num character varying, p_b_pan_num character varying, p_b_mandi_id integer, p_b_phone_num text) OWNER TO postgres;

--
-- Name: insert_business_type(character varying, text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_business_type(p_b_typename character varying, p_remarks text) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO business_type_table (b_typename, remarks)
    VALUES (p_b_typename, p_remarks);
END;
$$;


ALTER FUNCTION public.insert_business_type(p_b_typename character varying, p_remarks text) OWNER TO postgres;

--
-- Name: insert_category(character varying, integer, text, text, text); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.insert_category(IN cat_name character varying, IN super_cat integer DEFAULT NULL::integer, IN col1_text text DEFAULT NULL::text, IN col2_text text DEFAULT NULL::text, IN remarks_text text DEFAULT NULL::text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO master_category_table (category_name, super_cat_id, col1, col2, remarks, created_at, updated_at)
    VALUES (cat_name, super_cat, col1_text, col2_text, remarks_text, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
END;
$$;


ALTER PROCEDURE public.insert_category(IN cat_name character varying, IN super_cat integer, IN col1_text text, IN col2_text text, IN remarks_text text) OWNER TO postgres;

--
-- Name: insert_driver(character varying, integer, character varying, character varying, character varying, character varying, date, integer, date, character varying, integer, text, text, date); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_driver(p_driver_name character varying, p_driver_age integer, p_driver_license character varying, p_driver_number character varying, p_driver_address character varying, p_driver_status character varying, p_date_of_joining date, p_experience_years integer, p_license_expiry_date date, p_emergency_contact character varying, p_assigned_route_id integer, p_col1 text, p_col2 text, p_d_o_b date) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO master_driver_table (
        driver_name, driver_age, driver_license, driver_number, driver_address, driver_status,
        date_of_joining, experience_years, license_expiry_date, emergency_contact, assigned_route_id, 
        col1, col2, d_o_b, updated_at
    ) VALUES (
        p_driver_name, p_driver_age, p_driver_license, p_driver_number, p_driver_address, p_driver_status,
        p_date_of_joining, p_experience_years, p_license_expiry_date, p_emergency_contact, p_assigned_route_id, 
        p_col1, p_col2, p_d_o_b, CURRENT_TIMESTAMP
    );
END;
$$;


ALTER FUNCTION public.insert_driver(p_driver_name character varying, p_driver_age integer, p_driver_license character varying, p_driver_number character varying, p_driver_address character varying, p_driver_status character varying, p_date_of_joining date, p_experience_years integer, p_license_expiry_date date, p_emergency_contact character varying, p_assigned_route_id integer, p_col1 text, p_col2 text, p_d_o_b date) OWNER TO postgres;

--
-- Name: insert_driver(character varying, integer, character varying, character varying, character varying, character varying, date, integer, integer, date, character varying, integer, text, text, date, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.insert_driver(IN p_driver_name character varying, IN p_driver_age integer DEFAULT NULL::integer, IN p_driver_license character varying DEFAULT NULL::character varying, IN p_driver_number character varying DEFAULT NULL::character varying, IN p_driver_address character varying DEFAULT NULL::character varying, IN p_driver_status character varying DEFAULT NULL::character varying, IN p_date_of_joining date DEFAULT NULL::date, IN p_experience_years integer DEFAULT NULL::integer, IN p_vehicle_id integer DEFAULT NULL::integer, IN p_license_expiry_date date DEFAULT NULL::date, IN p_emergency_contact character varying DEFAULT NULL::character varying, IN p_assigned_route_id integer DEFAULT NULL::integer, IN p_col1 text DEFAULT NULL::text, IN p_col2 text DEFAULT NULL::text, IN p_d_o_b date DEFAULT NULL::date, IN p_violation integer DEFAULT NULL::integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO master_driver_table (
        driver_name, driver_age, driver_license, driver_number, driver_address,
        driver_status, date_of_joining, experience_years, vehicle_id,
        license_expiry_date, emergency_contact, assigned_route_id, col1, col2, d_o_b, violation, created_at
    ) VALUES (
        p_driver_name, p_driver_age, p_driver_license, p_driver_number, p_driver_address,
        p_driver_status, p_date_of_joining, p_experience_years, p_vehicle_id,
        p_license_expiry_date, p_emergency_contact, p_assigned_route_id, p_col1, p_col2, p_d_o_b, p_violation, CURRENT_TIMESTAMP
    );
END;
$$;


ALTER PROCEDURE public.insert_driver(IN p_driver_name character varying, IN p_driver_age integer, IN p_driver_license character varying, IN p_driver_number character varying, IN p_driver_address character varying, IN p_driver_status character varying, IN p_date_of_joining date, IN p_experience_years integer, IN p_vehicle_id integer, IN p_license_expiry_date date, IN p_emergency_contact character varying, IN p_assigned_route_id integer, IN p_col1 text, IN p_col2 text, IN p_d_o_b date, IN p_violation integer) OWNER TO postgres;

--
-- Name: insert_location(character varying); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.insert_location(IN p_location character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO master_location (location) 
    VALUES (p_location);
END;
$$;


ALTER PROCEDURE public.insert_location(IN p_location character varying) OWNER TO postgres;

--
-- Name: insert_master_mandi(character varying, character varying, character varying, character varying, character varying, text, text, integer, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.insert_master_mandi(IN p_mandi_location character varying, IN p_mandi_number character varying, IN p_mandi_incharge character varying, IN p_mandi_incharge_num character varying, IN p_mandi_pincode character varying, IN p_mandi_address text, IN p_remarks text, IN p_mandi_city integer, IN p_mandi_state integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO public.master_mandi_table (
        mandi_location, mandi_number, mandi_incharge,
        mandi_incharge_num, mandi_pincode, mandi_address,
        remarks, mandi_city, mandi_state
    ) 
    VALUES (
        p_mandi_location, p_mandi_number, p_mandi_incharge,
        p_mandi_incharge_num, p_mandi_pincode, p_mandi_address,
        p_remarks, p_mandi_city, p_mandi_state
    );
END;
$$;


ALTER PROCEDURE public.insert_master_mandi(IN p_mandi_location character varying, IN p_mandi_number character varying, IN p_mandi_incharge character varying, IN p_mandi_incharge_num character varying, IN p_mandi_pincode character varying, IN p_mandi_address text, IN p_remarks text, IN p_mandi_city integer, IN p_mandi_state integer) OWNER TO postgres;

--
-- Name: insert_master_product(integer, character varying, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.insert_master_product(IN p_category_id integer, IN p_product_name character varying, IN p_status integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO public.master_product (
        category_id, product_name, status
    ) 
    VALUES (
        p_category_id, p_product_name, COALESCE(p_status, 0)
    );
END;
$$;


ALTER PROCEDURE public.insert_master_product(IN p_category_id integer, IN p_product_name character varying, IN p_status integer) OWNER TO postgres;

--
-- Name: insert_master_state(character varying); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.insert_master_state(IN p_state character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO public.master_states (state) 
    VALUES (p_state);
END;
$$;


ALTER PROCEDURE public.insert_master_state(IN p_state character varying) OWNER TO postgres;

--
-- Name: insert_master_vehicle(integer, character varying, character varying, character varying, integer, integer, character varying, integer, date, character varying, text, text, text, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.insert_master_vehicle(IN p_insurance_id integer, IN p_vehicle_name character varying, IN p_vehicle_manufacture_year character varying, IN p_vehicle_warranty character varying, IN p_vehicle_make integer, IN p_vehicle_model integer, IN p_vehicle_registration_no character varying, IN p_vehicle_engine_type integer, IN p_vehicle_purchase_date date, IN p_vehicle_color character varying, IN p_col1 text, IN p_col2 text, IN p_col3 text, IN p_vehicle_insurance_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO public.master_vehicle_table (
        insurance_id, vehicle_name, vehicle_manufacture_year, vehicle_warranty,
        vehicle_make, vehicle_model, vehicle_registration_no, vehicle_engine_type,
        vehicle_purchase_date, vehicle_color, col1, col2, col3, vehicle_insurance_id
    ) 
    VALUES (
        p_insurance_id, p_vehicle_name, p_vehicle_manufacture_year, p_vehicle_warranty,
        p_vehicle_make, p_vehicle_model, p_vehicle_registration_no, p_vehicle_engine_type,
        p_vehicle_purchase_date, p_vehicle_color, p_col1, p_col2, p_col3, p_vehicle_insurance_id
    );
END;
$$;


ALTER PROCEDURE public.insert_master_vehicle(IN p_insurance_id integer, IN p_vehicle_name character varying, IN p_vehicle_manufacture_year character varying, IN p_vehicle_warranty character varying, IN p_vehicle_make integer, IN p_vehicle_model integer, IN p_vehicle_registration_no character varying, IN p_vehicle_engine_type integer, IN p_vehicle_purchase_date date, IN p_vehicle_color character varying, IN p_col1 text, IN p_col2 text, IN p_col3 text, IN p_vehicle_insurance_id integer) OWNER TO postgres;

--
-- Name: insert_master_violation(text, text, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.insert_master_violation(IN p_violation_name text, IN p_level_of_serious text, IN p_status integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO public.master_violation_table (violation_name, level_of_serious, status)
    VALUES (p_violation_name, p_level_of_serious, COALESCE(p_status, 1));
END;
$$;


ALTER PROCEDURE public.insert_master_violation(IN p_violation_name text, IN p_level_of_serious text, IN p_status integer) OWNER TO postgres;

--
-- Name: insert_user(integer, character varying, character varying, character varying, text, character varying, integer, integer, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_user(p_user_type_id integer, p_name character varying, p_mobile_num character varying, p_email character varying, p_address text, p_pincode character varying, p_location integer, p_state integer, p_status integer) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO user_table (user_type_id, name, mobile_num, email, address, pincode, location, state, status)
    VALUES (p_user_type_id, p_name, p_mobile_num, p_email, p_address, p_pincode, p_location, p_state, p_status);
END;
$$;


ALTER FUNCTION public.insert_user(p_user_type_id integer, p_name character varying, p_mobile_num character varying, p_email character varying, p_address text, p_pincode character varying, p_location integer, p_state integer, p_status integer) OWNER TO postgres;

--
-- Name: insert_user(integer, character varying, date, character varying, character varying, text, character varying, integer, character varying, character varying, character varying, date, character varying, character varying, integer, character varying, text, text, text, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.insert_user(IN p_user_type_id integer, IN p_name character varying, IN p_dt_of_commence_business date, IN p_mobile_num character varying, IN p_email character varying, IN p_address text, IN p_pincode character varying, IN p_location integer, IN p_business_license_no character varying, IN p_validity character varying, IN p_gst_no character varying, IN p_expiry_dt date, IN p_business_name character varying, IN p_business_type character varying, IN p_mandi_id integer, IN p_mandi_type_id character varying, IN p_remarks text, IN p_col1 text, IN p_col2 text, IN p_state integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO public.user_table (
        user_type_id, name, dt_of_commence_business, mobile_num, email, address,
        pincode, location, business_license_no, validity, gst_no, expiry_dt,
        business_name, business_type, mandi_id, mandi_type_id, remarks, col1, col2, state
    )
    VALUES (
        p_user_type_id, p_name, p_dt_of_commence_business, p_mobile_num, p_email, p_address,
        p_pincode, p_location, p_business_license_no, p_validity, p_gst_no, p_expiry_dt,
        p_business_name, p_business_type, p_mandi_id, p_mandi_type_id, p_remarks, p_col1, p_col2, p_state
    );
END;
$$;


ALTER PROCEDURE public.insert_user(IN p_user_type_id integer, IN p_name character varying, IN p_dt_of_commence_business date, IN p_mobile_num character varying, IN p_email character varying, IN p_address text, IN p_pincode character varying, IN p_location integer, IN p_business_license_no character varying, IN p_validity character varying, IN p_gst_no character varying, IN p_expiry_dt date, IN p_business_name character varying, IN p_business_type character varying, IN p_mandi_id integer, IN p_mandi_type_id character varying, IN p_remarks text, IN p_col1 text, IN p_col2 text, IN p_state integer) OWNER TO postgres;

--
-- Name: manage_user_bank_details(character varying, integer, integer, character, character varying, character, character varying, character varying, character varying, character varying, boolean); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.manage_user_bank_details(operation_type character varying, p_id integer DEFAULT NULL::integer, p_user_id integer DEFAULT NULL::integer, p_card_number character DEFAULT NULL::bpchar, p_upi_id character varying DEFAULT NULL::character varying, p_ifsc_code character DEFAULT NULL::bpchar, p_account_number character varying DEFAULT NULL::character varying, p_account_holder_name character varying DEFAULT NULL::character varying, p_bank_name character varying DEFAULT NULL::character varying, p_branch_name character varying DEFAULT NULL::character varying, p_status boolean DEFAULT NULL::boolean) RETURNS json
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF operation_type = 'INSERT' THEN
        INSERT INTO user_bank_details(
            user_id, card_number, upi_id, ifsc_code, account_number, 
            account_holder_name, bank_name, branch_name, status
        ) VALUES (
            p_user_id, p_card_number, p_upi_id, p_ifsc_code, p_account_number,
            p_account_holder_name, p_bank_name, p_branch_name, p_status
        );
        RETURN json_build_object('status', 'success', 'message', 'User bank detail added successfully');
    ELSE
        RETURN json_build_object('status', 'error', 'message', 'Unsupported operation');
    END IF;
END;
$$;


ALTER FUNCTION public.manage_user_bank_details(operation_type character varying, p_id integer, p_user_id integer, p_card_number character, p_upi_id character varying, p_ifsc_code character, p_account_number character varying, p_account_holder_name character varying, p_bank_name character varying, p_branch_name character varying, p_status boolean) OWNER TO postgres;

--
-- Name: manage_users(text, integer, integer, character varying, character varying); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.manage_users(IN op_type text, IN p_user_id integer DEFAULT NULL::integer, IN p_user_type_id integer DEFAULT NULL::integer, IN p_username character varying DEFAULT NULL::character varying, IN p_password character varying DEFAULT NULL::character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Insert Operation
    IF op_type = 'INSERT' THEN
        INSERT INTO users (user_type_id, username, password)
        VALUES (p_user_type_id, p_username, p_password);

    -- Update Operation
    ELSIF op_type = 'UPDATE' THEN
        UPDATE users
        SET 
            user_type_id = COALESCE(p_user_type_id, user_type_id),
            username = COALESCE(p_username, username),
            password = COALESCE(p_password, password)
        WHERE user_id = p_user_id;

    -- Delete Operation
    ELSIF op_type = 'DELETE' THEN
        DELETE FROM users WHERE user_id = p_user_id;

    -- Handle Invalid Operation
    ELSE
        RAISE EXCEPTION 'Invalid operation type: %', op_type;
    END IF;
END;
$$;


ALTER PROCEDURE public.manage_users(IN op_type text, IN p_user_id integer, IN p_user_type_id integer, IN p_username character varying, IN p_password character varying) OWNER TO postgres;

--
-- Name: set_default_pay_type(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.set_default_pay_type() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.pay_mode = 1 THEN
        NEW.pay_type = 5;
    END IF;
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.set_default_pay_type() OWNER TO postgres;

--
-- Name: sp_get_order_status(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.sp_get_order_status() RETURNS TABLE(order_id integer, order_status character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT * FROM order_status_table;
END;
$$;


ALTER FUNCTION public.sp_get_order_status() OWNER TO postgres;

--
-- Name: sp_get_orders(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.sp_get_orders() RETURNS TABLE(order_id bigint, date_of_order timestamp without time zone, order_status integer, order_status_name text, expected_delivery_date timestamp without time zone, actual_delivery_date timestamp without time zone, retailer_id bigint, retailer_name text, wholeseller_id bigint, wholeseller_name text, location_id integer, location_name text, state_id integer, state_name text, pincode text, address text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT
        o.order_id::BIGINT,  -- Explicit casting
        o.date_of_order,
        o.order_status,
        os.order_status AS order_status_name,
        o.expected_delivery_date,
        o.actual_delivery_date,
        o.retailer_id::BIGINT,  -- Explicit casting
        r.b_name AS retailer_name,
        o.wholeseller_id::BIGINT,  -- Explicit casting
        w.b_name AS wholeseller_name,
        o.location_id::INT,  -- Explicit casting
        ml.location AS location_name,
        o.state_id::INT,  -- Explicit casting
        ms.state AS state_name,
        o.pincode,
        o.address
    FROM order_table o
    LEFT JOIN order_status_table os ON o.order_status = os.order_status_id
    LEFT JOIN business_table r ON o.retailer_id = r.bid
    LEFT JOIN business_table w ON o.wholeseller_id = w.bid
    LEFT JOIN master_location ml ON o.location_id = ml.id
    LEFT JOIN master_states ms ON o.state_id = ms.id;
END;
$$;


ALTER FUNCTION public.sp_get_orders() OWNER TO postgres;

--
-- Name: sp_insert_order(integer, timestamp without time zone, timestamp without time zone, integer, integer, integer, integer, character varying, text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.sp_insert_order(p_order_status integer, p_expected_delivery_date timestamp without time zone, p_actual_delivery_date timestamp without time zone, p_retailer_id integer, p_wholeseller_id integer, p_location_id integer, p_state_id integer, p_pincode character varying, p_address text) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO order_table (
        date_of_order,
        order_status,
        expected_delivery_date,
        actual_delivery_date,
        retailer_id,
        wholeseller_id,
        location_id,
        state_id,
        pincode,
        address
    ) VALUES (
        CURRENT_TIMESTAMP,
        p_order_status,
        p_expected_delivery_date,
        p_actual_delivery_date,
        p_retailer_id,
        p_wholeseller_id,
        p_location_id,
        p_state_id,
        p_pincode,
        p_address
    );
END;
$$;


ALTER FUNCTION public.sp_insert_order(p_order_status integer, p_expected_delivery_date timestamp without time zone, p_actual_delivery_date timestamp without time zone, p_retailer_id integer, p_wholeseller_id integer, p_location_id integer, p_state_id integer, p_pincode character varying, p_address text) OWNER TO postgres;

--
-- Name: sp_insert_order_status(character varying); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.sp_insert_order_status(IN p_order_status character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO order_status_table (order_status)
    VALUES (p_order_status);
END;
$$;


ALTER PROCEDURE public.sp_insert_order_status(IN p_order_status character varying) OWNER TO postgres;

--
-- Name: sp_update_order(integer, integer, timestamp without time zone); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.sp_update_order(p_order_id integer, p_order_status integer, p_actual_delivery_date timestamp without time zone) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE order_table
    SET 
        order_status = p_order_status,
        actual_delivery_date = p_actual_delivery_date
    WHERE order_id = p_order_id;
END;
$$;


ALTER FUNCTION public.sp_update_order(p_order_id integer, p_order_status integer, p_actual_delivery_date timestamp without time zone) OWNER TO postgres;

--
-- Name: sp_update_order_status(integer, character varying); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.sp_update_order_status(IN p_order_id integer, IN p_order_status character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE order_status_table
    SET order_status = p_order_status
    WHERE order_id = p_order_id;
END;
$$;


ALTER PROCEDURE public.sp_update_order_status(IN p_order_id integer, IN p_order_status character varying) OWNER TO postgres;

--
-- Name: update_amt_of_order_item(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_amt_of_order_item() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE 
    product_price NUMERIC(10,2);
    tax_rate NUMERIC(5,2);
BEGIN
    -- Fetch price and tax_percentage from product_price_table
    SELECT price, tax_percentage INTO product_price, tax_rate 
    FROM product_price_table 
    WHERE product_id = NEW.product_id;

    -- If product exists, calculate the amt_of_order_item
    IF product_price IS NOT NULL THEN
        NEW.amt_of_order_item := NEW.quantity * product_price * (1 + tax_rate / 100);
    END IF;

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_amt_of_order_item() OWNER TO postgres;

--
-- Name: update_business(integer, integer, integer, integer, text, text, character varying, character varying, character varying, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.update_business(IN p_bid integer, IN p_b_typeid integer, IN p_b_location_id integer, IN p_b_state_id integer, IN p_b_address text, IN p_b_phone_num text, IN p_b_email character varying, IN p_b_gst_num character varying, IN p_b_pan_num character varying, IN p_b_mandi_id integer DEFAULT NULL::integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE business_table
    SET 
        b_typeid = p_b_typeid,
        b_location_id = p_b_location_id,
        b_state_id = p_b_state_id,
        b_mandi_id = p_b_mandi_id,
        b_address = p_b_address,
        b_phone_num = p_b_phone_num,
        b_email = p_b_email,
        b_gst_num = p_b_gst_num,
        b_pan_num = p_b_pan_num
    WHERE bid = p_bid;
END;
$$;


ALTER PROCEDURE public.update_business(IN p_bid integer, IN p_b_typeid integer, IN p_b_location_id integer, IN p_b_state_id integer, IN p_b_address text, IN p_b_phone_num text, IN p_b_email character varying, IN p_b_gst_num character varying, IN p_b_pan_num character varying, IN p_b_mandi_id integer) OWNER TO postgres;

--
-- Name: update_business(integer, integer, character varying, character varying, integer, text, text, character varying, character varying, character varying); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_business(_bid integer, _b_typeid integer, _b_name character varying, _b_location character varying, _b_mandiid integer DEFAULT NULL::integer, _b_address text DEFAULT NULL::text, _b_comt text DEFAULT NULL::text, _b_email character varying DEFAULT NULL::character varying, _b_gstnum character varying DEFAULT NULL::character varying, _b_pan character varying DEFAULT NULL::character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE business_table
    SET b_typeid = _b_typeid,
        b_name = _b_name,
        b_location = _b_location,
        b_mandiid = _b_mandiid,
        b_address = _b_address,
        b_comt = _b_comt,
        b_email = _b_email,
        b_gstnum = _b_gstnum,
        b_pan = _b_pan
    WHERE bid = _bid;
END;
$$;


ALTER FUNCTION public.update_business(_bid integer, _b_typeid integer, _b_name character varying, _b_location character varying, _b_mandiid integer, _b_address text, _b_comt text, _b_email character varying, _b_gstnum character varying, _b_pan character varying) OWNER TO postgres;

--
-- Name: update_business_type(integer, character varying, text); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_business_type(p_b_typeid integer, p_b_typename character varying, p_remarks text) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE business_type_table
    SET b_typename = p_b_typename,
        remarks = p_remarks
    WHERE b_typeid = p_b_typeid;
END;
$$;


ALTER FUNCTION public.update_business_type(p_b_typeid integer, p_b_typename character varying, p_remarks text) OWNER TO postgres;

--
-- Name: update_cash_payment(integer, character varying); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_cash_payment(payment_id integer, payment_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE cash_payment_list 
    SET payment_type = payment_name 
    WHERE id = payment_id;
END;
$$;


ALTER FUNCTION public.update_cash_payment(payment_id integer, payment_name character varying) OWNER TO postgres;

--
-- Name: update_category(integer, character varying, integer, text, text, text); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.update_category(IN cat_id integer, IN new_cat_name character varying, IN new_super_cat integer DEFAULT NULL::integer, IN new_col1 text DEFAULT NULL::text, IN new_col2 text DEFAULT NULL::text, IN new_remarks text DEFAULT NULL::text)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE master_category_table
    SET 
        category_name = new_cat_name,
        super_cat_id = new_super_cat,
        col1 = new_col1,
        col2 = new_col2,
        remarks = new_remarks,
        updated_at = CURRENT_TIMESTAMP
    WHERE category_id = cat_id;
END;
$$;


ALTER PROCEDURE public.update_category(IN cat_id integer, IN new_cat_name character varying, IN new_super_cat integer, IN new_col1 text, IN new_col2 text, IN new_remarks text) OWNER TO postgres;

--
-- Name: update_driver(integer, character varying, integer, character varying, character varying, character varying, character varying, date, integer, date, character varying, integer, text, text, date); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_driver(p_driver_id integer, p_driver_name character varying, p_driver_age integer, p_driver_license character varying, p_driver_number character varying, p_driver_address character varying, p_driver_status character varying, p_date_of_joining date, p_experience_years integer, p_license_expiry_date date, p_emergency_contact character varying, p_assigned_route_id integer, p_col1 text, p_col2 text, p_d_o_b date) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE master_driver_table
    SET
        driver_name = p_driver_name,
        driver_age = p_driver_age,
        driver_license = p_driver_license,
        driver_number = p_driver_number,
        driver_address = p_driver_address,
        driver_status = p_driver_status,
        date_of_joining = p_date_of_joining,
        experience_years = p_experience_years,
        license_expiry_date = p_license_expiry_date,
        emergency_contact = p_emergency_contact,
        assigned_route_id = p_assigned_route_id,
        col1 = p_col1,
        col2 = p_col2,
        d_o_b = p_d_o_b,
        updated_at = CURRENT_TIMESTAMP
    WHERE driver_id = p_driver_id;
END;
$$;


ALTER FUNCTION public.update_driver(p_driver_id integer, p_driver_name character varying, p_driver_age integer, p_driver_license character varying, p_driver_number character varying, p_driver_address character varying, p_driver_status character varying, p_date_of_joining date, p_experience_years integer, p_license_expiry_date date, p_emergency_contact character varying, p_assigned_route_id integer, p_col1 text, p_col2 text, p_d_o_b date) OWNER TO postgres;

--
-- Name: update_driver(integer, character varying, integer, character varying, character varying, character varying, character varying, date, integer, integer, date, character varying, integer, text, text, date, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.update_driver(IN p_driver_id integer, IN p_driver_name character varying DEFAULT NULL::character varying, IN p_driver_age integer DEFAULT NULL::integer, IN p_driver_license character varying DEFAULT NULL::character varying, IN p_driver_number character varying DEFAULT NULL::character varying, IN p_driver_address character varying DEFAULT NULL::character varying, IN p_driver_status character varying DEFAULT NULL::character varying, IN p_date_of_joining date DEFAULT NULL::date, IN p_experience_years integer DEFAULT NULL::integer, IN p_vehicle_id integer DEFAULT NULL::integer, IN p_license_expiry_date date DEFAULT NULL::date, IN p_emergency_contact character varying DEFAULT NULL::character varying, IN p_assigned_route_id integer DEFAULT NULL::integer, IN p_col1 text DEFAULT NULL::text, IN p_col2 text DEFAULT NULL::text, IN p_d_o_b date DEFAULT NULL::date, IN p_violation integer DEFAULT NULL::integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE master_driver_table  
    SET 
        driver_name = COALESCE(p_driver_name, driver_name),
        driver_age = COALESCE(p_driver_age, driver_age),
        driver_license = COALESCE(p_driver_license, driver_license),
        driver_number = COALESCE(p_driver_number, driver_number),
        driver_address = COALESCE(p_driver_address, driver_address),
        driver_status = COALESCE(p_driver_status, driver_status),
        date_of_joining = COALESCE(p_date_of_joining, date_of_joining),
        experience_years = COALESCE(p_experience_years, experience_years),
        vehicle_id = COALESCE(p_vehicle_id, vehicle_id),
        license_expiry_date = COALESCE(p_license_expiry_date, license_expiry_date),
        emergency_contact = COALESCE(p_emergency_contact, emergency_contact),
        assigned_route_id = COALESCE(p_assigned_route_id, assigned_route_id),
        col1 = COALESCE(p_col1, col1),
        col2 = COALESCE(p_col2, col2),
        d_o_b = COALESCE(p_d_o_b, d_o_b),
        violation = COALESCE(p_violation, violation),
        updated_at = CURRENT_TIMESTAMP
    WHERE driver_id = p_driver_id;
END;
$$;


ALTER PROCEDURE public.update_driver(IN p_driver_id integer, IN p_driver_name character varying, IN p_driver_age integer, IN p_driver_license character varying, IN p_driver_number character varying, IN p_driver_address character varying, IN p_driver_status character varying, IN p_date_of_joining date, IN p_experience_years integer, IN p_vehicle_id integer, IN p_license_expiry_date date, IN p_emergency_contact character varying, IN p_assigned_route_id integer, IN p_col1 text, IN p_col2 text, IN p_d_o_b date, IN p_violation integer) OWNER TO postgres;

--
-- Name: update_location(integer, character varying); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.update_location(IN p_id integer, IN p_location character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE master_location 
    SET location = p_location
    WHERE id = p_id;
END;
$$;


ALTER PROCEDURE public.update_location(IN p_id integer, IN p_location character varying) OWNER TO postgres;

--
-- Name: update_master_mandi(integer, character varying, character varying, character varying, character varying, character varying, text, text, integer, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.update_master_mandi(IN p_mandi_id integer, IN p_mandi_location character varying, IN p_mandi_number character varying, IN p_mandi_incharge character varying, IN p_mandi_incharge_num character varying, IN p_mandi_pincode character varying, IN p_mandi_address text, IN p_remarks text, IN p_mandi_city integer, IN p_mandi_state integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE public.master_mandi_table
    SET
        mandi_location = p_mandi_location,
        mandi_number = p_mandi_number,
        mandi_incharge = p_mandi_incharge,
        mandi_incharge_num = p_mandi_incharge_num,
        mandi_pincode = p_mandi_pincode,
        mandi_address = p_mandi_address,
        remarks = p_remarks,
        mandi_city = p_mandi_city,
        mandi_state = p_mandi_state,
        updated_at = CURRENT_TIMESTAMP
    WHERE mandi_id = p_mandi_id;
END;
$$;


ALTER PROCEDURE public.update_master_mandi(IN p_mandi_id integer, IN p_mandi_location character varying, IN p_mandi_number character varying, IN p_mandi_incharge character varying, IN p_mandi_incharge_num character varying, IN p_mandi_pincode character varying, IN p_mandi_address text, IN p_remarks text, IN p_mandi_city integer, IN p_mandi_state integer) OWNER TO postgres;

--
-- Name: update_master_product(integer, integer, character varying, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.update_master_product(IN p_product_id integer, IN p_category_id integer, IN p_product_name character varying, IN p_status integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE public.master_product
    SET
        category_id = p_category_id,
        product_name = p_product_name,
        status = p_status
    WHERE product_id = p_product_id;
END;
$$;


ALTER PROCEDURE public.update_master_product(IN p_product_id integer, IN p_category_id integer, IN p_product_name character varying, IN p_status integer) OWNER TO postgres;

--
-- Name: update_master_state(integer, character varying); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.update_master_state(IN p_id integer, IN p_state character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE public.master_states
    SET state = p_state
    WHERE id = p_id;
END;
$$;


ALTER PROCEDURE public.update_master_state(IN p_id integer, IN p_state character varying) OWNER TO postgres;

--
-- Name: update_master_vehicle(integer, integer, character varying, character varying, character varying, integer, integer, character varying, integer, date, character varying, text, text, text, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.update_master_vehicle(IN p_vehicle_id integer, IN p_insurance_id integer, IN p_vehicle_name character varying, IN p_vehicle_manufacture_year character varying, IN p_vehicle_warranty character varying, IN p_vehicle_make integer, IN p_vehicle_model integer, IN p_vehicle_registration_no character varying, IN p_vehicle_engine_type integer, IN p_vehicle_purchase_date date, IN p_vehicle_color character varying, IN p_col1 text, IN p_col2 text, IN p_col3 text, IN p_vehicle_insurance_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE public.master_vehicle_table
    SET
        insurance_id = p_insurance_id,
        vehicle_name = p_vehicle_name,
        vehicle_manufacture_year = p_vehicle_manufacture_year,
        vehicle_warranty = p_vehicle_warranty,
        vehicle_make = p_vehicle_make,
        vehicle_model = p_vehicle_model,
        vehicle_registration_no = p_vehicle_registration_no,
        vehicle_engine_type = p_vehicle_engine_type,
        vehicle_purchase_date = p_vehicle_purchase_date,
        vehicle_color = p_vehicle_color,
        col1 = p_col1,
        col2 = p_col2,
        col3 = p_col3,
        vehicle_insurance_id = p_vehicle_insurance_id,
        updated_at = CURRENT_TIMESTAMP
    WHERE vehicle_id = p_vehicle_id;
END;
$$;


ALTER PROCEDURE public.update_master_vehicle(IN p_vehicle_id integer, IN p_insurance_id integer, IN p_vehicle_name character varying, IN p_vehicle_manufacture_year character varying, IN p_vehicle_warranty character varying, IN p_vehicle_make integer, IN p_vehicle_model integer, IN p_vehicle_registration_no character varying, IN p_vehicle_engine_type integer, IN p_vehicle_purchase_date date, IN p_vehicle_color character varying, IN p_col1 text, IN p_col2 text, IN p_col3 text, IN p_vehicle_insurance_id integer) OWNER TO postgres;

--
-- Name: update_master_violation(integer, text, text, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.update_master_violation(IN p_id integer, IN p_violation_name text, IN p_level_of_serious text, IN p_status integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE public.master_violation_table
    SET violation_name = p_violation_name,
        level_of_serious = p_level_of_serious,
        status = p_status
    WHERE id = p_id;
END;
$$;


ALTER PROCEDURE public.update_master_violation(IN p_id integer, IN p_violation_name text, IN p_level_of_serious text, IN p_status integer) OWNER TO postgres;

--
-- Name: update_mode_of_payment(integer, character varying); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_mode_of_payment(p_id integer, p_payment_mode character varying) RETURNS TABLE(id integer, payment_mode character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE mode_of_payments_list
    SET payment_mode = p_payment_mode
    WHERE id = p_id
    RETURNING mode_of_payments_list.id, mode_of_payments_list.payment_mode;
END;
$$;


ALTER FUNCTION public.update_mode_of_payment(p_id integer, p_payment_mode character varying) OWNER TO postgres;

--
-- Name: update_order_history(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_order_history() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Insert updated order details into order_history_table
    INSERT INTO order_history_table (
        order_id, 
        date_of_order, 
        order_status, 
        expected_delivery_date, 
        actual_delivery_date, 
        retailer_id, 
        wholeseller_id, 
        location_id, 
        state_id, 
        pincode, 
        address, 
        delivery_completed_date
    )
    VALUES (
        NEW.order_id, 
        NEW.date_of_order, 
        NEW.order_status, 
        NEW.expected_delivery_date, 
        NEW.actual_delivery_date, 
        NEW.retailer_id, 
        NEW.wholeseller_id, 
        NEW.location_id, 
        NEW.state_id, 
        NEW.pincode, 
        NEW.address,
        CASE 
            WHEN NEW.order_status = 6 THEN CURRENT_TIMESTAMP 
            ELSE NULL 
        END
    );

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_order_history() OWNER TO postgres;

--
-- Name: update_total_order_amount(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_total_order_amount() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE order_table
    SET total_order_amount = (
        SELECT COALESCE(SUM(amt_of_order_item), 0)
        FROM order_item_table
        WHERE order_id = NEW.order_id
    )
    WHERE order_id = NEW.order_id;

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_total_order_amount() OWNER TO postgres;

--
-- Name: update_user(integer, character varying, character varying, character varying, text, character varying, integer, integer, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_user(p_user_id integer, p_name character varying, p_mobile_num character varying, p_email character varying, p_address text, p_pincode character varying, p_location integer, p_state integer, p_status integer) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE user_table
    SET name = p_name,
        mobile_num = p_mobile_num,
        email = p_email,
        address = p_address,
        pincode = p_pincode,
        location = p_location,
        state = p_state,
        status = p_status
    WHERE user_id = p_user_id;
END;
$$;


ALTER FUNCTION public.update_user(p_user_id integer, p_name character varying, p_mobile_num character varying, p_email character varying, p_address text, p_pincode character varying, p_location integer, p_state integer, p_status integer) OWNER TO postgres;

--
-- Name: update_user(integer, integer, character varying, date, character varying, character varying, text, character varying, integer, character varying, character varying, character varying, date, character varying, character varying, integer, character varying, text, text, text, integer); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.update_user(IN p_user_id integer, IN p_user_type_id integer, IN p_name character varying, IN p_dt_of_commence_business date, IN p_mobile_num character varying, IN p_email character varying, IN p_address text, IN p_pincode character varying, IN p_location integer, IN p_business_license_no character varying, IN p_validity character varying, IN p_gst_no character varying, IN p_expiry_dt date, IN p_business_name character varying, IN p_business_type character varying, IN p_mandi_id integer, IN p_mandi_type_id character varying, IN p_remarks text, IN p_col1 text, IN p_col2 text, IN p_state integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE public.user_table
    SET
        user_type_id = p_user_type_id,
        name = p_name,
        dt_of_commence_business = p_dt_of_commence_business,
        mobile_num = p_mobile_num,
        email = p_email,
        address = p_address,
        pincode = p_pincode,
        location = p_location,
        business_license_no = p_business_license_no,
        validity = p_validity,
        gst_no = p_gst_no,
        expiry_dt = p_expiry_dt,
        business_name = p_business_name,
        business_type = p_business_type,
        mandi_id = p_mandi_id,
        mandi_type_id = p_mandi_type_id,
        remarks = p_remarks,
        col1 = p_col1,
        col2 = p_col2,
        state = p_state
    WHERE user_id = p_user_id;
END;
$$;


ALTER PROCEDURE public.update_user(IN p_user_id integer, IN p_user_type_id integer, IN p_name character varying, IN p_dt_of_commence_business date, IN p_mobile_num character varying, IN p_email character varying, IN p_address text, IN p_pincode character varying, IN p_location integer, IN p_business_license_no character varying, IN p_validity character varying, IN p_gst_no character varying, IN p_expiry_dt date, IN p_business_name character varying, IN p_business_type character varying, IN p_mandi_id integer, IN p_mandi_type_id character varying, IN p_remarks text, IN p_col1 text, IN p_col2 text, IN p_state integer) OWNER TO postgres;

--
-- Name: user_table_audit_function(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.user_table_audit_function() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN
        INSERT INTO audit_log(user_id, action_type, changed_data)
        VALUES (NEW.user_id, 'INSERT', to_jsonb(NEW));
    ELSIF TG_OP = 'UPDATE' THEN
        INSERT INTO audit_log(user_id, action_type, changed_data)
        VALUES (NEW.user_id, 'UPDATE', to_jsonb(NEW));
    ELSIF TG_OP = 'DELETE' THEN
        INSERT INTO audit_log(user_id, action_type, changed_data)
        VALUES (OLD.user_id, 'DELETE', to_jsonb(OLD));
    END IF;
    RETURN NULL; -- Trigger function does not modify the table
END;
$$;


ALTER FUNCTION public.user_table_audit_function() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: business_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.business_table (
    bid bigint NOT NULL,
    b_typeid integer NOT NULL,
    b_location_id integer NOT NULL,
    b_state_id integer NOT NULL,
    b_mandi_id integer,
    b_address text NOT NULL,
    b_phone_num text,
    b_email character varying(255) NOT NULL,
    b_gst_num character varying(20) NOT NULL,
    b_pan_num character varying(10) NOT NULL,
    b_name character varying(255)
);


ALTER TABLE public.business_table OWNER TO postgres;

--
-- Name: business_type_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.business_type_table (
    b_typeid integer NOT NULL,
    b_typename character varying(255) NOT NULL,
    remarks text
);


ALTER TABLE public.business_type_table OWNER TO postgres;

--
-- Name: business_table_b_typeid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.business_table_b_typeid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.business_table_b_typeid_seq OWNER TO postgres;

--
-- Name: business_table_b_typeid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.business_table_b_typeid_seq OWNED BY public.business_type_table.b_typeid;


--
-- Name: business_table_bid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.business_table_bid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.business_table_bid_seq OWNER TO postgres;

--
-- Name: business_table_bid_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.business_table_bid_seq OWNED BY public.business_table.bid;


--
-- Name: cash_payment_list; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cash_payment_list (
    id integer NOT NULL,
    payment_type character varying(50) NOT NULL
);


ALTER TABLE public.cash_payment_list OWNER TO postgres;

--
-- Name: cash_payment_list_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cash_payment_list_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cash_payment_list_id_seq OWNER TO postgres;

--
-- Name: cash_payment_list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cash_payment_list_id_seq OWNED BY public.cash_payment_list.id;


--
-- Name: daily_price_update; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.daily_price_update (
    product_id integer,
    price character varying(50),
    status integer,
    unit_id integer,
    wholeseller_id integer
);


ALTER TABLE public.daily_price_update OWNER TO postgres;

--
-- Name: discount_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.discount_table (
    dis_id integer NOT NULL,
    unit_id integer,
    quantity_range character varying(50),
    discount_percentage character varying(10),
    amount numeric(10,2)
);


ALTER TABLE public.discount_table OWNER TO postgres;

--
-- Name: discount_table_dis_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.discount_table_dis_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.discount_table_dis_id_seq OWNER TO postgres;

--
-- Name: discount_table_dis_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.discount_table_dis_id_seq OWNED BY public.discount_table.dis_id;


--
-- Name: driver_violation_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.driver_violation_table (
    id integer NOT NULL,
    driver_id integer NOT NULL,
    violation_id integer NOT NULL,
    entry_date date DEFAULT CURRENT_DATE NOT NULL
);


ALTER TABLE public.driver_violation_table OWNER TO postgres;

--
-- Name: driver_violation_table_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.driver_violation_table_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.driver_violation_table_id_seq OWNER TO postgres;

--
-- Name: driver_violation_table_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.driver_violation_table_id_seq OWNED BY public.driver_violation_table.id;


--
-- Name: insurance_provider; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.insurance_provider (
    id integer NOT NULL,
    provider_name character varying(50)
);


ALTER TABLE public.insurance_provider OWNER TO postgres;

--
-- Name: insurance_provider_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.insurance_provider_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.insurance_provider_id_seq OWNER TO postgres;

--
-- Name: insurance_provider_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.insurance_provider_id_seq OWNED BY public.insurance_provider.id;


--
-- Name: insurance_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.insurance_table (
    insurance_id integer NOT NULL,
    vehicle_id integer,
    driver_id integer,
    produce_transit_to_retailer boolean DEFAULT false,
    committed_by_farm boolean DEFAULT false,
    committed_by_driver boolean DEFAULT false,
    committed_due_to_nature boolean DEFAULT false,
    product_safe boolean DEFAULT true,
    quantity_lost integer DEFAULT 0,
    quantity_left integer DEFAULT 0,
    insurance_policy_number character varying(50) NOT NULL,
    insurance_duration character varying(50),
    type_of_insurance integer,
    insurance_coverage_amount numeric(15,2) NOT NULL,
    insurance_provider integer NOT NULL,
    col1 text,
    col2 text,
    col3 text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.insurance_table OWNER TO postgres;

--
-- Name: insurance_table_insurance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.insurance_table_insurance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.insurance_table_insurance_id_seq OWNER TO postgres;

--
-- Name: insurance_table_insurance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.insurance_table_insurance_id_seq OWNED BY public.insurance_table.insurance_id;


--
-- Name: invoice_daily_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.invoice_daily_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    MAXVALUE 999999
    CACHE 1
    CYCLE;


ALTER SEQUENCE public.invoice_daily_seq OWNER TO postgres;

--
-- Name: invoice_details_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.invoice_details_table (
    id bigint NOT NULL,
    invoice_id bigint,
    order_item_id bigint,
    quantity numeric(10,2) NOT NULL,
    retailer_status integer DEFAULT 0
);


ALTER TABLE public.invoice_details_table OWNER TO postgres;

--
-- Name: invoice_details_table_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.invoice_details_table_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.invoice_details_table_id_seq OWNER TO postgres;

--
-- Name: invoice_details_table_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.invoice_details_table_id_seq OWNED BY public.invoice_details_table.id;


--
-- Name: invoice_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.invoice_table (
    id integer NOT NULL,
    invoice_number character varying(50),
    order_id bigint,
    total_amount numeric(10,2) NOT NULL,
    discount_amount numeric(10,2) DEFAULT 0,
    invoice_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    due_date date,
    pay_mode integer,
    pay_type integer,
    final_amount numeric(10,2) GENERATED ALWAYS AS ((total_amount - discount_amount)) STORED
);


ALTER TABLE public.invoice_table OWNER TO postgres;

--
-- Name: invoice_table_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.invoice_table_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.invoice_table_id_seq OWNER TO postgres;

--
-- Name: invoice_table_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.invoice_table_id_seq OWNED BY public.invoice_table.id;


--
-- Name: master_category_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.master_category_table (
    category_id integer NOT NULL,
    category_name character varying(255) NOT NULL,
    super_cat_id integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    col1 text,
    col2 text,
    remarks text,
    CONSTRAINT master_category_table_updated_at CHECK ((updated_at >= created_at))
);


ALTER TABLE public.master_category_table OWNER TO postgres;

--
-- Name: master_category_table_category_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.master_category_table_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.master_category_table_category_id_seq OWNER TO postgres;

--
-- Name: master_category_table_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.master_category_table_category_id_seq OWNED BY public.master_category_table.category_id;


--
-- Name: master_driver_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.master_driver_table (
    driver_id integer NOT NULL,
    driver_name character varying(255) NOT NULL,
    driver_age integer,
    driver_license character varying(50) NOT NULL,
    driver_number character varying(15) NOT NULL,
    driver_address character varying(255),
    driver_status character varying(50),
    date_of_joining date,
    experience_years integer,
    license_expiry_date date NOT NULL,
    emergency_contact character varying(15),
    assigned_route_id integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    col1 text,
    col2 text,
    d_o_b date,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT master_driver_table_driver_age_check CHECK ((driver_age >= 18)),
    CONSTRAINT master_driver_table_experience_years_check CHECK ((experience_years >= 0))
);


ALTER TABLE public.master_driver_table OWNER TO postgres;

--
-- Name: master_driver_table_driver_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.master_driver_table_driver_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.master_driver_table_driver_id_seq OWNER TO postgres;

--
-- Name: master_driver_table_driver_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.master_driver_table_driver_id_seq OWNED BY public.master_driver_table.driver_id;


--
-- Name: master_location; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.master_location (
    id integer NOT NULL,
    location character varying(50)
);


ALTER TABLE public.master_location OWNER TO postgres;

--
-- Name: master_location_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.master_location_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.master_location_id_seq OWNER TO postgres;

--
-- Name: master_location_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.master_location_id_seq OWNED BY public.master_location.id;


--
-- Name: master_mandi_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.master_mandi_table (
    mandi_id integer NOT NULL,
    mandi_location character varying(255) NOT NULL,
    mandi_number character varying(50) NOT NULL,
    mandi_incharge character varying(255) NOT NULL,
    mandi_incharge_num character varying(15) NOT NULL,
    mandi_pincode character varying(6) NOT NULL,
    mandi_address text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    remarks text,
    mandi_city integer,
    mandi_state integer,
    mandi_name character varying(255),
    CONSTRAINT master_mandi_table_updated_at CHECK ((updated_at >= created_at))
);


ALTER TABLE public.master_mandi_table OWNER TO postgres;

--
-- Name: master_mandi_table_mandi_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.master_mandi_table_mandi_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.master_mandi_table_mandi_id_seq OWNER TO postgres;

--
-- Name: master_mandi_table_mandi_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.master_mandi_table_mandi_id_seq OWNED BY public.master_mandi_table.mandi_id;


--
-- Name: master_product; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.master_product (
    product_id bigint NOT NULL,
    category_id integer,
    product_name character varying(100),
    status integer DEFAULT 0
);


ALTER TABLE public.master_product OWNER TO postgres;

--
-- Name: master_product_product_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.master_product_product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.master_product_product_id_seq OWNER TO postgres;

--
-- Name: master_product_product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.master_product_product_id_seq OWNED BY public.master_product.product_id;


--
-- Name: master_states; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.master_states (
    id integer NOT NULL,
    state character varying(50)
);


ALTER TABLE public.master_states OWNER TO postgres;

--
-- Name: master_states_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.master_states_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.master_states_id_seq OWNER TO postgres;

--
-- Name: master_states_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.master_states_id_seq OWNED BY public.master_states.id;


--
-- Name: master_vehicle_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.master_vehicle_table (
    vehicle_id integer NOT NULL,
    insurance_id integer,
    vehicle_name character varying(255) NOT NULL,
    vehicle_manufacture_year character varying(4),
    vehicle_warranty character varying(255),
    vehicle_make integer,
    vehicle_model integer,
    vehicle_registration_no character varying(50),
    vehicle_engine_type integer,
    vehicle_purchase_date date,
    vehicle_color character varying(50),
    col1 text,
    col2 text,
    col3 text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    vehicle_insurance_id integer,
    vehicle_make_name character varying(255),
    vehicle_model_name character varying(255),
    vehicle_engine_type_name character varying(255)
);


ALTER TABLE public.master_vehicle_table OWNER TO postgres;

--
-- Name: master_vechicle_table_vehicle_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.master_vechicle_table_vehicle_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.master_vechicle_table_vehicle_id_seq OWNER TO postgres;

--
-- Name: master_vechicle_table_vehicle_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.master_vechicle_table_vehicle_id_seq OWNED BY public.master_vehicle_table.vehicle_id;


--
-- Name: master_violation_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.master_violation_table (
    violation_name text,
    level_of_serious text,
    id integer NOT NULL
);


ALTER TABLE public.master_violation_table OWNER TO postgres;

--
-- Name: master_violation_table_new_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.master_violation_table_new_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.master_violation_table_new_id_seq OWNER TO postgres;

--
-- Name: master_violation_table_new_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.master_violation_table_new_id_seq OWNED BY public.master_violation_table.id;


--
-- Name: mode_of_payment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mode_of_payment (
    id integer NOT NULL,
    pay_mode integer,
    pay_type integer
);


ALTER TABLE public.mode_of_payment OWNER TO postgres;

--
-- Name: mode_of_payment_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mode_of_payment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.mode_of_payment_id_seq OWNER TO postgres;

--
-- Name: mode_of_payment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mode_of_payment_id_seq OWNED BY public.mode_of_payment.id;


--
-- Name: mode_of_payments_list; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mode_of_payments_list (
    id integer NOT NULL,
    payment_mode character varying(50) NOT NULL
);


ALTER TABLE public.mode_of_payments_list OWNER TO postgres;

--
-- Name: mode_of_payments_list_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mode_of_payments_list_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.mode_of_payments_list_id_seq OWNER TO postgres;

--
-- Name: mode_of_payments_list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mode_of_payments_list_id_seq OWNED BY public.mode_of_payments_list.id;


--
-- Name: order_history_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_history_table (
    order_id integer NOT NULL,
    date_of_order timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    order_status integer,
    expected_delivery_date timestamp without time zone,
    actual_delivery_date timestamp without time zone,
    retailer_id integer,
    wholeseller_id integer,
    location_id integer,
    state_id integer,
    pincode character varying(10),
    address text,
    delivery_completed_date timestamp without time zone,
    history_id integer NOT NULL
);


ALTER TABLE public.order_history_table OWNER TO postgres;

--
-- Name: order_history_table_history_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_history_table_history_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_history_table_history_id_seq OWNER TO postgres;

--
-- Name: order_history_table_history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_history_table_history_id_seq OWNED BY public.order_history_table.history_id;


--
-- Name: order_item_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_item_table (
    order_item_id bigint NOT NULL,
    order_id bigint,
    product_id bigint,
    quantity numeric(10,2),
    unit_id integer,
    amt_of_order_item numeric(10,3),
    order_item_status integer
);


ALTER TABLE public.order_item_table OWNER TO postgres;

--
-- Name: order_item_table_product_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_item_table_product_order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_item_table_product_order_id_seq OWNER TO postgres;

--
-- Name: order_item_table_product_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_item_table_product_order_id_seq OWNED BY public.order_item_table.order_item_id;


--
-- Name: order_status_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_status_table (
    order_status_id integer NOT NULL,
    order_status character varying(50) NOT NULL
);


ALTER TABLE public.order_status_table OWNER TO postgres;

--
-- Name: order_status_table_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_status_table_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_status_table_order_id_seq OWNER TO postgres;

--
-- Name: order_status_table_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_status_table_order_id_seq OWNED BY public.order_status_table.order_status_id;


--
-- Name: order_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_table (
    order_id bigint NOT NULL,
    date_of_order timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    order_status integer,
    expected_delivery_date timestamp without time zone,
    actual_delivery_date timestamp without time zone,
    retailer_id integer,
    wholeseller_id integer,
    location_id integer,
    state_id integer,
    pincode character varying(10),
    address text,
    total_order_amount numeric(10,2) DEFAULT 0
);


ALTER TABLE public.order_table OWNER TO postgres;

--
-- Name: order_table_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_table_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_table_order_id_seq OWNER TO postgres;

--
-- Name: order_table_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_table_order_id_seq OWNED BY public.order_table.order_id;


--
-- Name: product_price_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_price_table (
    id bigint NOT NULL,
    product_id bigint,
    price numeric(10,2),
    bid bigint,
    tax_percentage numeric(5,2) DEFAULT 0,
    CONSTRAINT product_price_table_tax_percentage_check CHECK (((tax_percentage >= (0)::numeric) AND (tax_percentage <= (100)::numeric)))
);


ALTER TABLE public.product_price_table OWNER TO postgres;

--
-- Name: product_price_table_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_price_table_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_price_table_id_seq OWNER TO postgres;

--
-- Name: product_price_table_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_price_table_id_seq OWNED BY public.product_price_table.id;


--
-- Name: retailer_rating_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.retailer_rating_table (
    id integer NOT NULL,
    order_id integer NOT NULL,
    creditworthiness_rating numeric(3,2),
    payment_disputes integer DEFAULT 0,
    order_frequency numeric(3,2),
    order_cancellation_rate numeric(3,2),
    urgent_order_requests numeric(3,2) DEFAULT 0,
    seasonal_dependency boolean DEFAULT false,
    communication_rating numeric(3,2),
    professionalism_rating numeric(3,2),
    negotiation_behavior_rating numeric(3,2),
    return_frequency numeric(3,2),
    return_validity_rating numeric(3,2),
    complaint_handling_rating numeric(3,2),
    overall_rating numeric(3,2) GENERATED ALWAYS AS (((((((creditworthiness_rating + order_frequency) + communication_rating) + professionalism_rating) + return_validity_rating) + complaint_handling_rating) / (6)::numeric)) STORED,
    comments text,
    rated_on timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT retailer_rating_table_communication_rating_check CHECK (((communication_rating >= (1)::numeric) AND (communication_rating <= (5)::numeric))),
    CONSTRAINT retailer_rating_table_complaint_handling_rating_check CHECK (((complaint_handling_rating >= (1)::numeric) AND (complaint_handling_rating <= (5)::numeric))),
    CONSTRAINT retailer_rating_table_creditworthiness_rating_check CHECK (((creditworthiness_rating >= (1)::numeric) AND (creditworthiness_rating <= (5)::numeric))),
    CONSTRAINT retailer_rating_table_negotiation_behavior_rating_check CHECK (((negotiation_behavior_rating >= (1)::numeric) AND (negotiation_behavior_rating <= (5)::numeric))),
    CONSTRAINT retailer_rating_table_order_cancellation_rate_check CHECK (((order_cancellation_rate >= (0)::numeric) AND (order_cancellation_rate <= (100)::numeric))),
    CONSTRAINT retailer_rating_table_order_frequency_check CHECK ((order_frequency >= (0)::numeric)),
    CONSTRAINT retailer_rating_table_professionalism_rating_check CHECK (((professionalism_rating >= (1)::numeric) AND (professionalism_rating <= (5)::numeric))),
    CONSTRAINT retailer_rating_table_return_frequency_check CHECK ((return_frequency >= (0)::numeric)),
    CONSTRAINT retailer_rating_table_return_validity_rating_check CHECK (((return_validity_rating >= (1)::numeric) AND (return_validity_rating <= (5)::numeric))),
    CONSTRAINT retailer_rating_table_urgent_order_requests_check CHECK ((urgent_order_requests >= (0)::numeric))
);


ALTER TABLE public.retailer_rating_table OWNER TO postgres;

--
-- Name: retailer_rating_table_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.retailer_rating_table_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.retailer_rating_table_id_seq OWNER TO postgres;

--
-- Name: retailer_rating_table_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.retailer_rating_table_id_seq OWNED BY public.retailer_rating_table.id;


--
-- Name: retailer_status; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.retailer_status (
    id integer NOT NULL,
    status_name character varying(20)
);


ALTER TABLE public.retailer_status OWNER TO postgres;

--
-- Name: stock_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.stock_table (
    stock_id integer NOT NULL,
    product_id integer NOT NULL,
    stock_in numeric(10,2) DEFAULT 0,
    stock_left numeric(10,2) DEFAULT 0,
    stock_sold numeric(10,2) DEFAULT 0,
    date_of_order timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    user_id integer,
    warehouse_id integer,
    status boolean DEFAULT true,
    updatedat timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    createdat timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    carry_fwd numeric(10,2) DEFAULT 0,
    source_id integer,
    col1 text,
    col2 text,
    remarks text
);


ALTER TABLE public.stock_table OWNER TO postgres;

--
-- Name: stock_table_stock_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.stock_table_stock_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.stock_table_stock_id_seq OWNER TO postgres;

--
-- Name: stock_table_stock_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.stock_table_stock_id_seq OWNED BY public.stock_table.stock_id;


--
-- Name: transporter_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transporter_table (
    transporter_id integer NOT NULL,
    transporter_name character varying(255) NOT NULL,
    registration_number character varying(100) NOT NULL,
    gst_num character varying(50),
    tax_identification_number character varying(50),
    address text,
    phone_number character varying(15) NOT NULL,
    email character varying(255),
    company_website character varying(255),
    regions_covered boolean DEFAULT false,
    insurance_details text,
    contact_person character varying(255),
    contact_phone character varying(15),
    contact_phone2 character varying(15),
    contact_email character varying(255),
    payment_terms text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.transporter_table OWNER TO postgres;

--
-- Name: transporter_table_transporter_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transporter_table_transporter_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.transporter_table_transporter_id_seq OWNER TO postgres;

--
-- Name: transporter_table_transporter_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transporter_table_transporter_id_seq OWNED BY public.transporter_table.transporter_id;


--
-- Name: type_of_insurance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.type_of_insurance (
    id integer NOT NULL,
    type_of_insurance text
);


ALTER TABLE public.type_of_insurance OWNER TO postgres;

--
-- Name: type_of_insurance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.type_of_insurance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.type_of_insurance_id_seq OWNER TO postgres;

--
-- Name: type_of_insurance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.type_of_insurance_id_seq OWNED BY public.type_of_insurance.id;


--
-- Name: units_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.units_table (
    id integer NOT NULL,
    unit_name character varying(50) NOT NULL
);


ALTER TABLE public.units_table OWNER TO postgres;

--
-- Name: units_table_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.units_table_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.units_table_id_seq OWNER TO postgres;

--
-- Name: units_table_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.units_table_id_seq OWNED BY public.units_table.id;


--
-- Name: user_bank_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_bank_details (
    id integer NOT NULL,
    user_id integer NOT NULL,
    card_number character(16),
    upi_id character varying(255),
    ifsc_code character(11),
    account_number character varying(20),
    account_holder_name character varying(255),
    bank_name character varying(255),
    branch_name character varying(255),
    date_of_creation timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    status boolean DEFAULT true
);


ALTER TABLE public.user_bank_details OWNER TO postgres;

--
-- Name: user_bank_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_bank_details_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_bank_details_id_seq OWNER TO postgres;

--
-- Name: user_bank_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_bank_details_id_seq OWNED BY public.user_bank_details.id;


--
-- Name: user_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_table (
    user_id integer NOT NULL,
    user_type_id integer,
    name character varying(255) NOT NULL,
    mobile_num character varying(15) NOT NULL,
    email character varying(255),
    address text,
    pincode character varying(10),
    location integer,
    state integer,
    status integer
);


ALTER TABLE public.user_table OWNER TO postgres;

--
-- Name: user_table_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_table_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_table_user_id_seq OWNER TO postgres;

--
-- Name: user_table_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_table_user_id_seq OWNED BY public.user_table.user_id;


--
-- Name: user_type_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_type_table (
    user_type_id integer NOT NULL,
    user_type character varying(50) NOT NULL,
    col1 text,
    col2 text,
    remarks text
);


ALTER TABLE public.user_type_table OWNER TO postgres;

--
-- Name: user_type_table_user_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_type_table_user_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_type_table_user_type_id_seq OWNER TO postgres;

--
-- Name: user_type_table_user_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_type_table_user_type_id_seq OWNED BY public.user_type_table.user_type_id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    user_id integer NOT NULL,
    user_type_id integer NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    machine_id character varying(255),
    doe timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    remarks text,
    col1 character varying(255),
    col2 character varying(255)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_user_id_seq OWNER TO postgres;

--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;


--
-- Name: vehicle_engine_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vehicle_engine_type (
    id integer NOT NULL,
    engine_type character varying(50)
);


ALTER TABLE public.vehicle_engine_type OWNER TO postgres;

--
-- Name: vehicle_engine_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vehicle_engine_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.vehicle_engine_type_id_seq OWNER TO postgres;

--
-- Name: vehicle_engine_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vehicle_engine_type_id_seq OWNED BY public.vehicle_engine_type.id;


--
-- Name: vehicle_insurance_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vehicle_insurance_table (
    vehicle_insurance_id integer NOT NULL,
    policy_number character varying(50) NOT NULL,
    insurance_provider character varying(100) NOT NULL,
    policy_type character varying(50) NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    premium_amount numeric(10,2) NOT NULL,
    insured_amount numeric(12,2) NOT NULL,
    status character varying(20) NOT NULL,
    policy_document bytea,
    chassis_number character varying(50),
    claims_made integer DEFAULT 0,
    last_claim_date date,
    date_of_creation timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    remarks text,
    col1 text,
    col2 text
);


ALTER TABLE public.vehicle_insurance_table OWNER TO postgres;

--
-- Name: vehicle_insurance_table_vehicle_insurance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vehicle_insurance_table_vehicle_insurance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.vehicle_insurance_table_vehicle_insurance_id_seq OWNER TO postgres;

--
-- Name: vehicle_insurance_table_vehicle_insurance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vehicle_insurance_table_vehicle_insurance_id_seq OWNED BY public.vehicle_insurance_table.vehicle_insurance_id;


--
-- Name: vehicle_make; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vehicle_make (
    id integer NOT NULL,
    make character varying(50)
);


ALTER TABLE public.vehicle_make OWNER TO postgres;

--
-- Name: vehicle_make_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vehicle_make_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.vehicle_make_id_seq OWNER TO postgres;

--
-- Name: vehicle_make_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vehicle_make_id_seq OWNED BY public.vehicle_make.id;


--
-- Name: vehicle_model; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vehicle_model (
    id integer NOT NULL,
    model character varying(50)
);


ALTER TABLE public.vehicle_model OWNER TO postgres;

--
-- Name: vehicle_model_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vehicle_model_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.vehicle_model_id_seq OWNER TO postgres;

--
-- Name: vehicle_model_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vehicle_model_id_seq OWNED BY public.vehicle_model.id;


--
-- Name: warehouse_list; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.warehouse_list (
    warehouse_id integer NOT NULL,
    user_id integer NOT NULL,
    warehouse_name character varying(255),
    warehouse_address text,
    warehouse_email character varying(255),
    warehouse_phone character varying(255),
    warehouse_pincode integer,
    warehouse_status boolean DEFAULT true,
    col1 character varying(255),
    col2 character varying(255),
    remarks text
);


ALTER TABLE public.warehouse_list OWNER TO postgres;

--
-- Name: warehouse_list_warehouse_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.warehouse_list_warehouse_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.warehouse_list_warehouse_id_seq OWNER TO postgres;

--
-- Name: warehouse_list_warehouse_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.warehouse_list_warehouse_id_seq OWNED BY public.warehouse_list.warehouse_id;


--
-- Name: wholeseller_rating_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wholeseller_rating_table (
    id integer NOT NULL,
    order_id integer NOT NULL,
    product_quality_rating numeric(3,2),
    product_authenticity boolean DEFAULT true,
    consistency_in_quality numeric(3,2),
    product_variety_rating numeric(3,2),
    stock_availability_rating numeric(3,2),
    new_product_introduction boolean DEFAULT false,
    order_processing_speed numeric(3,2),
    delivery_time_adherence numeric(3,2),
    urgent_order_handling numeric(3,2),
    order_fulfillment_accuracy numeric(3,2),
    packaging_quality_rating numeric(3,2),
    eco_friendly_packaging boolean DEFAULT false,
    storage_conditions_rating numeric(3,2),
    pricing_fairness_rating numeric(3,2),
    hidden_charges boolean DEFAULT false,
    discount_offers boolean DEFAULT false,
    flexible_payment_terms boolean DEFAULT false,
    responsiveness_rating numeric(3,2),
    complaint_resolution_rating numeric(3,2),
    return_refund_policy_rating numeric(3,2),
    replacement_handling_rating numeric(3,2),
    overall_rating numeric(3,2) GENERATED ALWAYS AS (((((((((product_quality_rating + consistency_in_quality) + order_processing_speed) + delivery_time_adherence) + packaging_quality_rating) + pricing_fairness_rating) + responsiveness_rating) + complaint_resolution_rating) / (8)::numeric)) STORED,
    comments text,
    rated_on timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.wholeseller_rating_table OWNER TO postgres;

--
-- Name: wholesaler_rating_table_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wholesaler_rating_table_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.wholesaler_rating_table_id_seq OWNER TO postgres;

--
-- Name: wholesaler_rating_table_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wholesaler_rating_table_id_seq OWNED BY public.wholeseller_rating_table.id;


--
-- Name: business_table bid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_table ALTER COLUMN bid SET DEFAULT nextval('public.business_table_bid_seq'::regclass);


--
-- Name: business_type_table b_typeid; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_type_table ALTER COLUMN b_typeid SET DEFAULT nextval('public.business_table_b_typeid_seq'::regclass);


--
-- Name: cash_payment_list id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cash_payment_list ALTER COLUMN id SET DEFAULT nextval('public.cash_payment_list_id_seq'::regclass);


--
-- Name: discount_table dis_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.discount_table ALTER COLUMN dis_id SET DEFAULT nextval('public.discount_table_dis_id_seq'::regclass);


--
-- Name: driver_violation_table id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.driver_violation_table ALTER COLUMN id SET DEFAULT nextval('public.driver_violation_table_id_seq'::regclass);


--
-- Name: insurance_provider id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_provider ALTER COLUMN id SET DEFAULT nextval('public.insurance_provider_id_seq'::regclass);


--
-- Name: insurance_table insurance_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_table ALTER COLUMN insurance_id SET DEFAULT nextval('public.insurance_table_insurance_id_seq'::regclass);


--
-- Name: invoice_details_table id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_details_table ALTER COLUMN id SET DEFAULT nextval('public.invoice_details_table_id_seq'::regclass);


--
-- Name: invoice_table id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_table ALTER COLUMN id SET DEFAULT nextval('public.invoice_table_id_seq'::regclass);


--
-- Name: master_category_table category_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_category_table ALTER COLUMN category_id SET DEFAULT nextval('public.master_category_table_category_id_seq'::regclass);


--
-- Name: master_driver_table driver_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_driver_table ALTER COLUMN driver_id SET DEFAULT nextval('public.master_driver_table_driver_id_seq'::regclass);


--
-- Name: master_location id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_location ALTER COLUMN id SET DEFAULT nextval('public.master_location_id_seq'::regclass);


--
-- Name: master_mandi_table mandi_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_mandi_table ALTER COLUMN mandi_id SET DEFAULT nextval('public.master_mandi_table_mandi_id_seq'::regclass);


--
-- Name: master_product product_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_product ALTER COLUMN product_id SET DEFAULT nextval('public.master_product_product_id_seq'::regclass);


--
-- Name: master_states id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_states ALTER COLUMN id SET DEFAULT nextval('public.master_states_id_seq'::regclass);


--
-- Name: master_vehicle_table vehicle_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_vehicle_table ALTER COLUMN vehicle_id SET DEFAULT nextval('public.master_vechicle_table_vehicle_id_seq'::regclass);


--
-- Name: master_violation_table id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_violation_table ALTER COLUMN id SET DEFAULT nextval('public.master_violation_table_new_id_seq'::regclass);


--
-- Name: mode_of_payment id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mode_of_payment ALTER COLUMN id SET DEFAULT nextval('public.mode_of_payment_id_seq'::regclass);


--
-- Name: mode_of_payments_list id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mode_of_payments_list ALTER COLUMN id SET DEFAULT nextval('public.mode_of_payments_list_id_seq'::regclass);


--
-- Name: order_history_table history_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_history_table ALTER COLUMN history_id SET DEFAULT nextval('public.order_history_table_history_id_seq'::regclass);


--
-- Name: order_item_table order_item_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_item_table ALTER COLUMN order_item_id SET DEFAULT nextval('public.order_item_table_product_order_id_seq'::regclass);


--
-- Name: order_status_table order_status_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_status_table ALTER COLUMN order_status_id SET DEFAULT nextval('public.order_status_table_order_id_seq'::regclass);


--
-- Name: order_table order_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_table ALTER COLUMN order_id SET DEFAULT nextval('public.order_table_order_id_seq'::regclass);


--
-- Name: product_price_table id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_price_table ALTER COLUMN id SET DEFAULT nextval('public.product_price_table_id_seq'::regclass);


--
-- Name: retailer_rating_table id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.retailer_rating_table ALTER COLUMN id SET DEFAULT nextval('public.retailer_rating_table_id_seq'::regclass);


--
-- Name: stock_table stock_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stock_table ALTER COLUMN stock_id SET DEFAULT nextval('public.stock_table_stock_id_seq'::regclass);


--
-- Name: transporter_table transporter_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transporter_table ALTER COLUMN transporter_id SET DEFAULT nextval('public.transporter_table_transporter_id_seq'::regclass);


--
-- Name: type_of_insurance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.type_of_insurance ALTER COLUMN id SET DEFAULT nextval('public.type_of_insurance_id_seq'::regclass);


--
-- Name: units_table id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.units_table ALTER COLUMN id SET DEFAULT nextval('public.units_table_id_seq'::regclass);


--
-- Name: user_bank_details id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_bank_details ALTER COLUMN id SET DEFAULT nextval('public.user_bank_details_id_seq'::regclass);


--
-- Name: user_table user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_table ALTER COLUMN user_id SET DEFAULT nextval('public.user_table_user_id_seq'::regclass);


--
-- Name: user_type_table user_type_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_type_table ALTER COLUMN user_type_id SET DEFAULT nextval('public.user_type_table_user_type_id_seq'::regclass);


--
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Name: vehicle_engine_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicle_engine_type ALTER COLUMN id SET DEFAULT nextval('public.vehicle_engine_type_id_seq'::regclass);


--
-- Name: vehicle_insurance_table vehicle_insurance_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicle_insurance_table ALTER COLUMN vehicle_insurance_id SET DEFAULT nextval('public.vehicle_insurance_table_vehicle_insurance_id_seq'::regclass);


--
-- Name: vehicle_make id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicle_make ALTER COLUMN id SET DEFAULT nextval('public.vehicle_make_id_seq'::regclass);


--
-- Name: vehicle_model id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicle_model ALTER COLUMN id SET DEFAULT nextval('public.vehicle_model_id_seq'::regclass);


--
-- Name: warehouse_list warehouse_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.warehouse_list ALTER COLUMN warehouse_id SET DEFAULT nextval('public.warehouse_list_warehouse_id_seq'::regclass);


--
-- Name: wholeseller_rating_table id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wholeseller_rating_table ALTER COLUMN id SET DEFAULT nextval('public.wholesaler_rating_table_id_seq'::regclass);


--
-- Data for Name: business_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.business_table VALUES (5, 2, 2, 2, 2, '456 Market Road', '8765432109', 'business2@example.com', 'GST987654321', 'PAN7654321', 'FreshMarket');
INSERT INTO public.business_table VALUES (6, 3, 1, 3, 2, '789 Commerce Lane', '7654321098', 'business3@example.com', 'GST567890123', 'PAN3456789', 'TradeHub');
INSERT INTO public.business_table VALUES (12, 1, 1, 1, 2, '123 Main Street', '9876543210', 'business15@example.com', 'GST0123456789', 'PAN1273457', 'SuperMart');


--
-- Data for Name: business_type_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.business_type_table VALUES (1, 'Retailer', 'Sells directly to consumers');
INSERT INTO public.business_type_table VALUES (2, 'Wholesaler', 'Sells in bulk to retailers');
INSERT INTO public.business_type_table VALUES (3, 'Transporter', 'Handles logistics and transportation');


--
-- Data for Name: cash_payment_list; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.cash_payment_list VALUES (1, 'Online Banking');
INSERT INTO public.cash_payment_list VALUES (2, 'UPI Payment');
INSERT INTO public.cash_payment_list VALUES (3, 'Mobile Wallets');
INSERT INTO public.cash_payment_list VALUES (4, 'Credit/Debit Cards');
INSERT INTO public.cash_payment_list VALUES (5, 'Credit/Pay after delivery');


--
-- Data for Name: daily_price_update; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.daily_price_update VALUES (1, '100.50', 1, 1, 5);
INSERT INTO public.daily_price_update VALUES (2, '250.00', 1, 2, 6);
INSERT INTO public.daily_price_update VALUES (3, '75.75', 0, 3, 12);


--
-- Data for Name: discount_table; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: driver_violation_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.driver_violation_table VALUES (1, 2, 2, '2024-02-10');
INSERT INTO public.driver_violation_table VALUES (2, 2, 1, '2024-02-15');
INSERT INTO public.driver_violation_table VALUES (3, 2, 2, '2024-02-18');


--
-- Data for Name: insurance_provider; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: insurance_table; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: invoice_details_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.invoice_details_table VALUES (1, 2, 6, 5.00, 0);


--
-- Data for Name: invoice_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.invoice_table VALUES (2, 'INV-20250224-0001', 1, 590.00, 500.00, '2025-02-24 18:25:44.460827', '2025-03-01', 2, 4, DEFAULT);
INSERT INTO public.invoice_table VALUES (3, 'INV-20250225-00003', 3, 1500.00, 100.00, '2025-02-25 14:14:21.29295', '2025-03-02', 1, 2, DEFAULT);


--
-- Data for Name: master_category_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.master_category_table VALUES (5, 'Roots', 2, '2025-01-01 14:23:15.363233', '2025-01-01 14:23:15.363233', 'Extra data for clothing', 'Additional info', 'Pending verification');
INSERT INTO public.master_category_table VALUES (6, 'Vegetables', 3, '2025-01-01 14:23:15.363233', '2025-01-01 14:23:15.363233', 'Extra data for furniture', 'Additional info', 'Verified');
INSERT INTO public.master_category_table VALUES (8, 'Fruits', 4, '2025-02-10 19:43:12.712312', '2025-02-10 19:43:12.712312', 'Extra info 1', 'Extra info 2', 'extra info 3');
INSERT INTO public.master_category_table VALUES (4, 'Leafy', 0, '2025-01-01 14:23:15.363233', '2025-01-01 14:23:15.363233', 'Extra data for electronics', 'Additional info', 'No remarks');


--
-- Data for Name: master_driver_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.master_driver_table VALUES (2, 'John Updated', 35, 'DL12345', '9876543210', '123 Main St', 'Inactive', '2023-01-01', 5, '2026-12-31', '9876543211', 2, '2025-02-07 11:31:28.041893', 'Additional Info 1', 'Additional Info 2', '1989-05-20', '2025-02-15 16:37:29.629968');


--
-- Data for Name: master_location; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.master_location VALUES (1, 'Chennai');
INSERT INTO public.master_location VALUES (2, 'coimbatore');


--
-- Data for Name: master_mandi_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.master_mandi_table VALUES (2, 'Location 1', 'M123', 'John Doe', '9876543210', '123456', 'Address 1', '2025-02-07 12:01:02.04624', '2025-02-07 12:01:02.04624', 'Remarks', 1, 1, NULL);


--
-- Data for Name: master_product; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.master_product VALUES (1, 4, 'New Product', 1);
INSERT INTO public.master_product VALUES (2, 5, 'product 2', 1);
INSERT INTO public.master_product VALUES (3, 6, 'product 2', 1);


--
-- Data for Name: master_states; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.master_states VALUES (1, 'tamilnadu');
INSERT INTO public.master_states VALUES (2, 'kerla');
INSERT INTO public.master_states VALUES (3, 'West bengal');


--
-- Data for Name: master_vehicle_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.master_vehicle_table VALUES (5, 102, 'Honda City', '2021', '3 years', 2, 3, 'TN10XY5678', 1, '2021-08-22', 'Black', 'Extra2', 'Info2', 'Details2', '2025-02-07 11:30:10.747771', '2025-02-07 11:30:10.747771', 202, NULL, NULL, NULL);
INSERT INTO public.master_vehicle_table VALUES (6, 103, 'Ford Endeavour', '2023', '7 years', 3, 1, 'TN22PQ9876', 2, '2023-02-10', 'Blue', 'Extra3', 'Info3', 'Details3', '2025-02-07 11:30:10.747771', '2025-02-07 11:30:10.747771', 203, NULL, NULL, NULL);


--
-- Data for Name: master_violation_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.master_violation_table VALUES ('Hit and run', 'High', 1);
INSERT INTO public.master_violation_table VALUES ('drunk and drive', 'medium', 2);


--
-- Data for Name: mode_of_payment; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.mode_of_payment VALUES (1, 2, 3);
INSERT INTO public.mode_of_payment VALUES (2, 2, 3);
INSERT INTO public.mode_of_payment VALUES (3, 1, 5);


--
-- Data for Name: mode_of_payments_list; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.mode_of_payments_list VALUES (1, 'credit_payment');
INSERT INTO public.mode_of_payments_list VALUES (2, 'cash_payment');


--
-- Data for Name: order_history_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.order_history_table VALUES (1, '2025-02-22 09:55:59.551139', 4, '2025-02-27 09:55:59.551139', '2025-02-22 09:58:54.219252', 5, 6, 1, 2, '123456', '123 Street, City', NULL, 1);
INSERT INTO public.order_history_table VALUES (1, '2025-02-22 09:55:59.551139', 6, '2025-02-27 09:55:59.551139', '2025-02-22 10:01:52.229039', 5, 6, 1, 2, '123456', '123 Street, City', '2025-02-22 10:01:52.229039', 2);
INSERT INTO public.order_history_table VALUES (1, '2025-02-22 09:55:59.551139', 6, '2025-02-27 09:55:59.551139', '2025-02-22 10:01:52.229039', 5, 6, 1, 2, '123456', '123 Street, City', '2025-02-24 11:29:28.140697', 3);
INSERT INTO public.order_history_table VALUES (1, '2025-02-22 09:55:59.551139', 6, '2025-02-27 09:55:59.551139', '2025-02-22 10:01:52.229039', 5, 6, 1, 2, '123456', '123 Street, City', '2025-02-24 11:59:18.925182', 4);
INSERT INTO public.order_history_table VALUES (1, '2025-02-22 09:55:59.551139', 6, '2025-02-27 09:55:59.551139', '2025-02-22 10:01:52.229039', 5, 6, 1, 2, '123456', '123 Street, City', '2025-02-24 12:09:45.416871', 5);
INSERT INTO public.order_history_table VALUES (1, '2025-02-22 09:55:59.551139', 6, '2025-02-27 09:55:59.551139', '2025-02-22 10:01:52.229039', 5, 6, 1, 2, '123456', '123 Street, City', '2025-02-24 21:12:04.580976', 6);
INSERT INTO public.order_history_table VALUES (3, '2025-02-23 10:30:00', 2, '2025-02-26 10:30:00', NULL, 12, 6, 1, 2, '987654', '789 Road, City', NULL, 8);
INSERT INTO public.order_history_table VALUES (3, '2025-02-23 10:30:00', 2, '2025-02-26 10:30:00', NULL, 12, 6, 1, 2, '987654', '789 Road, City', NULL, 9);
INSERT INTO public.order_history_table VALUES (4, '2025-02-24 11:15:00', 5, '2025-02-28 11:15:00', '2025-02-24 12:00:00', 5, 6, 2, 1, '876543', '1011 Lane, City', NULL, 10);
INSERT INTO public.order_history_table VALUES (4, '2025-02-24 11:15:00', 5, '2025-02-28 11:15:00', '2025-02-24 12:00:00', 5, 6, 2, 1, '876543', '1011 Lane, City', NULL, 11);


--
-- Data for Name: order_item_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.order_item_table VALUES (6, 1, 2, 5.00, 1, 590.000, NULL);
INSERT INTO public.order_item_table VALUES (4, 1, 2, 5.00, 1, 590.000, 2);
INSERT INTO public.order_item_table VALUES (7, 3, 3, 2.00, 1, 750.000, 1);
INSERT INTO public.order_item_table VALUES (8, 3, 1, 1.00, 2, 750.000, 1);
INSERT INTO public.order_item_table VALUES (9, 4, 2, 4.00, 1, 472.000, 4);
INSERT INTO public.order_item_table VALUES (10, 4, 1, 3.00, 1, 1000.000, 4);


--
-- Data for Name: order_status_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.order_status_table VALUES (1, 'Processing');
INSERT INTO public.order_status_table VALUES (2, 'Confirmed');
INSERT INTO public.order_status_table VALUES (3, 'Payment');
INSERT INTO public.order_status_table VALUES (4, 'Out for Delivery');
INSERT INTO public.order_status_table VALUES (5, 'Successful');
INSERT INTO public.order_status_table VALUES (6, 'Cancellation');
INSERT INTO public.order_status_table VALUES (7, 'Returned');


--
-- Data for Name: order_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.order_table VALUES (2, '2025-02-22 09:55:59.551139', 4, '2025-02-25 09:55:59.551139', NULL, 5, 6, 2, 1, '654321', '456 Avenue, City', 0.00);
INSERT INTO public.order_table VALUES (1, '2025-02-22 09:55:59.551139', 6, '2025-02-27 09:55:59.551139', '2025-02-22 10:01:52.229039', 5, 6, 1, 2, '123456', '123 Street, City', 1180.00);
INSERT INTO public.order_table VALUES (3, '2025-02-23 10:30:00', 2, '2025-02-26 10:30:00', NULL, 12, 6, 1, 2, '987654', '789 Road, City', 1500.00);
INSERT INTO public.order_table VALUES (4, '2025-02-24 11:15:00', 5, '2025-02-28 11:15:00', '2025-02-24 12:00:00', 5, 6, 2, 1, '876543', '1011 Lane, City', 1472.00);


--
-- Data for Name: product_price_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.product_price_table VALUES (2, 2, 100.00, 5, 18.00);


--
-- Data for Name: retailer_rating_table; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: retailer_status; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.retailer_status VALUES (0, 'initiated');
INSERT INTO public.retailer_status VALUES (1, 'Accepting');
INSERT INTO public.retailer_status VALUES (2, 'Rejecting');
INSERT INTO public.retailer_status VALUES (3, 'negotiation');


--
-- Data for Name: stock_table; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: transporter_table; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: type_of_insurance; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: units_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.units_table VALUES (1, 'kg');
INSERT INTO public.units_table VALUES (2, 'liter');
INSERT INTO public.units_table VALUES (3, 'tonnes');
INSERT INTO public.units_table VALUES (4, 'pieces');
INSERT INTO public.units_table VALUES (5, 'grams');


--
-- Data for Name: user_bank_details; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: user_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_table VALUES (12, 1, 'John Doe', '9876543210', 'johndoe@example.com', '123 Main St, City', '123456', 1, 1, 1);
INSERT INTO public.user_table VALUES (13, 2, 'Jane Smith', '9123456789', 'janesmith@example.com', '456 Market St, Town', '654321', 2, 2, 1);
INSERT INTO public.user_table VALUES (14, 1, 'John Doe', '9876543210', 'john.doe@example.com', '123 Street', '123456', 1, 2, 1);


--
-- Data for Name: user_type_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.user_type_table VALUES (1, 'Admin', 'Additional info for Admin', 'Extra data for Admin', 'Admin type for managing users and data');
INSERT INTO public.user_type_table VALUES (2, 'Retailer', 'Additional info for Retailer', 'Extra data for Retailer', 'Retail type for business owners and transactions');
INSERT INTO public.user_type_table VALUES (3, 'Distributor', 'Additional info for Distributor', 'Extra data for Distributor', 'Distributor type for managing supply chain operations');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: vehicle_engine_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.vehicle_engine_type VALUES (1, 'disel');
INSERT INTO public.vehicle_engine_type VALUES (2, 'petrol');


--
-- Data for Name: vehicle_insurance_table; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: vehicle_make; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.vehicle_make VALUES (1, 'Toyota');
INSERT INTO public.vehicle_make VALUES (2, 'Honda');
INSERT INTO public.vehicle_make VALUES (3, 'Ford');


--
-- Data for Name: vehicle_model; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.vehicle_model VALUES (1, 'Corolla');
INSERT INTO public.vehicle_model VALUES (2, 'Civic');
INSERT INTO public.vehicle_model VALUES (3, 'Mustang');


--
-- Data for Name: warehouse_list; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: wholeseller_rating_table; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Name: business_table_b_typeid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.business_table_b_typeid_seq', 3, true);


--
-- Name: business_table_bid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.business_table_bid_seq', 12, true);


--
-- Name: cash_payment_list_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cash_payment_list_id_seq', 5, true);


--
-- Name: discount_table_dis_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.discount_table_dis_id_seq', 1, false);


--
-- Name: driver_violation_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.driver_violation_table_id_seq', 3, true);


--
-- Name: insurance_provider_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.insurance_provider_id_seq', 1, false);


--
-- Name: insurance_table_insurance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.insurance_table_insurance_id_seq', 1, false);


--
-- Name: invoice_daily_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.invoice_daily_seq', 1, false);


--
-- Name: invoice_details_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.invoice_details_table_id_seq', 1, true);


--
-- Name: invoice_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.invoice_table_id_seq', 3, true);


--
-- Name: master_category_table_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.master_category_table_category_id_seq', 8, true);


--
-- Name: master_driver_table_driver_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.master_driver_table_driver_id_seq', 2, true);


--
-- Name: master_location_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.master_location_id_seq', 5, true);


--
-- Name: master_mandi_table_mandi_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.master_mandi_table_mandi_id_seq', 2, true);


--
-- Name: master_product_product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.master_product_product_id_seq', 1, true);


--
-- Name: master_states_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.master_states_id_seq', 3, true);


--
-- Name: master_vechicle_table_vehicle_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.master_vechicle_table_vehicle_id_seq', 6, true);


--
-- Name: master_violation_table_new_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.master_violation_table_new_id_seq', 2, true);


--
-- Name: mode_of_payment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mode_of_payment_id_seq', 3, true);


--
-- Name: mode_of_payments_list_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mode_of_payments_list_id_seq', 2, true);


--
-- Name: order_history_table_history_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_history_table_history_id_seq', 11, true);


--
-- Name: order_item_table_product_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_item_table_product_order_id_seq', 6, true);


--
-- Name: order_status_table_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_status_table_order_id_seq', 7, true);


--
-- Name: order_table_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_table_order_id_seq', 2, true);


--
-- Name: product_price_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_price_table_id_seq', 4, true);


--
-- Name: retailer_rating_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.retailer_rating_table_id_seq', 1, false);


--
-- Name: stock_table_stock_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.stock_table_stock_id_seq', 1, false);


--
-- Name: transporter_table_transporter_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transporter_table_transporter_id_seq', 1, false);


--
-- Name: type_of_insurance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.type_of_insurance_id_seq', 1, false);


--
-- Name: units_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.units_table_id_seq', 5, true);


--
-- Name: user_bank_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_bank_details_id_seq', 34, true);


--
-- Name: user_table_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_table_user_id_seq', 14, true);


--
-- Name: user_type_table_user_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_type_table_user_type_id_seq', 3, true);


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_user_id_seq', 1, false);


--
-- Name: vehicle_engine_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vehicle_engine_type_id_seq', 2, true);


--
-- Name: vehicle_insurance_table_vehicle_insurance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vehicle_insurance_table_vehicle_insurance_id_seq', 1, false);


--
-- Name: vehicle_make_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vehicle_make_id_seq', 3, true);


--
-- Name: vehicle_model_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vehicle_model_id_seq', 3, true);


--
-- Name: warehouse_list_warehouse_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.warehouse_list_warehouse_id_seq', 1, false);


--
-- Name: wholesaler_rating_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wholesaler_rating_table_id_seq', 1, false);


--
-- Name: business_table business_table_b_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_table
    ADD CONSTRAINT business_table_b_email_key UNIQUE (b_email);


--
-- Name: business_table business_table_b_gst_num_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_table
    ADD CONSTRAINT business_table_b_gst_num_key UNIQUE (b_gst_num);


--
-- Name: business_table business_table_b_pan_num_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_table
    ADD CONSTRAINT business_table_b_pan_num_key UNIQUE (b_pan_num);


--
-- Name: business_type_table business_table_b_typename_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_type_table
    ADD CONSTRAINT business_table_b_typename_key UNIQUE (b_typename);


--
-- Name: business_type_table business_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_type_table
    ADD CONSTRAINT business_table_pkey PRIMARY KEY (b_typeid);


--
-- Name: business_table business_table_pkey1; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_table
    ADD CONSTRAINT business_table_pkey1 PRIMARY KEY (bid);


--
-- Name: cash_payment_list cash_payment_list_payment_type_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cash_payment_list
    ADD CONSTRAINT cash_payment_list_payment_type_key UNIQUE (payment_type);


--
-- Name: cash_payment_list cash_payment_list_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cash_payment_list
    ADD CONSTRAINT cash_payment_list_pkey PRIMARY KEY (id);


--
-- Name: discount_table discount_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.discount_table
    ADD CONSTRAINT discount_table_pkey PRIMARY KEY (dis_id);


--
-- Name: driver_violation_table driver_violation_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.driver_violation_table
    ADD CONSTRAINT driver_violation_table_pkey PRIMARY KEY (id);


--
-- Name: insurance_provider insurance_provider_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_provider
    ADD CONSTRAINT insurance_provider_pkey PRIMARY KEY (id);


--
-- Name: insurance_table insurance_table_insurance_policy_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_table
    ADD CONSTRAINT insurance_table_insurance_policy_number_key UNIQUE (insurance_policy_number);


--
-- Name: insurance_table insurance_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_table
    ADD CONSTRAINT insurance_table_pkey PRIMARY KEY (insurance_id);


--
-- Name: invoice_details_table invoice_details_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_details_table
    ADD CONSTRAINT invoice_details_table_pkey PRIMARY KEY (id);


--
-- Name: invoice_table invoice_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_table
    ADD CONSTRAINT invoice_table_pkey PRIMARY KEY (id);


--
-- Name: master_category_table master_category_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_category_table
    ADD CONSTRAINT master_category_table_pkey PRIMARY KEY (category_id);


--
-- Name: master_driver_table master_driver_table_driver_license_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_driver_table
    ADD CONSTRAINT master_driver_table_driver_license_key UNIQUE (driver_license);


--
-- Name: master_driver_table master_driver_table_driver_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_driver_table
    ADD CONSTRAINT master_driver_table_driver_number_key UNIQUE (driver_number);


--
-- Name: master_driver_table master_driver_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_driver_table
    ADD CONSTRAINT master_driver_table_pkey PRIMARY KEY (driver_id);


--
-- Name: master_location master_location_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_location
    ADD CONSTRAINT master_location_pkey PRIMARY KEY (id);


--
-- Name: master_mandi_table master_mandi_table_mandi_incharge_num_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_mandi_table
    ADD CONSTRAINT master_mandi_table_mandi_incharge_num_key UNIQUE (mandi_incharge_num);


--
-- Name: master_mandi_table master_mandi_table_mandi_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_mandi_table
    ADD CONSTRAINT master_mandi_table_mandi_number_key UNIQUE (mandi_number);


--
-- Name: master_mandi_table master_mandi_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_mandi_table
    ADD CONSTRAINT master_mandi_table_pkey PRIMARY KEY (mandi_id);


--
-- Name: master_product master_product_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_product
    ADD CONSTRAINT master_product_pkey PRIMARY KEY (product_id);


--
-- Name: master_states master_states_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_states
    ADD CONSTRAINT master_states_pkey PRIMARY KEY (id);


--
-- Name: master_vehicle_table master_vechicle_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_vehicle_table
    ADD CONSTRAINT master_vechicle_table_pkey PRIMARY KEY (vehicle_id);


--
-- Name: master_vehicle_table master_vechicle_table_vehicle_registration_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_vehicle_table
    ADD CONSTRAINT master_vechicle_table_vehicle_registration_no_key UNIQUE (vehicle_registration_no);


--
-- Name: master_violation_table master_violation_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_violation_table
    ADD CONSTRAINT master_violation_table_pkey PRIMARY KEY (id);


--
-- Name: mode_of_payment mode_of_payment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mode_of_payment
    ADD CONSTRAINT mode_of_payment_pkey PRIMARY KEY (id);


--
-- Name: mode_of_payments_list mode_of_payments_list_payment_mode_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mode_of_payments_list
    ADD CONSTRAINT mode_of_payments_list_payment_mode_key UNIQUE (payment_mode);


--
-- Name: mode_of_payments_list mode_of_payments_list_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mode_of_payments_list
    ADD CONSTRAINT mode_of_payments_list_pkey PRIMARY KEY (id);


--
-- Name: order_history_table order_history_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_history_table
    ADD CONSTRAINT order_history_table_pkey PRIMARY KEY (history_id);


--
-- Name: order_item_table order_item_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_item_table
    ADD CONSTRAINT order_item_table_pkey PRIMARY KEY (order_item_id);


--
-- Name: order_status_table order_status_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_status_table
    ADD CONSTRAINT order_status_table_pkey PRIMARY KEY (order_status_id);


--
-- Name: order_table order_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_table
    ADD CONSTRAINT order_table_pkey PRIMARY KEY (order_id);


--
-- Name: product_price_table product_price_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_price_table
    ADD CONSTRAINT product_price_table_pkey PRIMARY KEY (id);


--
-- Name: retailer_rating_table retailer_rating_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.retailer_rating_table
    ADD CONSTRAINT retailer_rating_table_pkey PRIMARY KEY (id);


--
-- Name: retailer_status retailer_status_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.retailer_status
    ADD CONSTRAINT retailer_status_pkey PRIMARY KEY (id);


--
-- Name: stock_table stock_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stock_table
    ADD CONSTRAINT stock_table_pkey PRIMARY KEY (stock_id);


--
-- Name: transporter_table transporter_table_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transporter_table
    ADD CONSTRAINT transporter_table_email_key UNIQUE (email);


--
-- Name: transporter_table transporter_table_gst_num_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transporter_table
    ADD CONSTRAINT transporter_table_gst_num_key UNIQUE (gst_num);


--
-- Name: transporter_table transporter_table_phone_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transporter_table
    ADD CONSTRAINT transporter_table_phone_number_key UNIQUE (phone_number);


--
-- Name: transporter_table transporter_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transporter_table
    ADD CONSTRAINT transporter_table_pkey PRIMARY KEY (transporter_id);


--
-- Name: transporter_table transporter_table_registration_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transporter_table
    ADD CONSTRAINT transporter_table_registration_number_key UNIQUE (registration_number);


--
-- Name: type_of_insurance type_of_insurance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.type_of_insurance
    ADD CONSTRAINT type_of_insurance_pkey PRIMARY KEY (id);


--
-- Name: invoice_table unique_invoice_number; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_table
    ADD CONSTRAINT unique_invoice_number UNIQUE (invoice_number);


--
-- Name: order_status_table unique_order_status_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_status_table
    ADD CONSTRAINT unique_order_status_id UNIQUE (order_status_id);


--
-- Name: units_table units_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.units_table
    ADD CONSTRAINT units_table_pkey PRIMARY KEY (id);


--
-- Name: units_table units_table_unit_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.units_table
    ADD CONSTRAINT units_table_unit_name_key UNIQUE (unit_name);


--
-- Name: user_bank_details user_bank_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_bank_details
    ADD CONSTRAINT user_bank_details_pkey PRIMARY KEY (id);


--
-- Name: user_table user_table_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_table
    ADD CONSTRAINT user_table_email_key UNIQUE (email);


--
-- Name: user_table user_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_table
    ADD CONSTRAINT user_table_pkey PRIMARY KEY (user_id);


--
-- Name: user_type_table user_type_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_type_table
    ADD CONSTRAINT user_type_table_pkey PRIMARY KEY (user_type_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: vehicle_engine_type vehicle_engine_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicle_engine_type
    ADD CONSTRAINT vehicle_engine_type_pkey PRIMARY KEY (id);


--
-- Name: vehicle_insurance_table vehicle_insurance_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicle_insurance_table
    ADD CONSTRAINT vehicle_insurance_table_pkey PRIMARY KEY (vehicle_insurance_id);


--
-- Name: vehicle_make vehicle_make_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicle_make
    ADD CONSTRAINT vehicle_make_pkey PRIMARY KEY (id);


--
-- Name: vehicle_model vehicle_model_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicle_model
    ADD CONSTRAINT vehicle_model_pkey PRIMARY KEY (id);


--
-- Name: warehouse_list warehouse_list_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.warehouse_list
    ADD CONSTRAINT warehouse_list_pkey PRIMARY KEY (warehouse_id);


--
-- Name: wholeseller_rating_table wholesaler_rating_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wholeseller_rating_table
    ADD CONSTRAINT wholesaler_rating_table_pkey PRIMARY KEY (id);


--
-- Name: invoice_table invoice_number_trigger; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER invoice_number_trigger BEFORE INSERT ON public.invoice_table FOR EACH ROW EXECUTE FUNCTION public.generate_invoice_number();


--
-- Name: mode_of_payment trg_default_pay_type; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trg_default_pay_type BEFORE INSERT ON public.mode_of_payment FOR EACH ROW EXECUTE FUNCTION public.set_default_pay_type();


--
-- Name: invoice_table trigger_generate_invoice_number; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trigger_generate_invoice_number BEFORE INSERT ON public.invoice_table FOR EACH ROW EXECUTE FUNCTION public.generate_invoice_number();


--
-- Name: order_item_table trigger_update_amt; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trigger_update_amt BEFORE INSERT OR UPDATE ON public.order_item_table FOR EACH ROW EXECUTE FUNCTION public.update_amt_of_order_item();


--
-- Name: order_table trigger_update_order_history; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trigger_update_order_history BEFORE UPDATE ON public.order_table FOR EACH ROW EXECUTE FUNCTION public.update_order_history();


--
-- Name: order_item_table trigger_update_total_order_amount; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trigger_update_total_order_amount AFTER INSERT OR DELETE OR UPDATE ON public.order_item_table FOR EACH ROW EXECUTE FUNCTION public.update_total_order_amount();


--
-- Name: driver_violation_table driver_violation_table_driver_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.driver_violation_table
    ADD CONSTRAINT driver_violation_table_driver_id_fkey FOREIGN KEY (driver_id) REFERENCES public.master_driver_table(driver_id) ON DELETE CASCADE;


--
-- Name: driver_violation_table driver_violation_table_violation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.driver_violation_table
    ADD CONSTRAINT driver_violation_table_violation_id_fkey FOREIGN KEY (violation_id) REFERENCES public.master_violation_table(id) ON DELETE CASCADE;


--
-- Name: business_table fk_b_location; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_table
    ADD CONSTRAINT fk_b_location FOREIGN KEY (b_location_id) REFERENCES public.master_location(id) ON DELETE CASCADE;


--
-- Name: business_table fk_b_state; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_table
    ADD CONSTRAINT fk_b_state FOREIGN KEY (b_state_id) REFERENCES public.master_states(id) ON DELETE CASCADE;


--
-- Name: business_table fk_b_type; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_table
    ADD CONSTRAINT fk_b_type FOREIGN KEY (b_typeid) REFERENCES public.business_type_table(b_typeid) ON DELETE CASCADE;


--
-- Name: stock_table fk_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stock_table
    ADD CONSTRAINT fk_category FOREIGN KEY (warehouse_id) REFERENCES public.warehouse_list(warehouse_id);


--
-- Name: master_vehicle_table fk_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_vehicle_table
    ADD CONSTRAINT fk_category FOREIGN KEY (vehicle_make) REFERENCES public.vehicle_make(id);


--
-- Name: daily_price_update fk_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.daily_price_update
    ADD CONSTRAINT fk_category FOREIGN KEY (product_id) REFERENCES public.master_product(product_id);


--
-- Name: master_vehicle_table fk_engine; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_vehicle_table
    ADD CONSTRAINT fk_engine FOREIGN KEY (vehicle_engine_type) REFERENCES public.vehicle_engine_type(id);


--
-- Name: master_mandi_table fk_loc; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_mandi_table
    ADD CONSTRAINT fk_loc FOREIGN KEY (mandi_city) REFERENCES public.master_location(id);


--
-- Name: business_table fk_mandi_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_table
    ADD CONSTRAINT fk_mandi_id FOREIGN KEY (b_mandi_id) REFERENCES public.master_mandi_table(mandi_id);


--
-- Name: mode_of_payment fk_mode; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mode_of_payment
    ADD CONSTRAINT fk_mode FOREIGN KEY (pay_mode) REFERENCES public.mode_of_payments_list(id);


--
-- Name: wholeseller_rating_table fk_order; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wholeseller_rating_table
    ADD CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES public.order_table(order_id) ON DELETE CASCADE;


--
-- Name: retailer_rating_table fk_order; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.retailer_rating_table
    ADD CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES public.order_table(order_id) ON DELETE CASCADE;


--
-- Name: order_item_table fk_order_item_status; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_item_table
    ADD CONSTRAINT fk_order_item_status FOREIGN KEY (order_item_status) REFERENCES public.order_status_table(order_status_id);


--
-- Name: order_table fk_order_status; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_table
    ADD CONSTRAINT fk_order_status FOREIGN KEY (order_status) REFERENCES public.order_status_table(order_status_id) ON DELETE SET NULL;


--
-- Name: insurance_table fk_provider; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_table
    ADD CONSTRAINT fk_provider FOREIGN KEY (insurance_provider) REFERENCES public.insurance_provider(id);


--
-- Name: order_table fk_retailer; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_table
    ADD CONSTRAINT fk_retailer FOREIGN KEY (retailer_id) REFERENCES public.business_table(bid) ON DELETE SET NULL;


--
-- Name: user_table fk_state; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_table
    ADD CONSTRAINT fk_state FOREIGN KEY (state) REFERENCES public.master_states(id);


--
-- Name: master_mandi_table fk_state; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_mandi_table
    ADD CONSTRAINT fk_state FOREIGN KEY (mandi_state) REFERENCES public.master_states(id);


--
-- Name: invoice_details_table fk_status; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_details_table
    ADD CONSTRAINT fk_status FOREIGN KEY (retailer_status) REFERENCES public.retailer_status(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: insurance_table fk_type; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_table
    ADD CONSTRAINT fk_type FOREIGN KEY (type_of_insurance) REFERENCES public.type_of_insurance(id);


--
-- Name: mode_of_payment fk_type; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mode_of_payment
    ADD CONSTRAINT fk_type FOREIGN KEY (pay_type) REFERENCES public.cash_payment_list(id);


--
-- Name: daily_price_update fk_unit; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.daily_price_update
    ADD CONSTRAINT fk_unit FOREIGN KEY (unit_id) REFERENCES public.units_table(id);


--
-- Name: order_item_table fk_units; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_item_table
    ADD CONSTRAINT fk_units FOREIGN KEY (unit_id) REFERENCES public.units_table(id);


--
-- Name: user_bank_details fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_bank_details
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.user_table(user_id) ON DELETE CASCADE;


--
-- Name: users fk_user_type_table; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_user_type_table FOREIGN KEY (user_type_id) REFERENCES public.user_type_table(user_type_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: master_vehicle_table fk_vehicle_model; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_vehicle_table
    ADD CONSTRAINT fk_vehicle_model FOREIGN KEY (vehicle_model) REFERENCES public.vehicle_model(id);


--
-- Name: order_table fk_wholeseller; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_table
    ADD CONSTRAINT fk_wholeseller FOREIGN KEY (wholeseller_id) REFERENCES public.business_table(bid) ON DELETE SET NULL;


--
-- Name: insurance_table insurance_table_driver_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_table
    ADD CONSTRAINT insurance_table_driver_id_fkey FOREIGN KEY (driver_id) REFERENCES public.master_driver_table(driver_id) ON DELETE CASCADE;


--
-- Name: insurance_table insurance_table_vehicle_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_table
    ADD CONSTRAINT insurance_table_vehicle_id_fkey FOREIGN KEY (vehicle_id) REFERENCES public.master_vehicle_table(vehicle_id) ON DELETE CASCADE;


--
-- Name: invoice_details_table invoice_details_table_invoice_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_details_table
    ADD CONSTRAINT invoice_details_table_invoice_id_fkey FOREIGN KEY (invoice_id) REFERENCES public.invoice_table(id) ON DELETE CASCADE;


--
-- Name: invoice_details_table invoice_details_table_order_item_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_details_table
    ADD CONSTRAINT invoice_details_table_order_item_id_fkey FOREIGN KEY (order_item_id) REFERENCES public.order_item_table(order_item_id) ON DELETE CASCADE;


--
-- Name: invoice_table invoice_table_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_table
    ADD CONSTRAINT invoice_table_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.order_table(order_id) ON DELETE CASCADE;


--
-- Name: invoice_table invoice_table_pay_mode_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_table
    ADD CONSTRAINT invoice_table_pay_mode_fkey FOREIGN KEY (pay_mode) REFERENCES public.mode_of_payments_list(id);


--
-- Name: invoice_table invoice_table_pay_type_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_table
    ADD CONSTRAINT invoice_table_pay_type_fkey FOREIGN KEY (pay_type) REFERENCES public.cash_payment_list(id);


--
-- Name: master_product master_product_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.master_product
    ADD CONSTRAINT master_product_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.master_category_table(category_id) ON DELETE CASCADE;


--
-- Name: order_item_table order_item_table_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_item_table
    ADD CONSTRAINT order_item_table_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.order_table(order_id) ON DELETE CASCADE;


--
-- Name: order_item_table order_item_table_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_item_table
    ADD CONSTRAINT order_item_table_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.master_product(product_id) ON DELETE CASCADE;


--
-- Name: order_table order_table_location_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_table
    ADD CONSTRAINT order_table_location_id_fkey FOREIGN KEY (location_id) REFERENCES public.master_location(id) ON DELETE SET NULL;


--
-- Name: order_table order_table_state_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_table
    ADD CONSTRAINT order_table_state_id_fkey FOREIGN KEY (state_id) REFERENCES public.master_states(id) ON DELETE SET NULL;


--
-- Name: product_price_table product_price_table_bid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_price_table
    ADD CONSTRAINT product_price_table_bid_fkey FOREIGN KEY (bid) REFERENCES public.business_table(bid) ON DELETE CASCADE;


--
-- Name: product_price_table product_price_table_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_price_table
    ADD CONSTRAINT product_price_table_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.master_product(product_id) ON DELETE CASCADE;


--
-- Name: stock_table stock_table_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stock_table
    ADD CONSTRAINT stock_table_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.master_product(product_id) ON DELETE CASCADE;


--
-- Name: stock_table stock_table_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stock_table
    ADD CONSTRAINT stock_table_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE SET NULL;


--
-- Name: user_table user_table_user_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_table
    ADD CONSTRAINT user_table_user_type_id_fkey FOREIGN KEY (user_type_id) REFERENCES public.user_type_table(user_type_id) ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--

