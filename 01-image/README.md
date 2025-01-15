# Docker Image

**Docker Image** adalah blueprint atau template yang bersifat read-only untuk membuat dan menjalankan **Docker Container**. Image berisi semua yang dibutuhkan untuk menjalankan aplikasi, termasuk kode, dependensi, library, konfigurasi, dan sistem operasi dasar.

---

## Karakteristik Docker Image

1. **Layered Structure**:

   - Docker Image terdiri dari beberapa layer. Setiap perubahan atau instruksi dalam Dockerfile akan menambah layer baru.
   - Layer yang tidak berubah dapat digunakan kembali, sehingga proses build lebih efisien.

2. **Immutable**:

   - Docker Image bersifat read-only, sehingga tidak dapat dimodifikasi secara langsung setelah dibuat.

3. **Portable**:

   - Image dapat digunakan di berbagai sistem operasi yang mendukung Docker tanpa perubahan.

4. **Versioned**:
   - Docker Image menggunakan tag (seperti `nginx:1.21`) untuk mengelola versi, memungkinkan pengembang menentukan image yang sesuai dengan kebutuhan.

---

## Contoh Umum Docker Image

1. **Base Image**:

   - Image dasar yang digunakan untuk membangun aplikasi.
   - Contoh: `ubuntu`, `alpine`, `python`.

2. **Custom Image**:
   - Dibuat menggunakan **Dockerfile** dengan konfigurasi dan aplikasi yang spesifik.
   - Contoh: `my-app:latest` adalah custom image dari aplikasi Anda.

---

## Cara Kerja Docker Image

1. **Membuat Image**:

   - Menggunakan Dockerfile dengan instruksi tertentu.
   - Contoh:
     ```bash
     docker build -t my-app .
     ```

2. **Menyimpan Image**:

   - Image yang sudah dibuat dapat di-push ke Docker Hub atau registry lain untuk distribusi.
   - Contoh:
     ```bash
     docker push my-app:latest
     ```

3. **Menjalankan Image**:
   - Image dijalankan sebagai container.
   - Contoh:
     ```bash
     docker run -d my-app
     ```

---

## Kesimpulan

Docker Image adalah elemen inti dalam ekosistem Docker yang memungkinkan aplikasi dikemas dengan semua dependensinya. Dengan struktur berlapis, portabilitas tinggi, dan kemampuan versi, Docker Image mendukung proses pengembangan dan deployment yang lebih cepat dan efisien.
