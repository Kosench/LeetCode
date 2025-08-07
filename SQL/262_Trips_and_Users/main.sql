WITH unbanned_trips AS (
    SELECT
        t.request_at::DATE AS trip_date,
        t.status
    FROM Trips t
    JOIN Users u1 ON t.client_id = u1.users_id
    JOIN Users u2 ON t.driver_id = u2.users_id
    WHERE
        u1.banned = 'No'
        AND u2.banned = 'No'
        AND t.request_at BETWEEN '2013-10-01' AND '2013-10-03'
)
SELECT
    trip_date AS "Day",
    ROUND(
        COUNT(CASE WHEN status LIKE 'cancelled%' THEN 1 END) * 1.0 / COUNT(*),
        2
    ) AS "Cancellation Rate"
FROM unbanned_trips
GROUP BY trip_date
ORDER BY trip_date;
