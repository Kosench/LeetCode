SELECT
    d.name AS Department,
    e.name AS Employee,
    e.salary AS Salary
FROM (
    SELECT
        id,
        name,
        salary,
        departmentId,
        MAX(salary) OVER (PARTITION BY departmentId) AS max_salary_in_dept
    FROM Employee
) e
JOIN Department d ON e.departmentId = d.id
WHERE e.salary = e.max_salary_in_dept;
