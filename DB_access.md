## Roles and Permissions in PostgreSQL

### **wholeseller_user Permissions**

| Table Name              | Schema           | Privileges                     |
|-------------------------|-----------------|--------------------------------|
| invoice_details_table   | business_schema | INSERT, SELECT, UPDATE        |
| invoice_table          | business_schema | INSERT, SELECT, UPDATE        |
| order_table            | business_schema | SELECT                         |
| order_item_table       | business_schema | SELECT                         |

---

### **wholeseller_admin Permissions**

##### admin_schema

| Table Name          | Privileges           |
|---------------------|---------------------|
| master_mandi_table | INSERT, SELECT, UPDATE, REFERENCES |
| business_table     | INSERT              |

##### business_schema

| Table Name               | Privileges                                |
|--------------------------|------------------------------------------|
| daily_price_update       | INSERT, SELECT, UPDATE, DELETE, REFERENCES, TRIGGER |
| invoice_details_table    | INSERT, SELECT, UPDATE, DELETE, TRIGGER  |
| order_history_table      | INSERT, SELECT, UPDATE, DELETE          |
| stock_table              | INSERT, SELECT, UPDATE, DELETE          |
| invoice_table            | INSERT, SELECT, UPDATE, DELETE, TRIGGER  |
| order_table              | INSERT, SELECT, UPDATE, DELETE          |
| warehouse_list           | INSERT, SELECT, UPDATE, DELETE          |
| order_item_table         | INSERT, SELECT, UPDATE, DELETE          |


---

### **Retailer Permissions**

##### admin_schema

| Table Name              | Schema           | Privileges              |
|-------------------------|-----------------|-------------------------|
| master_category_table   | admin_schema    | SELECT                  |
| master_product         | admin_schema    | SELECT                  |
| business_table         | admin_schema    | SELECT                  |
| cash_payment_list      | admin_schema    | SELECT                  |
| mode_of_payments_list  | admin_schema    | SELECT                  |
|---------------------------------------------------------------------|

##### business_schema

| Table Name              | Schema           | Privileges              |
|-------------------------|-----------------|-------------------------|
| daily_price_update      | business_schema | SELECT                  |
| invoice_details_table   | business_schema | SELECT                  |
| stock_table            | business_schema | SELECT                  |
| invoice_table         | business_schema | SELECT                  |
| order_table            | business_schema | INSERT, SELECT          |
| order_item_table       | business_schema | INSERT, SELECT, UPDATE  |
|---------------------------------------------------------------------|

---

### **Admin Permissions**

| Table Name              | Privileges                                    |
|-------------------------|----------------------------------------------|
| master_category_table  | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| master_mandi_table      | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| master_driver_table     | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| business_type_table     | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| master_vehicle_table    | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| master_violation_table  | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| vehicle_engine_type     | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| vehicle_insurance_table | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| user_table              | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| mode_of_payment         | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| order_status_table      | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER|
| master_product         | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| units_table             | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| master_states          | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| business_table          | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| cash_payment_list       | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| master_location        | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| user_type_table         | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| users                   | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |
| warehouse_list          | SELECT                                       |
| mode_of_payments_list   | INSERT, SELECT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER |


