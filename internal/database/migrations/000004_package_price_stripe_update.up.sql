ALTER TABLE package_price
ADD stripe_product_id TEXT;
ALTER TABLE auth.users
ADD stripe_customer_id TEXT;