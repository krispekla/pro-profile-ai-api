ALTER TABLE package_price
ADD COLUMN stripe_product_id TEXT;
ALTER TABLE auth.users
ADD COLUMN stripe_customer_id TEXT,
    ADD COLUMN first_name TEXT,
    ADD COLUMN last_name TEXT;