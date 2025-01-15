# Container Resources

**Container Resources** merujuk pada alokasi sumber daya sistem (seperti CPU, memori, dan disk) untuk **Docker Container**. Docker memungkinkan pengelolaan sumber daya ini untuk memastikan container berjalan secara efisien dan tidak saling mengganggu, terutama di lingkungan multi-container.

---

## Kenapa Mengelola Container Resources?

1. **Efisiensi Sistem**:

   - Memastikan bahwa container tidak menggunakan lebih banyak sumber daya dari yang diperlukan.

2. **Stabilitas**:

   - Membatasi konsumsi sumber daya agar satu container tidak memengaruhi kinerja container lain atau sistem host.

3. **Isolasi**:

   - Memberikan kontrol granular terhadap bagaimana sumber daya sistem digunakan oleh container.

4. **Skalabilitas**:
   - Mengoptimalkan penggunaan sumber daya untuk mendukung lebih banyak container di satu mesin.

---

## Jenis Container Resources

1. **CPU**:

   - Docker memungkinkan Anda membatasi jumlah CPU yang dapat digunakan container.
   - Parameter:
     - `--cpus=<value>`: Membatasi jumlah CPU.
     - `--cpu-shares=<value>`: Memberikan prioritas CPU relatif terhadap container lain.

2. **Memori (RAM)**:

   - Anda dapat menentukan batas memori yang dapat digunakan oleh container.
   - Parameter:
     - `--memory=<value>`: Membatasi memori maksimum.
     - `--memory-swap=<value>`: Membatasi swap (kombinasi memori fisik dan swap).

3. **Disk I/O**:

   - Mengontrol kecepatan baca/tulis ke disk oleh container.
   - Parameter:
     - `--device-read-bps` dan `--device-write-bps`: Membatasi bandwidth baca/tulis disk.

4. **Jaringan**:
   - Mengelola bandwidth jaringan untuk container.
   - Parameter:
     - `--network`: Menentukan jenis jaringan yang digunakan oleh container.

## Contoh Penggunaan

1. **Membatasi CPU dan Memori**:

```bash
# Membatasi container hanya menggunakan 1.5 CPU dan 512 MB memori.
docker run -d --cpus="1.5" --memory="512m" nginx
```

2. **Mengatur Disk I/O:**

```bash
# Membatasi kecepatan baca 10 MB/s dan tulis 5 MB/s ke disk /dev/sda.
docker run -d --device-read-bps=/dev/sda:10mb --device-write-bps=/dev/sda:5mb nginx
```

3. **Prioritas CPU:**

```bash
# Memberikan prioritas CPU lebih tinggi untuk container ini dibandingkan container dengan --cpu-shares yang lebih rendah.
docker run -d --cpu-shares=1024 nginx
```

## Kesimpulan

Pengelolaan Container Resources adalah aspek penting dalam menggunakan Docker, terutama di lingkungan dengan banyak container. Dengan membatasi dan mengalokasikan sumber daya secara bijaksana, Anda dapat meningkatkan efisiensi, stabilitas, dan skalabilitas sistem secara keseluruhan.
