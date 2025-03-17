### **Table 1: `business_schema.daily_price_update`**  

#### **Columns & Data Types**  
- **`product_id`** (`integer`): References the product being priced.  
- **`price`** (`character varying(50)`): Price of the product (nullable).  
- **`status`** (`integer`): Status of the price update (nullable).  
- **`unit_id`** (`integer`): References the measurement unit.  
- **`wholeseller_id`** (`integer`): References the wholeseller who updates the price.  

#### **Indexes & Constraints**  
- **Foreign Key References**:  
  - **`product_id`** → `fk_category` → `admin_schema.master_product(product_id)` (ON DELETE CASCADE).  
  - **`unit_id`** → `fk_unit` → `admin_schema.units_table(id)`.  
  - **`wholeseller_id`** → `fk_wholeseller` → `admin_schema.business_table(bid)` (ON DELETE CASCADE).  

---

### **Table 2: `business_schema.invoice_details_table`**  

#### **Columns & Data Types**  
- **`id`** (`bigint`): Primary key, auto-incremented.  
- **`invoice_id`** (`bigint`): References the associated invoice.  
- **`order_item_id`** (`bigint`): References the order item.  
- **`quantity`** (`numeric(10,2)`): Quantity of the item (not null).  
- **`retailer_status`** (`integer`): Status of the invoice from the retailer's perspective (default: `0`).  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `invoice_details_table_pkey` → `id`.  
- **Foreign Key References**:  
  - **`retailer_status`** → `fk_status` → `retailer_status(id)` (ON UPDATE CASCADE, ON DELETE SET NULL).  
  - **`invoice_id`** → `invoice_details_table_invoice_id_fkey` → `business_schema.invoice_table(id)` (ON DELETE CASCADE).  
  - **`order_item_id`** → `invoice_details_table_order_item_id_fkey` → `business_schema.order_item_table(order_item_id)` (ON DELETE CASCADE).  

---

### **Table 3: `business_schema.invoice_table`**  

#### **Columns & Data Types**  
- **`id`** (`integer`): Primary key, auto-incremented.  
- **`invoice_number`** (`character varying(50)`): Unique invoice identifier.  
- **`order_id`** (`bigint`): References the associated order.  
- **`total_amount`** (`numeric(10,2)`): Total amount for the invoice (not null).  
- **`discount_amount`** (`numeric(10,2)`): Discount applied (default: `0`).  
- **`invoice_date`** (`timestamp without time zone`): Timestamp of invoice creation (default: `CURRENT_TIMESTAMP`).  
- **`due_date`** (`date`): Payment due date.  
- **`pay_mode`** (`integer`): Payment mode reference.  
- **`pay_type`** (`integer`): Payment type reference.  
- **`final_amount`** (`numeric(10,2)`): **Stored generated column** calculated as `(total_amount - discount_amount)`.  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `invoice_table_pkey` → `id`.  
- **Unique Constraint**:  
  - `unique_invoice_number` → `invoice_number`.  
- **Foreign Key References**:  
  - **`order_id`** → `invoice_table_order_id_fkey` → `business_schema.order_table(order_id)` (ON DELETE CASCADE).  
  - **`pay_mode`** → `invoice_table_pay_mode_fkey` → `admin_schema.mode_of_payments_list(id)`.  
  - **`pay_type`** → `invoice_table_pay_type_fkey` → `admin_schema.cash_payment_list(id)`.  
- **Referenced By**:  
  - **`business_schema.invoice_details_table`** → `invoice_details_table_invoice_id_fkey` → `business_schema.invoice_table(id)` (ON DELETE CASCADE).  

#### **Triggers**  
- **`invoice_number_trigger`** → `BEFORE INSERT` → Calls `generate_invoice_number()`.  
- **`trigger_generate_invoice_number`** → `BEFORE INSERT` → Calls `generate_invoice_number()`.  

---
### **Table 4: `business_schema.mode_of_payment`**  

#### **Columns & Data Types**  
- **`id`** (`integer`): Primary key, auto-incremented.  
- **`pay_mode`** (`integer`): References the mode of payment.  
- **`pay_type`** (`integer`): References the type of payment.  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `mode_of_payment_pkey` → `id`.  
- **Foreign Key References**:  
  - **`pay_mode`** → `fk_mode` → `admin_schema.mode_of_payments_list(id)`.  
  - **`pay_type`** → `fk_type` → `admin_schema.cash_payment_list(id)`.  

