create table users (
    "id" serial primary key,
    "email" varchar(256) not null unique,
    "password_hash" text not null,
    "first_name" varchar(256) null,
    "last_name" varchar(256) null,
    "middle_name" varchar(256) null,
    "telephone" varchar(15) null unique,
    "address" varchar(512) null,
    "date_birthday" date null,
    "country_birth" varchar(64) null,
    is_blocked boolean not null default false,
    is_active boolean not null default true,
    create_at timestamp default current_timestamp
);
