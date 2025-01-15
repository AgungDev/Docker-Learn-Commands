# Melihat Resources Container
docker container stats

# Mengatur Limit Resource Container
docker create --name smallenginx --memory 100m --cpus 0.5 -p 8081:80 nginx