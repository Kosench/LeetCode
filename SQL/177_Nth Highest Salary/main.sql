CREATE OR REPLACE FUNCTION NthHighestSalary(N INT) RETURNS TABLE (Salary INT) AS $$
BEGIN
RETURN QUERY (
    -- Write your PostgreSQL query statement below.
    SELECT(
        CASE WHEN N > 0 THEN (
            SELECT DISTINCT e.salary
            FROM Employee e
            ORDER BY e.salary DESC
            OFFSET N-1
            LIMIT 1
        )
        ELSE NULL
        END
    )
  );
END;
$$ LANGUAGE plpgsql;