### **Table 1: `admin_schema.business_table`**  

#### **Columns & Data Types**  
- **`bid`** (`bigint`): Primary key, auto-incremented.  
- **`b_typeid`** (`integer`): Foreign key referencing `business_type_table`.  
- **`b_location_id`** (`integer`): Foreign key referencing `master_location`.  
- **`b_state_id`** (`integer`): Foreign key referencing `master_states`.  
- **`b_mandi_id`** (`integer`): Foreign key referencing `master_mandi_table`.  
- **`b_address`** (`text`): Business address.  
- **`b_phone_num`** (`text`): Contact number (optional).  
- **`b_email`** (`varchar(255)`): Unique business email.  
- **`b_gst_num`** (`varchar(20)`): Unique GST number.  
- **`b_pan_num`** (`varchar(10)`): Unique PAN number.  
- **`b_name`** (`varchar(255)`): Business name (optional).  

---

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `business_table_pkey1` → `bid`.  
- **Unique Constraints**:  
  - `business_table_b_email_key` → Ensures unique `b_email`.  
  - `business_table_b_gst_num_key` → Ensures unique `b_gst_num`.  
  - `business_table_b_pan_num_key` → Ensures unique `b_pan_num`.  

---

#### **Foreign Key Constraints**  
- **`b_location_id`** → References `admin_schema.master_location(id)`, **ON DELETE CASCADE**.  
- **`b_state_id`** → References `admin_schema.master_states(id)`, **ON DELETE CASCADE**.  
- **`b_typeid`** → References `admin_schema.business_type_table(b_typeid)`, **ON DELETE CASCADE**.  
- **`b_mandi_id`** → References `admin_schema.master_mandi_table(mandi_id)`.  

---

#### **Referenced By (Other Tables Depending on `business_table`)**  
- **`order_table.retailer_id`** → References `business_table.bid`, **ON DELETE SET NULL**.  
- **`order_table.wholeseller_id`** → References `business_table.bid`, **ON DELETE SET NULL**.  
- **`public.product_price_table.bid`** → References `business_table.bid`, **ON DELETE CASCADE**.  
- **`public.product_reviews.user_id`** → References `business_table.bid`, **ON DELETE CASCADE**.  

---

#### **Row-Level Security Policy**  
- **Policy Name**: `insert_wholeseller_admin`  
- **Applies To**: `INSERT` operations.  
- **Allowed Role**: `wholeseller_admin`.  
- **Condition**: `b_typeid = 5` (Only allows inserting businesses of type 5).  

---

### **Table 2: `admin_schema.business_type_table`**  

#### **Columns & Data Types**  
- **`b_typeid`** (`integer`): Primary key, auto-incremented.  
- **`b_typename`** (`varchar(255)`): Unique business type name.  
- **`remarks`** (`text`): Additional comments or descriptions (optional).  

---

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `business_table_pkey` → `b_typeid`.  
- **Unique Constraints**:  
  - `business_table_b_typename_key` → Ensures unique `b_typename`.  

---

#### **Referenced By (Other Tables Depending on `business_type_table`)**  
- **`admin_schema.business_table.b_typeid`** → References `business_type_table.b_typeid`, **ON DELETE CASCADE**.  

---

### **Table 3: `admin_schema.cash_payment_list`**  

#### **Columns & Data Types**  
- **`id`** (`integer`): Primary key, auto-incremented.  
- **`payment_type`** (`varchar(50)`): Unique payment type name.  

---

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `cash_payment_list_pkey` → `id`.  
- **Unique Constraints**:  
  - `cash_payment_list_payment_type_key` → Ensures unique `payment_type`.  

---

#### **Referenced By (Other Tables Depending on `cash_payment_list`)**  
- **`admin_schema.mode_of_payment.pay_type`** → References `cash_payment_list.id`.  
- **`business_schema.invoice_table.pay_type`** → References `cash_payment_list.id`.  

---

### **Table 4: `admin_schema.master_category_table`**  

#### **Columns & Data Types**  
- **`category_id`** (`integer`): Primary key, auto-incremented.  
- **`category_name`** (`varchar(255)`): Category name (not null).  
- **`super_cat_id`** (`integer`): Parent category reference (nullable).  
- **`img_path`** (`text`): Image path for the category.  
- **`col1`** (`text`): Additional column for future use.  
- **`col2`** (`text`): Additional column for future use.  
- **`remarks`** (`text`): Optional remarks about the category.  

