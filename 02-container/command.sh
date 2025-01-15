# Melihat Contaier berjalan
docker ps -a
docker container ls -a

# Membuat Container
docker container create --name webService nginx
docker create --name webService nginx

# Menjalanka Container
docker container start webService
docker start id/name

# Menghentika Container
docker container stop webService
docker stop id/name

# Menghapus container
docker container rm webService
docker rm id/name

# Melihat logs
docker container logs webService
docker logs -f webService # listning

# Masuk ke Container
docker container exec -i -t id/name /bin/bash
docker exec -i -t webService /bin/bash
