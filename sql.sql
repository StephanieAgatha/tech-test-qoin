--saya menggunakan postgresql karena Skalabilitas dan Performa: PostgreSQL dikenal dengan skalabilitas dan performanya yang baik,
--yang membuatnya cocok untuk aplikasi dengan beban kerja data yang besar.

-- table customers
create table customers (
                           id serial primary key,
                           nama varchar(100),
                           no_hp varchar(13),
                           email varchar(100)
);

-- table menu
create table menu (
                      id serial primary key,
                      menu_name varchar(255),
                      price int,
                      stock int
);

-- table order
create table food_order (
                            id serial primary key,
                            customer_id int REFERENCES customers(id),
                            total_price int
);

-- table order_details
create table order_details (
                               id serial primary key,
                               order_id int REFERENCES food_order(id),
                               menu_id int REFERENCES menu(id),
                               total_menu int
);

-- table payments
create table payments (
                          ID serial primary key,
                          Order_ID int UNIQUE REFERENCES food_order(id),
                          Amount int,
                          Payment_Date DATE
);

-- table reports
create table reports (
                         id serial primary key,
                         report_date DATE,
                         detail_report varchar(250),
                         jenis_laporan varchar(50)
);

-- laporan income mingguan
select DATE_TRUNC('week', payments.Payment_Date) AS week, SUM(food_order.total_price) AS Total_Income
from payments
         join food_order ON payments.Order_ID = food_order.id
group by week
order by week;

-- laporan income bulanan
select DATE_TRUNC('month', payments.Payment_Date) AS Month, SUM(food_order.total_price) AS Total_Income
from payments
    join food_order ON payments.Order_ID = food_order.id
group by Month
order by Month;

-- laporan stock
select menu.menu_name, menu.stock - SUM(order_details.total_menu) AS Remaining_Stock
from menu
         join order_details ON menu.id = order_details.menu_id
group by menu.menu_name, menu.stock;

-- table ratings
CREATE TABLE ratings (
                         id serial primary key,
                         customer_id int REFERENCES customers(id),
                         menu_id int REFERENCES menu(id),
                         rating int CHECK (rating >= 1 AND rating <= 5),
                         review varchar(255)
);

--struk
SELECT
    customers.nama AS cust_name,
    menu.menu_name AS menu_name,
    menu.price AS price,
    order_details.total_menu AS quantity,
    (menu.price * order_details.total_menu) AS Total_Price
FROM
    food_order
JOIN
    customers ON food_order.customer_id = customers.id 
JOIN
    order_details ON food_order.id = order_details.order_id 
JOIN
    menu ON order_details.menu_id = menu.id 
WHERE
    customers.id = 1 AND menu.id = 1;
