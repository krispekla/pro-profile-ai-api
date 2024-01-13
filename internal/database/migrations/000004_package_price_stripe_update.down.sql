ALTER TABLE package_price DROP COLUMN stripe_product_id;
ALTER TABLE auth.users DROP COLUMN stripe_customer_id,
    DROP COLUMN first_name,
    DROP COLUMN last_name;
ALTER TABLE package_order DROP COLUMN stripe_payment_intent_id;