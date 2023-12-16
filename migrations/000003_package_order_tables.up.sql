CREATE TABLE IF NOT EXISTS package (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    cover_img_url TEXT NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS package_price (
    id SERIAL PRIMARY KEY,
    package_id INT REFERENCES package (id) NOT NULL,
    amount INT NOT NULL,
    currency TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS coupon (
    id SERIAL PRIMARY KEY,
    package_id INT REFERENCES package (id) NOT NULL,
    code TEXT NOT NULL,
    amount INT,
    currency TEXT,
    percentage INT CHECK (
        percentage IS NULL
        OR (
            percentage > 0
            AND percentage <= 100
        )
    ),
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CHECK (
        (
            amount IS NOT NULL
            AND currency IS NOT NULL
        )
        OR percentage IS NOT NULL
    )
);
CREATE TABLE IF NOT EXISTS package_example_img (
    id SERIAL PRIMARY KEY,
    package_id INT REFERENCES package (id) NOT NULL,
    img_url TEXT NOT NULL
);
DO $$ BEGIN IF NOT EXISTS (
    SELECT 1
    FROM pg_type
    WHERE typname = 'order_status'
) THEN CREATE TYPE order_status AS ENUM (
    'created',
    'pending',
    'paid',
    'cancelled',
    'refunded'
);
END IF;
END $$;
CREATE TABLE IF NOT EXISTS package_order (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES auth.users (id) NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    total_amount INT NOT NULL,
    currency TEXT NOT NULL,
    status order_status NOT NULL,
    coupon_id INT REFERENCES coupon (id)
);
CREATE TABLE IF NOT EXISTS package_order_item (
    id SERIAL PRIMARY KEY,
    package_order_id INT REFERENCES package_order (id) NOT NULL,
    package_id INT REFERENCES package (id) NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS generated_package (
    id SERIAL PRIMARY KEY,
    package_order_item_id INT REFERENCES package_order_item (id) NOT NULL,
    character_id INT REFERENCES character (id) NOT NULL,
    cover_img_url TEXT NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS generated_package_img (
    id SERIAL PRIMARY KEY,
    generated_package_id INT REFERENCES generated_package (id) NOT NULL,
    img_url TEXT NOT NULL,
    model_id INT REFERENCES character_model (id) NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);