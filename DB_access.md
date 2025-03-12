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