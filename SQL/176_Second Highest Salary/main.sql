select(select distinct salary as SecondHighestSalary
       from employee
       order by salary desc
       offset 1 limit 1)