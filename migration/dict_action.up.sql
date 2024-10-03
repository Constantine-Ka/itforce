CREATE TABLE IF NOT EXISTS public.dict_action(
    id     SERIAL PRIMARY KEY,
    eng varchar(100) constraint eng unique,
    rus varchar(100)
);
INSERT INTO public.dict_action (eng, rus) VALUES ('Sel','Продажа'),('Purchase','Покупка') ON CONFLICT(eng) DO NOTHING;