#### **Triggers**  
- **`trg_default_pay_type`** → `BEFORE INSERT` → Calls `set_default_pay_type()`.  
---
### **Table 5: `business_schema.order_item_table`**  

#### **Columns & Data Types**  
- **`order_item_id`** (`bigint`): Primary key, auto-incremented.  
- **`order_id`** (`bigint`): References the order.  
- **`product_id`** (`bigint`): References the product.  
- **`quantity`** (`numeric(10,2)`): The number of units ordered.  
- **`unit_id`** (`integer`): The measurement unit.  
- **`amt_of_order_item`** (`numeric(10,3)`): Total amount for this order item.  
- **`order_item_status`** (`integer`): References the order status.  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `order_item_table_pkey` → `order_item_id`.  
- **Foreign Key References**:  
  - **`order_item_status`** → `fk_order_item_status` → `admin_schema.order_status_table(order_status_id)`.  
  - **`unit_id`** → `fk_units` → `admin_schema.units_table(id)`.  
  - **`order_id`** → `order_item_table_order_id_fkey` → `business_schema.order_table(order_id) ON DELETE CASCADE`.  
  - **`product_id`** → `order_item_table_product_id_fkey` → `admin_schema.master_product(product_id) ON DELETE CASCADE`.  

#### **Referenced By**  
- **`business_schema.invoice_details_table`**  
  - Constraint: `invoice_details_table_order_item_id_fkey`  
  - References: `business_schema.order_item_table(order_item_id) ON DELETE CASCADE`.  

#### **Triggers**  
- **`trigger_update_amt`** → `BEFORE INSERT OR UPDATE` → Calls `update_amt_of_order_item()`.  
- **`trigger_update_total_order_amount`** → `AFTER INSERT, DELETE, OR UPDATE` → Calls `update_total_order_amount()`.  
---
### **Table 6: `business_schema.order_table`**  

#### **Columns & Data Types**  
- **`order_id`** (`bigint`): Primary key, auto-incremented.  
- **`date_of_order`** (`timestamp`): Defaults to `CURRENT_TIMESTAMP`.  
- **`order_status`** (`integer`): References order status.  
- **`expected_delivery_date`** (`timestamp`): Estimated delivery date.  
- **`actual_delivery_date`** (`timestamp`): Actual delivery date.  
- **`retailer_id`** (`integer`): References retailer.  
- **`wholeseller_id`** (`integer`): References wholesaler.  
- **`location_id`** (`integer`): Location of the order.  
- **`state_id`** (`integer`): References state.  
- **`pincode`** (`varchar(10)`): Pincode for delivery.  
- **`address`** (`text`): Delivery address.  
- **`total_order_amount`** (`numeric(10,2)`): Defaults to `0`.  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `order_table_pkey` → `order_id`.  
- **Foreign Key References**:  
  - **`order_status`** → `fk_order_status` → `admin_schema.order_status_table(order_status_id) ON DELETE SET NULL`.  
  - **`retailer_id`** → `fk_retailer` → `admin_schema.business_table(bid) ON DELETE SET NULL`.  
  - **`wholeseller_id`** → `fk_wholeseller` → `admin_schema.business_table(bid) ON DELETE SET NULL`.  
  - **`location_id`** → `order_table_location_id_fkey` → `admin_schema.master_location(id) ON DELETE SET NULL`.  
  - **`state_id`** → `order_table_state_id_fkey` → `admin_schema.master_states(id) ON DELETE SET NULL`.  

#### **Referenced By**  
- **`wholeseller_rating_table`**  
  - Constraint: `fk_order`  
  - References: `business_schema.order_table(order_id) ON DELETE CASCADE`.  
- **`retailer_rating_table`**  
  - Constraint: `fk_order`  
  - References: `business_schema.order_table(order_id) ON DELETE CASCADE`.  
- **`business_schema.invoice_table`**  
  - Constraint: `invoice_table_order_id_fkey`  
  - References: `business_schema.order_table(order_id) ON DELETE CASCADE`.  
- **`business_schema.order_item_table`**  
  - Constraint: `order_item_table_order_id_fkey`  
  - References: `business_schema.order_table(order_id) ON DELETE CASCADE`.  

#### **Triggers**  
- **`trigger_update_order_history`** → `BEFORE UPDATE` → Calls `update_order_history()`.  
---
### **Table 7: `business_schema.stock_table`**  

