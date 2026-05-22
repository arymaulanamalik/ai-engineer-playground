# Belajar AI Engineer — Context Window (Memori Jangka Pendek LLM)

## 1. Apa itu Context Window?

**Context Window** adalah **jumlah maksimum token yang bisa "diingat" model dalam satu waktu**.

Cara paling gampang memahami:

> **Context Window = RAM / memori kerja sementara LLM**

LLM hanya bisa "melihat" sejumlah token tertentu saat menjawab.

Contoh:

```text
Model A → 8k tokens
Model B → 32k tokens
Model C → 128k tokens
Model D → 1M tokens
```

Kalau token melebihi batas itu:

- sebagian teks lama bisa dibuang
- prompt bisa terpotong
- model bisa lupa konteks

---

# 2. Analogi paling gampang

Bayangkan kamu sedang baca buku, tapi hanya boleh membuka **1 halaman**.

Kalau pertanyaan jawabannya ada di halaman sebelumnya yang sudah ditutup:

> kamu bisa lupa.

LLM juga begitu.

Dia **tidak benar-benar punya ingatan permanen**, hanya melihat isi **context window saat ini**.

---

# 3. Apa saja yang masuk ke Context Window?

Banyak orang pikir hanya prompt user.  
Padahal **semuanya ikut dihitung**.

Contoh:

```text
[System Prompt]
Kamu adalah AI assistant.

[Chat History]
User: Halo
Assistant: Halo

User: Jelaskan Docker

[Retrieved Docs dari RAG]
Dokumen 1...
Dokumen 2...

[Current User Prompt]
Buat ringkasannya.
```

Semua itu masuk ke context window.

---

# 4. Formula sederhananya

```text
Total Context =
System Prompt
+ Chat History
+ Retrieved Documents
+ User Prompt
+ Expected Output
```

Semua dalam **token**.

---

# 5. Kenapa penting untuk AI Engineer?

Karena kalau context penuh, masalah bisa muncul:

### Model lupa pesan lama

Contoh:

```text
User awal: jawab pakai Bahasa Indonesia
...
50 pesan kemudian
Model jawab Bahasa Inggris
```

Kenapa?

Karena instruksi awal bisa terdorong keluar.

---

### RAG gagal

Kalau terlalu banyak dokumen dimasukkan:

- context overflow
- dokumen penting terpotong
- jawaban jadi jelek

---

### Biaya naik

Semakin besar context:

- token input naik
- biaya API naik

---

### Latensi naik

Semakin banyak token:

- respon lebih lambat

---

# 6. Contoh nyata

Misalnya model punya:

```text
128k context window
```

Lalu kamu kirim:

```text
System prompt       = 2k
Chat history        = 30k
RAG docs            = 80k
User prompt         = 2k
Expected output     = 5k
```

Total:

```text
119k
```

Masih aman.

---

Kalau jadi:

```text
140k
```

Bisa:

- request ditolak
- sebagian teks dipotong

---

# 7. Context Window ≠ Memory Permanen

Ini penting.

Banyak orang salah paham:

> "Model ingat percakapan saya."

Sebenarnya:

> model hanya membaca ulang semua pesan yang masih ada di context.

Kalau pesan lama keluar dari context:

> model "lupa"

---

# 8. Teknik mengelola Context

## A. Summarize history

Jangan simpan semua chat.

Ubah jadi ringkasan:

```text
User sedang membangun aplikasi RAG dengan Golang.
Sudah memakai pgvector.
Sedang debugging embedding mismatch.
```

Lebih hemat token.

---

## B. Chunking dokumen

Untuk RAG:

jangan kirim PDF penuh.

Potong kecil:

```text
500–1000 tokens per chunk
```

Lebih efisien.

---

## C. Retrieval yang relevan

Jangan masukkan semua dokumen.

Ambil top-k yang paling relevan.

Contoh:

```text
Top 3 chunks
```

bukan:

```text
Top 50 chunks
```

---

## D. Prompt yang ringkas

Jangan verbose.

Contoh buruk:

```text
Tolong dengan hormat bantu saya...
```

Lebih baik:

```text
Jelaskan...
```

---

## E. Batasi output

Gunakan:

```text
Jawab maksimal 200 kata.
```

Karena output juga memakan context.

---

# 9. Hubungan dengan RAG

Context window sangat penting untuk RAG.

Flow:

```text
User bertanya
    ↓
Cari dokumen relevan
    ↓
Masukkan ke context
    ↓
LLM membaca
    ↓
Jawab
```

Kalau context kecil:

- dokumen terbatas

Kalau context besar:

- bisa baca lebih banyak

Tapi:

> lebih besar ≠ selalu lebih baik

Karena:

- lebih mahal
- lebih lambat
- noise bisa meningkat

---

# 10. Rule of Thumb untuk AI Engineer

### Chatbot sederhana

```text
8k–32k cukup
```

---

### Coding assistant

```text
32k–128k
```

---

### RAG dokumen panjang

```text
128k+
```

---

### Multi-document analysis

```text
200k–1M
```

---

# 11. Inti yang harus diingat

> Context Window adalah memori jangka pendek LLM.

Ringkasnya:

```text
Semua token masuk ke context:
- system prompt
- history
- RAG docs
- user prompt
- output
```

Kalau penuh:

- model lupa
- jawaban turun kualitas
- biaya naik

---

# Formula penting

```text
Good AI engineering =
Good prompt
+ Good retrieval
+ Good context management
```