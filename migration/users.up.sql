create table if not exists public.users
(
    id     SERIAL PRIMARY KEY,
    name   varchar(500),
    balance NUMERIC(10, 2)
    );

alter table public.users
    owner to docker;

create index if not exists city
    on public.users (balance);

INSERT INTO users (name, balance) VALUES ('user1', 5000.55 ),('user2',1000)


