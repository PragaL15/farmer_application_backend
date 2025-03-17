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

