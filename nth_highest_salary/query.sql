CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    declare
        offst int default 0;
        set offst = N - 1;
  RETURN (
      # Write your MySQL query statement below.
      with highest_salary as (
        select
            distinct(e.salary) as salary
        from Employee as e
        order by e.salary desc
        limit 1
        offset offst
      )
      
      select
        case 
            when h.salary is not null then h.salary
            else null
        end
      from highest_salary as h
            
  );
END