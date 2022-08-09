#!/bin/bash
echo "Running postgres container...";
docker run --add-host database:127.0.0.1 --name eshop_test -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres;
echo "Connecting create_db file...";
docker cp ./create_db.sql eshop_test:/create_db.sql;
echo "Execute create_db file...";
docker exec -it eshop_test psql -U postgres -f /create_db.sql;

echo "Connecting create_table file...";
docker cp ./create_table.sql eshop_test:/create_table.sql;
echo "Execute create_table file...";
docker exec -it eshop_test psql -U postgres -d eshop_test -f /create_table.sql;
echo "Migration done!";
sleep(2);
