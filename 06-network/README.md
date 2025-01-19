# Docker Network

**Docker Network** adalah mekanisme yang disediakan oleh Docker untuk memungkinkan komunikasi antar container, antara container dan host, serta antara container dan jaringan eksternal. Dengan Docker Network, Anda dapat mengelola bagaimana container saling terhubung dan berinteraksi.

---

## Kenapa Menggunakan Docker Network?

1. **Isolasi Jaringan**:

   - Setiap container dapat diisolasi dari container lain sesuai kebutuhan.

2. **Komunikasi Antar Container**:

   - Mendukung aplikasi multi-container untuk berkomunikasi dengan lancar.

3. **Custom Topology**:

   - Mendukung pembuatan jaringan dengan konfigurasi khusus sesuai kebutuhan aplikasi.

4. **Keamanan**:
   - Jaringan yang terisolasi meningkatkan keamanan aplikasi dengan membatasi akses.

---

## Jenis Docker Network

1. **Bridge Network** (default):

   - Jaringan lokal yang digunakan oleh container di host yang sama.
   - Contoh:
     ```bash
     docker network create my-bridge-network
     docker run --network my-bridge-network nginx
     ```

2. **Host Network**:

   - Container berbagi jaringan dengan host.
   - Contoh:
     ```bash
     docker run --network host nginx
     ```

3. **None Network**:

   - Container berjalan tanpa akses jaringan.
   - Contoh:
     ```bash
     docker run --network none nginx
     ```

4. **Overlay Network**:
   - Digunakan untuk container yang berjalan di beberapa host Docker dalam mode swarm.
   - Contoh:
     ```bash
     docker network create -d overlay my-overlay-network
     ```
