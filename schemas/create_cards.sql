create table if not exists cards (
    id varchar(25) not null unique,
    title varchar(25),
    price float(10, 2),
    description_ varchar(100),
    count_ int,
    picture_path varchar(50),
);