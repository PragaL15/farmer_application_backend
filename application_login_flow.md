# Role-Based Access Control (RBAC) System

## Tables

### 1. `roles_table`
Stores the roles in the application. Each role has a unique `role_id` and a `role_name`.

| Column     | Type               | Description                     |
|------------|--------------------|---------------------------------|
| `role_id`  | `integer`          | Unique ID for the role.         |
| `role_name`| `varchar(50)`      | Name of the role (e.g., Admin). |

---

### 2. `user_table`
Stores user details and maps each user to a role using `role_id`.

| Column        | Type               | Description                          |
|---------------|--------------------|--------------------------------------|
| `user_id`     | `integer`          | Unique ID for the user.              |
| `name`        | `varchar(255)`     | Full name of the user.               |
| `mobile_num`  | `varchar(15)`      | Mobile number of the user.           |
| `email`       | `varchar(255)`     | Email address of the user.           |
| `address`     | `text`             | Physical address of the user.        |
| `pincode`     | `varchar(10)`      | Pincode of the user’s location.      |
| `location`    | `integer`          | Location ID of the user.             |
| `state`       | `integer`          | State ID of the user.                |
| `active_status`| `integer`         | Status of the user (0 = inactive, 1 = active). |
| `role_id`     | `integer`          | Role ID of the user (default = 5).   |

---

### 3. `business_table`
Stores business-related data and is linked to a role via `role_id`.

| Column               | Type               | Description                          |
|----------------------|--------------------|--------------------------------------|
| `bid`                | `integer`          | Unique ID for the business.          |
| `b_type_id`          | `integer`          | Type ID of the business.             |
| `b_location_id`      | `integer`          | Location ID of the business.         |
| `b_state_id`         | `integer`          | State ID of the business.            |
| `b_mandi_id`         | `integer`          | Mandi ID of the business.            |
| `b_address`          | `varchar(255)`     | Address of the business.             |
| `b_phone_num`        | `varchar(15)`      | Phone number of the business.        |
| `b_email`            | `varchar(255)`     | Email address of the business.       |
| `b_gst_num`          | `varchar(255)`     | GST number of the business.          |
| `b_pan_num`          | `varchar(255)`     | PAN number of the business.          |
| `b_person_name`      | `varchar(255)`     | Contact person’s name.               |
| `active_status`      | `integer`          | Status of the business (0 = inactive, 1 = active). |
| `b_registration_num` | `varchar(255)`     | Registration number of the business. |
| `b_owner_name`       | `varchar(255)`     | Owner’s name of the business.        |
| `b_established_year` | `integer`          | Year the business was established.   |
| `b_shop_name`        | `varchar(255)`     | Name of the business shop.           |

---

### 4. `permissions_table`
Defines what actions each role can perform on the `business_table`.

| `role_id` | `resource_name`   | `can_create` | `can_read` | `can_update` | `can_delete` |
|-----------|-------------------|--------------|------------|--------------|--------------|
| 1         | `business_table`  | ✅           | ✅         | ✅           | ✅           |
| 2         | `business_table`  | ✅           | ✅         | ✅           | ❌           |
| 3         | `business_table`  | ❌           | ✅         | ✅           | ❌           |
| 4         | `business_table`  | ❌           | ✅         | ❌           | ❌           |
| 5         | `business_table`  | ❌           | ✅         | ❌           | ❌           |

---
## How It Works: Step-by-Step

### Step 1: User Logs In
1. The user enters their **username** and **password** (like logging into a game).
2. The system checks if the username and password are correct by looking at the `business_user_table` (like a list of players).
3. If the username and password are correct, the system finds the user’s **role** (like their job title) from the `users`.

---

### Step 2: Check What the User Can Do
1. The system looks at the `permissions_table` to see what the user’s role is allowed to do.
   - For example, if the user is a **Retailer**, the system checks if they can **read**, **update**, **create**, or **delete** records in the `business_table`.
2. The `permissions_table` is like a **rulebook** that says:
   - **Admins** can do everything.
   - **Retailers** can read and update but can’t create or delete.
   - **Users** can only read.

---

### Step 3: Adjust the App’s Buttons
1. Based on the permissions, the system shows or hides buttons in the app.
   - For example:
     - If the user is a **Retailer**, they will see **View** and **Edit** buttons but not **Delete**.
     - If the user is an **Admin**, they will see all buttons: **Create**, **Edit**, and **Delete**.

---

### Step 4: Handle User Actions
1. When the user tries to do something (like update a record), the system checks the **rulebook** again to see if they’re allowed.
2. If the user has permission:
   - The system allows the action and does what the user asked (like updating a record).
3. If the user does **not** have permission:
   - The system blocks the action and shows an error like **"Access Denied."**

---

### Step 5: Log the Action
1. If the action is allowed, the system writes it down in a **log** (like a diary) for security purposes.
   - For example:
     - If a Retailer updates a record, the system writes:
       - **Who**: Retailer
       - **What**: Updated a record
       - **When**: Today at 3:00 PM

---

## Example Scenarios

### Example 1: Admin Logs In
1. **User**: `admin1` (Admin).
2. **Permissions**:
   - Can **create**, **read**, **update**, and **delete** records.
3. **UI**:
   - Shows all buttons: **Create**, **Edit**, **Delete**.
4. **Actions**:
   - Can do anything in the app.

---

### Example 2: Retailer Logs In
1. **User**: `retailer1` (Retailer).
2. **Permissions**:
   - Can **read** and **update** records.
   - Cannot **create** or **delete** records.
3. **UI**:
   - Shows **View** and **Edit** buttons.
   - Hides **Create** and **Delete** buttons.
4. **Actions**:
   - Can view and edit records but can’t create or delete them.

---

### Example 3: User Logs In
1. **User**: `user1` (User).
2. **Permissions**:
   - Can only **read** records.
   - Cannot **create**, **update**, or **delete** records.
3. **UI**:
   - Shows only the **View** button.
   - Hides **Create**, **Edit**, and **Delete** buttons.
4. **Actions**:
   - Can only look at records. Can’t make any changes.

---

## Why Is This Important?

This system ensures that:
- **Admins** are like **game developers**—they can do anything.
- **Retailers** are like **moderators**—they can help but can’t change everything.
- **Users** are like **players**—they can play the game but can’t change the rules.

It keeps the app **safe** and **organized** so that everyone can only do what they’re supposed to do.

---

## How this works along backend?

1. Whenever an API is called, the backend checks the user’s role_id and permissions.
2. If the user has permission, the backend performs the action.
3. If the user does not have permission, the backend rejects the request and sends an error.