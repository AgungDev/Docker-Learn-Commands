# Docker Container

**Docker Container** adalah instansi yang berjalan dari **Docker Image**. Container adalah unit perangkat lunak yang ringan dan portabel yang mengemas aplikasi beserta semua dependensinya dalam lingkungan yang terisolasi. Container memungkinkan aplikasi berjalan konsisten di berbagai sistem tanpa memerlukan konfigurasi tambahan.

---

## Karakteristik Docker Container

1. **Isolasi**:

   - Setiap container berjalan secara terpisah dari container lain dan sistem host.
   - Container hanya memiliki akses ke sumber daya yang diizinkan (CPU, memori, jaringan).

2. **Ringan**:

   - Container berbagi kernel sistem operasi host, sehingga lebih efisien dibandingkan mesin virtual.

3. **Portabilitas**:

   - Container dapat dijalankan di berbagai platform yang mendukung Docker tanpa modifikasi.

4. **Ephemeral**:
   - Secara default, container bersifat sementara (ephemeral), tetapi data dapat dipertahankan menggunakan **volume** atau penyimpanan eksternal.

---

## Fungsi Docker Container

- Menjalankan aplikasi dengan konfigurasi yang konsisten.
- Mendukung pengembangan, pengujian, dan deployment secara cepat.
- Menyediakan lingkungan pengembangan terisolasi.

---

## Perintah Dasar Docker Container

**Menjalankan Container**:

```bash
docker run -d -p 8080:80 nginx
```