---

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `master_category_table_pkey` → `category_id`.  

---

### **Table 5: `admin_schema.master_driver_table`**  

#### **Columns & Data Types**  
- **`driver_id`** (`integer`): Primary key, auto-incremented.  
- **`driver_name`** (`varchar(255)`): Driver's full name (not null).  
- **`driver_license`** (`varchar(50)`): Unique license number (not null).  
- **`driver_number`** (`varchar(15)`): Unique contact number (not null).  
- **`driver_address`** (`varchar(255)`): Driver's address (nullable).  
- **`driver_status`** (`varchar(50)`): Driver's current status (nullable).  
- **`date_of_joining`** (`date`): Date the driver joined (nullable).  
- **`license_expiry_date`** (`date`): License validity end date (not null).  
- **`emergency_contact`** (`varchar(15)`): Emergency contact number (nullable).  

---

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `master_driver_table_pkey` → `driver_id`.  
- **Unique Constraints**:  
  - `master_driver_table_driver_license_key` → `driver_license`.  
  - `master_driver_table_driver_number_key` → `driver_number`.  


#### **Foreign Key References**  
- Referenced in **`driver_violation_table`**:  
  - `driver_violation_table_driver_id_fkey` → `driver_id` (ON DELETE CASCADE).  
- Referenced in **`insurance_table`**:  
  - `insurance_table_driver_id_fkey` → `driver_id` (ON DELETE CASCADE).  

---

### **Table 6: `admin_schema.master_location`**  

#### **Columns & Data Types**  
- **`id`** (`integer`): Primary key, auto-incremented.  
- **`location`** (`varchar(50)`): Name of the location (nullable).  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `master_location_pkey` → `id`.  


#### **Foreign Key References**  
- **Referenced in `admin_schema.business_table`**:  
  - `fk_b_location` → `b_location_id` (ON DELETE CASCADE).  
- **Referenced in `admin_schema.master_mandi_table`**:  
  - `fk_loc` → `mandi_city`.  
- **Referenced in `business_schema.order_table`**:  
  - `order_table_location_id_fkey` → `location_id` (ON DELETE SET NULL).  

---

### **Table 7: `admin_schema.master_states`**  

#### **Columns & Data Types**  
- **`id`** (`integer`): Primary key, auto-incremented.  
- **`state`** (`varchar(50)`): Name of the state (nullable).  

---

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `master_states_pkey` → `id`.  

#### **Foreign Key References**  
- **Referenced in `admin_schema.business_table`**:  
  - `fk_b_state` → `b_state_id` (ON DELETE CASCADE).  
- **Referenced in `admin_schema.master_mandi_table`**:  
  - `fk_state` → `mandi_state`.  
- **Referenced in `admin_schema.user_table`**:  
  - `fk_state` → `state`.  
- **Referenced in `business_schema.order_table`**:  
  - `order_table_state_id_fkey` → `state_id` (ON DELETE SET NULL).  

---
### **Table 8: `admin_schema.master_mandi_table`**  

#### **Columns & Data Types**  
- **`mandi_id`** (`integer`): Primary key, auto-incremented.  
- **`mandi_location`** (`varchar(255)`): Name of the mandi location (not null).  
- **`mandi_number`** (`varchar(50)`): Unique identifier for the mandi (not null).  
- **`mandi_incharge`** (`varchar(255)`): Name of the mandi in-charge (not null).  
- **`mandi_incharge_num`** (`varchar(15)`): Contact number of the in-charge (not null).  
- **`mandi_pincode`** (`varchar(6)`): Pincode of the mandi location (not null).  
- **`mandi_address`** (`text`): Address of the mandi (nullable).   
- **`remarks`** (`text`): Additional remarks or notes (nullable).  
- **`mandi_city`** (`integer`): Reference to the city of the mandi (nullable).  
- **`mandi_state`** (`integer`): Reference to the state of the mandi (nullable).  
- **`mandi_name`** (`varchar(255)`): Name of the mandi (nullable).  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `master_mandi_table_pkey` → `mandi_id`.  
- **Unique Constraints**:  
  - `master_mandi_table_mandi_incharge_num_key` → `mandi_incharge_num`.  
  - `master_mandi_table_mandi_number_key` → `mandi_number`.  

