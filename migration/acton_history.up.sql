create table if not exists public.action_history
(
    id     SERIAL PRIMARY KEY,
    user_id  integer,
    action   integer,
    amount   NUMERIC(10, 2),
    ts TIMESTAMP default NOW() not null
);

alter table public.action_history
    owner to docker;

create index if not exists user_id
    on public.action_history (user_id);

create index if not exists action
    on public.action_history (action);



