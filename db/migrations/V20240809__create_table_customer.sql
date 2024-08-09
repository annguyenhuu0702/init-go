create table customer (
    "id" uuid primary key default uuid_generate_v4(),
    "name" varchar(255) not null,
    "phone_number" VARCHAR(255) not null unique,
    "email" VARCHAR(255) null,
    "created_at" timestamptz default current_timestamp,
    "updated_at" timestamptz default current_timestamp
);