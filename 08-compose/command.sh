### Postgres SQL

# 1. Buat volume untuk data dan backup
docker volume create postgres_volume
docker volume create postgres_backup_volume

# 2. Jalankan container postgres:alpine
docker run -d \
  --name default_postgres_db \
  -e POSTGRES_PASSWORD=admin \
  -p 5432:5432 \
  --cpus="0.5" \
  --memory="100m" \
  -v postgres_volume:/var/lib/postgresql/data \
  -v postgres_backup_volume:/var/lib/postgresql/backup \
  postgres:alpine