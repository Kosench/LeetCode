SELECT w2.id
FROM Weather w1
JOIN Weather w2 ON w2.recordDate = w1.recordDate + INTERVAL '1 day'
WHERE w2.temperature > w1.temperature;

--solution 2
SELECT id
FROM (
    SELECT
        id,
        temperature,
        LAG(temperature) OVER (ORDER BY recordDate) as prev_temp
    FROM Weather
) t
WHERE temperature > prev_temp;
