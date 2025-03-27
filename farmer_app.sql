--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: admin_schema; Type: SCHEMA; Schema: -; Owner: admin
--

CREATE SCHEMA admin_schema;


ALTER SCHEMA admin_schema OWNER TO admin;

--
-- Name: business_schema; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA business_schema;


ALTER SCHEMA business_schema OWNER TO postgres;

--
-- Name: users_schema; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA users_schema;


ALTER SCHEMA users_schema OWNER TO postgres;

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
-- Name: create_business_branch(integer, character varying, integer, integer, integer, character varying, character varying, character varying, integer, integer, character varying, character varying, integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.create_business_branch(p_bid integer, p_b_shop_name character varying, p_b_type_id integer, p_b_location integer, p_b_state integer, p_b_address character varying, p_b_email character varying, p_b_number character varying, p_active_status integer DEFAULT 0, p_b_mandi_id integer DEFAULT NULL::integer, p_b_gst_num character varying DEFAULT NULL::character varying, p_b_pan_num character varying DEFAULT NULL::character varying, p_b_privilege_user integer DEFAULT NULL::integer, p_b_established_year character varying DEFAULT NULL::character varying) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE
    new_branch_id INT;
BEGIN
    INSERT INTO admin_schema.business_branch_table (
        bid, b_shop_name, b_type_id, b_location, b_state, b_mandi_id,
        b_address, b_email, b_number, b_gst_num, b_pan_num,
        b_privilege_user, b_established_year, active_status, created_at, updated_at
    ) VALUES (
        p_bid, p_b_shop_name, p_b_type_id, p_b_location, p_b_state, p_b_mandi_id,
        p_b_address, p_b_email, p_b_number, p_b_gst_num, p_b_pan_num,
        p_b_privilege_user, p_b_established_year, p_active_status, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
    ) RETURNING b_branch_id INTO new_branch_id;

    RETURN new_branch_id;
END;
$$;


ALTER FUNCTION admin_schema.create_business_branch(p_bid integer, p_b_shop_name character varying, p_b_type_id integer, p_b_location integer, p_b_state integer, p_b_address character varying, p_b_email character varying, p_b_number character varying, p_active_status integer, p_b_mandi_id integer, p_b_gst_num character varying, p_b_pan_num character varying, p_b_privilege_user integer, p_b_established_year character varying) OWNER TO postgres;

--
-- Name: create_business_branch(integer, character varying, integer, integer, integer, character varying, character varying, character varying, integer, character varying, character varying, integer, character varying, integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.create_business_branch(p_bid integer, p_b_shop_name character varying, p_b_type_id integer, p_b_location integer, p_b_state integer, p_b_address character varying, p_b_email character varying, p_b_number character varying, p_b_mandi_id integer DEFAULT NULL::integer, p_b_gst_num character varying DEFAULT NULL::character varying, p_b_pan_num character varying DEFAULT NULL::character varying, p_b_privilege_user integer DEFAULT NULL::integer, p_b_established_year character varying DEFAULT NULL::character varying, p_active_status integer DEFAULT 0) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE
    new_branch_id INT;
BEGIN
    INSERT INTO admin_schema.business_branch_table (
        bid, b_shop_name, b_type_id, b_location, b_state, b_mandi_id,
        b_address, b_email, b_number, b_gst_num, b_pan_num,
        b_privilege_user, b_established_year, active_status
    ) VALUES (
        p_bid, p_b_shop_name, p_b_type_id, p_b_location, p_b_state, p_b_mandi_id,
        p_b_address, p_b_email, p_b_number, p_b_gst_num, p_b_pan_num,
        p_b_privilege_user, p_b_established_year, p_active_status
    ) RETURNING b_branch_id INTO new_branch_id;

    RETURN new_branch_id;
END;
$$;


ALTER FUNCTION admin_schema.create_business_branch(p_bid integer, p_b_shop_name character varying, p_b_type_id integer, p_b_location integer, p_b_state integer, p_b_address character varying, p_b_email character varying, p_b_number character varying, p_b_mandi_id integer, p_b_gst_num character varying, p_b_pan_num character varying, p_b_privilege_user integer, p_b_established_year character varying, p_active_status integer) OWNER TO postgres;

--
-- Name: get_all_business_branches(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_business_branches() RETURNS TABLE(b_branch_id integer, bid integer, b_shop_name character varying, b_type_id integer, b_location integer, b_state integer, b_mandi_id integer, b_address character varying, b_email character varying, b_number character varying, b_gst_num character varying, b_pan_num character varying, b_privilege_user integer, b_established_year character varying, active_status integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        b_branch_id, bid, b_shop_name, b_type_id, b_location, b_state, 
        b_mandi_id, b_address, b_email, b_number, b_gst_num, 
        b_pan_num, b_privilege_user, b_established_year, active_status
    FROM admin_schema.business_branch_table;
END;
$$;


ALTER FUNCTION admin_schema.get_all_business_branches() OWNER TO postgres;

--
-- Name: get_all_business_users(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_business_users() RETURNS TABLE(b_id integer, user_name character varying, active_status integer, created_at timestamp without time zone, updated_at timestamp without time zone, is_locked boolean)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT b_id, user_name, active_status, created_at, updated_at, is_locked
    FROM admin_schema.business_user_table;
END;
$$;


ALTER FUNCTION admin_schema.get_all_business_users() OWNER TO postgres;

--
-- Name: get_all_businesses(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_businesses() RETURNS TABLE(bid bigint, b_person_name character varying, b_registration_num character varying, b_owner_name character varying, active_status integer, created_at timestamp with time zone, updated_at timestamp with time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        bid, b_person_name, b_registration_num, b_owner_name, 
        active_status, created_at, updated_at 
    FROM admin_schema.business_table;
END;
$$;


ALTER FUNCTION admin_schema.get_all_businesses() OWNER TO postgres;

--
-- Name: get_all_cash_payments(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_cash_payments() RETURNS TABLE(id integer, payment_type character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT id, payment_type FROM admin_schema.cash_payment_list;
END;
$$;


ALTER FUNCTION admin_schema.get_all_cash_payments() OWNER TO postgres;

--
-- Name: get_all_category_regional_names(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_category_regional_names() RETURNS TABLE(category_regional_id integer, language_id integer, category_id integer, category_regional_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        crn.category_regional_id, 
        crn.language_id, 
        crn.category_id, 
        crn.category_regional_name
    FROM admin_schema.category_regional_name AS crn;
END;
$$;


ALTER FUNCTION admin_schema.get_all_category_regional_names() OWNER TO postgres;

--
-- Name: get_all_cities(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_cities() RETURNS TABLE(id integer, city_shortnames character varying, city_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT mc.id, mc.city_shortnames, mc.city_name
    FROM admin_schema.master_city mc;
END;
$$;


ALTER FUNCTION admin_schema.get_all_cities() OWNER TO postgres;

--
-- Name: get_all_languages(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_languages() RETURNS TABLE(id integer, language_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT ml.id, ml.language_name
    FROM admin_schema.master_language ml;
END;
$$;


ALTER FUNCTION admin_schema.get_all_languages() OWNER TO postgres;

--
-- Name: get_all_mandi(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_mandi() RETURNS TABLE(mandi_id integer, mandi_location character varying, mandi_incharge character varying, mandi_incharge_num character varying, mandi_pincode character varying, mandi_address text, mandi_state integer, mandi_name character varying, mandi_shortnames character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT * FROM admin_schema.master_mandi_table;
END;
$$;


ALTER FUNCTION admin_schema.get_all_mandi() OWNER TO postgres;

--
-- Name: get_all_order_statuses(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_order_statuses() RETURNS TABLE(order_status_id integer, order_status character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT order_status_id, order_status FROM admin_schema.order_status_table ORDER BY order_status_id;
END;
$$;


ALTER FUNCTION admin_schema.get_all_order_statuses() OWNER TO postgres;

--
-- Name: get_all_payment_modes(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_payment_modes() RETURNS TABLE(id integer, payment_mode character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT id, payment_mode FROM admin_schema.mode_of_payments_list ORDER BY id;
END;
$$;


ALTER FUNCTION admin_schema.get_all_payment_modes() OWNER TO postgres;

--
-- Name: get_all_product_categories(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_product_categories() RETURNS TABLE(category_id integer, category_name character varying, super_cat_id integer, img_path text, active_status integer, category_regional_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        mpct.category_id, 
        mpct.category_name, 
        mpct.super_cat_id, 
        mpct.img_path, 
        mpct.active_status, 
        mpct.category_regional_id
    FROM admin_schema.master_product_category_table AS mpct;
END;
$$;


ALTER FUNCTION admin_schema.get_all_product_categories() OWNER TO postgres;

--
-- Name: get_all_product_regional_names(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_product_regional_names() RETURNS TABLE(product_regional_id integer, language_id integer, product_id integer, product_name character varying, product_regional_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT prn.product_regional_id, prn.language_id, prn.product_id, 
           p.product_name, prn.product_regional_name
    FROM admin_schema.product_regional_name prn
    JOIN admin_schema.master_product p ON prn.product_id = p.product_id;
END;
$$;


ALTER FUNCTION admin_schema.get_all_product_regional_names() OWNER TO postgres;

--
-- Name: get_all_products(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_products() RETURNS TABLE(product_id bigint, category_id integer, category_name character varying, product_name character varying, image_path character varying, active_status integer, product_regional_id integer, product_regional_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        p.product_id,
        p.category_id,
        c.category_name,
        p.product_name,
        p.image_path,
        p.active_status,
        p.product_regional_id,
        prn.product_regional_name
    FROM admin_schema.master_product p
    LEFT JOIN admin_schema.master_product_category_table c 
        ON p.category_id = c.category_id
    LEFT JOIN admin_schema.product_regional_name prn 
        ON p.product_regional_id = prn.product_regional_id;
END;
$$;


ALTER FUNCTION admin_schema.get_all_products() OWNER TO postgres;

--
-- Name: get_all_states(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_states() RETURNS TABLE(state_id integer, state character varying, state_shortnames character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT master_states.id, master_states.state, master_states.state_shortnames 
    FROM admin_schema.master_states;
END;
$$;


ALTER FUNCTION admin_schema.get_all_states() OWNER TO postgres;

--
-- Name: get_all_units(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_units() RETURNS TABLE(id integer, unit_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT id, unit_name FROM admin_schema.units_table ORDER BY id;
END;
$$;


ALTER FUNCTION admin_schema.get_all_units() OWNER TO postgres;

--
-- Name: get_all_users(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_users() RETURNS TABLE(user_id integer, name text, mobile_num text, email text, address text, pincode text, location integer, state integer, active_status integer, role_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT u.user_id, u.name::TEXT, u.mobile_num::TEXT, u.email::TEXT, u.address::TEXT, u.pincode::TEXT, 
           u.location, u.state, u.active_status, u.role_id
    FROM admin_schema.user_table u
    ORDER BY u.user_id;
END;
$$;


ALTER FUNCTION admin_schema.get_all_users() OWNER TO postgres;

--
-- Name: get_all_vehicles(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_all_vehicles() RETURNS TABLE(vehicle_id integer, vehicle_name character varying, vehicle_registration_no character varying, vehicle_manufacture_year character varying, vehicle_warranty character varying, vehicle_make_name character varying, vehicle_model_name character varying, vehicle_engine_type_name character varying, vehicle_purchase_date date, vehicle_color character varying, created_at timestamp without time zone, updated_at timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        v.vehicle_id,
        v.vehicle_name,
        v.vehicle_registration_no,
        v.vehicle_manufacture_year,
        v.vehicle_warranty,
        v.vehicle_make_name,
        v.vehicle_model_name,
        v.vehicle_engine_type_name,
        v.vehicle_purchase_date,
        v.vehicle_color,
        v.created_at,
        v.updated_at
    FROM admin_schema.master_vehicle_table v;
END;
$$;


ALTER FUNCTION admin_schema.get_all_vehicles() OWNER TO postgres;

--
-- Name: get_business_branch_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_business_branch_by_id(p_b_branch_id integer) RETURNS TABLE(b_branch_id integer, bid integer, b_shop_name character varying, b_type_id integer, b_location integer, b_state integer, b_mandi_id integer, b_address character varying, b_email character varying, b_number character varying, b_gst_num character varying, b_pan_num character varying, b_privilege_user integer, b_established_year character varying, active_status integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        b_branch_id, bid, b_shop_name, b_type_id, b_location, b_state, 
        b_mandi_id, b_address, b_email, b_number, b_gst_num, 
        b_pan_num, b_privilege_user, b_established_year, active_status
    FROM admin_schema.business_branch_table
    WHERE b_branch_id = p_b_branch_id;
END;
$$;


ALTER FUNCTION admin_schema.get_business_branch_by_id(p_b_branch_id integer) OWNER TO postgres;

--
-- Name: get_business_by_id(bigint); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_business_by_id(p_bid bigint) RETURNS TABLE(bid bigint, b_person_name character varying, b_registration_num character varying, b_owner_name character varying, active_status integer, created_at timestamp with time zone, updated_at timestamp with time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT 
        bid, b_person_name, b_registration_num, b_owner_name, 
        active_status, created_at, updated_at 
    FROM admin_schema.business_table
    WHERE bid = p_bid;
END;
$$;


ALTER FUNCTION admin_schema.get_business_by_id(p_bid bigint) OWNER TO postgres;

--
-- Name: get_business_category(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_business_category(p_b_category_id integer DEFAULT NULL::integer) RETURNS TABLE(b_category_id integer, b_category_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT bc.b_category_id, bc.b_category_name
    FROM admin_schema.business_category_table bc
    WHERE p_b_category_id IS NULL OR bc.b_category_id = p_b_category_id;
END;
$$;


ALTER FUNCTION admin_schema.get_business_category(p_b_category_id integer) OWNER TO postgres;

--
-- Name: get_business_user_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_business_user_by_id(p_b_id integer) RETURNS TABLE(b_id integer, user_name character varying, active_status integer, created_at timestamp without time zone, updated_at timestamp without time zone, is_locked boolean)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT b_id, user_name, active_status, created_at, updated_at, is_locked
    FROM admin_schema.business_user_table
    WHERE b_id = p_b_id;
END;
$$;


ALTER FUNCTION admin_schema.get_business_user_by_id(p_b_id integer) OWNER TO postgres;

--
-- Name: get_cash_payment_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_cash_payment_by_id(p_id integer) RETURNS TABLE(id integer, payment_type character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT id, payment_type
    FROM admin_schema.cash_payment_list
    WHERE id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.get_cash_payment_by_id(p_id integer) OWNER TO postgres;

--
-- Name: get_category_regional_name_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_category_regional_name_by_id(p_id integer) RETURNS TABLE(category_regional_id integer, language_id integer, category_id integer, category_regional_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT category_regional_id, language_id, category_id, category_regional_name
    FROM admin_schema.category_regional_name
    WHERE category_regional_id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.get_category_regional_name_by_id(p_id integer) OWNER TO postgres;

--
-- Name: get_city_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_city_by_id(p_id integer) RETURNS TABLE(id integer, city_shortnames character varying, city_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT mc.id, mc.city_shortnames, mc.city_name
    FROM admin_schema.master_city mc
    WHERE mc.id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.get_city_by_id(p_id integer) OWNER TO postgres;

--
-- Name: get_language_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_language_by_id(p_id integer) RETURNS TABLE(id integer, language_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT ml.id, ml.language_name
    FROM admin_schema.master_language ml
    WHERE ml.id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.get_language_by_id(p_id integer) OWNER TO postgres;

--
-- Name: get_mandi_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_mandi_by_id(p_mandi_id integer) RETURNS TABLE(mandi_id integer, mandi_location character varying, mandi_incharge character varying, mandi_incharge_num character varying, mandi_pincode character varying, mandi_address text, mandi_state integer, mandi_name character varying, mandi_shortnames character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT * FROM admin_schema.master_mandi_table WHERE mandi_id = p_mandi_id;
END;
$$;


ALTER FUNCTION admin_schema.get_mandi_by_id(p_mandi_id integer) OWNER TO postgres;

--
-- Name: get_order_status_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_order_status_by_id(p_order_status_id integer) RETURNS TABLE(order_status_id integer, order_status character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT order_status_id, order_status 
    FROM admin_schema.order_status_table 
    WHERE order_status_id = p_order_status_id;
END;
$$;


ALTER FUNCTION admin_schema.get_order_status_by_id(p_order_status_id integer) OWNER TO postgres;

--
-- Name: get_payment_mode_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_payment_mode_by_id(p_id integer) RETURNS TABLE(id integer, payment_mode character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT id, payment_mode FROM admin_schema.mode_of_payments_list WHERE id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.get_payment_mode_by_id(p_id integer) OWNER TO postgres;

--
-- Name: get_product_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_product_by_id(p_product_id integer) RETURNS TABLE(product_id integer, category_id integer, category_name character varying, product_name character varying, image_path character varying, active_status integer, product_regional_id integer, product_regional_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        p.product_id, 
        p.category_id, 
        c.category_name, 
        p.product_name, 
        p.image_path, 
        p.active_status, 
        p.product_regional_id, 
        prn.product_regional_name
    FROM admin_schema.master_product p
    JOIN admin_schema.master_category c ON p.category_id = c.category_id
    JOIN admin_schema.product_regional_name prn ON p.product_regional_id = prn.product_regional_id
    WHERE p.product_id = p_product_id;
END;
$$;


ALTER FUNCTION admin_schema.get_product_by_id(p_product_id integer) OWNER TO postgres;

--
-- Name: get_product_category_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_product_category_by_id(p_category_id integer) RETURNS TABLE(category_id integer, category_name character varying, super_cat_id integer, img_path text, active_status integer, category_regional_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        category_id, category_name, super_cat_id, img_path, active_status, category_regional_id
    FROM admin_schema.master_product_category_table
    WHERE category_id = p_category_id;
END;
$$;


ALTER FUNCTION admin_schema.get_product_category_by_id(p_category_id integer) OWNER TO postgres;

--
-- Name: get_product_regional_name_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_product_regional_name_by_id(p_product_regional_id integer) RETURNS TABLE(product_regional_id integer, language_id integer, product_id integer, product_name character varying, product_regional_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT prn.product_regional_id, prn.language_id, prn.product_id, 
           p.product_name, prn.product_regional_name
    FROM admin_schema.product_regional_name prn
    JOIN admin_schema.master_product p ON prn.product_id = p.product_id
    WHERE prn.product_regional_id = p_product_regional_id;
END;
$$;


ALTER FUNCTION admin_schema.get_product_regional_name_by_id(p_product_regional_id integer) OWNER TO postgres;

--
-- Name: get_products_by_category(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_products_by_category(p_category_id integer) RETURNS TABLE(product_id integer, category_id integer, category_name character varying, product_name character varying, image_path character varying, active_status integer, product_regional_id integer, product_regional_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        p.product_id, 
        p.category_id, 
        c.category_name, 
        p.product_name, 
        p.image_path, 
        p.active_status, 
        p.product_regional_id, 
        prn.product_regional_name
    FROM admin_schema.master_product p
    JOIN admin_schema.master_category c ON p.category_id = c.category_id
    JOIN admin_schema.product_regional_name prn ON p.product_regional_id = prn.product_regional_id
    WHERE p.category_id = p_category_id;
END;
$$;


ALTER FUNCTION admin_schema.get_products_by_category(p_category_id integer) OWNER TO postgres;

--
-- Name: get_state_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_state_by_id(p_id integer) RETURNS TABLE(id integer, state character varying, state_shortnames character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT id, state, state_shortnames FROM admin_schema.master_states
    WHERE id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.get_state_by_id(p_id integer) OWNER TO postgres;

--
-- Name: get_unit_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_unit_by_id(p_id integer) RETURNS TABLE(id integer, unit_name character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT id, unit_name 
    FROM admin_schema.units_table 
    WHERE id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.get_unit_by_id(p_id integer) OWNER TO postgres;

--
-- Name: get_user_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_user_by_id(p_user_id integer) RETURNS TABLE(user_id integer, name character varying, mobile_num character varying, email character varying, address text, pincode character varying, location integer, state integer, active_status integer, role_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT user_id, name, mobile_num, email, address, pincode, location, state, active_status, role_id 
    FROM admin_schema.user_table 
    WHERE user_id = p_user_id;
END;
$$;


ALTER FUNCTION admin_schema.get_user_by_id(p_user_id integer) OWNER TO postgres;

--
-- Name: get_vehicle_by_id(integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.get_vehicle_by_id(p_vehicle_id integer) RETURNS TABLE(vehicle_id integer, vehicle_name character varying, vehicle_registration_no character varying, vehicle_manufacture_year character varying, vehicle_warranty character varying, vehicle_make_name character varying, vehicle_model_name character varying, vehicle_engine_type_name character varying, vehicle_purchase_date date, vehicle_color character varying, created_at timestamp without time zone, updated_at timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY 
    SELECT 
        v.vehicle_id,
        v.vehicle_name,
        v.vehicle_registration_no,
        v.vehicle_manufacture_year,
        v.vehicle_warranty,
        v.vehicle_make_name,
        v.vehicle_model_name,
        v.vehicle_engine_type_name,
        v.vehicle_purchase_date,
        v.vehicle_color,
        v.created_at,
        v.updated_at
    FROM admin_schema.master_vehicle_table v
    WHERE v.vehicle_id = p_vehicle_id;
END;
$$;


ALTER FUNCTION admin_schema.get_vehicle_by_id(p_vehicle_id integer) OWNER TO postgres;

--
-- Name: insert_business(character varying, character varying, character varying, integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_business(p_b_person_name character varying, p_b_registration_num character varying, p_b_owner_name character varying, p_active_status integer) RETURNS bigint
    LANGUAGE plpgsql
    AS $$
DECLARE
    new_bid BIGINT;
BEGIN
    INSERT INTO admin_schema.business_table (b_person_name, b_registration_num, b_owner_name, active_status)
    VALUES (p_b_person_name, p_b_registration_num, p_b_owner_name, p_active_status)
    RETURNING bid INTO new_bid;

    RETURN new_bid;
END;
$$;


ALTER FUNCTION admin_schema.insert_business(p_b_person_name character varying, p_b_registration_num character varying, p_b_owner_name character varying, p_active_status integer) OWNER TO postgres;

--
-- Name: insert_business_category(character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_business_category(p_b_category_name character varying) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE
    new_b_category_id INT;
BEGIN
    INSERT INTO admin_schema.business_category_table (b_category_name)
    VALUES (p_b_category_name)
    RETURNING b_category_id INTO new_b_category_id;

    RETURN new_b_category_id;
END;
$$;


ALTER FUNCTION admin_schema.insert_business_category(p_b_category_name character varying) OWNER TO postgres;

--
-- Name: insert_business_user(); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_business_user() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Insert into business_user_table using values from business_branch_table
    INSERT INTO admin_schema.business_user_table (b_id, user_name, password, active_status)
    VALUES (NEW.bid, NEW.b_shop_name, NEW.b_number, 1);

    RETURN NEW;
END;
$$;


ALTER FUNCTION admin_schema.insert_business_user() OWNER TO postgres;

--
-- Name: insert_cash_payment(character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_cash_payment(p_payment_type character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.cash_payment_list (payment_type)
    VALUES (p_payment_type)
    ON CONFLICT (payment_type) DO NOTHING;
END;
$$;


ALTER FUNCTION admin_schema.insert_cash_payment(p_payment_type character varying) OWNER TO postgres;

--
-- Name: insert_category_regional_name(integer, integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_category_regional_name(p_language_id integer, p_category_id integer, p_category_regional_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.category_regional_name (language_id, category_id, category_regional_name)
    VALUES (p_language_id, p_category_id, p_category_regional_name);
END;
$$;


ALTER FUNCTION admin_schema.insert_category_regional_name(p_language_id integer, p_category_id integer, p_category_regional_name character varying) OWNER TO postgres;

--
-- Name: insert_city(character varying, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_city(p_city_shortnames character varying, p_city_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.master_city (city_shortnames, city_name)
    VALUES (p_city_shortnames, p_city_name);
END;
$$;


ALTER FUNCTION admin_schema.insert_city(p_city_shortnames character varying, p_city_name character varying) OWNER TO postgres;

--
-- Name: insert_language(character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_language(p_language_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.master_language (language_name)
    VALUES (p_language_name);
END;
$$;


ALTER FUNCTION admin_schema.insert_language(p_language_name character varying) OWNER TO postgres;

--
-- Name: insert_mandi(character varying, character varying, character varying, character varying, text, integer, character varying, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_mandi(p_mandi_location character varying, p_mandi_incharge character varying, p_mandi_incharge_num character varying, p_mandi_pincode character varying, p_mandi_address text DEFAULT NULL::text, p_mandi_state integer DEFAULT NULL::integer, p_mandi_name character varying DEFAULT NULL::character varying, p_mandi_shortnames character varying DEFAULT NULL::character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.master_mandi_table (
        mandi_location, mandi_incharge, mandi_incharge_num, 
        mandi_pincode, mandi_address, mandi_state, 
        mandi_name, mandi_shortnames
    ) VALUES (
        p_mandi_location, p_mandi_incharge, p_mandi_incharge_num, 
        p_mandi_pincode, p_mandi_address, p_mandi_state, 
        p_mandi_name, p_mandi_shortnames
    );
END;
$$;


ALTER FUNCTION admin_schema.insert_mandi(p_mandi_location character varying, p_mandi_incharge character varying, p_mandi_incharge_num character varying, p_mandi_pincode character varying, p_mandi_address text, p_mandi_state integer, p_mandi_name character varying, p_mandi_shortnames character varying) OWNER TO postgres;

--
-- Name: insert_order_status(character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_order_status(p_order_status character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.order_status_table (order_status) 
    VALUES (p_order_status);
END;
$$;


ALTER FUNCTION admin_schema.insert_order_status(p_order_status character varying) OWNER TO postgres;

--
-- Name: insert_payment_mode(character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_payment_mode(p_payment_mode character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.mode_of_payments_list (payment_mode) VALUES (p_payment_mode);
END;
$$;


ALTER FUNCTION admin_schema.insert_payment_mode(p_payment_mode character varying) OWNER TO postgres;

--
-- Name: insert_product_category(character varying, integer, text, integer, integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_product_category(p_category_name character varying, p_super_cat_id integer DEFAULT NULL::integer, p_img_path text DEFAULT NULL::text, p_active_status integer DEFAULT 1, p_category_regional_id integer DEFAULT NULL::integer) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.master_product_category_table 
    (category_name, super_cat_id, img_path, active_status, category_regional_id)
    VALUES (p_category_name, p_super_cat_id, p_img_path, p_active_status, p_category_regional_id);
END;
$$;


ALTER FUNCTION admin_schema.insert_product_category(p_category_name character varying, p_super_cat_id integer, p_img_path text, p_active_status integer, p_category_regional_id integer) OWNER TO postgres;

--
-- Name: insert_product_regional_name(integer, integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_product_regional_name(p_language_id integer, p_product_id integer, p_product_regional_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.product_regional_name (language_id, product_id, product_regional_name)
    VALUES (p_language_id, p_product_id, p_product_regional_name);
END;
$$;


ALTER FUNCTION admin_schema.insert_product_regional_name(p_language_id integer, p_product_id integer, p_product_regional_name character varying) OWNER TO postgres;

--
-- Name: insert_state(character varying, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_state(p_state character varying, p_state_shortnames character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.master_states (state, state_shortnames)
    VALUES (p_state, p_state_shortnames);
END;
$$;


ALTER FUNCTION admin_schema.insert_state(p_state character varying, p_state_shortnames character varying) OWNER TO postgres;

--
-- Name: insert_unit(character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_unit(p_unit_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.units_table (unit_name) 
    VALUES (p_unit_name);
END;
$$;


ALTER FUNCTION admin_schema.insert_unit(p_unit_name character varying) OWNER TO postgres;

--
-- Name: insert_user(character varying, character varying, character varying, text, character varying, integer, integer, integer, integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_user(p_name character varying, p_mobile_num character varying, p_email character varying, p_address text, p_pincode character varying, p_location integer, p_state integer, p_active_status integer DEFAULT 0, p_role_id integer DEFAULT 5) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.user_table (
        name, mobile_num, email, address, pincode, location, state, active_status, role_id
    ) VALUES (
        p_name, p_mobile_num, p_email, p_address, p_pincode, p_location, p_state, p_active_status, p_role_id
    );
END;
$$;


ALTER FUNCTION admin_schema.insert_user(p_name character varying, p_mobile_num character varying, p_email character varying, p_address text, p_pincode character varying, p_location integer, p_state integer, p_active_status integer, p_role_id integer) OWNER TO postgres;

--
-- Name: insert_vehicle(character varying, character varying, character varying, character varying, integer, integer, integer, date, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.insert_vehicle(p_vehicle_name character varying, p_vehicle_registration_no character varying, p_vehicle_manufacture_year character varying, p_vehicle_warranty character varying, p_vehicle_make integer, p_vehicle_model integer, p_vehicle_engine_type integer, p_vehicle_purchase_date date, p_vehicle_color character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.master_vehicle_table (
        vehicle_name, vehicle_registration_no, vehicle_manufacture_year,
        vehicle_warranty, vehicle_make, vehicle_model, vehicle_engine_type,
        vehicle_purchase_date, vehicle_color, created_at, updated_at
    ) VALUES (
        p_vehicle_name, p_vehicle_registration_no, p_vehicle_manufacture_year,
        p_vehicle_warranty, p_vehicle_make, p_vehicle_model, p_vehicle_engine_type,
        p_vehicle_purchase_date, p_vehicle_color, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
    );
END;
$$;


ALTER FUNCTION admin_schema.insert_vehicle(p_vehicle_name character varying, p_vehicle_registration_no character varying, p_vehicle_manufacture_year character varying, p_vehicle_warranty character varying, p_vehicle_make integer, p_vehicle_model integer, p_vehicle_engine_type integer, p_vehicle_purchase_date date, p_vehicle_color character varying) OWNER TO postgres;

--
-- Name: update_business(bigint, character varying, character varying, character varying, integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_business(p_bid bigint, p_b_person_name character varying, p_b_registration_num character varying, p_b_owner_name character varying, p_active_status integer) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.business_table
    SET b_person_name = p_b_person_name,
        b_registration_num = p_b_registration_num,
        b_owner_name = p_b_owner_name,
        active_status = p_active_status,
        updated_at = NOW()
    WHERE bid = p_bid;

    RETURN FOUND;
END;
$$;


ALTER FUNCTION admin_schema.update_business(p_bid bigint, p_b_person_name character varying, p_b_registration_num character varying, p_b_owner_name character varying, p_active_status integer) OWNER TO postgres;

--
-- Name: update_business_branch(integer, character varying, integer, integer, integer, character varying, character varying, character varying, integer, integer, character varying, character varying, integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_business_branch(p_b_branch_id integer, p_b_shop_name character varying, p_b_type_id integer, p_b_location integer, p_b_state integer, p_b_address character varying, p_b_email character varying, p_b_number character varying, p_active_status integer DEFAULT 0, p_b_mandi_id integer DEFAULT NULL::integer, p_b_gst_num character varying DEFAULT NULL::character varying, p_b_pan_num character varying DEFAULT NULL::character varying, p_b_privilege_user integer DEFAULT NULL::integer, p_b_established_year character varying DEFAULT NULL::character varying) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.business_branch_table
    SET
        b_shop_name = p_b_shop_name,
        b_type_id = p_b_type_id,
        b_location = p_b_location,
        b_state = p_b_state,
        b_mandi_id = p_b_mandi_id,
        b_address = p_b_address,
        b_email = p_b_email,
        b_number = p_b_number,
        b_gst_num = p_b_gst_num,
        b_pan_num = p_b_pan_num,
        b_privilege_user = p_b_privilege_user,
        b_established_year = p_b_established_year,
        active_status = p_active_status,
        updated_at = CURRENT_TIMESTAMP
    WHERE b_branch_id = p_b_branch_id;

    RETURN FOUND;
END;
$$;


ALTER FUNCTION admin_schema.update_business_branch(p_b_branch_id integer, p_b_shop_name character varying, p_b_type_id integer, p_b_location integer, p_b_state integer, p_b_address character varying, p_b_email character varying, p_b_number character varying, p_active_status integer, p_b_mandi_id integer, p_b_gst_num character varying, p_b_pan_num character varying, p_b_privilege_user integer, p_b_established_year character varying) OWNER TO postgres;

--
-- Name: update_business_branch(integer, character varying, integer, integer, integer, character varying, character varying, character varying, integer, character varying, character varying, integer, character varying, integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_business_branch(p_b_branch_id integer, p_b_shop_name character varying, p_b_type_id integer, p_b_location integer, p_b_state integer, p_b_address character varying, p_b_email character varying, p_b_number character varying, p_b_mandi_id integer DEFAULT NULL::integer, p_b_gst_num character varying DEFAULT NULL::character varying, p_b_pan_num character varying DEFAULT NULL::character varying, p_b_privilege_user integer DEFAULT NULL::integer, p_b_established_year character varying DEFAULT NULL::character varying, p_active_status integer DEFAULT 0) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.business_branch_table
    SET
        b_shop_name = p_b_shop_name,
        b_type_id = p_b_type_id,
        b_location = p_b_location,
        b_state = p_b_state,
        b_mandi_id = p_b_mandi_id,
        b_address = p_b_address,
        b_email = p_b_email,
        b_number = p_b_number,
        b_gst_num = p_b_gst_num,
        b_pan_num = p_b_pan_num,
        b_privilege_user = p_b_privilege_user,
        b_established_year = p_b_established_year,
        active_status = p_active_status,
        updated_at = CURRENT_TIMESTAMP
    WHERE b_branch_id = p_b_branch_id;

    IF FOUND THEN
        RETURN TRUE;
    ELSE
        RETURN FALSE;
    END IF;
END;
$$;


ALTER FUNCTION admin_schema.update_business_branch(p_b_branch_id integer, p_b_shop_name character varying, p_b_type_id integer, p_b_location integer, p_b_state integer, p_b_address character varying, p_b_email character varying, p_b_number character varying, p_b_mandi_id integer, p_b_gst_num character varying, p_b_pan_num character varying, p_b_privilege_user integer, p_b_established_year character varying, p_active_status integer) OWNER TO postgres;

--
-- Name: update_business_category(integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_business_category(p_b_category_id integer, p_b_category_name character varying) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.business_category_table
    SET b_category_name = p_b_category_name
    WHERE b_category_id = p_b_category_id;

    RETURN FOUND;  -- Returns TRUE if row was updated, FALSE if not
END;
$$;


ALTER FUNCTION admin_schema.update_business_category(p_b_category_id integer, p_b_category_name character varying) OWNER TO postgres;

--
-- Name: update_business_user(integer, character varying, integer, boolean); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_business_user(p_b_id integer, p_user_name character varying, p_active_status integer, p_is_locked boolean) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.business_user_table
    SET user_name = p_user_name,
        active_status = p_active_status,
        updated_at = CURRENT_TIMESTAMP,
        is_locked = p_is_locked
    WHERE b_id = p_b_id;
END;
$$;


ALTER FUNCTION admin_schema.update_business_user(p_b_id integer, p_user_name character varying, p_active_status integer, p_is_locked boolean) OWNER TO postgres;

--
-- Name: update_cash_payment(integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_cash_payment(p_id integer, p_payment_type character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.cash_payment_list
    SET payment_type = p_payment_type
    WHERE id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.update_cash_payment(p_id integer, p_payment_type character varying) OWNER TO postgres;

--
-- Name: update_category_regional_name(integer, integer, integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_category_regional_name(p_category_regional_id integer, p_language_id integer, p_category_id integer, p_category_regional_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.category_regional_name
    SET language_id = p_language_id,
        category_id = p_category_id,
        category_regional_name = p_category_regional_name
    WHERE category_regional_id = p_category_regional_id;
END;
$$;


ALTER FUNCTION admin_schema.update_category_regional_name(p_category_regional_id integer, p_language_id integer, p_category_id integer, p_category_regional_name character varying) OWNER TO postgres;

--
-- Name: update_city(integer, character varying, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_city(p_id integer, p_city_shortnames character varying, p_city_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.master_city
    SET city_shortnames = p_city_shortnames,
        city_name = p_city_name
    WHERE id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.update_city(p_id integer, p_city_shortnames character varying, p_city_name character varying) OWNER TO postgres;

--
-- Name: update_language(integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_language(p_id integer, p_language_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.master_language
    SET language_name = p_language_name
    WHERE id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.update_language(p_id integer, p_language_name character varying) OWNER TO postgres;

--
-- Name: update_mandi(integer, character varying, character varying, character varying, character varying, text, integer, character varying, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_mandi(p_mandi_id integer, p_mandi_location character varying, p_mandi_incharge character varying, p_mandi_incharge_num character varying, p_mandi_pincode character varying, p_mandi_address text DEFAULT NULL::text, p_mandi_state integer DEFAULT NULL::integer, p_mandi_name character varying DEFAULT NULL::character varying, p_mandi_shortnames character varying DEFAULT NULL::character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.master_mandi_table
    SET 
        mandi_location = p_mandi_location,
        mandi_incharge = p_mandi_incharge,
        mandi_incharge_num = p_mandi_incharge_num,
        mandi_pincode = p_mandi_pincode,
        mandi_address = p_mandi_address,
        mandi_state = p_mandi_state,
        mandi_name = p_mandi_name,
        mandi_shortnames = p_mandi_shortnames
    WHERE mandi_id = p_mandi_id;
END;
$$;


ALTER FUNCTION admin_schema.update_mandi(p_mandi_id integer, p_mandi_location character varying, p_mandi_incharge character varying, p_mandi_incharge_num character varying, p_mandi_pincode character varying, p_mandi_address text, p_mandi_state integer, p_mandi_name character varying, p_mandi_shortnames character varying) OWNER TO postgres;

--
-- Name: update_order_status(integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_order_status(p_order_status_id integer, p_order_status character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.order_status_table 
    SET order_status = p_order_status 
    WHERE order_status_id = p_order_status_id;
END;
$$;


ALTER FUNCTION admin_schema.update_order_status(p_order_status_id integer, p_order_status character varying) OWNER TO postgres;

--
-- Name: update_payment_mode(integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_payment_mode(p_id integer, p_payment_mode character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.mode_of_payments_list SET payment_mode = p_payment_mode WHERE id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.update_payment_mode(p_id integer, p_payment_mode character varying) OWNER TO postgres;

--
-- Name: update_product(bigint, integer, character varying, character varying, integer, integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_product(p_product_id bigint, p_category_id integer, p_product_name character varying, p_image_path character varying, p_active_status integer, p_product_regional_id integer) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.master_product
    SET 
        category_id = p_category_id,
        product_name = p_product_name,
        image_path = p_image_path,
        active_status = p_active_status,
        product_regional_id = p_product_regional_id
    WHERE product_id = p_product_id;
END;
$$;


ALTER FUNCTION admin_schema.update_product(p_product_id bigint, p_category_id integer, p_product_name character varying, p_image_path character varying, p_active_status integer, p_product_regional_id integer) OWNER TO postgres;

--
-- Name: update_product_category(integer, character varying, integer, text, integer, integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_product_category(p_category_id integer, p_category_name character varying, p_super_cat_id integer DEFAULT NULL::integer, p_img_path text DEFAULT NULL::text, p_active_status integer DEFAULT 1, p_category_regional_id integer DEFAULT NULL::integer) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.master_product_category_table
    SET 
        category_name = p_category_name,
        super_cat_id = p_super_cat_id,
        img_path = p_img_path,
        active_status = p_active_status,
        category_regional_id = p_category_regional_id
    WHERE category_id = p_category_id;
END;
$$;


ALTER FUNCTION admin_schema.update_product_category(p_category_id integer, p_category_name character varying, p_super_cat_id integer, p_img_path text, p_active_status integer, p_category_regional_id integer) OWNER TO postgres;

--
-- Name: update_product_regional_name(integer, integer, integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_product_regional_name(p_product_regional_id integer, p_language_id integer, p_product_id integer, p_product_regional_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.product_regional_name
    SET language_id = p_language_id,
        product_id = p_product_id,
        product_regional_name = p_product_regional_name
    WHERE product_regional_id = p_product_regional_id;
END;
$$;


ALTER FUNCTION admin_schema.update_product_regional_name(p_product_regional_id integer, p_language_id integer, p_product_id integer, p_product_regional_name character varying) OWNER TO postgres;

--
-- Name: update_state(integer, character varying, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_state(p_id integer, p_state character varying, p_state_shortnames character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.master_states 
    SET state = p_state, state_shortnames = p_state_shortnames
    WHERE id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.update_state(p_id integer, p_state character varying, p_state_shortnames character varying) OWNER TO postgres;

--
-- Name: update_unit(integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_unit(p_id integer, p_unit_name character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.units_table 
    SET unit_name = p_unit_name 
    WHERE id = p_id;
END;
$$;


ALTER FUNCTION admin_schema.update_unit(p_id integer, p_unit_name character varying) OWNER TO postgres;

--
-- Name: update_user(integer, character varying, character varying, character varying, text, character varying, integer, integer, integer, integer); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_user(p_user_id integer, p_name character varying, p_mobile_num character varying, p_email character varying, p_address text, p_pincode character varying, p_location integer, p_state integer, p_active_status integer, p_role_id integer) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.user_table 
    SET 
        name = p_name, 
        mobile_num = p_mobile_num, 
        email = p_email, 
        address = p_address, 
        pincode = p_pincode, 
        location = p_location, 
        state = p_state, 
        active_status = p_active_status, 
        role_id = p_role_id
    WHERE user_id = p_user_id;
END;
$$;


ALTER FUNCTION admin_schema.update_user(p_user_id integer, p_name character varying, p_mobile_num character varying, p_email character varying, p_address text, p_pincode character varying, p_location integer, p_state integer, p_active_status integer, p_role_id integer) OWNER TO postgres;

--
-- Name: update_vehicle(integer, character varying, character varying, character varying, character varying, integer, integer, integer, date, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.update_vehicle(p_vehicle_id integer, p_vehicle_name character varying, p_vehicle_registration_no character varying, p_vehicle_manufacture_year character varying, p_vehicle_warranty character varying, p_vehicle_make integer, p_vehicle_model integer, p_vehicle_engine_type integer, p_vehicle_purchase_date date, p_vehicle_color character varying) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.master_vehicle_table
    SET 
        vehicle_name = p_vehicle_name,
        vehicle_registration_no = p_vehicle_registration_no,
        vehicle_manufacture_year = p_vehicle_manufacture_year,
        vehicle_warranty = p_vehicle_warranty,
        vehicle_make = p_vehicle_make,
        vehicle_model = p_vehicle_model,
        vehicle_engine_type = p_vehicle_engine_type,
        vehicle_purchase_date = p_vehicle_purchase_date,
        vehicle_color = p_vehicle_color,
        updated_at = CURRENT_TIMESTAMP
    WHERE vehicle_id = p_vehicle_id;
END;
$$;


ALTER FUNCTION admin_schema.update_vehicle(p_vehicle_id integer, p_vehicle_name character varying, p_vehicle_registration_no character varying, p_vehicle_manufacture_year character varying, p_vehicle_warranty character varying, p_vehicle_make integer, p_vehicle_model integer, p_vehicle_engine_type integer, p_vehicle_purchase_date date, p_vehicle_color character varying) OWNER TO postgres;

--
-- Name: upsert_business_branch(integer, integer, character varying, integer, integer, integer, character varying, character varying, character varying, integer, character varying, character varying, integer, character varying); Type: FUNCTION; Schema: admin_schema; Owner: postgres
--

CREATE FUNCTION admin_schema.upsert_business_branch(p_b_branch_id integer, p_bid integer, p_b_shop_name character varying, p_b_type_id integer, p_b_location integer, p_b_state integer, p_b_address character varying, p_b_email character varying, p_b_number character varying, p_b_mandi_id integer DEFAULT NULL::integer, p_b_gst_num character varying DEFAULT NULL::character varying, p_b_pan_num character varying DEFAULT NULL::character varying, p_b_privilege_user integer DEFAULT NULL::integer, p_b_established_year character varying DEFAULT NULL::character varying) RETURNS integer
    LANGUAGE plpgsql
    AS $$
DECLARE 
    new_branch_id INT;
BEGIN
    IF p_b_branch_id IS NULL THEN
        -- Insert new branch
        INSERT INTO admin_schema.business_branch_table (
            bid, b_shop_name, b_type_id, b_location, b_state, b_mandi_id, 
            b_address, b_email, b_number, b_gst_num, b_pan_num, 
            b_privilege_user, b_established_year, created_at, updated_at
        ) VALUES (
            p_bid, p_b_shop_name, p_b_type_id, p_b_location, p_b_state, p_b_mandi_id, 
            p_b_address, p_b_email, p_b_number, p_b_gst_num, p_b_pan_num, 
            p_b_privilege_user, p_b_established_year, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
        ) RETURNING b_branch_id INTO new_branch_id;

    ELSE
        -- Update existing branch
        UPDATE admin_schema.business_branch_table
        SET 
            bid = p_bid,
            b_shop_name = p_b_shop_name,
            b_type_id = p_b_type_id,
            b_location = p_b_location,
            b_state = p_b_state,
            b_mandi_id = p_b_mandi_id,
            b_address = p_b_address,
            b_email = p_b_email,
            b_number = p_b_number,
            b_gst_num = p_b_gst_num,
            b_pan_num = p_b_pan_num,
            b_privilege_user = p_b_privilege_user,
            b_established_year = p_b_established_year,
            updated_at = CURRENT_TIMESTAMP
        WHERE b_branch_id = p_b_branch_id
        RETURNING b_branch_id INTO new_branch_id;
    END IF;

    RETURN new_branch_id;
END;
$$;


ALTER FUNCTION admin_schema.upsert_business_branch(p_b_branch_id integer, p_bid integer, p_b_shop_name character varying, p_b_type_id integer, p_b_location integer, p_b_state integer, p_b_address character varying, p_b_email character varying, p_b_number character varying, p_b_mandi_id integer, p_b_gst_num character varying, p_b_pan_num character varying, p_b_privilege_user integer, p_b_established_year character varying) OWNER TO postgres;

--
-- Name: get_business_by_id(bigint); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_business_by_id(p_bid bigint) RETURNS TABLE(bid bigint, b_person_name text, b_registration_num text, b_owner_name text, active_status integer, created_at timestamp with time zone, updated_at timestamp with time zone, b_category_id integer, b_type_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT * FROM admin_schema.business_table WHERE bid = p_bid;
END;
$$;


ALTER FUNCTION public.get_business_by_id(p_bid bigint) OWNER TO postgres;

--
-- Name: get_businesses(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_businesses() RETURNS TABLE(bid bigint, b_person_name text, b_registration_num text, b_owner_name text, active_status integer, created_at timestamp with time zone, updated_at timestamp with time zone, b_category_id integer, b_type_id integer)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY SELECT * FROM admin_schema.business_table;
END;
$$;


ALTER FUNCTION public.get_businesses() OWNER TO postgres;

--
-- Name: insert_business(text, text, text, integer, integer, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_business(p_b_person_name text, p_b_registration_num text, p_b_owner_name text, p_active_status integer, p_b_category_id integer, p_b_type_id integer) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO admin_schema.business_table (
        b_person_name, 
        b_registration_num, 
        b_owner_name, 
        active_status, 
        b_category_id, 
        b_type_id
    ) VALUES (
        p_b_person_name, 
        p_b_registration_num, 
        p_b_owner_name, 
        p_active_status, 
        p_b_category_id, 
        p_b_type_id
    );
END;
$$;


ALTER FUNCTION public.insert_business(p_b_person_name text, p_b_registration_num text, p_b_owner_name text, p_active_status integer, p_b_category_id integer, p_b_type_id integer) OWNER TO postgres;

--
-- Name: insert_business_user(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.insert_business_user() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Insert user record when new business is created
    INSERT INTO admin_schema.business_user_table (b_id, user_name, password)
    SELECT NEW.bid, b_shop_name, b_number
    FROM admin_schema.business_branch_table
    WHERE bid = NEW.bid
    LIMIT 1; -- Ensure only one entry per business

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.insert_business_user() OWNER TO postgres;

--
-- Name: insert_master_product(integer, character varying, integer, character varying, character varying, character varying, character varying, character varying); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.insert_master_product(IN p_category_id integer, IN p_product_name character varying, IN p_status integer, IN p_image_path character varying DEFAULT NULL::character varying, IN p_regional_name1 character varying DEFAULT NULL::character varying, IN p_regional_name2 character varying DEFAULT NULL::character varying, IN p_regional_name3 character varying DEFAULT NULL::character varying, IN p_regional_name4 character varying DEFAULT NULL::character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO master_product (
        category_id, product_name, status, image_path, 
        regional_name1, regional_name2, regional_name3, regional_name4
    ) VALUES (
        p_category_id, p_product_name, p_status, p_image_path, 
        p_regional_name1, p_regional_name2, p_regional_name3, p_regional_name4
    );
END;
$$;


ALTER PROCEDURE public.insert_master_product(IN p_category_id integer, IN p_product_name character varying, IN p_status integer, IN p_image_path character varying, IN p_regional_name1 character varying, IN p_regional_name2 character varying, IN p_regional_name3 character varying, IN p_regional_name4 character varying) OWNER TO postgres;

--
-- Name: set_business_user_credentials(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.set_business_user_credentials() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Fetch b_shop_name and b_number from business_branch_table based on bid
    SELECT b_shop_name, b_number INTO NEW.user_name, NEW.password
    FROM admin_schema.business_branch_table
    WHERE bid = NEW.b_id
    LIMIT 1; -- Ensuring one row is picked

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.set_business_user_credentials() OWNER TO postgres;

--
-- Name: update_business(bigint, text, text, text, integer, integer, integer); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_business(p_bid bigint, p_b_person_name text, p_b_registration_num text, p_b_owner_name text, p_active_status integer, p_b_category_id integer, p_b_type_id integer) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    UPDATE admin_schema.business_table
    SET 
        b_person_name = p_b_person_name,
        b_registration_num = p_b_registration_num,
        b_owner_name = p_b_owner_name,
        active_status = p_active_status,
        b_category_id = p_b_category_id,
        b_type_id = p_b_type_id,
        updated_at = NOW()
    WHERE bid = p_bid;
END;
$$;


ALTER FUNCTION public.update_business(p_bid bigint, p_b_person_name text, p_b_registration_num text, p_b_owner_name text, p_active_status integer, p_b_category_id integer, p_b_type_id integer) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: business_branch_table; Type: TABLE; Schema: admin_schema; Owner: postgres
--

CREATE TABLE admin_schema.business_branch_table (
    b_branch_id integer NOT NULL,
    bid integer NOT NULL,
    b_shop_name character varying(255) NOT NULL,
    b_type_id integer NOT NULL,
    b_location integer NOT NULL,
    b_state integer NOT NULL,
    b_mandi_id integer,
    b_address character varying(255) NOT NULL,
    b_email character varying(50) NOT NULL,
    b_number character varying(10) NOT NULL,
    b_gst_num character varying(15),
    b_pan_num character varying(10),
    b_privilege_user integer,
    b_established_year character varying(4),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    active_status integer DEFAULT 0
);


ALTER TABLE admin_schema.business_branch_table OWNER TO postgres;

--
-- Name: business_branch_table_b_branch_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: postgres
--

CREATE SEQUENCE admin_schema.business_branch_table_b_branch_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.business_branch_table_b_branch_id_seq OWNER TO postgres;

--
-- Name: business_branch_table_b_branch_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: postgres
--

ALTER SEQUENCE admin_schema.business_branch_table_b_branch_id_seq OWNED BY admin_schema.business_branch_table.b_branch_id;


--
-- Name: business_category_table; Type: TABLE; Schema: admin_schema; Owner: postgres
--

CREATE TABLE admin_schema.business_category_table (
    b_category_id integer NOT NULL,
    b_category_name character varying(25) NOT NULL
);


ALTER TABLE admin_schema.business_category_table OWNER TO postgres;

--
-- Name: business_category_table_b_category_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: postgres
--

CREATE SEQUENCE admin_schema.business_category_table_b_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.business_category_table_b_category_id_seq OWNER TO postgres;

--
-- Name: business_category_table_b_category_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: postgres
--

ALTER SEQUENCE admin_schema.business_category_table_b_category_id_seq OWNED BY admin_schema.business_category_table.b_category_id;


--
-- Name: business_table; Type: TABLE; Schema: admin_schema; Owner: postgres
--

CREATE TABLE admin_schema.business_table (
    bid bigint NOT NULL,
    b_person_name character varying(255),
    b_registration_num character varying(50),
    b_owner_name character varying(255),
    active_status integer,
    created_at timestamp with time zone,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    b_category_id integer,
    b_type_id integer
);


ALTER TABLE admin_schema.business_table OWNER TO postgres;

--
-- Name: business_type_table; Type: TABLE; Schema: admin_schema; Owner: postgres
--

CREATE TABLE admin_schema.business_type_table (
    type_id integer NOT NULL,
    type_name character varying(50) NOT NULL
);


ALTER TABLE admin_schema.business_type_table OWNER TO postgres;

--
-- Name: business_type_table_type_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: postgres
--

CREATE SEQUENCE admin_schema.business_type_table_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.business_type_table_type_id_seq OWNER TO postgres;

--
-- Name: business_type_table_type_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: postgres
--

ALTER SEQUENCE admin_schema.business_type_table_type_id_seq OWNED BY admin_schema.business_type_table.type_id;


--
-- Name: business_user_table; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.business_user_table (
    b_id integer NOT NULL,
    user_name character varying(255) NOT NULL,
    password character varying(20) NOT NULL,
    active_status integer DEFAULT 1,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    password_reset_token character varying(255),
    is_locked boolean DEFAULT false
);


ALTER TABLE admin_schema.business_user_table OWNER TO admin;

--
-- Name: cash_payment_list; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.cash_payment_list (
    id integer NOT NULL,
    payment_type character varying(50) NOT NULL
);


ALTER TABLE admin_schema.cash_payment_list OWNER TO admin;

--
-- Name: cash_payment_list_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.cash_payment_list_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.cash_payment_list_id_seq OWNER TO admin;

--
-- Name: cash_payment_list_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.cash_payment_list_id_seq OWNED BY admin_schema.cash_payment_list.id;


--
-- Name: category_regional_name; Type: TABLE; Schema: admin_schema; Owner: postgres
--

CREATE TABLE admin_schema.category_regional_name (
    category_regional_id integer NOT NULL,
    language_id integer NOT NULL,
    category_id integer NOT NULL,
    category_regional_name character varying(100) NOT NULL
);


ALTER TABLE admin_schema.category_regional_name OWNER TO postgres;

--
-- Name: category_regional_name_category_regional_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: postgres
--

CREATE SEQUENCE admin_schema.category_regional_name_category_regional_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.category_regional_name_category_regional_id_seq OWNER TO postgres;

--
-- Name: category_regional_name_category_regional_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: postgres
--

ALTER SEQUENCE admin_schema.category_regional_name_category_regional_id_seq OWNED BY admin_schema.category_regional_name.category_regional_id;


--
-- Name: indian_driver_licenses_type; Type: TABLE; Schema: admin_schema; Owner: postgres
--

CREATE TABLE admin_schema.indian_driver_licenses_type (
    license_type_id integer NOT NULL,
    license_type_name character varying(100) NOT NULL,
    license_type_abbreviation character varying(20) NOT NULL,
    minimum_age integer NOT NULL,
    allowed_vehicles text NOT NULL,
    license_valid_duration integer NOT NULL,
    commercial_use boolean NOT NULL,
    national_permit_required boolean NOT NULL,
    medical_certificate_required boolean NOT NULL
);


ALTER TABLE admin_schema.indian_driver_licenses_type OWNER TO postgres;

--
-- Name: indian_driver_licenses_type_license_type_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: postgres
--

CREATE SEQUENCE admin_schema.indian_driver_licenses_type_license_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.indian_driver_licenses_type_license_type_id_seq OWNER TO postgres;

--
-- Name: indian_driver_licenses_type_license_type_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: postgres
--

ALTER SEQUENCE admin_schema.indian_driver_licenses_type_license_type_id_seq OWNED BY admin_schema.indian_driver_licenses_type.license_type_id;


--
-- Name: master_product_category_table; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.master_product_category_table (
    category_id integer NOT NULL,
    category_name character varying(255) NOT NULL,
    super_cat_id integer,
    img_path text,
    active_status integer DEFAULT 1,
    category_regional_id integer
);


ALTER TABLE admin_schema.master_product_category_table OWNER TO admin;

--
-- Name: master_category_table_category_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.master_category_table_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.master_category_table_category_id_seq OWNER TO admin;

--
-- Name: master_category_table_category_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.master_category_table_category_id_seq OWNED BY admin_schema.master_product_category_table.category_id;


--
-- Name: master_city; Type: TABLE; Schema: admin_schema; Owner: postgres
--

CREATE TABLE admin_schema.master_city (
    id integer NOT NULL,
    city_shortnames character varying(5) NOT NULL,
    city_name character varying(50) NOT NULL
);


ALTER TABLE admin_schema.master_city OWNER TO postgres;

--
-- Name: master_city_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: postgres
--

CREATE SEQUENCE admin_schema.master_city_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.master_city_id_seq OWNER TO postgres;

--
-- Name: master_city_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: postgres
--

ALTER SEQUENCE admin_schema.master_city_id_seq OWNED BY admin_schema.master_city.id;


--
-- Name: master_driver_table; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.master_driver_table (
    driver_id integer NOT NULL,
    driver_name character varying(255) NOT NULL,
    driver_license character varying(50) NOT NULL,
    driver_number character varying(15) NOT NULL,
    driver_address character varying(255),
    date_of_joining date,
    emergency_contact character varying(15),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    driver_dob date,
    driver_license_type integer
);


ALTER TABLE admin_schema.master_driver_table OWNER TO admin;

--
-- Name: master_driver_table_driver_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.master_driver_table_driver_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.master_driver_table_driver_id_seq OWNER TO admin;

--
-- Name: master_driver_table_driver_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.master_driver_table_driver_id_seq OWNED BY admin_schema.master_driver_table.driver_id;


--
-- Name: master_language; Type: TABLE; Schema: admin_schema; Owner: postgres
--

CREATE TABLE admin_schema.master_language (
    id integer NOT NULL,
    language_name character varying(50) NOT NULL
);


ALTER TABLE admin_schema.master_language OWNER TO postgres;

--
-- Name: master_language_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: postgres
--

CREATE SEQUENCE admin_schema.master_language_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.master_language_id_seq OWNER TO postgres;

--
-- Name: master_language_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: postgres
--

ALTER SEQUENCE admin_schema.master_language_id_seq OWNED BY admin_schema.master_language.id;


--
-- Name: master_location; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.master_location (
    id integer NOT NULL,
    location character varying(50),
    city_shortnames integer,
    state integer
);


ALTER TABLE admin_schema.master_location OWNER TO admin;

--
-- Name: master_location_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.master_location_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.master_location_id_seq OWNER TO admin;

--
-- Name: master_location_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.master_location_id_seq OWNED BY admin_schema.master_location.id;


--
-- Name: master_mandi_table; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.master_mandi_table (
    mandi_id integer NOT NULL,
    mandi_location character varying(255) NOT NULL,
    mandi_incharge character varying(255) NOT NULL,
    mandi_incharge_num character varying(15) NOT NULL,
    mandi_pincode character varying(6) NOT NULL,
    mandi_address text,
    mandi_state integer,
    mandi_name character varying(255),
    mandi_shortnames character varying(10)
);


ALTER TABLE admin_schema.master_mandi_table OWNER TO admin;

--
-- Name: master_mandi_table_mandi_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.master_mandi_table_mandi_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.master_mandi_table_mandi_id_seq OWNER TO admin;

--
-- Name: master_mandi_table_mandi_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.master_mandi_table_mandi_id_seq OWNED BY admin_schema.master_mandi_table.mandi_id;


--
-- Name: master_product; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.master_product (
    product_id bigint NOT NULL,
    category_id integer,
    product_name character varying(100),
    image_path character varying(100),
    active_status integer DEFAULT 1,
    product_regional_id integer
);


ALTER TABLE admin_schema.master_product OWNER TO admin;

--
-- Name: master_product_utf8_product_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.master_product_utf8_product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.master_product_utf8_product_id_seq OWNER TO admin;

--
-- Name: master_product_utf8_product_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.master_product_utf8_product_id_seq OWNED BY admin_schema.master_product.product_id;


--
-- Name: master_states; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.master_states (
    id integer NOT NULL,
    state character varying(50),
    state_shortnames character varying(5)
);


ALTER TABLE admin_schema.master_states OWNER TO admin;

--
-- Name: master_states_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.master_states_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.master_states_id_seq OWNER TO admin;

--
-- Name: master_states_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.master_states_id_seq OWNED BY admin_schema.master_states.id;


--
-- Name: master_vehicle_table; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.master_vehicle_table (
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


ALTER TABLE admin_schema.master_vehicle_table OWNER TO admin;

--
-- Name: master_vechicle_table_vehicle_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.master_vechicle_table_vehicle_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.master_vechicle_table_vehicle_id_seq OWNER TO admin;

--
-- Name: master_vechicle_table_vehicle_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.master_vechicle_table_vehicle_id_seq OWNED BY admin_schema.master_vehicle_table.vehicle_id;


--
-- Name: master_violation_table; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.master_violation_table (
    violation_name text,
    level_of_serious text,
    id integer NOT NULL
);


ALTER TABLE admin_schema.master_violation_table OWNER TO admin;

--
-- Name: master_violation_table_new_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.master_violation_table_new_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.master_violation_table_new_id_seq OWNER TO admin;

--
-- Name: master_violation_table_new_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.master_violation_table_new_id_seq OWNED BY admin_schema.master_violation_table.id;


--
-- Name: mode_of_payments_list; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.mode_of_payments_list (
    id integer NOT NULL,
    payment_mode character varying(50) NOT NULL
);


ALTER TABLE admin_schema.mode_of_payments_list OWNER TO admin;

--
-- Name: mode_of_payments_list_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.mode_of_payments_list_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.mode_of_payments_list_id_seq OWNER TO admin;

--
-- Name: mode_of_payments_list_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.mode_of_payments_list_id_seq OWNED BY admin_schema.mode_of_payments_list.id;


--
-- Name: order_status_table; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.order_status_table (
    order_status_id integer NOT NULL,
    order_status character varying(50) NOT NULL
);


ALTER TABLE admin_schema.order_status_table OWNER TO admin;

--
-- Name: order_status_table_order_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.order_status_table_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.order_status_table_order_id_seq OWNER TO admin;

--
-- Name: order_status_table_order_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.order_status_table_order_id_seq OWNED BY admin_schema.order_status_table.order_status_id;


--
-- Name: permission_audit_log; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.permission_audit_log (
    log_id integer NOT NULL,
    changed_by integer,
    changed_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    role_id integer,
    table_name character varying(100),
    action character varying(50)
);


ALTER TABLE admin_schema.permission_audit_log OWNER TO admin;

--
-- Name: permission_audit_log_log_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.permission_audit_log_log_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.permission_audit_log_log_id_seq OWNER TO admin;

--
-- Name: permission_audit_log_log_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.permission_audit_log_log_id_seq OWNED BY admin_schema.permission_audit_log.log_id;


--
-- Name: product_regional_name; Type: TABLE; Schema: admin_schema; Owner: postgres
--

CREATE TABLE admin_schema.product_regional_name (
    product_regional_id integer NOT NULL,
    language_id integer NOT NULL,
    product_id integer NOT NULL,
    product_regional_name character varying(100) NOT NULL
);


ALTER TABLE admin_schema.product_regional_name OWNER TO postgres;

--
-- Name: product_regional_name_product_regional_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: postgres
--

CREATE SEQUENCE admin_schema.product_regional_name_product_regional_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.product_regional_name_product_regional_id_seq OWNER TO postgres;

--
-- Name: product_regional_name_product_regional_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: postgres
--

ALTER SEQUENCE admin_schema.product_regional_name_product_regional_id_seq OWNED BY admin_schema.product_regional_name.product_regional_id;


--
-- Name: roles_table; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.roles_table (
    role_id integer NOT NULL,
    role_name character varying(50) NOT NULL
);


ALTER TABLE admin_schema.roles_table OWNER TO admin;

--
-- Name: roles_table_role_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.roles_table_role_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.roles_table_role_id_seq OWNER TO admin;

--
-- Name: roles_table_role_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.roles_table_role_id_seq OWNED BY admin_schema.roles_table.role_id;


--
-- Name: units_table; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.units_table (
    id integer NOT NULL,
    unit_name character varying(50) NOT NULL
);


ALTER TABLE admin_schema.units_table OWNER TO admin;

--
-- Name: units_table_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.units_table_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.units_table_id_seq OWNER TO admin;

--
-- Name: units_table_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.units_table_id_seq OWNED BY admin_schema.units_table.id;


--
-- Name: user_table; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.user_table (
    user_id integer NOT NULL,
    name character varying(255) NOT NULL,
    mobile_num character varying(15) NOT NULL,
    email character varying(255),
    address text,
    pincode character varying(10),
    location integer,
    state integer,
    active_status integer DEFAULT 0,
    role_id integer DEFAULT 5
);


ALTER TABLE admin_schema.user_table OWNER TO admin;

--
-- Name: user_table_user_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.user_table_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.user_table_user_id_seq OWNER TO admin;

--
-- Name: user_table_user_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.user_table_user_id_seq OWNED BY admin_schema.user_table.user_id;


--
-- Name: users; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.users (
    user_id integer NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    role_id integer
);


ALTER TABLE admin_schema.users OWNER TO admin;

--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.users_user_id_seq OWNER TO admin;

--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.users_user_id_seq OWNED BY admin_schema.users.user_id;


--
-- Name: vehicle_engine_type; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.vehicle_engine_type (
    id integer NOT NULL,
    engine_type character varying(50)
);


ALTER TABLE admin_schema.vehicle_engine_type OWNER TO admin;

--
-- Name: vehicle_engine_type_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.vehicle_engine_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.vehicle_engine_type_id_seq OWNER TO admin;

--
-- Name: vehicle_engine_type_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.vehicle_engine_type_id_seq OWNED BY admin_schema.vehicle_engine_type.id;


--
-- Name: vehicle_insurance_table; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.vehicle_insurance_table (
    vehicle_insurance_id integer NOT NULL,
    policy_number character varying(100) NOT NULL,
    insurance_provider integer NOT NULL,
    policy_type integer NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    premium_amount numeric(10,2) NOT NULL,
    insured_amount numeric(12,2) NOT NULL,
    chassis_number character varying(100),
    claims_made integer DEFAULT 0,
    last_claim_date date,
    vehicle_id integer,
    renewal_date date,
    payment_status integer,
    policy_holder_name character varying(100),
    policy_holder_contact character varying(100)
);


ALTER TABLE admin_schema.vehicle_insurance_table OWNER TO admin;

--
-- Name: vehicle_insurance_table_vehicle_insurance_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.vehicle_insurance_table_vehicle_insurance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.vehicle_insurance_table_vehicle_insurance_id_seq OWNER TO admin;

--
-- Name: vehicle_insurance_table_vehicle_insurance_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.vehicle_insurance_table_vehicle_insurance_id_seq OWNED BY admin_schema.vehicle_insurance_table.vehicle_insurance_id;


--
-- Name: vehicle_make; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.vehicle_make (
    id integer NOT NULL,
    make character varying(50)
);


ALTER TABLE admin_schema.vehicle_make OWNER TO admin;

--
-- Name: vehicle_make_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.vehicle_make_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.vehicle_make_id_seq OWNER TO admin;

--
-- Name: vehicle_make_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.vehicle_make_id_seq OWNED BY admin_schema.vehicle_make.id;


--
-- Name: vehicle_model; Type: TABLE; Schema: admin_schema; Owner: admin
--

CREATE TABLE admin_schema.vehicle_model (
    id integer NOT NULL,
    model character varying(50)
);


ALTER TABLE admin_schema.vehicle_model OWNER TO admin;

--
-- Name: vehicle_model_id_seq; Type: SEQUENCE; Schema: admin_schema; Owner: admin
--

CREATE SEQUENCE admin_schema.vehicle_model_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE admin_schema.vehicle_model_id_seq OWNER TO admin;

--
-- Name: vehicle_model_id_seq; Type: SEQUENCE OWNED BY; Schema: admin_schema; Owner: admin
--

ALTER SEQUENCE admin_schema.vehicle_model_id_seq OWNED BY admin_schema.vehicle_model.id;


--
-- Name: daily_price_update; Type: TABLE; Schema: business_schema; Owner: postgres
--

CREATE TABLE business_schema.daily_price_update (
    product_id integer,
    price numeric(10,2),
    unit_id integer,
    wholeseller_id integer,
    currency character varying(10) DEFAULT 'INR'::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    remarks character varying(255)
);


ALTER TABLE business_schema.daily_price_update OWNER TO postgres;

--
-- Name: invoice_details_table; Type: TABLE; Schema: business_schema; Owner: postgres
--

CREATE TABLE business_schema.invoice_details_table (
    id bigint NOT NULL,
    invoice_id bigint,
    order_item_id bigint,
    quantity numeric(10,2) NOT NULL,
    retailer_approval_status integer DEFAULT 0,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE business_schema.invoice_details_table OWNER TO postgres;

--
-- Name: invoice_details_table_id_seq; Type: SEQUENCE; Schema: business_schema; Owner: postgres
--

CREATE SEQUENCE business_schema.invoice_details_table_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE business_schema.invoice_details_table_id_seq OWNER TO postgres;

--
-- Name: invoice_details_table_id_seq; Type: SEQUENCE OWNED BY; Schema: business_schema; Owner: postgres
--

ALTER SEQUENCE business_schema.invoice_details_table_id_seq OWNED BY business_schema.invoice_details_table.id;


--
-- Name: invoice_table; Type: TABLE; Schema: business_schema; Owner: postgres
--

CREATE TABLE business_schema.invoice_table (
    id integer NOT NULL,
    invoice_number character varying(50),
    order_id bigint,
    total_amount numeric(10,2) NOT NULL,
    discount_amount numeric(10,2) DEFAULT 0,
    invoice_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    due_date date,
    pay_mode integer,
    pay_type integer,
    tax_amount numeric(10,2),
    payment_date date,
    currency character varying(10) DEFAULT 'INR'::character varying,
    final_amount numeric(10,2) GENERATED ALWAYS AS (((total_amount - discount_amount) + tax_amount)) STORED
);


ALTER TABLE business_schema.invoice_table OWNER TO postgres;

--
-- Name: invoice_table_id_seq; Type: SEQUENCE; Schema: business_schema; Owner: postgres
--

CREATE SEQUENCE business_schema.invoice_table_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE business_schema.invoice_table_id_seq OWNER TO postgres;

--
-- Name: invoice_table_id_seq; Type: SEQUENCE OWNED BY; Schema: business_schema; Owner: postgres
--

ALTER SEQUENCE business_schema.invoice_table_id_seq OWNED BY business_schema.invoice_table.id;


--
-- Name: mode_of_payment; Type: TABLE; Schema: business_schema; Owner: admin
--

CREATE TABLE business_schema.mode_of_payment (
    id integer NOT NULL,
    pay_mode integer,
    pay_type integer
);


ALTER TABLE business_schema.mode_of_payment OWNER TO admin;

--
-- Name: mode_of_payment_id_seq; Type: SEQUENCE; Schema: business_schema; Owner: admin
--

CREATE SEQUENCE business_schema.mode_of_payment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE business_schema.mode_of_payment_id_seq OWNER TO admin;

--
-- Name: mode_of_payment_id_seq; Type: SEQUENCE OWNED BY; Schema: business_schema; Owner: admin
--

ALTER SEQUENCE business_schema.mode_of_payment_id_seq OWNED BY business_schema.mode_of_payment.id;


--
-- Name: order_history_table; Type: TABLE; Schema: business_schema; Owner: postgres
--

CREATE TABLE business_schema.order_history_table (
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


ALTER TABLE business_schema.order_history_table OWNER TO postgres;

--
-- Name: order_history_table_history_id_seq; Type: SEQUENCE; Schema: business_schema; Owner: postgres
--

CREATE SEQUENCE business_schema.order_history_table_history_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE business_schema.order_history_table_history_id_seq OWNER TO postgres;

--
-- Name: order_history_table_history_id_seq; Type: SEQUENCE OWNED BY; Schema: business_schema; Owner: postgres
--

ALTER SEQUENCE business_schema.order_history_table_history_id_seq OWNED BY business_schema.order_history_table.history_id;


--
-- Name: order_item_table; Type: TABLE; Schema: business_schema; Owner: postgres
--

CREATE TABLE business_schema.order_item_table (
    order_item_id bigint NOT NULL,
    order_id bigint,
    product_id bigint,
    quantity numeric(10,2),
    unit_id integer,
    discount_amount numeric(10,2),
    tax_amount numeric(10,2)
);


ALTER TABLE business_schema.order_item_table OWNER TO postgres;

--
-- Name: order_item_table_product_order_id_seq; Type: SEQUENCE; Schema: business_schema; Owner: postgres
--

CREATE SEQUENCE business_schema.order_item_table_product_order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE business_schema.order_item_table_product_order_id_seq OWNER TO postgres;

--
-- Name: order_item_table_product_order_id_seq; Type: SEQUENCE OWNED BY; Schema: business_schema; Owner: postgres
--

ALTER SEQUENCE business_schema.order_item_table_product_order_id_seq OWNED BY business_schema.order_item_table.order_item_id;


--
-- Name: order_table; Type: TABLE; Schema: business_schema; Owner: postgres
--

CREATE TABLE business_schema.order_table (
    order_id bigint NOT NULL,
    date_of_order timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    order_status integer,
    expected_delivery_date date,
    actual_delivery_date date,
    retailer_id integer,
    wholeseller_id integer,
    location_id integer,
    state_id integer,
    pincode character varying(6),
    address character varying(200),
    total_order_amount numeric(10,2) DEFAULT 0,
    discount_amount numeric(10,2),
    tax_amount numeric(10,2),
    final_amount numeric(10,2),
    created_by timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE business_schema.order_table OWNER TO postgres;

--
-- Name: order_table_order_id_seq; Type: SEQUENCE; Schema: business_schema; Owner: postgres
--

CREATE SEQUENCE business_schema.order_table_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE business_schema.order_table_order_id_seq OWNER TO postgres;

--
-- Name: order_table_order_id_seq; Type: SEQUENCE OWNED BY; Schema: business_schema; Owner: postgres
--

ALTER SEQUENCE business_schema.order_table_order_id_seq OWNED BY business_schema.order_table.order_id;


--
-- Name: stock_table; Type: TABLE; Schema: business_schema; Owner: postgres
--

CREATE TABLE business_schema.stock_table (
    stock_id integer NOT NULL,
    product_id integer NOT NULL,
    stock_in numeric(10,2) DEFAULT 0,
    stock_left numeric(10,2) DEFAULT 0,
    stock_sold numeric(10,2) DEFAULT 0,
    date_of_order date DEFAULT CURRENT_TIMESTAMP,
    warehouse_id integer,
    updatedat timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    createdat timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    supplier_id integer,
    expiry_date date,
    manufacturing_date date,
    minimum_stock_level numeric(10,2),
    maximum_stock_level numeric(10,2)
);


ALTER TABLE business_schema.stock_table OWNER TO postgres;

--
-- Name: stock_table_stock_id_seq; Type: SEQUENCE; Schema: business_schema; Owner: postgres
--

CREATE SEQUENCE business_schema.stock_table_stock_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE business_schema.stock_table_stock_id_seq OWNER TO postgres;

--
-- Name: stock_table_stock_id_seq; Type: SEQUENCE OWNED BY; Schema: business_schema; Owner: postgres
--

ALTER SEQUENCE business_schema.stock_table_stock_id_seq OWNED BY business_schema.stock_table.stock_id;


--
-- Name: warehouse_list; Type: TABLE; Schema: business_schema; Owner: postgres
--

CREATE TABLE business_schema.warehouse_list (
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


ALTER TABLE business_schema.warehouse_list OWNER TO postgres;

--
-- Name: warehouse_list_warehouse_id_seq; Type: SEQUENCE; Schema: business_schema; Owner: postgres
--

CREATE SEQUENCE business_schema.warehouse_list_warehouse_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE business_schema.warehouse_list_warehouse_id_seq OWNER TO postgres;

--
-- Name: warehouse_list_warehouse_id_seq; Type: SEQUENCE OWNED BY; Schema: business_schema; Owner: postgres
--

ALTER SEQUENCE business_schema.warehouse_list_warehouse_id_seq OWNED BY business_schema.warehouse_list.warehouse_id;


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
-- Name: product_reviews; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_reviews (
    review_id integer NOT NULL,
    product_id integer NOT NULL,
    user_id integer NOT NULL,
    rating integer,
    review_text text,
    CONSTRAINT product_reviews_rating_check CHECK (((rating >= 1) AND (rating <= 5)))
);


ALTER TABLE public.product_reviews OWNER TO postgres;

--
-- Name: product_reviews_review_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_reviews_review_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_reviews_review_id_seq OWNER TO postgres;

--
-- Name: product_reviews_review_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_reviews_review_id_seq OWNED BY public.product_reviews.review_id;


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
-- Name: business_branch_table b_branch_id; Type: DEFAULT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.business_branch_table ALTER COLUMN b_branch_id SET DEFAULT nextval('admin_schema.business_branch_table_b_branch_id_seq'::regclass);


--
-- Name: business_category_table b_category_id; Type: DEFAULT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.business_category_table ALTER COLUMN b_category_id SET DEFAULT nextval('admin_schema.business_category_table_b_category_id_seq'::regclass);


--
-- Name: business_type_table type_id; Type: DEFAULT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.business_type_table ALTER COLUMN type_id SET DEFAULT nextval('admin_schema.business_type_table_type_id_seq'::regclass);


--
-- Name: cash_payment_list id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.cash_payment_list ALTER COLUMN id SET DEFAULT nextval('admin_schema.cash_payment_list_id_seq'::regclass);


--
-- Name: category_regional_name category_regional_id; Type: DEFAULT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.category_regional_name ALTER COLUMN category_regional_id SET DEFAULT nextval('admin_schema.category_regional_name_category_regional_id_seq'::regclass);


--
-- Name: indian_driver_licenses_type license_type_id; Type: DEFAULT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.indian_driver_licenses_type ALTER COLUMN license_type_id SET DEFAULT nextval('admin_schema.indian_driver_licenses_type_license_type_id_seq'::regclass);


--
-- Name: master_city id; Type: DEFAULT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.master_city ALTER COLUMN id SET DEFAULT nextval('admin_schema.master_city_id_seq'::regclass);


--
-- Name: master_driver_table driver_id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_driver_table ALTER COLUMN driver_id SET DEFAULT nextval('admin_schema.master_driver_table_driver_id_seq'::regclass);


--
-- Name: master_language id; Type: DEFAULT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.master_language ALTER COLUMN id SET DEFAULT nextval('admin_schema.master_language_id_seq'::regclass);


--
-- Name: master_location id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_location ALTER COLUMN id SET DEFAULT nextval('admin_schema.master_location_id_seq'::regclass);


--
-- Name: master_mandi_table mandi_id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_mandi_table ALTER COLUMN mandi_id SET DEFAULT nextval('admin_schema.master_mandi_table_mandi_id_seq'::regclass);


--
-- Name: master_product product_id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_product ALTER COLUMN product_id SET DEFAULT nextval('admin_schema.master_product_utf8_product_id_seq'::regclass);


--
-- Name: master_product_category_table category_id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_product_category_table ALTER COLUMN category_id SET DEFAULT nextval('admin_schema.master_category_table_category_id_seq'::regclass);


--
-- Name: master_states id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_states ALTER COLUMN id SET DEFAULT nextval('admin_schema.master_states_id_seq'::regclass);


--
-- Name: master_vehicle_table vehicle_id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_vehicle_table ALTER COLUMN vehicle_id SET DEFAULT nextval('admin_schema.master_vechicle_table_vehicle_id_seq'::regclass);


--
-- Name: master_violation_table id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_violation_table ALTER COLUMN id SET DEFAULT nextval('admin_schema.master_violation_table_new_id_seq'::regclass);


--
-- Name: mode_of_payments_list id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.mode_of_payments_list ALTER COLUMN id SET DEFAULT nextval('admin_schema.mode_of_payments_list_id_seq'::regclass);


--
-- Name: order_status_table order_status_id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.order_status_table ALTER COLUMN order_status_id SET DEFAULT nextval('admin_schema.order_status_table_order_id_seq'::regclass);


--
-- Name: permission_audit_log log_id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.permission_audit_log ALTER COLUMN log_id SET DEFAULT nextval('admin_schema.permission_audit_log_log_id_seq'::regclass);


--
-- Name: product_regional_name product_regional_id; Type: DEFAULT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.product_regional_name ALTER COLUMN product_regional_id SET DEFAULT nextval('admin_schema.product_regional_name_product_regional_id_seq'::regclass);


--
-- Name: roles_table role_id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.roles_table ALTER COLUMN role_id SET DEFAULT nextval('admin_schema.roles_table_role_id_seq'::regclass);


--
-- Name: units_table id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.units_table ALTER COLUMN id SET DEFAULT nextval('admin_schema.units_table_id_seq'::regclass);


--
-- Name: user_table user_id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.user_table ALTER COLUMN user_id SET DEFAULT nextval('admin_schema.user_table_user_id_seq'::regclass);


--
-- Name: users user_id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.users ALTER COLUMN user_id SET DEFAULT nextval('admin_schema.users_user_id_seq'::regclass);


--
-- Name: vehicle_engine_type id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.vehicle_engine_type ALTER COLUMN id SET DEFAULT nextval('admin_schema.vehicle_engine_type_id_seq'::regclass);


--
-- Name: vehicle_insurance_table vehicle_insurance_id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.vehicle_insurance_table ALTER COLUMN vehicle_insurance_id SET DEFAULT nextval('admin_schema.vehicle_insurance_table_vehicle_insurance_id_seq'::regclass);


--
-- Name: vehicle_make id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.vehicle_make ALTER COLUMN id SET DEFAULT nextval('admin_schema.vehicle_make_id_seq'::regclass);


--
-- Name: vehicle_model id; Type: DEFAULT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.vehicle_model ALTER COLUMN id SET DEFAULT nextval('admin_schema.vehicle_model_id_seq'::regclass);


--
-- Name: invoice_details_table id; Type: DEFAULT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.invoice_details_table ALTER COLUMN id SET DEFAULT nextval('business_schema.invoice_details_table_id_seq'::regclass);


--
-- Name: invoice_table id; Type: DEFAULT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.invoice_table ALTER COLUMN id SET DEFAULT nextval('business_schema.invoice_table_id_seq'::regclass);


--
-- Name: mode_of_payment id; Type: DEFAULT; Schema: business_schema; Owner: admin
--

ALTER TABLE ONLY business_schema.mode_of_payment ALTER COLUMN id SET DEFAULT nextval('business_schema.mode_of_payment_id_seq'::regclass);


--
-- Name: order_history_table history_id; Type: DEFAULT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_history_table ALTER COLUMN history_id SET DEFAULT nextval('business_schema.order_history_table_history_id_seq'::regclass);


--
-- Name: order_item_table order_item_id; Type: DEFAULT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_item_table ALTER COLUMN order_item_id SET DEFAULT nextval('business_schema.order_item_table_product_order_id_seq'::regclass);


--
-- Name: order_table order_id; Type: DEFAULT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_table ALTER COLUMN order_id SET DEFAULT nextval('business_schema.order_table_order_id_seq'::regclass);


--
-- Name: stock_table stock_id; Type: DEFAULT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.stock_table ALTER COLUMN stock_id SET DEFAULT nextval('business_schema.stock_table_stock_id_seq'::regclass);


--
-- Name: warehouse_list warehouse_id; Type: DEFAULT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.warehouse_list ALTER COLUMN warehouse_id SET DEFAULT nextval('business_schema.warehouse_list_warehouse_id_seq'::regclass);


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
-- Name: product_price_table id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_price_table ALTER COLUMN id SET DEFAULT nextval('public.product_price_table_id_seq'::regclass);


--
-- Name: product_reviews review_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_reviews ALTER COLUMN review_id SET DEFAULT nextval('public.product_reviews_review_id_seq'::regclass);


--
-- Name: retailer_rating_table id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.retailer_rating_table ALTER COLUMN id SET DEFAULT nextval('public.retailer_rating_table_id_seq'::regclass);


--
-- Name: transporter_table transporter_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transporter_table ALTER COLUMN transporter_id SET DEFAULT nextval('public.transporter_table_transporter_id_seq'::regclass);


--
-- Name: type_of_insurance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.type_of_insurance ALTER COLUMN id SET DEFAULT nextval('public.type_of_insurance_id_seq'::regclass);


--
-- Name: user_bank_details id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_bank_details ALTER COLUMN id SET DEFAULT nextval('public.user_bank_details_id_seq'::regclass);


--
-- Name: wholeseller_rating_table id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wholeseller_rating_table ALTER COLUMN id SET DEFAULT nextval('public.wholesaler_rating_table_id_seq'::regclass);


--
-- Data for Name: business_branch_table; Type: TABLE DATA; Schema: admin_schema; Owner: postgres
--

INSERT INTO admin_schema.business_branch_table VALUES (4, 104, 'JohnShop', 1, 1, 1, NULL, '123 Street', 'john@example.com', '9876543210', NULL, NULL, NULL, NULL, '2025-03-27 11:55:01.194804', '2025-03-27 11:55:01.194804', 0);


--
-- Data for Name: business_category_table; Type: TABLE DATA; Schema: admin_schema; Owner: postgres
--



--
-- Data for Name: business_table; Type: TABLE DATA; Schema: admin_schema; Owner: postgres
--

INSERT INTO admin_schema.business_table VALUES (5, 'FreshMarket', NULL, 'UNKNOWN', 1, '2025-03-25 20:55:46.181233+05:30', '2025-03-25 20:55:46.181233+05:30', NULL, NULL);
INSERT INTO admin_schema.business_table VALUES (6, 'TradeHub', NULL, 'UNKNOWN', 1, '2025-03-25 20:55:46.181233+05:30', '2025-03-25 20:55:46.181233+05:30', NULL, NULL);
INSERT INTO admin_schema.business_table VALUES (12, 'SuperMart', NULL, 'UNKNOWN', 1, '2025-03-25 20:55:46.181233+05:30', '2025-03-25 20:55:46.181233+05:30', NULL, NULL);
INSERT INTO admin_schema.business_table VALUES (101, 'John Doe', 'REG123', 'Jane Doe', 1, NULL, '2025-03-27 11:46:40.100715+05:30', NULL, NULL);
INSERT INTO admin_schema.business_table VALUES (103, 'John Doe', 'REG789', 'Jane Doe', 1, NULL, '2025-03-27 11:51:51.001148+05:30', NULL, NULL);
INSERT INTO admin_schema.business_table VALUES (104, 'John Doe', 'REG104', 'Jane Doe', 1, NULL, '2025-03-27 11:54:56.79385+05:30', NULL, NULL);


--
-- Data for Name: business_type_table; Type: TABLE DATA; Schema: admin_schema; Owner: postgres
--

INSERT INTO admin_schema.business_type_table VALUES (1, 'retailer');
INSERT INTO admin_schema.business_type_table VALUES (2, 'wholeseller_admin');
INSERT INTO admin_schema.business_type_table VALUES (3, 'wholeseller_user');


--
-- Data for Name: business_user_table; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.business_user_table VALUES (104, 'JohnShop', '9876543210', 1, '2025-03-27 11:55:01.194804', '2025-03-27 11:55:01.194804', NULL, false);


--
-- Data for Name: cash_payment_list; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.cash_payment_list VALUES (1, 'Online Banking');
INSERT INTO admin_schema.cash_payment_list VALUES (2, 'UPI Payment');
INSERT INTO admin_schema.cash_payment_list VALUES (3, 'Mobile Wallets');
INSERT INTO admin_schema.cash_payment_list VALUES (4, 'Credit/Debit Cards');
INSERT INTO admin_schema.cash_payment_list VALUES (5, 'Credit/Pay after delivery');


--
-- Data for Name: category_regional_name; Type: TABLE DATA; Schema: admin_schema; Owner: postgres
--



--
-- Data for Name: indian_driver_licenses_type; Type: TABLE DATA; Schema: admin_schema; Owner: postgres
--



--
-- Data for Name: master_city; Type: TABLE DATA; Schema: admin_schema; Owner: postgres
--



--
-- Data for Name: master_driver_table; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.master_driver_table VALUES (2, 'John Updated', 'DL12345', '9876543210', '123 Main St', '2023-01-01', '9876543211', '2025-02-07 11:31:28.041893', '1989-05-20', NULL);


--
-- Data for Name: master_language; Type: TABLE DATA; Schema: admin_schema; Owner: postgres
--



--
-- Data for Name: master_location; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.master_location VALUES (1, 'Chennai', NULL, NULL);
INSERT INTO admin_schema.master_location VALUES (2, 'coimbatore', NULL, NULL);


--
-- Data for Name: master_mandi_table; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.master_mandi_table VALUES (2, 'Location 1', 'John Doe', '9876543210', '123456', 'Address 1', 1, NULL, NULL);


--
-- Data for Name: master_product; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.master_product VALUES (1, 4, 'New Product', NULL, 1, NULL);
INSERT INTO admin_schema.master_product VALUES (2, 5, 'product 2', NULL, 1, NULL);
INSERT INTO admin_schema.master_product VALUES (3, 6, 'product 2', NULL, 1, NULL);


--
-- Data for Name: master_product_category_table; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.master_product_category_table VALUES (5, 'Roots', 2, NULL, 1, NULL);
INSERT INTO admin_schema.master_product_category_table VALUES (6, 'Vegetables', 3, NULL, 1, NULL);
INSERT INTO admin_schema.master_product_category_table VALUES (8, 'Fruits', 4, NULL, 1, NULL);
INSERT INTO admin_schema.master_product_category_table VALUES (4, 'Leafy', 0, NULL, 1, NULL);
INSERT INTO admin_schema.master_product_category_table VALUES (9, 'Organic', NULL, NULL, 1, NULL);
INSERT INTO admin_schema.master_product_category_table VALUES (10, 'Organic Vegetables', 9, NULL, 1, NULL);
INSERT INTO admin_schema.master_product_category_table VALUES (11, 'Organic Fruits', 9, NULL, 1, NULL);


--
-- Data for Name: master_states; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.master_states VALUES (1, 'tamilnadu', NULL);
INSERT INTO admin_schema.master_states VALUES (2, 'kerla', NULL);
INSERT INTO admin_schema.master_states VALUES (3, 'West bengal', NULL);


--
-- Data for Name: master_vehicle_table; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.master_vehicle_table VALUES (5, 102, 'Honda City', '2021', '3 years', 2, 3, 'TN10XY5678', 1, '2021-08-22', 'Black', 'Extra2', 'Info2', 'Details2', '2025-02-07 11:30:10.747771', '2025-02-07 11:30:10.747771', 202, NULL, NULL, NULL);
INSERT INTO admin_schema.master_vehicle_table VALUES (6, 103, 'Ford Endeavour', '2023', '7 years', 3, 1, 'TN22PQ9876', 2, '2023-02-10', 'Blue', 'Extra3', 'Info3', 'Details3', '2025-02-07 11:30:10.747771', '2025-02-07 11:30:10.747771', 203, NULL, NULL, NULL);


--
-- Data for Name: master_violation_table; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.master_violation_table VALUES ('Hit and run', 'High', 1);
INSERT INTO admin_schema.master_violation_table VALUES ('drunk and drive', 'medium', 2);


--
-- Data for Name: mode_of_payments_list; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.mode_of_payments_list VALUES (1, 'credit_payment');
INSERT INTO admin_schema.mode_of_payments_list VALUES (2, 'cash_payment');


--
-- Data for Name: order_status_table; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.order_status_table VALUES (1, 'Processing');
INSERT INTO admin_schema.order_status_table VALUES (2, 'Confirmed');
INSERT INTO admin_schema.order_status_table VALUES (3, 'Payment');
INSERT INTO admin_schema.order_status_table VALUES (4, 'Out for Delivery');
INSERT INTO admin_schema.order_status_table VALUES (5, 'Successful');
INSERT INTO admin_schema.order_status_table VALUES (6, 'Cancellation');
INSERT INTO admin_schema.order_status_table VALUES (7, 'Returned');
INSERT INTO admin_schema.order_status_table VALUES (8, 'Processing');
INSERT INTO admin_schema.order_status_table VALUES (9, 'return');


--
-- Data for Name: permission_audit_log; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.permission_audit_log VALUES (2, 1, '2025-03-18 17:48:53.421015', 4, 'invoice_table', 'Updated Permissions');


--
-- Data for Name: product_regional_name; Type: TABLE DATA; Schema: admin_schema; Owner: postgres
--



--
-- Data for Name: roles_table; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.roles_table VALUES (1, 'Admin');
INSERT INTO admin_schema.roles_table VALUES (2, 'Wholeseller_Admin');
INSERT INTO admin_schema.roles_table VALUES (3, 'Retailer');
INSERT INTO admin_schema.roles_table VALUES (4, 'Wholeseller_User');
INSERT INTO admin_schema.roles_table VALUES (5, 'User');


--
-- Data for Name: units_table; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.units_table VALUES (1, 'kg');
INSERT INTO admin_schema.units_table VALUES (2, 'liter');
INSERT INTO admin_schema.units_table VALUES (3, 'tonnes');
INSERT INTO admin_schema.units_table VALUES (4, 'pieces');
INSERT INTO admin_schema.units_table VALUES (5, 'grams');


--
-- Data for Name: user_table; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.user_table VALUES (12, 'John Doe', '9876543210', 'johndoe@example.com', '123 Main St, City', '123456', 1, 1, 1, 1);
INSERT INTO admin_schema.user_table VALUES (13, 'Jane Smith', '9123456789', 'janesmith@example.com', '456 Market St, Town', '654321', 2, 2, 0, 1);
INSERT INTO admin_schema.user_table VALUES (14, 'John Doe', '9876543210', 'john.doe@example.com', '123 Street', '123456', 1, 2, 0, 1);


--
-- Data for Name: users; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.users VALUES (12, 'John Doe', '9876543210', NULL);


--
-- Data for Name: vehicle_engine_type; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.vehicle_engine_type VALUES (1, 'disel');
INSERT INTO admin_schema.vehicle_engine_type VALUES (2, 'petrol');


--
-- Data for Name: vehicle_insurance_table; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--



--
-- Data for Name: vehicle_make; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.vehicle_make VALUES (1, 'Toyota');
INSERT INTO admin_schema.vehicle_make VALUES (2, 'Honda');
INSERT INTO admin_schema.vehicle_make VALUES (3, 'Ford');


--
-- Data for Name: vehicle_model; Type: TABLE DATA; Schema: admin_schema; Owner: admin
--

INSERT INTO admin_schema.vehicle_model VALUES (1, 'Corolla');
INSERT INTO admin_schema.vehicle_model VALUES (2, 'Civic');
INSERT INTO admin_schema.vehicle_model VALUES (3, 'Mustang');


--
-- Data for Name: daily_price_update; Type: TABLE DATA; Schema: business_schema; Owner: postgres
--

INSERT INTO business_schema.daily_price_update VALUES (1, 100.50, 1, 5, 'INR', '2025-03-25 22:18:27.268885', '2025-03-25 22:18:27.268885', NULL);
INSERT INTO business_schema.daily_price_update VALUES (2, 250.00, 2, 6, 'INR', '2025-03-25 22:18:27.268885', '2025-03-25 22:18:27.268885', NULL);
INSERT INTO business_schema.daily_price_update VALUES (3, 75.75, 3, 12, 'INR', '2025-03-25 22:18:27.268885', '2025-03-25 22:18:27.268885', NULL);


--
-- Data for Name: invoice_details_table; Type: TABLE DATA; Schema: business_schema; Owner: postgres
--

INSERT INTO business_schema.invoice_details_table VALUES (1, 2, 6, 5.00, 0, '2025-03-25 23:01:40.615709', '2025-03-25 23:01:40.615709');


--
-- Data for Name: invoice_table; Type: TABLE DATA; Schema: business_schema; Owner: postgres
--

INSERT INTO business_schema.invoice_table VALUES (2, 'INV-20250224-0001', 1, 590.00, 500.00, '2025-02-24 18:25:44.460827', '2025-03-01', 2, 4, NULL, NULL, 'INR', DEFAULT);
INSERT INTO business_schema.invoice_table VALUES (3, 'INV-20250225-00003', 3, 1500.00, 100.00, '2025-02-25 14:14:21.29295', '2025-03-02', 1, 2, NULL, NULL, 'INR', DEFAULT);


--
-- Data for Name: mode_of_payment; Type: TABLE DATA; Schema: business_schema; Owner: admin
--

INSERT INTO business_schema.mode_of_payment VALUES (1, 2, 3);
INSERT INTO business_schema.mode_of_payment VALUES (2, 2, 3);
INSERT INTO business_schema.mode_of_payment VALUES (3, 1, 5);


--
-- Data for Name: order_history_table; Type: TABLE DATA; Schema: business_schema; Owner: postgres
--

INSERT INTO business_schema.order_history_table VALUES (1, '2025-02-22 09:55:59.551139', 4, '2025-02-27 09:55:59.551139', '2025-02-22 09:58:54.219252', 5, 6, 1, 2, '123456', '123 Street, City', NULL, 1);
INSERT INTO business_schema.order_history_table VALUES (1, '2025-02-22 09:55:59.551139', 6, '2025-02-27 09:55:59.551139', '2025-02-22 10:01:52.229039', 5, 6, 1, 2, '123456', '123 Street, City', '2025-02-22 10:01:52.229039', 2);
INSERT INTO business_schema.order_history_table VALUES (1, '2025-02-22 09:55:59.551139', 6, '2025-02-27 09:55:59.551139', '2025-02-22 10:01:52.229039', 5, 6, 1, 2, '123456', '123 Street, City', '2025-02-24 11:29:28.140697', 3);
INSERT INTO business_schema.order_history_table VALUES (1, '2025-02-22 09:55:59.551139', 6, '2025-02-27 09:55:59.551139', '2025-02-22 10:01:52.229039', 5, 6, 1, 2, '123456', '123 Street, City', '2025-02-24 11:59:18.925182', 4);
INSERT INTO business_schema.order_history_table VALUES (1, '2025-02-22 09:55:59.551139', 6, '2025-02-27 09:55:59.551139', '2025-02-22 10:01:52.229039', 5, 6, 1, 2, '123456', '123 Street, City', '2025-02-24 12:09:45.416871', 5);
INSERT INTO business_schema.order_history_table VALUES (1, '2025-02-22 09:55:59.551139', 6, '2025-02-27 09:55:59.551139', '2025-02-22 10:01:52.229039', 5, 6, 1, 2, '123456', '123 Street, City', '2025-02-24 21:12:04.580976', 6);
INSERT INTO business_schema.order_history_table VALUES (3, '2025-02-23 10:30:00', 2, '2025-02-26 10:30:00', NULL, 12, 6, 1, 2, '987654', '789 Road, City', NULL, 8);
INSERT INTO business_schema.order_history_table VALUES (3, '2025-02-23 10:30:00', 2, '2025-02-26 10:30:00', NULL, 12, 6, 1, 2, '987654', '789 Road, City', NULL, 9);
INSERT INTO business_schema.order_history_table VALUES (4, '2025-02-24 11:15:00', 5, '2025-02-28 11:15:00', '2025-02-24 12:00:00', 5, 6, 2, 1, '876543', '1011 Lane, City', NULL, 10);
INSERT INTO business_schema.order_history_table VALUES (4, '2025-02-24 11:15:00', 5, '2025-02-28 11:15:00', '2025-02-24 12:00:00', 5, 6, 2, 1, '876543', '1011 Lane, City', NULL, 11);


--
-- Data for Name: order_item_table; Type: TABLE DATA; Schema: business_schema; Owner: postgres
--

INSERT INTO business_schema.order_item_table VALUES (6, 1, 2, 5.00, 1, NULL, NULL);
INSERT INTO business_schema.order_item_table VALUES (4, 1, 2, 5.00, 1, NULL, NULL);
INSERT INTO business_schema.order_item_table VALUES (7, 3, 3, 2.00, 1, NULL, NULL);
INSERT INTO business_schema.order_item_table VALUES (8, 3, 1, 1.00, 2, NULL, NULL);
INSERT INTO business_schema.order_item_table VALUES (9, 4, 2, 4.00, 1, NULL, NULL);
INSERT INTO business_schema.order_item_table VALUES (10, 4, 1, 3.00, 1, NULL, NULL);


--
-- Data for Name: order_table; Type: TABLE DATA; Schema: business_schema; Owner: postgres
--

INSERT INTO business_schema.order_table VALUES (2, '2025-02-22 09:55:59.551139', 4, '2025-02-25', NULL, 5, 6, 2, 1, '654321', '456 Avenue, City', 0.00, NULL, NULL, NULL, '2025-03-25 23:07:59.479305', '2025-03-25 23:07:59.479305');
INSERT INTO business_schema.order_table VALUES (1, '2025-02-22 09:55:59.551139', 6, '2025-02-27', '2025-02-22', 5, 6, 1, 2, '123456', '123 Street, City', 1180.00, NULL, NULL, NULL, '2025-03-25 23:07:59.479305', '2025-03-25 23:07:59.479305');
INSERT INTO business_schema.order_table VALUES (3, '2025-02-23 10:30:00', 2, '2025-02-26', NULL, 12, 6, 1, 2, '987654', '789 Road, City', 1500.00, NULL, NULL, NULL, '2025-03-25 23:07:59.479305', '2025-03-25 23:07:59.479305');
INSERT INTO business_schema.order_table VALUES (4, '2025-02-24 11:15:00', 5, '2025-02-28', '2025-02-24', 5, 6, 2, 1, '876543', '1011 Lane, City', 1472.00, NULL, NULL, NULL, '2025-03-25 23:07:59.479305', '2025-03-25 23:07:59.479305');


--
-- Data for Name: stock_table; Type: TABLE DATA; Schema: business_schema; Owner: postgres
--



--
-- Data for Name: warehouse_list; Type: TABLE DATA; Schema: business_schema; Owner: postgres
--



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
-- Data for Name: product_price_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.product_price_table VALUES (2, 2, 100.00, 5, 18.00);


--
-- Data for Name: product_reviews; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.product_reviews VALUES (1, 1, 5, 5, 'Excellent product! Very high quality.');
INSERT INTO public.product_reviews VALUES (2, 2, 6, 4, 'Good value for money, but delivery was slow.');
INSERT INTO public.product_reviews VALUES (3, 1, 12, 3, 'Average product, could be better.');


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
-- Data for Name: transporter_table; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: type_of_insurance; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: user_bank_details; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: wholeseller_rating_table; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Name: business_branch_table_b_branch_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: postgres
--

SELECT pg_catalog.setval('admin_schema.business_branch_table_b_branch_id_seq', 4, true);


--
-- Name: business_category_table_b_category_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: postgres
--

SELECT pg_catalog.setval('admin_schema.business_category_table_b_category_id_seq', 1, false);


--
-- Name: business_type_table_type_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: postgres
--

SELECT pg_catalog.setval('admin_schema.business_type_table_type_id_seq', 3, true);


--
-- Name: cash_payment_list_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.cash_payment_list_id_seq', 5, true);


--
-- Name: category_regional_name_category_regional_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: postgres
--

SELECT pg_catalog.setval('admin_schema.category_regional_name_category_regional_id_seq', 1, false);


--
-- Name: indian_driver_licenses_type_license_type_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: postgres
--

SELECT pg_catalog.setval('admin_schema.indian_driver_licenses_type_license_type_id_seq', 1, false);


--
-- Name: master_category_table_category_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.master_category_table_category_id_seq', 11, true);


--
-- Name: master_city_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: postgres
--

SELECT pg_catalog.setval('admin_schema.master_city_id_seq', 1, false);


--
-- Name: master_driver_table_driver_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.master_driver_table_driver_id_seq', 2, true);


--
-- Name: master_language_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: postgres
--

SELECT pg_catalog.setval('admin_schema.master_language_id_seq', 1, false);


--
-- Name: master_location_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.master_location_id_seq', 5, true);


--
-- Name: master_mandi_table_mandi_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.master_mandi_table_mandi_id_seq', 2, true);


--
-- Name: master_product_utf8_product_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.master_product_utf8_product_id_seq', 1, false);


--
-- Name: master_states_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.master_states_id_seq', 3, true);


--
-- Name: master_vechicle_table_vehicle_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.master_vechicle_table_vehicle_id_seq', 6, true);


--
-- Name: master_violation_table_new_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.master_violation_table_new_id_seq', 2, true);


--
-- Name: mode_of_payments_list_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.mode_of_payments_list_id_seq', 2, true);


--
-- Name: order_status_table_order_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.order_status_table_order_id_seq', 9, true);


--
-- Name: permission_audit_log_log_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.permission_audit_log_log_id_seq', 2, true);


--
-- Name: product_regional_name_product_regional_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: postgres
--

SELECT pg_catalog.setval('admin_schema.product_regional_name_product_regional_id_seq', 1, false);


--
-- Name: roles_table_role_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.roles_table_role_id_seq', 4, true);


--
-- Name: units_table_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.units_table_id_seq', 5, true);


--
-- Name: user_table_user_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.user_table_user_id_seq', 14, true);


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.users_user_id_seq', 1, false);


--
-- Name: vehicle_engine_type_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.vehicle_engine_type_id_seq', 2, true);


--
-- Name: vehicle_insurance_table_vehicle_insurance_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.vehicle_insurance_table_vehicle_insurance_id_seq', 1, false);


--
-- Name: vehicle_make_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.vehicle_make_id_seq', 3, true);


--
-- Name: vehicle_model_id_seq; Type: SEQUENCE SET; Schema: admin_schema; Owner: admin
--

SELECT pg_catalog.setval('admin_schema.vehicle_model_id_seq', 3, true);


--
-- Name: invoice_details_table_id_seq; Type: SEQUENCE SET; Schema: business_schema; Owner: postgres
--

SELECT pg_catalog.setval('business_schema.invoice_details_table_id_seq', 1, true);


--
-- Name: invoice_table_id_seq; Type: SEQUENCE SET; Schema: business_schema; Owner: postgres
--

SELECT pg_catalog.setval('business_schema.invoice_table_id_seq', 3, true);


--
-- Name: mode_of_payment_id_seq; Type: SEQUENCE SET; Schema: business_schema; Owner: admin
--

SELECT pg_catalog.setval('business_schema.mode_of_payment_id_seq', 3, true);


--
-- Name: order_history_table_history_id_seq; Type: SEQUENCE SET; Schema: business_schema; Owner: postgres
--

SELECT pg_catalog.setval('business_schema.order_history_table_history_id_seq', 11, true);


--
-- Name: order_item_table_product_order_id_seq; Type: SEQUENCE SET; Schema: business_schema; Owner: postgres
--

SELECT pg_catalog.setval('business_schema.order_item_table_product_order_id_seq', 6, true);


--
-- Name: order_table_order_id_seq; Type: SEQUENCE SET; Schema: business_schema; Owner: postgres
--

SELECT pg_catalog.setval('business_schema.order_table_order_id_seq', 2, true);


--
-- Name: stock_table_stock_id_seq; Type: SEQUENCE SET; Schema: business_schema; Owner: postgres
--

SELECT pg_catalog.setval('business_schema.stock_table_stock_id_seq', 1, false);


--
-- Name: warehouse_list_warehouse_id_seq; Type: SEQUENCE SET; Schema: business_schema; Owner: postgres
--

SELECT pg_catalog.setval('business_schema.warehouse_list_warehouse_id_seq', 1, false);


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
-- Name: product_price_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_price_table_id_seq', 4, true);


--
-- Name: product_reviews_review_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_reviews_review_id_seq', 3, true);


--
-- Name: retailer_rating_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.retailer_rating_table_id_seq', 1, false);


--
-- Name: transporter_table_transporter_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transporter_table_transporter_id_seq', 1, false);


--
-- Name: type_of_insurance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.type_of_insurance_id_seq', 1, false);


--
-- Name: user_bank_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_bank_details_id_seq', 34, true);


--
-- Name: wholesaler_rating_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wholesaler_rating_table_id_seq', 1, false);


--
-- Name: business_branch_table business_branch_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.business_branch_table
    ADD CONSTRAINT business_branch_table_pkey PRIMARY KEY (b_branch_id);


--
-- Name: business_category_table business_category_table_b_category_name_key; Type: CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.business_category_table
    ADD CONSTRAINT business_category_table_b_category_name_key UNIQUE (b_category_name);


--
-- Name: business_category_table business_category_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.business_category_table
    ADD CONSTRAINT business_category_table_pkey PRIMARY KEY (b_category_id);


--
-- Name: business_table business_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.business_table
    ADD CONSTRAINT business_table_pkey PRIMARY KEY (bid);


--
-- Name: business_type_table business_type_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.business_type_table
    ADD CONSTRAINT business_type_table_pkey PRIMARY KEY (type_id);


--
-- Name: business_type_table business_type_table_type_name_key; Type: CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.business_type_table
    ADD CONSTRAINT business_type_table_type_name_key UNIQUE (type_name);


--
-- Name: business_user_table business_user_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.business_user_table
    ADD CONSTRAINT business_user_table_pkey PRIMARY KEY (b_id);


--
-- Name: cash_payment_list cash_payment_list_payment_type_key; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.cash_payment_list
    ADD CONSTRAINT cash_payment_list_payment_type_key UNIQUE (payment_type);


--
-- Name: cash_payment_list cash_payment_list_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.cash_payment_list
    ADD CONSTRAINT cash_payment_list_pkey PRIMARY KEY (id);


--
-- Name: category_regional_name category_regional_name_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.category_regional_name
    ADD CONSTRAINT category_regional_name_pkey PRIMARY KEY (category_regional_id);


--
-- Name: indian_driver_licenses_type indian_driver_licenses_type_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.indian_driver_licenses_type
    ADD CONSTRAINT indian_driver_licenses_type_pkey PRIMARY KEY (license_type_id);


--
-- Name: master_product_category_table master_category_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_product_category_table
    ADD CONSTRAINT master_category_table_pkey PRIMARY KEY (category_id);


--
-- Name: master_city master_city_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.master_city
    ADD CONSTRAINT master_city_pkey PRIMARY KEY (id);


--
-- Name: master_driver_table master_driver_table_driver_license_key; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_driver_table
    ADD CONSTRAINT master_driver_table_driver_license_key UNIQUE (driver_license);


--
-- Name: master_driver_table master_driver_table_driver_number_key; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_driver_table
    ADD CONSTRAINT master_driver_table_driver_number_key UNIQUE (driver_number);


--
-- Name: master_driver_table master_driver_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_driver_table
    ADD CONSTRAINT master_driver_table_pkey PRIMARY KEY (driver_id);


--
-- Name: master_language master_language_language_name_key; Type: CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.master_language
    ADD CONSTRAINT master_language_language_name_key UNIQUE (language_name);


--
-- Name: master_language master_language_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.master_language
    ADD CONSTRAINT master_language_pkey PRIMARY KEY (id);


--
-- Name: master_location master_location_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_location
    ADD CONSTRAINT master_location_pkey PRIMARY KEY (id);


--
-- Name: master_mandi_table master_mandi_table_mandi_incharge_num_key; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_mandi_table
    ADD CONSTRAINT master_mandi_table_mandi_incharge_num_key UNIQUE (mandi_incharge_num);


--
-- Name: master_mandi_table master_mandi_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_mandi_table
    ADD CONSTRAINT master_mandi_table_pkey PRIMARY KEY (mandi_id);


--
-- Name: master_product master_product_utf8_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_product
    ADD CONSTRAINT master_product_utf8_pkey PRIMARY KEY (product_id);


--
-- Name: master_states master_states_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_states
    ADD CONSTRAINT master_states_pkey PRIMARY KEY (id);


--
-- Name: master_vehicle_table master_vechicle_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_vehicle_table
    ADD CONSTRAINT master_vechicle_table_pkey PRIMARY KEY (vehicle_id);


--
-- Name: master_vehicle_table master_vechicle_table_vehicle_registration_no_key; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_vehicle_table
    ADD CONSTRAINT master_vechicle_table_vehicle_registration_no_key UNIQUE (vehicle_registration_no);


--
-- Name: master_violation_table master_violation_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_violation_table
    ADD CONSTRAINT master_violation_table_pkey PRIMARY KEY (id);


--
-- Name: mode_of_payments_list mode_of_payments_list_payment_mode_key; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.mode_of_payments_list
    ADD CONSTRAINT mode_of_payments_list_payment_mode_key UNIQUE (payment_mode);


--
-- Name: mode_of_payments_list mode_of_payments_list_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.mode_of_payments_list
    ADD CONSTRAINT mode_of_payments_list_pkey PRIMARY KEY (id);


--
-- Name: order_status_table order_status_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.order_status_table
    ADD CONSTRAINT order_status_table_pkey PRIMARY KEY (order_status_id);


--
-- Name: permission_audit_log permission_audit_log_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.permission_audit_log
    ADD CONSTRAINT permission_audit_log_pkey PRIMARY KEY (log_id);


--
-- Name: product_regional_name product_regional_name_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.product_regional_name
    ADD CONSTRAINT product_regional_name_pkey PRIMARY KEY (product_regional_id);


--
-- Name: roles_table roles_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.roles_table
    ADD CONSTRAINT roles_table_pkey PRIMARY KEY (role_id);


--
-- Name: roles_table roles_table_role_name_key; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.roles_table
    ADD CONSTRAINT roles_table_role_name_key UNIQUE (role_name);


--
-- Name: order_status_table unique_order_status_id; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.order_status_table
    ADD CONSTRAINT unique_order_status_id UNIQUE (order_status_id);


--
-- Name: units_table units_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.units_table
    ADD CONSTRAINT units_table_pkey PRIMARY KEY (id);


--
-- Name: units_table units_table_unit_name_key; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.units_table
    ADD CONSTRAINT units_table_unit_name_key UNIQUE (unit_name);


--
-- Name: user_table user_table_email_key; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.user_table
    ADD CONSTRAINT user_table_email_key UNIQUE (email);


--
-- Name: user_table user_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.user_table
    ADD CONSTRAINT user_table_pkey PRIMARY KEY (user_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: vehicle_engine_type vehicle_engine_type_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.vehicle_engine_type
    ADD CONSTRAINT vehicle_engine_type_pkey PRIMARY KEY (id);


--
-- Name: vehicle_insurance_table vehicle_insurance_table_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.vehicle_insurance_table
    ADD CONSTRAINT vehicle_insurance_table_pkey PRIMARY KEY (vehicle_insurance_id);


--
-- Name: vehicle_make vehicle_make_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.vehicle_make
    ADD CONSTRAINT vehicle_make_pkey PRIMARY KEY (id);


--
-- Name: vehicle_model vehicle_model_pkey; Type: CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.vehicle_model
    ADD CONSTRAINT vehicle_model_pkey PRIMARY KEY (id);


--
-- Name: invoice_details_table invoice_details_table_pkey; Type: CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.invoice_details_table
    ADD CONSTRAINT invoice_details_table_pkey PRIMARY KEY (id);


--
-- Name: invoice_table invoice_table_pkey; Type: CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.invoice_table
    ADD CONSTRAINT invoice_table_pkey PRIMARY KEY (id);


--
-- Name: mode_of_payment mode_of_payment_pkey; Type: CONSTRAINT; Schema: business_schema; Owner: admin
--

ALTER TABLE ONLY business_schema.mode_of_payment
    ADD CONSTRAINT mode_of_payment_pkey PRIMARY KEY (id);


--
-- Name: order_history_table order_history_table_pkey; Type: CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_history_table
    ADD CONSTRAINT order_history_table_pkey PRIMARY KEY (history_id);


--
-- Name: order_item_table order_item_table_pkey; Type: CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_item_table
    ADD CONSTRAINT order_item_table_pkey PRIMARY KEY (order_item_id);


--
-- Name: order_table order_table_pkey; Type: CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_table
    ADD CONSTRAINT order_table_pkey PRIMARY KEY (order_id);


--
-- Name: stock_table stock_table_pkey; Type: CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.stock_table
    ADD CONSTRAINT stock_table_pkey PRIMARY KEY (stock_id);


--
-- Name: invoice_table unique_invoice_number; Type: CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.invoice_table
    ADD CONSTRAINT unique_invoice_number UNIQUE (invoice_number);


--
-- Name: warehouse_list warehouse_list_pkey; Type: CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.warehouse_list
    ADD CONSTRAINT warehouse_list_pkey PRIMARY KEY (warehouse_id);


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
-- Name: product_price_table product_price_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_price_table
    ADD CONSTRAINT product_price_table_pkey PRIMARY KEY (id);


--
-- Name: product_reviews product_reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_reviews
    ADD CONSTRAINT product_reviews_pkey PRIMARY KEY (review_id);


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
-- Name: user_bank_details user_bank_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_bank_details
    ADD CONSTRAINT user_bank_details_pkey PRIMARY KEY (id);


--
-- Name: wholeseller_rating_table wholesaler_rating_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wholeseller_rating_table
    ADD CONSTRAINT wholesaler_rating_table_pkey PRIMARY KEY (id);


--
-- Name: business_table business_user_auto_insert; Type: TRIGGER; Schema: admin_schema; Owner: postgres
--

CREATE TRIGGER business_user_auto_insert AFTER INSERT ON admin_schema.business_table FOR EACH ROW EXECUTE FUNCTION public.insert_business_user();


--
-- Name: business_user_table business_user_insert_trigger; Type: TRIGGER; Schema: admin_schema; Owner: admin
--

CREATE TRIGGER business_user_insert_trigger BEFORE INSERT ON admin_schema.business_user_table FOR EACH ROW EXECUTE FUNCTION public.set_business_user_credentials();


--
-- Name: business_branch_table trg_insert_business_user; Type: TRIGGER; Schema: admin_schema; Owner: postgres
--

CREATE TRIGGER trg_insert_business_user AFTER INSERT ON admin_schema.business_branch_table FOR EACH ROW EXECUTE FUNCTION admin_schema.insert_business_user();


--
-- Name: business_table fk_b_cat; Type: FK CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.business_table
    ADD CONSTRAINT fk_b_cat FOREIGN KEY (b_category_id) REFERENCES admin_schema.business_category_table(b_category_id);


--
-- Name: business_table fk_b_type; Type: FK CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.business_table
    ADD CONSTRAINT fk_b_type FOREIGN KEY (b_type_id) REFERENCES admin_schema.business_table(bid);


--
-- Name: business_branch_table fk_business; Type: FK CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.business_branch_table
    ADD CONSTRAINT fk_business FOREIGN KEY (bid) REFERENCES admin_schema.business_table(bid) ON DELETE CASCADE;


--
-- Name: business_user_table fk_business_user; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.business_user_table
    ADD CONSTRAINT fk_business_user FOREIGN KEY (b_id) REFERENCES admin_schema.business_table(bid) ON DELETE CASCADE;


--
-- Name: master_vehicle_table fk_category; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_vehicle_table
    ADD CONSTRAINT fk_category FOREIGN KEY (vehicle_make) REFERENCES admin_schema.vehicle_make(id);


--
-- Name: category_regional_name fk_category; Type: FK CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.category_regional_name
    ADD CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES admin_schema.master_product_category_table(category_id) ON DELETE CASCADE;


--
-- Name: master_product_category_table fk_category_regional; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_product_category_table
    ADD CONSTRAINT fk_category_regional FOREIGN KEY (category_regional_id) REFERENCES admin_schema.category_regional_name(category_regional_id) ON DELETE CASCADE;


--
-- Name: master_location fk_city; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_location
    ADD CONSTRAINT fk_city FOREIGN KEY (city_shortnames) REFERENCES admin_schema.master_city(id) ON DELETE CASCADE;


--
-- Name: master_vehicle_table fk_engine; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_vehicle_table
    ADD CONSTRAINT fk_engine FOREIGN KEY (vehicle_engine_type) REFERENCES admin_schema.vehicle_engine_type(id);


--
-- Name: category_regional_name fk_language; Type: FK CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.category_regional_name
    ADD CONSTRAINT fk_language FOREIGN KEY (language_id) REFERENCES admin_schema.master_language(id) ON DELETE CASCADE;


--
-- Name: product_regional_name fk_language; Type: FK CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.product_regional_name
    ADD CONSTRAINT fk_language FOREIGN KEY (language_id) REFERENCES admin_schema.master_language(id) ON DELETE CASCADE;


--
-- Name: master_driver_table fk_license_type; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_driver_table
    ADD CONSTRAINT fk_license_type FOREIGN KEY (driver_license_type) REFERENCES admin_schema.indian_driver_licenses_type(license_type_id) ON DELETE CASCADE;


--
-- Name: product_regional_name fk_product; Type: FK CONSTRAINT; Schema: admin_schema; Owner: postgres
--

ALTER TABLE ONLY admin_schema.product_regional_name
    ADD CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES admin_schema.master_product(product_id) ON DELETE CASCADE;


--
-- Name: master_product fk_product_regional; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_product
    ADD CONSTRAINT fk_product_regional FOREIGN KEY (product_regional_id) REFERENCES admin_schema.product_regional_name(product_regional_id) ON DELETE CASCADE;


--
-- Name: user_table fk_state; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.user_table
    ADD CONSTRAINT fk_state FOREIGN KEY (state) REFERENCES admin_schema.master_states(id);


--
-- Name: master_mandi_table fk_state; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_mandi_table
    ADD CONSTRAINT fk_state FOREIGN KEY (mandi_state) REFERENCES admin_schema.master_states(id);


--
-- Name: master_location fk_state; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_location
    ADD CONSTRAINT fk_state FOREIGN KEY (state) REFERENCES admin_schema.master_states(id) ON DELETE CASCADE;


--
-- Name: user_table fk_user_role; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.user_table
    ADD CONSTRAINT fk_user_role FOREIGN KEY (role_id) REFERENCES admin_schema.roles_table(role_id) ON DELETE SET NULL;


--
-- Name: master_vehicle_table fk_vehicle_model; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.master_vehicle_table
    ADD CONSTRAINT fk_vehicle_model FOREIGN KEY (vehicle_model) REFERENCES admin_schema.vehicle_model(id);


--
-- Name: permission_audit_log permission_audit_log_changed_by_fkey; Type: FK CONSTRAINT; Schema: admin_schema; Owner: admin
--

ALTER TABLE ONLY admin_schema.permission_audit_log
    ADD CONSTRAINT permission_audit_log_changed_by_fkey FOREIGN KEY (changed_by) REFERENCES admin_schema.roles_table(role_id);


--
-- Name: stock_table fk_category; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.stock_table
    ADD CONSTRAINT fk_category FOREIGN KEY (warehouse_id) REFERENCES business_schema.warehouse_list(warehouse_id);


--
-- Name: daily_price_update fk_category; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.daily_price_update
    ADD CONSTRAINT fk_category FOREIGN KEY (product_id) REFERENCES admin_schema.master_product(product_id) ON DELETE CASCADE;


--
-- Name: mode_of_payment fk_mode; Type: FK CONSTRAINT; Schema: business_schema; Owner: admin
--

ALTER TABLE ONLY business_schema.mode_of_payment
    ADD CONSTRAINT fk_mode FOREIGN KEY (pay_mode) REFERENCES admin_schema.mode_of_payments_list(id);


--
-- Name: order_table fk_order_status; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_table
    ADD CONSTRAINT fk_order_status FOREIGN KEY (order_status) REFERENCES admin_schema.order_status_table(order_status_id) ON DELETE SET NULL;


--
-- Name: invoice_details_table fk_status; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.invoice_details_table
    ADD CONSTRAINT fk_status FOREIGN KEY (retailer_approval_status) REFERENCES public.retailer_status(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: mode_of_payment fk_type; Type: FK CONSTRAINT; Schema: business_schema; Owner: admin
--

ALTER TABLE ONLY business_schema.mode_of_payment
    ADD CONSTRAINT fk_type FOREIGN KEY (pay_type) REFERENCES admin_schema.cash_payment_list(id);


--
-- Name: daily_price_update fk_unit; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.daily_price_update
    ADD CONSTRAINT fk_unit FOREIGN KEY (unit_id) REFERENCES admin_schema.units_table(id);


--
-- Name: order_item_table fk_units; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_item_table
    ADD CONSTRAINT fk_units FOREIGN KEY (unit_id) REFERENCES admin_schema.units_table(id);


--
-- Name: daily_price_update fk_wholeseller; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.daily_price_update
    ADD CONSTRAINT fk_wholeseller FOREIGN KEY (wholeseller_id) REFERENCES admin_schema.business_table(bid) ON DELETE CASCADE;


--
-- Name: order_table fk_wholeseller; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_table
    ADD CONSTRAINT fk_wholeseller FOREIGN KEY (wholeseller_id) REFERENCES admin_schema.business_table(bid) ON DELETE CASCADE;


--
-- Name: invoice_details_table invoice_details_table_invoice_id_fkey; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.invoice_details_table
    ADD CONSTRAINT invoice_details_table_invoice_id_fkey FOREIGN KEY (invoice_id) REFERENCES business_schema.invoice_table(id) ON DELETE CASCADE;


--
-- Name: invoice_details_table invoice_details_table_order_item_id_fkey; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.invoice_details_table
    ADD CONSTRAINT invoice_details_table_order_item_id_fkey FOREIGN KEY (order_item_id) REFERENCES business_schema.order_item_table(order_item_id) ON DELETE CASCADE;


--
-- Name: invoice_table invoice_table_order_id_fkey; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.invoice_table
    ADD CONSTRAINT invoice_table_order_id_fkey FOREIGN KEY (order_id) REFERENCES business_schema.order_table(order_id) ON DELETE CASCADE;


--
-- Name: invoice_table invoice_table_pay_mode_fkey; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.invoice_table
    ADD CONSTRAINT invoice_table_pay_mode_fkey FOREIGN KEY (pay_mode) REFERENCES admin_schema.mode_of_payments_list(id);


--
-- Name: invoice_table invoice_table_pay_type_fkey; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.invoice_table
    ADD CONSTRAINT invoice_table_pay_type_fkey FOREIGN KEY (pay_type) REFERENCES admin_schema.cash_payment_list(id);


--
-- Name: order_item_table order_item_table_order_id_fkey; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_item_table
    ADD CONSTRAINT order_item_table_order_id_fkey FOREIGN KEY (order_id) REFERENCES business_schema.order_table(order_id) ON DELETE CASCADE;


--
-- Name: order_item_table order_item_table_product_id_fkey; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_item_table
    ADD CONSTRAINT order_item_table_product_id_fkey FOREIGN KEY (product_id) REFERENCES admin_schema.master_product(product_id) ON DELETE CASCADE;


--
-- Name: order_table order_table_location_id_fkey; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_table
    ADD CONSTRAINT order_table_location_id_fkey FOREIGN KEY (location_id) REFERENCES admin_schema.master_location(id) ON DELETE SET NULL;


--
-- Name: order_table order_table_state_id_fkey; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.order_table
    ADD CONSTRAINT order_table_state_id_fkey FOREIGN KEY (state_id) REFERENCES admin_schema.master_states(id) ON DELETE SET NULL;


--
-- Name: stock_table stock_table_product_id_fkey; Type: FK CONSTRAINT; Schema: business_schema; Owner: postgres
--

ALTER TABLE ONLY business_schema.stock_table
    ADD CONSTRAINT stock_table_product_id_fkey FOREIGN KEY (product_id) REFERENCES admin_schema.master_product(product_id) ON DELETE CASCADE;


--
-- Name: driver_violation_table driver_violation_table_driver_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.driver_violation_table
    ADD CONSTRAINT driver_violation_table_driver_id_fkey FOREIGN KEY (driver_id) REFERENCES admin_schema.master_driver_table(driver_id) ON DELETE CASCADE;


--
-- Name: driver_violation_table driver_violation_table_violation_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.driver_violation_table
    ADD CONSTRAINT driver_violation_table_violation_id_fkey FOREIGN KEY (violation_id) REFERENCES admin_schema.master_violation_table(id) ON DELETE CASCADE;


--
-- Name: wholeseller_rating_table fk_order; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wholeseller_rating_table
    ADD CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES business_schema.order_table(order_id) ON DELETE CASCADE;


--
-- Name: retailer_rating_table fk_order; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.retailer_rating_table
    ADD CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES business_schema.order_table(order_id) ON DELETE CASCADE;


--
-- Name: insurance_table fk_provider; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_table
    ADD CONSTRAINT fk_provider FOREIGN KEY (insurance_provider) REFERENCES public.insurance_provider(id);


--
-- Name: insurance_table fk_type; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_table
    ADD CONSTRAINT fk_type FOREIGN KEY (type_of_insurance) REFERENCES public.type_of_insurance(id);


--
-- Name: user_bank_details fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_bank_details
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES admin_schema.user_table(user_id) ON DELETE CASCADE;


--
-- Name: insurance_table insurance_table_driver_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_table
    ADD CONSTRAINT insurance_table_driver_id_fkey FOREIGN KEY (driver_id) REFERENCES admin_schema.master_driver_table(driver_id) ON DELETE CASCADE;


--
-- Name: insurance_table insurance_table_vehicle_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.insurance_table
    ADD CONSTRAINT insurance_table_vehicle_id_fkey FOREIGN KEY (vehicle_id) REFERENCES admin_schema.master_vehicle_table(vehicle_id) ON DELETE CASCADE;


--
-- Name: product_price_table product_price_table_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_price_table
    ADD CONSTRAINT product_price_table_product_id_fkey FOREIGN KEY (product_id) REFERENCES admin_schema.master_product(product_id) ON DELETE CASCADE;


--
-- Name: product_reviews product_reviews_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_reviews
    ADD CONSTRAINT product_reviews_product_id_fkey FOREIGN KEY (product_id) REFERENCES admin_schema.master_product(product_id) ON DELETE CASCADE;


--
-- Name: business_user_table; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.business_user_table ENABLE ROW LEVEL SECURITY;

--
-- Name: cash_payment_list; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.cash_payment_list ENABLE ROW LEVEL SECURITY;

--
-- Name: master_driver_table; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.master_driver_table ENABLE ROW LEVEL SECURITY;

--
-- Name: master_location; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.master_location ENABLE ROW LEVEL SECURITY;

--
-- Name: master_mandi_table; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.master_mandi_table ENABLE ROW LEVEL SECURITY;

--
-- Name: master_product; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.master_product ENABLE ROW LEVEL SECURITY;

--
-- Name: master_product_category_table; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.master_product_category_table ENABLE ROW LEVEL SECURITY;

--
-- Name: master_states; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.master_states ENABLE ROW LEVEL SECURITY;

--
-- Name: master_vehicle_table; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.master_vehicle_table ENABLE ROW LEVEL SECURITY;

--
-- Name: master_violation_table; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.master_violation_table ENABLE ROW LEVEL SECURITY;

--
-- Name: mode_of_payments_list; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.mode_of_payments_list ENABLE ROW LEVEL SECURITY;

--
-- Name: order_status_table; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.order_status_table ENABLE ROW LEVEL SECURITY;

--
-- Name: roles_table; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.roles_table ENABLE ROW LEVEL SECURITY;

--
-- Name: units_table; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.units_table ENABLE ROW LEVEL SECURITY;

--
-- Name: user_table; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.user_table ENABLE ROW LEVEL SECURITY;

--
-- Name: users; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.users ENABLE ROW LEVEL SECURITY;

--
-- Name: vehicle_engine_type; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.vehicle_engine_type ENABLE ROW LEVEL SECURITY;

--
-- Name: vehicle_insurance_table; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.vehicle_insurance_table ENABLE ROW LEVEL SECURITY;

--
-- Name: vehicle_make; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.vehicle_make ENABLE ROW LEVEL SECURITY;

--
-- Name: vehicle_model; Type: ROW SECURITY; Schema: admin_schema; Owner: admin
--

ALTER TABLE admin_schema.vehicle_model ENABLE ROW LEVEL SECURITY;

--
-- Name: SCHEMA admin_schema; Type: ACL; Schema: -; Owner: admin
--

GRANT USAGE ON SCHEMA admin_schema TO wholeseller_admin;


--
-- Name: SCHEMA business_schema; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON SCHEMA business_schema TO admin;
GRANT ALL ON SCHEMA business_schema TO business;
GRANT USAGE ON SCHEMA business_schema TO wholeseller_admin;


--
-- Name: SCHEMA users_schema; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON SCHEMA users_schema TO admin;
GRANT ALL ON SCHEMA users_schema TO users;


--
-- Name: TABLE business_branch_table; Type: ACL; Schema: admin_schema; Owner: postgres
--

GRANT ALL ON TABLE admin_schema.business_branch_table TO admin;


--
-- Name: TABLE business_category_table; Type: ACL; Schema: admin_schema; Owner: postgres
--

GRANT ALL ON TABLE admin_schema.business_category_table TO admin;


--
-- Name: TABLE business_table; Type: ACL; Schema: admin_schema; Owner: postgres
--

GRANT ALL ON TABLE admin_schema.business_table TO admin;


--
-- Name: TABLE business_type_table; Type: ACL; Schema: admin_schema; Owner: postgres
--

GRANT ALL ON TABLE admin_schema.business_type_table TO admin;


--
-- Name: TABLE cash_payment_list; Type: ACL; Schema: admin_schema; Owner: admin
--

GRANT SELECT ON TABLE admin_schema.cash_payment_list TO retailer;


--
-- Name: TABLE category_regional_name; Type: ACL; Schema: admin_schema; Owner: postgres
--

GRANT ALL ON TABLE admin_schema.category_regional_name TO admin;


--
-- Name: TABLE indian_driver_licenses_type; Type: ACL; Schema: admin_schema; Owner: postgres
--

GRANT ALL ON TABLE admin_schema.indian_driver_licenses_type TO admin;


--
-- Name: TABLE master_product_category_table; Type: ACL; Schema: admin_schema; Owner: admin
--

GRANT SELECT ON TABLE admin_schema.master_product_category_table TO retailer;


--
-- Name: TABLE master_city; Type: ACL; Schema: admin_schema; Owner: postgres
--

GRANT ALL ON TABLE admin_schema.master_city TO admin;


--
-- Name: TABLE master_language; Type: ACL; Schema: admin_schema; Owner: postgres
--

GRANT ALL ON TABLE admin_schema.master_language TO admin;


--
-- Name: TABLE master_mandi_table; Type: ACL; Schema: admin_schema; Owner: admin
--

GRANT SELECT,INSERT,REFERENCES,UPDATE ON TABLE admin_schema.master_mandi_table TO wholeseller_admin;


--
-- Name: TABLE master_product; Type: ACL; Schema: admin_schema; Owner: admin
--

GRANT SELECT ON TABLE admin_schema.master_product TO retailer;


--
-- Name: TABLE mode_of_payments_list; Type: ACL; Schema: admin_schema; Owner: admin
--

GRANT SELECT ON TABLE admin_schema.mode_of_payments_list TO retailer;


--
-- Name: TABLE product_regional_name; Type: ACL; Schema: admin_schema; Owner: postgres
--

GRANT ALL ON TABLE admin_schema.product_regional_name TO admin;


--
-- Name: TABLE daily_price_update; Type: ACL; Schema: business_schema; Owner: postgres
--

GRANT SELECT ON TABLE business_schema.daily_price_update TO retailer;
GRANT SELECT,INSERT,REFERENCES,DELETE,TRIGGER,UPDATE ON TABLE business_schema.daily_price_update TO wholeseller_admin;


--
-- Name: TABLE invoice_details_table; Type: ACL; Schema: business_schema; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE business_schema.invoice_details_table TO wholeseller_user;
GRANT SELECT ON TABLE business_schema.invoice_details_table TO retailer;
GRANT SELECT,INSERT,DELETE,TRIGGER,UPDATE ON TABLE business_schema.invoice_details_table TO wholeseller_admin;


--
-- Name: TABLE invoice_table; Type: ACL; Schema: business_schema; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE business_schema.invoice_table TO wholeseller_user;
GRANT SELECT ON TABLE business_schema.invoice_table TO retailer;
GRANT SELECT,INSERT,DELETE,TRIGGER,UPDATE ON TABLE business_schema.invoice_table TO wholeseller_admin;


--
-- Name: TABLE order_history_table; Type: ACL; Schema: business_schema; Owner: postgres
--

GRANT SELECT,INSERT,DELETE,UPDATE ON TABLE business_schema.order_history_table TO wholeseller_admin;


--
-- Name: TABLE order_item_table; Type: ACL; Schema: business_schema; Owner: postgres
--

GRANT SELECT,INSERT,UPDATE ON TABLE business_schema.order_item_table TO retailer;
GRANT SELECT ON TABLE business_schema.order_item_table TO wholeseller_user;
GRANT SELECT,INSERT,DELETE,UPDATE ON TABLE business_schema.order_item_table TO wholeseller_admin;


--
-- Name: TABLE order_table; Type: ACL; Schema: business_schema; Owner: postgres
--

GRANT SELECT,INSERT ON TABLE business_schema.order_table TO retailer;
GRANT SELECT ON TABLE business_schema.order_table TO wholeseller_user;
GRANT SELECT,INSERT,DELETE,UPDATE ON TABLE business_schema.order_table TO wholeseller_admin;


--
-- Name: TABLE stock_table; Type: ACL; Schema: business_schema; Owner: postgres
--

GRANT SELECT ON TABLE business_schema.stock_table TO retailer;
GRANT SELECT,INSERT,DELETE,UPDATE ON TABLE business_schema.stock_table TO wholeseller_admin;


--
-- Name: TABLE warehouse_list; Type: ACL; Schema: business_schema; Owner: postgres
--

GRANT SELECT ON TABLE business_schema.warehouse_list TO admin;
GRANT SELECT,INSERT,DELETE,UPDATE ON TABLE business_schema.warehouse_list TO wholeseller_admin;


--
-- Name: DEFAULT PRIVILEGES FOR TABLES; Type: DEFAULT ACL; Schema: admin_schema; Owner: admin
--

ALTER DEFAULT PRIVILEGES FOR ROLE admin IN SCHEMA admin_schema GRANT ALL ON TABLES TO admin;


--
-- Name: DEFAULT PRIVILEGES FOR TABLES; Type: DEFAULT ACL; Schema: admin_schema; Owner: postgres
--

ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA admin_schema GRANT ALL ON TABLES TO admin;


--
-- PostgreSQL database dump complete
--

