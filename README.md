###
use test; select * from checkList;
show databases;
show tables;

###
docker volume ls
docker volume rm golang-docker-tutorial_mysql_data

###
docker exec -it test_db mysql -uroot -p

###