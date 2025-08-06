--solution 1
DELETE FROM Person
WHERE id NOT IN (
    SELECT MIN(id)
    FROM (SELECT * FROM Person) AS temp
    GROUP BY email
);

--solution 2
WITH keep_ids AS (
    SELECT MIN(id) AS min_id
    FROM Person
    GROUP BY email
)
DELETE FROM Person
WHERE id NOT IN (SELECT min_id FROM keep_ids);
