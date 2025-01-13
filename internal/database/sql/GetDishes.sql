SELECT *
FROM dishes
WHERE calories <= $1 AND proteins <= $2 AND fats <= $3 AND carbohydrates <= $4
ORDER BY RANDOM() LIMIT 5;
