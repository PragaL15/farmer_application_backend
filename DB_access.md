### Roles and Permissions in PostgreSQL

#### **wholeseller_user Permissions**

| Table Name              | Privileges         |
|-------------------------|--------------------|
| invoice_details_table   | INSERT, SELECT, UPDATE |
| order_item_table        | SELECT |
| invoice_table           | INSERT, SELECT, UPDATE |
| order_table             | SELECT |

#### **wholeseller_admin Permissions**

| Table Name              | Privileges         |
|-------------------------|--------------------|
| daily_price_update      | INSERT, SELECT, UPDATE, REFERENCES, TRIGGER |
| master_mandi_table      | INSERT, SELECT, UPDATE, REFERENCES |
| invoice_details_table   | INSERT, SELECT, UPDATE |
| business_table          | INSERT |
| invoice_table           | INSERT, SELECT, UPDATE |
| order_item_table        |SELECT|
| order_table             |SELECT|

#### **Retailer Permissions**

| Grantee  | Privilege Type | Table Name                |
|----------|---------------|---------------------------|
| retailer | SELECT        | master_category_table     |
| retailer | SELECT        | daily_price_update        |
| retailer | SELECT , UPDATE      | invoice_details_table     |
| retailer | SELECT        | stock_table               |
| retailer | SELECT        | master_product            |
| retailer | SELECT        | business_table            |
| retailer | SELECT , INSERT       | order_table               |
| retailer | INSERT,SELECT, UPDATE        | order_item_table          |
| retailer | SELECT        | cash_payment_list         |
| retailer | SELECT        | mode_of_payments_list     |
| retailer | SELECT        | invoice_table             |

#### **admin Permissions**

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


