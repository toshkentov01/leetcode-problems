with info as (
    select
        row_number() over (order by l1.id) as row_num,
        l1.num as number
    from Logs as l1
)

select 
    distinct i1.number as ConsecutiveNums
from
    info as i1, info as i2, info as i3
where
    i1.row_num = i2.row_num - 1 and
    i2.row_num = i3.row_num - 1 and
    i1.number = i2.number and
    i2.number = i3.number



-- The reason why I used the window function is that, if I checked according to id, it would not work for condition where ids are not consecutive