#### **Columns & Data Types**  
- **`stock_id`** (`integer`): Primary key, auto-incremented.  
- **`product_id`** (`integer`): References the product.  
- **`stock_in`** (`numeric(10,2)`): Quantity of stock received (default: `0`).  
- **`stock_left`** (`numeric(10,2)`): Remaining stock after sales (default: `0`).  
- **`stock_sold`** (`numeric(10,2)`): Quantity sold (default: `0`).  
- **`date_of_order`** (`timestamp`): Defaults to `CURRENT_TIMESTAMP`.  
- **`user_id`** (`integer`): References the user who manages the stock.  
- **`warehouse_id`** (`integer`): References the warehouse where stock is stored.  
- **`status`** (`boolean`): Stock status (`true` by default).   
- **`carry_fwd`** (`numeric(10,2)`): Stock carried forward (default: `0`).  
- **`source_id`** (`integer`): Reference to the stock source.  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `stock_table_pkey` → `stock_id`.  
- **Foreign Key References**:  
  - **`product_id`** → `stock_table_product_id_fkey` → `admin_schema.master_product(product_id) ON DELETE CASCADE`.  
  - **`user_id`** → `stock_table_user_id_fkey` → `admin_schema.users(user_id) ON DELETE SET NULL`.  
  - **`warehouse_id`** → `fk_category` → `business_schema.warehouse_list(warehouse_id)`.  

---
### **Table 8: `business_schema.warehouse_list`**  

#### **Columns & Data Types**  
- **`warehouse_id`** (`integer`): Primary key, auto-incremented.  
- **`user_id`** (`integer`): ID of the user who manages the warehouse (mandatory).  
- **`warehouse_name`** (`varchar(255)`): Name of the warehouse.  
- **`warehouse_address`** (`text`): Address of the warehouse.  
- **`warehouse_email`** (`varchar(255)`): Email of the warehouse.  
- **`warehouse_phone`** (`varchar(255)`): Contact phone number.  
- **`warehouse_pincode`** (`integer`): Pincode of the warehouse.  
- **`warehouse_status`** (`boolean`): Indicates if the warehouse is active (`true` by default).  
- **`col1`** (`varchar(255)`): Additional field (for metadata or notes).  
- **`col2`** (`varchar(255)`): Additional field.  
- **`remarks`** (`text`): Additional remarks or comments.  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `warehouse_list_pkey` → `warehouse_id`.  

#### **Foreign Key References**  
- **Referenced by**:  
  - **`stock_table`** → `fk_category` → `business_schema.stock_table(warehouse_id)`.  

#### **Relationships**  
- A **Warehouse** can store multiple **Stocks**, as referenced in `stock_table`.  
- Each Warehouse is managed by a **User** (though the user relationship is not explicitly defined in this schema).  
---
### **Table 9: `business_schema.order_history_table`**  

#### **Columns & Data Types**  
- **`history_id`** (`integer`): Primary key, auto-incremented.  
- **`order_id`** (`integer`): ID of the order (mandatory).  
- **`date_of_order`** (`timestamp`): The date when the order was placed (default: `CURRENT_TIMESTAMP`).  
- **`order_status`** (`integer`): Status of the order.  
- **`expected_delivery_date`** (`timestamp`): Expected date of delivery.  
- **`actual_delivery_date`** (`timestamp`): Actual date of delivery.  
- **`retailer_id`** (`integer`): ID of the retailer who placed the order.  
- **`wholeseller_id`** (`integer`): ID of the wholesaler fulfilling the order.  
- **`location_id`** (`integer`): ID of the location associated with the order.  
- **`state_id`** (`integer`): ID of the state related to the order.  
- **`pincode`** (`varchar(10)`): Pincode of the delivery location.  
- **`address`** (`text`): Delivery address.  
- **`delivery_completed_date`** (`timestamp`): Timestamp when the order was marked as delivered.  

#### **Indexes & Constraints**  
- **Primary Key**:  
  - `order_history_table_pkey` → `history_id`.  

#### **Trigger: `update_order_history`**  
- **Function:**  
  The trigger function **`update_order_history`** automatically logs order updates into `order_history_table` whenever an update occurs on `order_table`.  

- **Logic:**  
  - Captures order updates and stores them in `order_history_table`.  
  - If `order_status = 6 (Delivered)`, records `CURRENT_TIMESTAMP` in `delivery_completed_date`.

- **Trigger Execution:**  
  - Trigger executes **before update** on `business_schema.order_table`.  

### Triggers used

1. `update_order_history()` -  The update_order_history trigger function ensures that every change to the order_table is logged in order_history_table, preserving a complete record of order updates. This helps in tracking modifications and monitoring order progress, especially when an order is marked as delivered (status = 6).

2. 