#### **Foreign Key References**  
- **References `admin_schema.master_location`**:  
  - `fk_loc` → `mandi_city`.  
- **References `admin_schema.master_states`**:  
  - `fk_state` → `mandi_state`.  
- **Referenced by `admin_schema.business_table`**:  
  - `fk_mandi_id` → `b_mandi_id` (ON DELETE CASCADE).  

---

### **Table 9: `admin_schema.master_product`**  

#### **Columns & Data Types**  
- **`product_id`** (`bigint`): Primary key, auto-incremented.  
- **`category_id`** (`integer`): References the product category (nullable).  
- **`product_name`** (`varchar(255)`): Name of the product (nullable).  
- **`status`** (`integer`): Status of the product (nullable).  
- **`image_path`** (`varchar(255)`): Path to the product image (nullable).  
- **`regional_name1`** (`varchar(255)`, `en-US-x-icu`): Regional language name 1.  
- **`regional_name2`** (`varchar(255)`, `en-US-x-icu`): Regional language name 2.  
- **`regional_name3`** (`varchar(255)`, `en-US-x-icu`): Regional language name 3.  
- **`regional_name4`** (`varchar(255)`, `en-US-x-icu`): Regional language name 4.  


#### **Indexes & Constraints**  
- **Primary Key**:  
  - `master_product_utf8_pkey` → `product_id`.  

#### **Foreign Key References**  
- **Referenced by other tables**:  
  - **`business_schema.daily_price_update`** → `fk_category` → `product_id` (ON DELETE CASCADE).  
  - **`business_schema.order_item_table`** → `order_item_table_product_id_fkey` → `product_id` (ON DELETE CASCADE).  
  - **`product_price_table`** → `product_price_table_product_id_fkey` → `product_id` (ON DELETE CASCADE).  
  - **`product_reviews`** → `product_reviews_product_id_fkey` → `product_id` (ON DELETE CASCADE).  
  - **`business_schema.stock_table`** → `stock_table_product_id_fkey` → `product_id` (ON DELETE CASCADE).  

---

### **Table 10: `admin_schema.master_vehicle_table`**  

#### **Columns & Data Types**  
- **`vehicle_id`** (`integer`): Primary key, auto-incremented.  
- **`insurance_id`** (`integer`): Reference to vehicle insurance (nullable).  
- **`vehicle_name`** (`varchar(255)`): Name of the vehicle (not null).  
- **`vehicle_manufacture_year`** (`varchar(4)`): Year of manufacture (nullable).  
- **`vehicle_warranty`** (`varchar(255)`): Warranty details (nullable).  
- **`vehicle_make`** (`integer`): Foreign key to `admin_schema.vehicle_make` (nullable).  
- **`vehicle_model`** (`integer`): Foreign key to `admin_schema.vehicle_model` (nullable).  
- **`vehicle_registration_no`** (`varchar(50)`): Unique registration number.  
- **`vehicle_engine_type`** (`integer`): Foreign key to `admin_schema.vehicle_engine_type` (nullable).  
- **`vehicle_purchase_date`** (`date`): Purchase date (nullable).  
- **`vehicle_color`** (`varchar(50)`): Vehicle color (nullable).   
- **`vehicle_insurance_id`** (`integer`): Another reference to insurance (nullable).  
- **`vehicle_make_name`** (`varchar(255)`): Name of the vehicle manufacturer (nullable).  
- **`vehicle_model_name`** (`varchar(255)`): Name of the vehicle model (nullable).  
- **`vehicle_engine_type_name`** (`varchar(255)`): Name of the vehicle's engine type (nullable).  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `master_vechicle_table_pkey` → `vehicle_id`.  
- **Unique Constraint**:  
  - `master_vechicle_table_vehicle_registration_no_key` → `vehicle_registration_no`.  

#### **Foreign Key References**  
- **This table references**:  
  - `admin_schema.vehicle_make` → `vehicle_make` (`fk_category`).  
  - `admin_schema.vehicle_engine_type` → `vehicle_engine_type` (`fk_engine`).  
  - `admin_schema.vehicle_model` → `vehicle_model` (`fk_vehicle_model`).  

