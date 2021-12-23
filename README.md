### Database commands
show databases;

show tables;

use rip_db; select * from checkList;
### Docker database volume delete
docker volume ls

docker volume rm rip_mysql_data
### Docker enter in bd container
docker exec -it Rip_db mysql -uroot -p
###