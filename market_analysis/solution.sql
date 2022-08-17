select
    u.user_id as buyer_id,
    u.join_date,
    (
        select count(o.buyer_id) from Orders as o where year(o.order_date) = '2019' and o.buyer_id = u.user_id
    ) as 'orders_in_2019'

from
    Users as u