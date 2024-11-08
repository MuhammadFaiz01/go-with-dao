CREATE TABLE IF NOT EXISTS kelas (
    id SERIAL PRIMARY KEY,
    nama_kelas VARCHAR(50) NOT NULL,
    tingkat INT NOT NULL,
    id_person INT,
    FOREIGN KEY (id_person) REFERENCES person(id) ON DELETE SET NULL
);