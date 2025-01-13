-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE dishes (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    calories INT NOT NULL,
    proteins FLOAT NOT NULL,
    fats FLOAT NOT NULL,
    carbohydrates FLOAT NOT NULL
);

INSERT INTO dishes (id, name, calories, proteins, fats, carbohydrates) VALUES
    (uuid_generate_v4(), 'Grilled Chicken Breast', 165, 31.0, 3.6, 0.0),
    (uuid_generate_v4(), 'Beef Steak', 271, 25.0, 18.0, 0.0),
    (uuid_generate_v4(), 'Salmon Fillet', 206, 22.0, 13.0, 0.0),
    (uuid_generate_v4(), 'Greek Salad', 120, 3.0, 9.0, 7.0),
    (uuid_generate_v4(), 'Oatmeal with Fruits', 150, 5.0, 3.0, 27.0),
    (uuid_generate_v4(), 'Avocado Toast', 200, 6.0, 10.0, 18.0),
    (uuid_generate_v4(), 'Vegetable Stir Fry', 180, 6.0, 8.0, 22.0),
    (uuid_generate_v4(), 'Protein Shake', 220, 25.0, 2.0, 10.0),
    (uuid_generate_v4(), 'Turkey Wrap', 250, 20.0, 6.0, 30.0),
    (uuid_generate_v4(), 'Egg Omelette', 190, 14.0, 14.0, 2.0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE dishes;
DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
