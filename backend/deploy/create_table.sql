create table if not exists cards (
    id varchar(25) not null unique,
    title varchar(50),
    price decimal(10, 2),
    description varchar(500),
    count int,
    picturepath varchar(100)
);