# Docker Bind Mounts

**Bind Mounts** adalah salah satu cara Docker untuk berbagi direktori di sistem host dengan container. Bind Mounts memungkinkan container mengakses data secara langsung dari direktori host, sehingga setiap perubahan pada host langsung terlihat di container, dan sebaliknya.

---

## Karakteristik Bind Mounts

1. **Direct Mapping**:

   - Menghubungkan direktori atau file di host ke direktori dalam container.

2. **High Performance**:

   - Karena data diakses langsung dari host, bind mounts cenderung memiliki performa yang tinggi.

3. **Flexibility**:

   - Cocok untuk pengembangan aplikasi di mana Anda perlu memperbarui file secara langsung tanpa membangun ulang container.

4. **No Isolation**:
   - Tidak ada isolasi antara host dan container, sehingga perubahan data langsung terlihat di kedua sisi.

---

## Kapan Menggunakan Bind Mounts?

- Untuk mengembangkan aplikasi dengan perubahan kode secara real-time.
- Ketika Anda ingin container memiliki akses langsung ke file atau direktori di host, seperti file log atau konfigurasi.

---

## Cara Menggunakan Bind Mounts

Bind Mounts didefinisikan menggunakan flag `-v` atau `--mount` saat menjalankan container. lalu tambahkan paramete `type`, `source`, `destination`, dan `readonly` (kalau mengunakan ini, kita hanya bisa membaca tidak bisa menulis)

### Sintaks Dasar:

```bash
docker run -v /path/on/host:/path/in/container [image_name]
```
