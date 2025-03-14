## wholeseller_admin , wholeseller_user


| Table Name              | `wholeseller_admin` (Update Prices) | `wholeseller_user` (Manage Orders & Invoices) |
|-------------------------|---------------------------------|------------------------------------|
| `daily_price_update`    | ✅ `INSERT, UPDATE` | ❌ No Access |
| `master_mandi_table`    | ✅ `SELECT, UPDATE` | ❌ No Access |
| `invoice_details_table` | ✅ `SELECT, INSERT` | ✅ `SELECT, INSERT, UPDATE` |
| `invoice_table`         | ✅ `SELECT, INSERT` | ✅ `SELECT, INSERT, UPDATE` |
| `order_table`           | ❌ No Access | ✅ `SELECT` |
| `order_item_table`      | ❌ No Access | ✅ `SELECT` |

---

## retailers 

| Table Name                  | `retailer` (View & Place Orders) |
|-----------------------------|--------------------------------|
| `master_category_table`     | ✅ `SELECT` |
| `master_product`            | ✅ `SELECT` |
| `daily_price_update`        | ✅ `SELECT` |
| `order_table`               | ✅ `INSERT, SELECT` |
| `order_item_table`          | ✅ `INSERT, SELECT`,`UPDATE` |
| `master_states`             | ✅ `SELECT` |
| `master_location`           | ✅ `SELECT` |
| `business_table`            | ✅ `SELECT` |
| `cash_payment_list`         | ✅ `SELECT` |
| `business_type_table`       | ✅ `SELECT` |
| `mode_of_payments_list`     | ✅ `SELECT` |
| `stock_table`               | ✅ `SELECT` |
| `invoice_table`             | ✅ `SELECT` |
| `invoice_details_table`     | ✅ `SELECT` |

