CREATE TYPE gender AS ENUM ('male', 'female', 'other');
CREATE TYPE hair_color AS ENUM (
    'black',
    'brown',
    'blond',
    'gray',
    'red',
    'nohair',
    'other'
);
CREATE TYPE eye_color AS ENUM ('brown', 'blue', 'green', 'gray');
CREATE TYPE ethnicity AS ENUM ('light', 'medium', 'dark', 'other');
CREATE TYPE age AS ENUM ('child', 'adult', 'elderly');
CREATE TABLE IF NOT EXISTS character (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES user_account (id) NOT NULL,
    name TEXT NOT NULL,
    gender gender NOT NULL,
    hair_color hair_color NOT NULL,
    eye_color eye_color NOT NULL,
    ethnicity ethnicity NOT NULL,
    age age NOT NULL,
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