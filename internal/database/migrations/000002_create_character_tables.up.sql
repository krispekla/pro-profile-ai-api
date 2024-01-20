DO $$ BEGIN IF NOT EXISTS (
    SELECT 1
    FROM pg_type
    WHERE typname = 'gender'
) THEN CREATE TYPE gender AS ENUM ('male', 'female', 'other');
END IF;
END $$;
DO $$ BEGIN IF NOT EXISTS (
    SELECT 1
    FROM pg_type
    WHERE typname = 'hair_color'
) THEN CREATE TYPE hair_color AS ENUM (
    'black',
    'brown',
    'blond',
    'gray',
    'red',
    'nohair',
    'other'
);
END IF;
END $$;
DO $$ BEGIN IF NOT EXISTS (
    SELECT 1
    FROM pg_type
    WHERE typname = 'eye_color'
) THEN CREATE TYPE eye_color AS ENUM ('brown', 'blue', 'green', 'gray');
END IF;
END $$;
DO $$ BEGIN IF NOT EXISTS (
    SELECT 1
    FROM pg_type
    WHERE typname = 'ethnicity'
) THEN CREATE TYPE ethnicity AS ENUM ('light', 'medium', 'dark', 'other');
END IF;
END $$;
DO $$ BEGIN IF NOT EXISTS (
    SELECT 1
    FROM pg_type
    WHERE typname = 'age'
) THEN CREATE TYPE age AS ENUM ('child', 'adult', 'elderly');
END IF;
END $$;
CREATE TABLE IF NOT EXISTS character (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES auth.users (id) NOT NULL,
    name TEXT,
    gender gender NOT NULL DEFAULT 'other',
    hair_color hair_color NOT NULL DEFAULT 'other',
    eye_color eye_color NOT NULL DEFAULT 'brown',
    ethnicity ethnicity NOT NULL DEFAULT 'other',
    age age NOT NULL DEFAULT 'adult',
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS character_img (
    id SERIAL PRIMARY KEY,
    character_id INT REFERENCES character (id) NOT NULL,
    img_url TEXT NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS character_model (
    id SERIAL PRIMARY KEY,
    character_id INT NOT NULL,
    model_url TEXT NOT NULL
);