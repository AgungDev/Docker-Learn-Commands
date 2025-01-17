# bind mounts
docker create --name mq --mount "type=bind,source=/media/fun5i/Database/Docker/Learn/04-bind-mounts/example,destination=/var/lib/mysql" -p 5555:3306 -e MYSQL_ROOT_PASSWORD=admin mysql