- **Referenced by other tables**:  
  - **`insurance_table`** → `insurance_table_vehicle_id_fkey` → `vehicle_id` (ON DELETE CASCADE).  

---

### **Table 11: `admin_schema.master_violation_table`**  

#### **Columns & Data Types**  
- **`violation_name`** (`text`): Name of the violation (nullable).  
- **`level_of_serious`** (`text`): Severity level of the violation (nullable).  
- **`id`** (`integer`): Primary key, auto-incremented (`nextval(...)`).  


#### **Indexes & Constraints**  
- **Primary Key**:  
  - `master_violation_table_pkey` → `id`.  


#### **Foreign Key References**  
- **Referenced by other tables**:  
  - **`driver_violation_table`** → `driver_violation_table_violation_id_fkey` → `violation_id` (ON DELETE CASCADE).  

---
### **Table 12: `admin_schema.mode_of_payments_list`**  

#### **Columns & Data Types**  
- **`id`** (`integer`): Primary key, auto-incremented (`nextval('admin_schema.mode_of_payments_list_id_seq'::regclass)`).  
- **`payment_mode`** (`character varying(50)`): Mode of payment (NOT NULL).  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `mode_of_payments_list_pkey` → `id`.  
- **Unique Constraint**:  
  - `mode_of_payments_list_payment_mode_key` → `payment_mode`.  

#### **Foreign Key References**  
- **Referenced by other tables**:  
  - **`business_schema.mode_of_payment`** → `fk_mode` → `pay_mode` (References `id` in `mode_of_payments_list`).  
  - **`business_schema.invoice_table`** → `invoice_table_pay_mode_fkey` → `pay_mode` (References `id` in `mode_of_payments_list`).  

---
### **Table 13: `admin_schema.order_status_table`**  

#### **Columns & Data Types**  
- **`order_status_id`** (`integer`): Primary key, auto-incremented (`nextval('admin_schema.order_status_table_order_id_seq'::regclass)`).  
- **`order_status`** (`character varying(50)`): Status of the order (NOT NULL).  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `order_status_table_pkey` → `order_status_id`.  
- **Unique Constraint**:  
  - `unique_order_status_id` → `order_status_id`.  

#### **Foreign Key References**  
- **Referenced by other tables**:  
  - **`business_schema.order_item_table`** → `fk_order_item_status` → `order_item_status` (References `order_status_id` in `order_status_table`).  
  - **`business_schema.order_table`** → `fk_order_status` → `order_status` (References `order_status_id` in `order_status_table`, ON DELETE SET NULL).  
---

### **Table 14: `admin_schema.units_table`**  

#### **Columns & Data Types**  
- **`id`** (`integer`): Primary key, auto-incremented (`nextval('admin_schema.units_table_id_seq'::regclass)`).  
- **`unit_name`** (`character varying(50)`): Name of the unit (NOT NULL).  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `units_table_pkey` → `id`.  
- **Unique Constraint**:  
  - `units_table_unit_name_key` → `unit_name`.  

#### **Foreign Key References**  
- **Referenced by other tables**:  
  - **`business_schema.daily_price_update`** → `fk_unit` → `unit_id` (References `id` in `units_table`).  
  - **`business_schema.order_item_table`** → `fk_units` → `unit_id` (References `id` in `units_table`).  

---
### **Table 15: `admin_schema.user_table`**  

#### **Columns & Data Types**  
- **`user_id`** (`integer`): Primary key, auto-incremented (`nextval('admin_schema.user_table_user_id_seq'::regclass)`).  
- **`user_type_id`** (`integer`): References user type (nullable).  
- **`name`** (`character varying(255)`): Name of the user (NOT NULL).  
- **`mobile_num`** (`character varying(15)`): Mobile number of the user (NOT NULL).  
- **`email`** (`character varying(255)`): Email of the user (nullable).  
- **`address`** (`text`): Address of the user (nullable).  
- **`pincode`** (`character varying(10)`): Pincode of the user's location (nullable).  
- **`location`** (`integer`): Location ID (nullable).  
- **`state`** (`integer`): State ID (nullable).  
- **`status`** (`integer`): Status of the user (nullable).  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `user_table_pkey` → `user_id`.  
- **Unique Constraint**:  
  - `user_table_email_key` → `email`.  

