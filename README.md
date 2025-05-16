Anggota Kelompok 6:
1. Muhammad Fathul Akbar (09021282328111)
2. Kenz Raki Abdurrazak (09021282328117)
3. Wahyu Trisno Aji (09021282328122)
4. Githa Evelyn (09021282328124)
5. Kian Fikram Ayyubi (09021282328098)

Databasenya bisa diimport dari file /static/siresa.sql

Berikut adalah password dari akun-akun yang sudah terdaftar di database:
-- Mahasiswa
('Ahmad Fauzan', '22102001', SHA2('password123', 256), 'mahasiswa'),
('Dinda Putri', '22102002', SHA2('putri123', 256), 'mahasiswa'),
('Reza Akbar', '22102003', SHA2('reza456', 256), 'mahasiswa'),
('Siti Rahma', '22102004', SHA2('rahma789', 256), 'mahasiswa'),
('Yusuf Hadi', '22102005', SHA2('yusuf000', 256), 'mahasiswa'),

-- Dosen
('Dr. Budi Santoso', '19800123', SHA2('dosenbudi', 256), 'dosen'),
('Dr. Lina Marlina', '19791230', SHA2('linadosen', 256), 'dosen'),
('Prof. Rudi Hartono', '19670615', SHA2('rudi1234', 256), 'dosen'),

-- Admin
('Admin SIAKAD', 'admin001', SHA2('adminpass', 256), 'admin'),
('Operator Akademik', 'admin002', SHA2('operator123', 256), 'admin');
