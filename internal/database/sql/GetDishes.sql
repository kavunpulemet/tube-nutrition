WITH target_dishes AS (
    SELECT *
    FROM dishes
    ORDER BY
        CASE
            WHEN $1 = 'gain' THEN calories
            WHEN $1 = 'lose' THEN -calories
            ELSE NULL
        END DESC NULLS LAST,
    RANDOM()
    LIMIT 20
    )
SELECT *
FROM target_dishes
ORDER BY RANDOM()
LIMIT 5;