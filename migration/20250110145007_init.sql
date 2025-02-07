-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE dishes (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    calories INT NOT NULL,
    proteins FLOAT NOT NULL,
    fats FLOAT NOT NULL,
    carbohydrates FLOAT NOT NULL
);

INSERT INTO dishes (id, name, description, calories, proteins, fats, carbohydrates) VALUES
    (uuid_generate_v4(), 'Grilled Chicken Breast', 'Juicy grilled chicken breast seasoned with herbs.', 165, 31.0, 3.6, 0.0),
    (uuid_generate_v4(), 'Beef Steak', 'Tender beef steak grilled to perfection.', 271, 25.0, 18.0, 0.0),
    (uuid_generate_v4(), 'Baked Salmon', 'Oven-baked salmon with lemon and dill.', 206, 22.0, 12.0, 0.0),
    (uuid_generate_v4(), 'Quinoa Salad', 'Quinoa mixed with fresh vegetables and olive oil.', 120, 4.1, 3.5, 21.3),
    (uuid_generate_v4(), 'Vegetable Stir Fry', 'A mix of stir-fried vegetables with soy sauce.', 98, 3.0, 5.0, 11.0),
    (uuid_generate_v4(), 'Greek Yogurt with Berries', 'Low-fat Greek yogurt topped with fresh berries.', 150, 10.0, 2.0, 20.0),
    (uuid_generate_v4(), 'Avocado Toast', 'Whole-grain toast with mashed avocado and a pinch of salt.', 180, 5.0, 10.0, 20.0),
    (uuid_generate_v4(), 'Egg White Omelette', 'Fluffy egg white omelette with spinach and mushrooms.', 110, 20.0, 3.0, 2.0),
    (uuid_generate_v4(), 'Grilled Shrimp', 'Shrimp grilled with garlic and olive oil.', 120, 23.0, 2.0, 1.0),
    (uuid_generate_v4(), 'Brown Rice Bowl', 'Steamed brown rice with steamed broccoli and chicken.', 250, 20.0, 5.0, 35.0),
    (uuid_generate_v4(), 'Lentil Soup', 'Hearty lentil soup with carrots and celery.', 180, 12.0, 3.0, 30.0),
    (uuid_generate_v4(), 'Turkey Meatballs', 'Lean turkey meatballs with marinara sauce.', 200, 22.0, 7.0, 8.0),
    (uuid_generate_v4(), 'Zucchini Noodles', 'Spiralized zucchini with tomato basil sauce.', 90, 2.0, 1.0, 8.0),
    (uuid_generate_v4(), 'Stuffed Bell Peppers', 'Bell peppers stuffed with quinoa, beans, and spices.', 190, 7.0, 4.0, 30.0),
    (uuid_generate_v4(), 'Cottage Cheese Bowl', 'Low-fat cottage cheese with sliced cucumber and herbs.', 140, 15.0, 5.0, 5.0),
    (uuid_generate_v4(), 'Chickpea Salad', 'Chickpeas mixed with parsley, lemon juice, and olive oil.', 210, 8.0, 7.0, 30.0),
    (uuid_generate_v4(), 'Tuna Salad', 'Canned tuna with lettuce, cucumber, and olive oil dressing.', 200, 25.0, 7.0, 2.0),
    (uuid_generate_v4(), 'Roasted Sweet Potatoes', 'Oven-roasted sweet potato wedges with paprika.', 180, 2.0, 4.0, 36.0),
    (uuid_generate_v4(), 'Spinach Smoothie', 'Spinach blended with banana, almond milk, and chia seeds.', 160, 3.0, 5.0, 25.0),
    (uuid_generate_v4(), 'Grilled Turkey Breast', 'Grilled turkey breast with rosemary seasoning.', 135, 28.0, 1.0, 0.0),
    (uuid_generate_v4(), 'Broccoli Soup', 'Creamy broccoli soup with almond milk.', 100, 5.0, 3.0, 12.0),
    (uuid_generate_v4(), 'Vegetable Frittata', 'Egg-based frittata with mixed vegetables.', 220, 14.0, 12.0, 12.0),
    (uuid_generate_v4(), 'Edamame Snack', 'Steamed edamame beans with a sprinkle of salt.', 120, 11.0, 5.0, 10.0),
    (uuid_generate_v4(), 'Cauliflower Rice', 'Grated cauliflower sautéed with garlic and herbs.', 50, 2.0, 1.0, 8.0),
    (uuid_generate_v4(), 'Turkey Lettuce Wraps', 'Ground turkey wrapped in fresh lettuce leaves.', 180, 20.0, 8.0, 4.0),
    (uuid_generate_v4(), 'Grilled Asparagus', 'Asparagus grilled with olive oil and lemon juice.', 60, 3.0, 2.0, 5.0),
    (uuid_generate_v4(), 'Oven-Baked Cod', 'Cod fillet baked with herbs and lemon.', 90, 20.0, 1.0, 0.0),
    (uuid_generate_v4(), 'Pumpkin Soup', 'Smooth pumpkin soup with a hint of ginger.', 130, 3.0, 4.0, 20.0),
    (uuid_generate_v4(), 'Tofu Stir Fry', 'Stir-fried tofu with bell peppers and soy sauce.', 220, 15.0, 12.0, 18.0),
    (uuid_generate_v4(), 'Roasted Brussels Sprouts', 'Brussels sprouts roasted with olive oil and balsamic glaze.', 110, 4.0, 6.0, 10.0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE dishes;
DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
