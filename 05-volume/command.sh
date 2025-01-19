# Melihat Volume
docker volume ls

# membuat volume
docker volume create contohvolume
docker create --name mq --mount "type=volume,source=contohvolume,destination=/var/lib/mysql" -p 5555:3306 -e MYSQL_ROOT_PASSWORD=admin mysql
docker create --name mq --mount "type=volume,source=mq,destination=/var/lib/mysql" \
 -p 5555:3306 -e MYSQL_ROOT_PASSWORD=admin mysql

# menghapus volume
docker volume rm contohvolume
docker volume rm -f $(docker volume ls -q) # remove all volume

# backup volume
docker run --rm --name ubuntubackup \
  --mount "type=bind,source=$(pwd)/05-volume/example,destination=/backup" \
  --mount "type=volume,source=mq,destination=/data" \
  ubuntu:latest tar cvf /backup/backup.tar.gz /data

# backup restore
docker volume create restoremq
docker run --rm --name ubuntubackup --mount "type=bind,source=$(pwd)/05-volume/example,destination=/backup" \
 --mount "type=volume,source=restoremq,destination=/data" ubuntu:latest \
  bash -c "cd /data && tar xvf /backup/backup.tar.gz --strip 1"
docker create --name mq2 --mount "type=volume,source=restoremq,destination=/var/lib/mysql" -p 5556:3306 -e MYSQL_ROOT_PASSWORD=admin mysql
