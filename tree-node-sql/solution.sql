with parents as (
    select
        p_id as parent_id
    from Tree
)

select 
    t.id as id,
    case when p_id is null then
        'Root'
    when t.id in (select parent_id from parents) then
        'Inner'
    else
        'Leaf'
    end as type
from 
    Tree as t