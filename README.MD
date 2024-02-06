# Final Course Project - Task Tracker Plus
### Test Case Examples & ENDPOINT

#### Test Case 1

**Input**:

```http
POST /api/v1/user/register HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "fullname": "John Doe",
    "email": "johndoe@example.com",
    "password": "password123"
}
```

**Expected Output / Behavior**:

- Jika registrasi berhasil, server harus mengembalikan kode status HTTP `201` Created dan respons JSON dengan pesan sukses:

  ```json
  {
    "message": "register success"
  }
  ```

- Jika terjadi kesalahan validasi data registrasi, server harus mengembalikan kode status HTTP `400` Bad Request dan respons JSON dengan pesan kesalahan:

  ```json
  {
    "error": "invalid data"
  }
  ```

- Jika terjadi kesalahan server saat menyimpan data pengguna, server harus mengembalikan kode status HTTP `500` Internal Server Error dan respons JSON dengan pesan kesalahan:

  ```json
  {
    "error": "internal server error"
  }
  ```

#### Test Case 2

**Input**:

```http
POST /api/v1/user/login HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "email": "johndoe@example.com",
    "password": "password123"
}
```

**Expected Output / Behavior**:

- Jika login berhasil, server harus mengembalikan kode status HTTP `200` OK dan respons JSON dengan pesan sukses:

  ```json
  {
    "message": "login success"
  }
  ```

- Jika terjadi kesalahan validasi data login, server harus mengembalikan kode status HTTP `400` Bad Request dan respons JSON dengan pesan kesalahan:

  ```json
  {
    "error": "email or password is empty"
  }
  ```

- Jika terjadi kesalahan server saat melakukan login, server harus mengembalikan kode status HTTP `500` Internal Server Error dan respons JSON dengan pesan kesalahan:

  ```json
  {
      "error": "internal server error"
  }
  ```

#### Test Case 3

**Input**:

```http
POST /api/v1/task/add HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "title": "Task 1",
    "deadline": "2023-06-30",
    "priority": 2,
    "category_id": "1",
    "status": "In Progress"
}
```

**Expected Output / Behavior**:

- Jika permintaan berhasil, server harus mengembalikan kode status HTTP `200` OK dan respons JSON dengan pesan sukses:

  ```json
  {
    "message": "add task success"
  }
  ```

- Jika permintaan gagal karena kesalahan validasi, server harus mengembalikan kode status HTTP `400` Bad Request dan respons JSON dengan pesan kesalahan:

  ```json
  {
    "error": "invalid data"
  }
  ```

- Jika terjadi kesalahan server saat menyimpan tugas, server harus mengembalikan kode status HTTP `500` Internal Server Error dan respons JSON dengan pesan kesalahan:

  ```json
  {
    "error": "internal server error"
  }
  ```

#### Test Case 4

**Input**:

```http
GET /api/v1/task/list HTTP/1.1
Host: localhost:8080
```

**Expected Output / Behavior**:

- Jika permintaan berhasil, server harus mengembalikan kode status HTTP `200` OK dan respons JSON dengan daftar tugas.

  ```json
  [
      {
          "id": 1,
          "title": "Task 1",
          "deadline": "2023-05-30",
          "priority": 2,
          "category_id": "1",
          "status": "In Progress"
      },
      {
          "id": 2,
          "title": "Task 2",
          "deadline": "2023-06-10",
          "priority": 1,
          "category_id": "2",
          "status": "Completed"
      },
      ...
  ]
  ```

- Jika terjadi kesalahan saat mendapatkan daftar tugas, server harus mengembalikan kode status HTTP `500` Internal Server Error dan respons JSON dengan pesan kesalahan.

  ```json
  {
      "error": "[error messages]"
  }
  ```

#### Test Case 5

**Input**:

```http
PUT /api/v1/task/update/{id} HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "title": "Task 1",
    "deadline": "2023-06-30",
    "priority": 2,
    "category_id": "1",
    "status": "In Progress"
}
```

**Expected Output / Behavior**:

- Jika permintaan berhasil dan ID tugas valid, server harus mengembalikan kode status HTTP `200` OK dan respons JSON dengan pesan sukses.

  ```json
  {
      "message": "update Task success"
  }
  ```

- Jika permintaan gagal karena ID tugas tidak valid, server harus mengembalikan kode status HTTP `400` Bad Request dan respons JSON dengan pesan kesalahan.

  ```json
  {
      "error": "Invalid task ID"
  }
  ```

- Jika permintaan gagal karena format data tugas tidak sesuai yang diharapkan, server harus mengembalikan kode status HTTP `400` Bad Request dan respons JSON dengan pesan kesalahan.

  ```json
  {
      "error": "[error messages]"
  }
  ```

- Jika terjadi kesalahan saat memperbarui data tugas, server harus mengembalikan kode status HTTP `500` Internal Server Error dan respons JSON dengan pesan kesalahan.

  ```json
  {
      "error": "[error messages]"
  }
  ```

#### Test Case 6

**Input**:

```http
DELETE /api/v1/task/delete/{id} HTTP/1.1
Host: localhost:8080
```

**Expected Output / Behavior**:

- Jika permintaan berhasil dan ID tugas valid, server harus mengembalikan kode status HTTP `200` OK dan respons JSON dengan pesan sukses.

   ```json
   {
       "message": "delete Task success"
   }
   ```

- Jika permintaan gagal karena ID tugas tidak valid, server harus mengembalikan kode status HTTP `400` Bad Request dan respons JSON dengan pesan kesalahan.

   ```json
   {
       "error": "Invalid task ID"
   }
   ```

- Jika terjadi kesalahan saat menghapus tugas, server harus mengembalikan kode status HTTP `500` Internal Server Error dan respons JSON dengan pesan kesalahan.

   ```json
   {
       "error": "[error messages]"
   }
   ```
