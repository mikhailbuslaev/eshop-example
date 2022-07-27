create table if not exists cards (
    id varchar(25) not null unique,
    title varchar(25),
    price decimal(10, 2),
    description varchar(100),
    count int,
    picturepath varchar(50)
);