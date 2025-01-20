# Introduction to Docker

Docker adalah platform open-source yang dirancang untuk membantu pengembang membangun, mengirimkan, dan menjalankan aplikasi dalam **container**. Container adalah unit perangkat lunak yang ringan dan portabel yang mengemas aplikasi beserta semua dependensi dan lingkungan yang diperlukan untuk menjalankannya, memastikan bahwa aplikasi dapat berjalan dengan konsisten di berbagai lingkungan.

---

## Kenapa Menggunakan Docker?

1. **Portabilitas**: Docker container dapat dijalankan di berbagai platform, termasuk Windows, macOS, dan Linux, tanpa perubahan kode.
2. **Isolasi Lingkungan**: Setiap container berjalan secara independen, sehingga mencegah konflik antar aplikasi yang menggunakan dependensi berbeda.
3. **Efisiensi**: Docker menggunakan sumber daya sistem lebih ringan dibandingkan mesin virtual (VM).
4. **Pengembangan yang Konsisten**: Dengan Docker, aplikasi akan berjalan sama di lingkungan pengembang, pengujian, dan produksi.
5. **Integrasi dengan CI/CD**: Docker mempermudah proses otomatisasi deployment melalui pipeline CI/CD.

---

## Komponen Utama Docker

1. **Docker Engine**:

   - Mesin inti yang bertugas membangun dan menjalankan container.
   - Terdiri dari **Docker Daemon**, **REST API**, dan **CLI**.

2. **Image**:

   - Blueprint atau template dari container. Docker image dibuat dari Dockerfile dan bersifat read-only.
   - Contoh: `nginx`, `mysql`, `python`.

3. **Container**:

   - Instance yang berjalan dari sebuah Docker image. Container bersifat portabel dan dapat dijalankan di berbagai platform.

4. **Dockerfile**:

   - File konfigurasi yang berisi instruksi untuk membuat Docker image.
   - Contoh: Menentukan sistem operasi dasar, dependensi, dan perintah yang akan dijalankan.

5. **Docker Hub**:
   - Repositori publik untuk menyimpan dan berbagi Docker image. Pengembang dapat menarik (pull) atau mendorong (push) image ke/dari Docker Hub.

---

## Cara Kerja Docker

1. **Build**:

   - Dockerfile digunakan untuk membangun image.
   - Contoh: `docker build -t my-app .`

2. **Run**:

   - Image yang sudah dibuat dijalankan sebagai container.
   - Contoh: `docker run -d -p 8080:80 my-app`

3. **Ship**:
   - Container dapat dipaketkan dan didistribusikan ke server atau cloud.

---

## Perbandingan Docker dengan Virtual Machine (VM)

| **Fitur**                  | **Docker**                 | **Virtual Machine**                 |
| -------------------------- | -------------------------- | ----------------------------------- |
| **Ukuran**                 | Ringan (MB)                | Berat (GB)                          |
| **Kecepatan Booting**      | Cepat (detik)              | Lambat (menit)                      |
| **Isolasi**                | Proses                     | Sistem Operasi                      |
| **Penggunaan Sumber Daya** | Efisien, berbagi kernel OS | Membutuhkan sumber daya lebih besar |
| **Portabilitas**           | Sangat mudah               | Tidak semudah Docker                |

---
