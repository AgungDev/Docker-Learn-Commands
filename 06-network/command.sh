# Melihat docker netowrk
docker network ls

# Membuat docker netowrk
docker network create --driver bridge contohnetwork

# Menhapus network
docker network rm contohnetwork

# Contoh
docker network create databaseMysql
docker run --name mq -e MYSQL_ROOT_PASSWORD=admin -d mysql
docker network connect databaseMysql mq
docker run --name phpmyadmin --network databaseMysql -e PMA_HOST=mq \
 -p 8082:80 -d phpmyadmin