#### **Foreign Key References**  
- **References other tables**:  
  - **`admin_schema.master_states`** → `fk_state` → `state`.  
  - **`admin_schema.user_type_table`** → `user_table_user_type_id_fkey` → `user_type_id` (ON DELETE SET NULL).  

- **Referenced by other tables**:  
  - **`user_bank_details`** → `fk_user` → `user_id` (ON DELETE CASCADE).  
---

### **Table 16: `admin_schema.user_type_table`**  

#### **Columns & Data Types**  
- **`user_type_id`** (`integer`): Primary key, auto-incremented (`nextval('admin_schema.user_type_table_user_type_id_seq'::regclass)`).  
- **`user_type`** (`character varying(50)`): Type of user (NOT NULL).    

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `user_type_table_pkey` → `user_type_id`.  

#### **Foreign Key References**  
- **Referenced by other tables**:  
  - **`admin_schema.users`** → `fk_user_type_table` → `user_type_id` (ON UPDATE CASCADE, ON DELETE CASCADE).  
  - **`admin_schema.user_table`** → `user_table_user_type_id_fkey` → `user_type_id` (ON DELETE SET NULL).  
---

### **Table 17: `admin_schema.users`**  

#### **Columns & Data Types**  
- **`user_id`** (`integer`): Primary key, auto-incremented (`nextval('admin_schema.users_user_id_seq'::regclass)`).  
- **`user_type_id`** (`integer`): References `user_type_table.user_type_id` (NOT NULL).  
- **`username`** (`character varying(255)`): Username of the user (NOT NULL).  
- **`password`** (`character varying(255)`): Password of the user (NOT NULL).  
- **`machine_id`** (`character varying(255)`): Machine identifier (nullable).  
- **`doe`** (`timestamp without time zone`): Date of entry (DEFAULT: `CURRENT_TIMESTAMP`).  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `users_pkey` → `user_id`.  

#### **Foreign Key References**  
- **Referenced by other tables**:  
  - **`business_schema.stock_table`** → `stock_table_user_id_fkey` → `user_id` (ON DELETE SET NULL).  

- **References other tables**:  
  - **`admin_schema.user_type_table`** → `fk_user_type_table` → `user_type_id` (ON UPDATE CASCADE, ON DELETE CASCADE).  
---
### **Table 18: `admin_schema.vehicle_engine_type`**  

#### **Columns & Data Types**  
- **`id`** (`integer`): Primary key, auto-incremented (`nextval('admin_schema.vehicle_engine_type_id_seq'::regclass)`).  
- **`engine_type`** (`character varying(50)`): Type of vehicle engine (nullable).  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `vehicle_engine_type_pkey` → `id`.  

#### **Foreign Key References**  
- **Referenced by other tables**:  
  - **`admin_schema.master_vehicle_table`** → `fk_engine` → `vehicle_engine_type` (REFERENCES `admin_schema.vehicle_engine_type(id)`).  
---
### **Table 19: `admin_schema.vehicle_make`**  

#### **Columns & Data Types**  
- **`id`** (`integer`): Primary key, auto-incremented (`nextval('admin_schema.vehicle_make_id_seq'::regclass)`).  
- **`make`** (`character varying(50)`): Vehicle make/brand (nullable).  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `vehicle_make_pkey` → `id`.  

#### **Foreign Key References**  
- **Referenced by other tables**:  
  - **`admin_schema.master_vehicle_table`** → `fk_category` → `vehicle_make` (REFERENCES `admin_schema.vehicle_make(id)`).  
---
### **Table 20: `admin_schema.vehicle_model`**  

#### **Columns & Data Types**  
- **`id`** (`integer`): Primary key, auto-incremented (`nextval('admin_schema.vehicle_model_id_seq'::regclass)`).  
- **`model`** (`character varying(50)`): Vehicle model name (nullable).  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `vehicle_model_pkey` → `id`.  

#### **Foreign Key References**  
- **Referenced by other tables**:  
  - **`admin_schema.master_vehicle_table`** → `fk_vehicle_model` → `vehicle_model` (REFERENCES `admin_schema.vehicle_model(id)`).  
---

