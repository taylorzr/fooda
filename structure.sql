-- create database fooda
-- psql fooda -f structure.sql

drop table if exists users cascade;
drop table if exists orders cascade;

create table users (
  id bigserial primary key,
  email varchar(255) not null,
  handle varchar(255)
);

create table orders (
  id bigserial primary key,
  user_id bigserial references users,
  created_at timestamp with time zone not null default now(),
  delivery_date date not null,
  restaurant varchar(255) not null
);

insert into users(
  email,
  handle
) values (
  'taylorzr@gmail.com',
  'nothotdog'
);

insert into orders(user_id, delivery_date, restaurant)
values (1, '2014-11-15', 'Taco Bell');

insert into orders(user_id, delivery_date, restaurant)
values (1, current_date, 'Burger King');

select * from users join orders on orders.user_id=users.id;
