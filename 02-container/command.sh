# Melihat Contaier berjalan
docker ps -a
docker container ls -a

# Membuat Container
docker container create --name webService nginx
docker create --name webService nginx

# Port Forwarding
docker create --name webService --publish 8080:80 nginx
docker create --name webService -p 8080:80 nginx:latest

# Add Environment variabel
docker create --name aa -p 27017:27017 --env MONGO_INITDB_ROOT_USERNAME=admin --env MONGO_INITDB_ROOT_PASSWORD=admin mongo
docker create --name bb -p 5432:5432 --env POSTGRES_PASSWORD=admin postgres


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

