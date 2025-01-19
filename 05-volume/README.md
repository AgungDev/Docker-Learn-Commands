# Docker Volume

**Docker Volume** adalah mekanisme penyimpanan data yang dikelola oleh Docker untuk menyimpan dan berbagi data antar container. Volume memungkinkan data bertahan meskipun container dihentikan atau dihapus, membuatnya ideal untuk kebutuhan penyimpanan yang persisten.

---

## Kenapa Menggunakan Docker Volume?

1. **Persistensi Data**:

   - Data tetap tersedia meskipun container dihapus atau dibuat ulang.

2. **Isolasi**:

   - Volume terisolasi dari sistem file host, memberikan keamanan lebih baik dibanding bind mounts.

3. **Portabilitas**:

   - Volume dapat dengan mudah dipindahkan atau digunakan kembali di berbagai container dan host.

4. **Pengelolaan oleh Docker**:
   - Docker mengelola lokasi dan pengaturan volume, mempermudah pengguna.

---

## Fitur Utama Docker Volume

- **Berbagi Data**: Volume dapat digunakan bersama oleh beberapa container.
- **Data Konsisten**: Volume menjaga data tetap konsisten meskipun container dihapus.
- **Mendukung Backup**: Volume dapat dengan mudah di-backup dan di-restore.

---

## Cara Menggunakan Docker Volume

```bash
# Membuat volume
docker volume create my-volume
# Menjalankan Volume
docker run -d -v my-volume:/data nginx
# Melihat volume
docker volume ls
# Menghapus volume
docker volume rm my-volume
```

## Perbedaan Docker Volume dan Bind Mounts

Docker Volume | Bind Mounts
Dikelola sepenuhnya oleh Docker | Menggunakan direktori host
Isolasi lebih baik | Tidak ada isolasi
Portabel antar host | Bergantung pada struktur direktori host
Cocok untuk produksi | Cocok untuk pengembangan
