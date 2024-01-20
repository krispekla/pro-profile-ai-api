ALTER TABLE package_price
ADD COLUMN stripe_product_id TEXT;
ALTER TABLE auth.users
ADD COLUMN stripe_customer_id TEXT,
    ADD COLUMN first_name TEXT,
    ADD COLUMN last_name TEXT;
ALTER TABLE package_order
ADD COLUMN stripe_payment_intent_id TEXT,
    ADD COLUMN stripe_checkout_session_id TEXT;
ALTER TABLE coupon
ADD COLUMN stripe_coupon_id TEXT;