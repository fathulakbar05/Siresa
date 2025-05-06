-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 24 Apr 2025 pada 15.28
-- Versi server: 10.4.32-MariaDB
-- Versi PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `siresa`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `jadwal_reservasi`
--

CREATE TABLE `jadwal_reservasi` (
  `id_reservasi` int(11) NOT NULL,
  `id_user` int(11) DEFAULT NULL,
  `id_ruangan` int(11) DEFAULT NULL,
  `tanggal` date DEFAULT NULL,
  `jam_mulai` time DEFAULT NULL,
  `jam_selesai` time DEFAULT NULL,
  `status` enum('menunggu','disetujui','ditolak') DEFAULT 'menunggu',
  `waktu_pengajuan` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `jadwal_reservasi`
--

INSERT INTO `jadwal_reservasi` (`id_reservasi`, `id_user`, `id_ruangan`, `tanggal`, `jam_mulai`, `jam_selesai`, `status`, `waktu_pengajuan`) VALUES
(1, 3, 12, '2025-05-01', '08:00:00', '09:00:00', 'disetujui', '2025-04-24 10:48:48'),
(2, 5, 45, '2025-05-02', '10:00:00', '11:00:00', 'ditolak', '2025-04-24 10:48:48'),
(3, 2, 30, '2025-05-05', '13:00:00', '14:00:00', 'disetujui', '2025-04-24 10:48:48'),
(4, 7, 68, '2025-05-06', '14:00:00', '15:00:00', 'ditolak', '2025-04-24 10:48:48'),
(5, 1, 5, '2025-05-07', '15:00:00', '16:00:00', 'disetujui', '2025-04-24 10:48:48'),
(6, 4, 20, '2025-05-08', '09:00:00', '10:00:00', 'menunggu', '2025-04-24 10:48:48'),
(7, 9, 50, '2025-05-09', '11:00:00', '12:00:00', 'menunggu', '2025-04-24 10:48:48'),
(8, 6, 33, '2025-05-12', '12:00:00', '13:00:00', 'menunggu', '2025-04-24 10:48:48'),
(9, 10, 60, '2025-05-13', '13:00:00', '14:00:00', 'menunggu', '2025-04-24 10:48:48'),
(10, 8, 2, '2025-05-14', '14:00:00', '15:00:00', 'menunggu', '2025-04-24 10:48:48');

-- --------------------------------------------------------

--
-- Struktur dari tabel `notifikasi`
--

CREATE TABLE `notifikasi` (
  `id_notifikasi` int(11) NOT NULL,
  `id_user` int(11) DEFAULT NULL,
  `id_ruangan` int(11) DEFAULT NULL,
  `id_reservasi` int(11) DEFAULT NULL,
  `status_reservasi` enum('disetujui','ditolak') NOT NULL,
  `waktu_dibuat` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `notifikasi`
--

INSERT INTO `notifikasi` (`id_notifikasi`, `id_user`, `id_ruangan`, `id_reservasi`, `status_reservasi`, `waktu_dibuat`) VALUES
(1, 3, 12, 1, 'disetujui', '2025-05-01 02:05:00'),
(2, 5, 45, 2, 'ditolak', '2025-05-02 04:15:00'),
(3, 2, 30, 3, 'disetujui', '2025-05-05 07:10:00'),
(4, 7, 68, 4, 'ditolak', '2025-05-06 08:20:00'),
(5, 1, 5, 5, 'disetujui', '2025-05-07 09:05:00');

-- --------------------------------------------------------

--
-- Struktur dari tabel `ruangan`
--

CREATE TABLE `ruangan` (
  `id_ruangan` int(11) NOT NULL,
  `nama_ruangan` varchar(100) DEFAULT NULL,
  `lokasi_gedung` varchar(100) DEFAULT NULL,
  `lantai` varchar(20) DEFAULT NULL,
  `kapasitas` int(11) DEFAULT NULL,
  `fasilitas` text DEFAULT NULL,
  `ukuran` varchar(50) DEFAULT NULL,
  `tipe` enum('ruang kelas','ruang lab') DEFAULT NULL,
  `layout` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `ruangan`
--

INSERT INTO `ruangan` (`id_ruangan`, `nama_ruangan`, `lokasi_gedung`, `lantai`, `kapasitas`, `fasilitas`, `ukuran`, `tipe`, `layout`) VALUES
(1, 'A1.1', 'Gedung A', '1', 35, 'AC, proyektor, whiteboard', '8x6 meter', 'ruang lab', 'layout lab standard'),
(2, 'A1.2', 'Gedung A', '1', 30, 'AC, tv, whiteboard', '7x6 meter', 'ruang lab', 'layout lab U-shape'),
(3, 'A1.3', 'Gedung A', '1', 40, 'AC, proyektor, whiteboard', '8x7 meter', 'ruang lab', 'layout lab island'),
(4, 'A1.4', 'Gedung A', '1', 25, 'AC, whiteboard', '7x5 meter', 'ruang lab', 'layout lab standard'),
(5, 'A2.1', 'Gedung A', '2', 38, 'AC, tv, whiteboard', '8x6 meter', 'ruang lab', 'layout lab U-shape'),
(6, 'A2.2', 'Gedung A', '2', 22, 'AC, whiteboard', '6x6 meter', 'ruang lab', 'layout lab compact'),
(7, 'A2.3', 'Gedung A', '2', 30, 'AC, proyektor', '7x6 meter', 'ruang lab', 'layout lab standard'),
(8, 'A2.4', 'Gedung A', '2', 32, 'AC, proyektor, tv', '7x6 meter', 'ruang lab', 'layout lab island'),
(9, 'A3.1', 'Gedung A', '3', 40, 'AC, whiteboard, tv', '8x6 meter', 'ruang lab', 'layout lab standard'),
(10, 'A3.2', 'Gedung A', '3', 36, 'AC, proyektor, whiteboard', '8x6 meter', 'ruang lab', 'layout lab compact'),
(11, 'B1.1', 'Gedung B', '1', 28, 'AC, proyektor, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(12, 'B1.2', 'Gedung B', '1', 22, 'AC, whiteboard', '6x5 meter', 'ruang kelas', 'layout kelas U-shape'),
(13, 'B1.3', 'Gedung B', '1', 30, 'AC, tv, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(14, 'B1.4', 'Gedung B', '1', 35, 'AC, proyektor', '8x6 meter', 'ruang kelas', 'layout kelas theater'),
(15, 'B2.1', 'Gedung B', '2', 27, 'AC, whiteboard', '6x6 meter', 'ruang kelas', 'layout kelas compact'),
(16, 'B2.2', 'Gedung B', '2', 34, 'AC, proyektor, tv', '8x6 meter', 'ruang kelas', 'layout kelas U-shape'),
(17, 'B2.3', 'Gedung B', '2', 26, 'AC, tv', '7x5 meter', 'ruang kelas', 'layout kelas baris'),
(18, 'B2.4', 'Gedung B', '2', 40, 'AC, whiteboard, tv', '8x6 meter', 'ruang kelas', 'layout kelas theater'),
(19, 'B3.1', 'Gedung B', '3', 24, 'AC, proyektor, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas compact'),
(20, 'B3.2', 'Gedung B', '3', 36, 'AC, whiteboard', '8x6 meter', 'ruang kelas', 'layout kelas baris'),
(21, 'C1.1', 'Gedung C', '1', 25, 'AC, proyektor, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(22, 'C1.2', 'Gedung C', '1', 30, 'AC, tv, whiteboard', '8x6 meter', 'ruang kelas', 'layout kelas U-shape'),
(23, 'C1.3', 'Gedung C', '1', 22, 'AC, whiteboard', '6x5 meter', 'ruang kelas', 'layout kelas compact'),
(24, 'C1.4', 'Gedung C', '1', 35, 'AC, proyektor, tv', '8x7 meter', 'ruang kelas', 'layout kelas theater'),
(25, 'C2.1', 'Gedung C', '2', 28, 'AC, tv', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(26, 'C2.2', 'Gedung C', '2', 33, 'AC, proyektor, whiteboard', '8x6 meter', 'ruang kelas', 'layout kelas U-shape'),
(27, 'C2.3', 'Gedung C', '2', 26, 'AC, whiteboard', '7x5 meter', 'ruang kelas', 'layout kelas compact'),
(28, 'C2.4', 'Gedung C', '2', 31, 'AC, tv, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(29, 'C3.1', 'Gedung C', '3', 36, 'AC, proyektor, whiteboard', '8x6 meter', 'ruang kelas', 'layout kelas theater'),
(30, 'C3.2', 'Gedung C', '3', 24, 'AC, whiteboard', '6x5 meter', 'ruang kelas', 'layout kelas compact'),
(31, 'C3.3', 'Gedung C', '3', 29, 'AC, proyektor, tv', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(32, 'C3.4', 'Gedung C', '3', 40, 'AC, whiteboard, tv', '8x6 meter', 'ruang kelas', 'layout kelas U-shape'),
(33, 'D1.1', 'Gedung D', '1', 28, 'AC, tv, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(34, 'D1.2', 'Gedung D', '1', 32, 'AC, proyektor', '8x6 meter', 'ruang kelas', 'layout kelas U-shape'),
(35, 'D1.3', 'Gedung D', '1', 23, 'AC, whiteboard', '6x5 meter', 'ruang kelas', 'layout kelas compact'),
(36, 'D1.4', 'Gedung D', '1', 35, 'AC, proyektor, whiteboard', '8x7 meter', 'ruang kelas', 'layout kelas theater'),
(37, 'D2.1', 'Gedung D', '2', 27, 'AC, tv', '7x5 meter', 'ruang kelas', 'layout kelas baris'),
(38, 'D2.2', 'Gedung D', '2', 34, 'AC, proyektor, whiteboard', '8x6 meter', 'ruang kelas', 'layout kelas U-shape'),
(39, 'D2.3', 'Gedung D', '2', 21, 'AC, whiteboard', '6x5 meter', 'ruang kelas', 'layout kelas compact'),
(40, 'D2.4', 'Gedung D', '2', 30, 'AC, tv, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(41, 'D3.1', 'Gedung D', '3', 38, 'AC, proyektor, tv', '8x6 meter', 'ruang kelas', 'layout kelas theater'),
(42, 'D3.2', 'Gedung D', '3', 26, 'AC, whiteboard', '6x6 meter', 'ruang kelas', 'layout kelas compact'),
(43, 'D3.3', 'Gedung D', '3', 29, 'AC, tv, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(44, 'D3.4', 'Gedung D', '3', 33, 'AC, proyektor, whiteboard', '8x7 meter', 'ruang kelas', 'layout kelas U-shape'),
(45, 'E1.1', 'Gedung E', '1', 24, 'AC, whiteboard', '6x5 meter', 'ruang kelas', 'layout kelas compact'),
(46, 'E1.2', 'Gedung E', '1', 29, 'AC, tv, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(47, 'E1.3', 'Gedung E', '1', 34, 'AC, proyektor, tv', '8x6 meter', 'ruang kelas', 'layout kelas theater'),
(48, 'E1.4', 'Gedung E', '1', 26, 'AC, proyektor, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas U-shape'),
(49, 'E2.1', 'Gedung E', '2', 31, 'AC, tv, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(50, 'E2.2', 'Gedung E', '2', 27, 'AC, whiteboard', '6x6 meter', 'ruang kelas', 'layout kelas compact'),
(51, 'E2.3', 'Gedung E', '2', 38, 'AC, proyektor, whiteboard', '8x6 meter', 'ruang kelas', 'layout kelas U-shape'),
(52, 'E2.4', 'Gedung E', '2', 22, 'AC, tv', '6x5 meter', 'ruang kelas', 'layout kelas baris'),
(53, 'E3.1', 'Gedung E', '3', 36, 'AC, whiteboard', '8x6 meter', 'ruang kelas', 'layout kelas compact'),
(54, 'E3.2', 'Gedung E', '3', 25, 'AC, proyektor, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(55, 'E3.3', 'Gedung E', '3', 40, 'AC, tv, whiteboard', '8x7 meter', 'ruang kelas', 'layout kelas theater'),
(56, 'E3.4', 'Gedung E', '3', 30, 'AC, proyektor', '7x6 meter', 'ruang kelas', 'layout kelas U-shape'),
(57, 'F1.1', 'Gedung F', '1', 33, 'AC, whiteboard', '7x5 meter', 'ruang kelas', 'layout kelas compact'),
(58, 'F1.2', 'Gedung F', '1', 28, 'AC, tv, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(59, 'F1.3', 'Gedung F', '1', 35, 'AC, proyektor, tv', '8x7 meter', 'ruang kelas', 'layout kelas theater'),
(60, 'F1.4', 'Gedung F', '1', 26, 'AC, whiteboard', '7x5 meter', 'ruang kelas', 'layout kelas U-shape'),
(61, 'F2.1', 'Gedung F', '2', 30, 'AC, tv, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas baris'),
(62, 'F2.2', 'Gedung F', '2', 24, 'AC, proyektor', '6x5 meter', 'ruang kelas', 'layout kelas compact'),
(63, 'F2.3', 'Gedung F', '2', 38, 'AC, whiteboard, tv', '8x6 meter', 'ruang kelas', 'layout kelas U-shape'),
(64, 'F2.4', 'Gedung F', '2', 22, 'AC, tv', '6x6 meter', 'ruang kelas', 'layout kelas baris'),
(65, 'F3.1', 'Gedung F', '3', 29, 'AC, whiteboard', '7x6 meter', 'ruang kelas', 'layout kelas compact'),
(66, 'F3.2', 'Gedung F', '3', 36, 'AC, proyektor, whiteboard', '8x6 meter', 'ruang kelas', 'layout kelas baris'),
(67, 'F3.3', 'Gedung F', '3', 40, 'AC, tv, whiteboard', '8x7 meter', 'ruang kelas', 'layout kelas theater'),
(68, 'F3.4', 'Gedung F', '3', 31, 'AC, proyektor', '7x6 meter', 'ruang kelas', 'layout kelas U-shape');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id_user` int(11) NOT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `nim_nip` varchar(30) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role` enum('mahasiswa','dosen','admin') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id_user`, `nama`, `nim_nip`, `password`, `role`) VALUES
(1, 'Ahmad Fauzan', '22102001', 'ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f', 'mahasiswa'),
(2, 'Dinda Putri', '22102002', 'e98ad157372606220a0aab2d0a4af25dae5897e7d97fa097bb589b39c197cf65', 'mahasiswa'),
(3, 'Reza Akbar', '22102003', '3fc7b23ea7b8ca6b9632eee9db2391b6b270f840481bc7f47128d2bdab3f50ed', 'mahasiswa'),
(4, 'Siti Rahma', '22102004', '38e7ddb50f4546d6f07936c4ed7a67cd38745db1e5630e5eb363bf0752cbb4a6', 'mahasiswa'),
(5, 'Yusuf Hadi', '22102005', 'a0e987329b857365cf9affc617744815ad5813f9126acd695dcc0964f38b64fe', 'mahasiswa'),
(6, 'Dr. Budi Santoso', '19800123', 'dc841da2fe4f20d5adefd1edb32ba8b8de4714f5b556a2e5863b5e4b65ee46f3', 'dosen'),
(7, 'Dr. Lina Marlina', '19791230', 'c276ebd4bdba34e44e8c14ede35d2d9d5103c05f3dcd0976274f362b916cf764', 'dosen'),
(8, 'Prof. Rudi Hartono', '19670615', '42c05d31c68f603acbd364dca42662fa9fa8ea1e80b9c326cd67f5613e94e917', 'dosen'),
(9, 'Admin SIAKAD', 'admin001', '713bfda78870bf9d1b261f565286f85e97ee614efe5f0faf7c34e7ca4f65baca', 'admin'),
(10, 'Operator Akademik', 'admin002', 'ec6e1c25258002eb1c67d15c7f45da7945fa4c58778fd7d88faa5e53e3b4698d', 'admin');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `jadwal_reservasi`
--
ALTER TABLE `jadwal_reservasi`
  ADD PRIMARY KEY (`id_reservasi`),
  ADD KEY `id_user` (`id_user`),
  ADD KEY `id_ruangan` (`id_ruangan`);

--
-- Indeks untuk tabel `notifikasi`
--
ALTER TABLE `notifikasi`
  ADD PRIMARY KEY (`id_notifikasi`),
  ADD KEY `id_user` (`id_user`),
  ADD KEY `id_ruangan` (`id_ruangan`),
  ADD KEY `id_reservasi` (`id_reservasi`);

--
-- Indeks untuk tabel `ruangan`
--
ALTER TABLE `ruangan`
  ADD PRIMARY KEY (`id_ruangan`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_user`),
  ADD UNIQUE KEY `nim_nip` (`nim_nip`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `jadwal_reservasi`
--
ALTER TABLE `jadwal_reservasi`
  MODIFY `id_reservasi` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT untuk tabel `notifikasi`
--
ALTER TABLE `notifikasi`
  MODIFY `id_notifikasi` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `ruangan`
--
ALTER TABLE `ruangan`
  MODIFY `id_ruangan` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=69;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id_user` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `jadwal_reservasi`
--
ALTER TABLE `jadwal_reservasi`
  ADD CONSTRAINT `jadwal_reservasi_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id_user`),
  ADD CONSTRAINT `jadwal_reservasi_ibfk_2` FOREIGN KEY (`id_ruangan`) REFERENCES `ruangan` (`id_ruangan`);

--
-- Ketidakleluasaan untuk tabel `notifikasi`
--
ALTER TABLE `notifikasi`
  ADD CONSTRAINT `notifikasi_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id_user`),
  ADD CONSTRAINT `notifikasi_ibfk_2` FOREIGN KEY (`id_ruangan`) REFERENCES `ruangan` (`id_ruangan`),
  ADD CONSTRAINT `notifikasi_ibfk_3` FOREIGN KEY (`id_reservasi`) REFERENCES `jadwal_reservasi` (`id_reservasi`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
