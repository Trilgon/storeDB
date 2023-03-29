create table if not exists public.goods
(
    goods_id integer     not null
        primary key,
    name     varchar(40) not null,
    price    integer     not null,
    quantity integer     not null
);

alter table public.goods
    owner to postgres;

create table if not exists public.carts
(
    cart_id integer not null
        primary key,
    total   integer not null
);

alter table public.carts
    owner to postgres;

create table if not exists public.orders
(
    order_id    integer                  not null
        primary key,
    total       integer                  not null,
    order_time  timestamp with time zone not null,
    finish_time timestamp with time zone
);

alter table public.orders
    owner to postgres;

create table if not exists public.goods_to_orders
(
    order_id integer not null
        constraint fk_goods_to_orders__order_id
            references public.orders,
    goods_id integer not null
        constraint fk_goods_to_orders__goods_id
            references public.goods,
    quantity integer not null,
    primary key (order_id, goods_id)
);

alter table public.goods_to_orders
    owner to postgres;

create table if not exists public.goods_to_carts
(
    cart_id  integer not null
        constraint fk_goods_to_carts__cart_id
            references public.carts,
    goods_id integer not null
        constraint fk_goods_to_carts__goods_id
            references public.goods,
    quantity integer not null,
    primary key (cart_id, goods_id)
);

alter table public.goods_to_carts
    owner to postgres;

