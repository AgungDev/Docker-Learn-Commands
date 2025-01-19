# Buat Image 
docker build -t laundry contoh

# Sample run my app
docker run --name tes -it laundry
docker start -i tes # -i interaktif
