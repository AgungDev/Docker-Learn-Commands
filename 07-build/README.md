# Docker Build

**Docker Build** adalah proses untuk membuat image Docker berdasarkan file konfigurasi bernama **Dockerfile**. Dockerfile berisi instruksi berurutan yang mendefinisikan bagaimana image dibuat, termasuk sistem operasi dasar, dependensi, konfigurasi, dan aplikasi yang akan dijalankan di dalam container.

---

## Kenapa Menggunakan Docker Build?

1. **Reproducibility**:

   - Memastikan lingkungan aplikasi dapat dibuat ulang dengan konsistensi.

2. **Portabilitas**:

   - Image Docker yang dihasilkan dapat dijalankan di mana saja selama Docker tersedia.

3. **Automasi**:
   - Proses build otomatis berdasarkan Dockerfile mempermudah pengelolaan aplikasi.

---

## Instruksi Dockerfile

Dockerfile adalah file teks yang berisi sekumpulan instruksi untuk membangun image Docker. Berikut adalah deskripsi dari setiap instruksi yang dapat digunakan dalam Dockerfile:

1. **FROM**
   Menentukan image dasar yang akan digunakan untuk membangun image baru. Setiap Dockerfile harus dimulai dengan instruksi ini.

2. **RUN**
   Menjalankan perintah selama proses build untuk menginstal perangkat lunak, memperbarui sistem, atau melakukan operasi lainnya dalam image.

3. **CMD**
   Menentukan perintah default yang akan dijalankan saat container dimulai. Instruksi ini dapat di-override saat container dijalankan.

4. **ENTRYPOINT**
   Menentukan perintah utama untuk container, sering digunakan untuk membuat container lebih fleksibel dengan argumen tambahan.

5. **COPY**
   Menyalin file atau direktori dari host ke image pada lokasi yang ditentukan.

6. **ADD**
   Seperti `COPY`, tetapi mendukung URL dan file arsip yang akan diekstrak secara otomatis ke dalam image.

7. **WORKDIR**
   Menetapkan direktori kerja untuk instruksi selanjutnya dalam Dockerfile dan untuk perintah yang dijalankan di dalam container.

8. **ENV**
   Menetapkan variabel lingkungan yang akan tersedia di dalam container.

9. **ARG**
   Mendefinisikan variabel build-time yang hanya tersedia selama proses build image.

10. **EXPOSE**
    Menentukan port yang akan dibuka oleh container untuk komunikasi.

11. **VOLUME**
    Mendefinisikan titik mount untuk berbagi data antara container dan host.

12. **USER**
    Menentukan pengguna yang akan digunakan untuk menjalankan perintah di dalam container.

13. **LABEL**
    Menambahkan metadata ke image, seperti informasi maintainer atau versi aplikasi.

14. **STOPSIGNAL**
    Menentukan sinyal sistem yang akan dikirimkan untuk menghentikan container.

15. **HEALTHCHECK**
    Menginstruksikan Docker untuk memeriksa kesehatan container secara berkala.

16. **ONBUILD**
    Menentukan instruksi yang akan dijalankan saat image ini digunakan sebagai base image untuk Dockerfile lain.

17. **SHELL**
    Mengubah shell default yang digunakan oleh perintah seperti `RUN`, `CMD`, atau `ENTRYPOINT`.

---

## Kesimpulan

Dengan memahami dan menggunakan instruksi Dockerfile ini, Anda dapat membangun image yang efisien, terstruktur, dan sesuai dengan kebutuhan aplikasi Anda.
