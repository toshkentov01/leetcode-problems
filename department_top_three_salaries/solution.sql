with maxSalariesByDepartment as (
    select
        e.name as name,
        e.salary as salary,
        d.name as department,
    
        dense_rank() over(
            partition by d.name
            order by e.salary desc
        ) as 'rank'
    
    from Employee as e
    join Department as d on d.id = e.departmentId
)

select
    mxb.department as Department,
    mxb.name as Employee,
    mxb.salary as Salary
from
    maxSalariesByDepartment as mxb
where 3 >= mxb.rank