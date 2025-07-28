Решение 1
SELECT c.name AS Customers,
    o.id
FROM Customers c
LEFT JOIN Orders o ON c.id = o.customerId
WHERE o.id IS NULL;

Решение 2
SELECT
    Customers.name AS Customers
FROM Customers
WHERE Customers.id NOT IN (
    SELECT Orders.customerId
    FROM Orders
)
