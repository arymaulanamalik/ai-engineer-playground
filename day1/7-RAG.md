# Belajar AI Engineer — RAG (Retrieval Augmented Generation)

## 1. Apa itu RAG?

**RAG** adalah teknik untuk membuat LLM bisa menjawab berdasarkan **data kita sendiri**.

Kepanjangannya:

```text
Retrieval Augmented Generation
```

Artinya:

- **Retrieval** → ambil informasi yang relevan
- **Augmented** → tambahkan ke prompt
- **Generation** → LLM menghasilkan jawaban

Definisi sederhana:

> **RAG = kasih “contekan” ke LLM sebelum dia menjawab**

---

# 2. Kenapa perlu RAG?

Karena LLM punya masalah:

### Hallucination

Kadang model:

- terdengar yakin
- tapi salah

---

### Knowledge lama

Model bisa tidak tahu data terbaru.

Misalnya:

```text
Berapa saldo user?
```

LLM tidak tahu.

---

### Tidak tahu data private

Contoh:

- dokumen perusahaan
- SOP internal
- source code internal
- database bisnis

LLM tidak punya akses.

---

Solusi:

> Cari data dulu → kasih ke model → baru jawab

Itulah RAG.

---

# 3. Analogi paling gampang

Bayangkan ujian **open book**.

Tanpa buku:

```text
Tebak berdasarkan ingatan
```

Dengan buku:

```text
Cari halaman relevan
Baca
Jawab
```

RAG = model boleh buka buku dulu.

---

# 4. Flow besar RAG

```text
User bertanya
      ↓
Convert query ke embedding
      ↓
Cari dokumen relevan di Vector DB
      ↓
Ambil top-k chunks
      ↓
Masukkan ke prompt
      ↓
LLM menjawab berdasarkan context
```

---

# 5. Komponen utama RAG

## A. Documents

Sumber data:

- PDF
- DOCX
- website
- database
- source code
- markdown
- CSV

---

## B. Chunking

Dokumen dipotong kecil-kecil.

Contoh:

```text
Chunk 1
Chunk 2
Chunk 3
```

Biasanya:

```text
500–1000 tokens
```

Kenapa?

Karena embedding lebih efektif.

---

## C. Embedding

Setiap chunk diubah jadi vector.

```text
Chunk
 ↓
Embedding model
 ↓
Vector
```

---

## D. Vector Database

Tempat menyimpan embedding.

Contoh:

- pgvector
- Pinecone
- Weaviate
- Qdrant
- Milvus

---

## E. Retriever

Mencari chunk paling relevan.

Contoh:

```text
Top 3 nearest chunks
```

---

## F. LLM

LLM membaca chunk itu lalu menjawab.

---

# 6. Ingestion Pipeline (masukkan data)

Ini proses saat menyiapkan knowledge base.

```text
Load document
    ↓
Clean text
    ↓
Chunking
    ↓
Embedding
    ↓
Store in Vector DB
```

Disebut:

> **Indexing / Ingestion**

---

# 7. Query Pipeline (saat user bertanya)

Saat runtime:

```text
User question
    ↓
Embedding query
    ↓
Similarity search
    ↓
Retrieve chunks
    ↓
Build prompt
    ↓
LLM answer
```

---

# 8. Contoh nyata

Misalnya kamu punya:

```text
employee_handbook.pdf
```

Isi:

```text
Cuti tahunan adalah 12 hari.
```

User bertanya:

```text
Berapa jatah cuti?
```

LLM sendiri mungkin tidak tahu.

Dengan RAG:

```text
Retriever menemukan chunk:
"Cuti tahunan adalah 12 hari."
```

Lalu prompt menjadi:

```text
Gunakan konteks berikut:
Cuti tahunan adalah 12 hari.

Pertanyaan:
Berapa jatah cuti?
```

LLM jawab:

```text
Jatah cuti tahunan adalah 12 hari.
```

Lebih akurat.

---

# 9. Kenapa RAG lebih bagus daripada Fine-tuning?

## RAG

### Kelebihan

- data bisa update cepat
- tidak perlu retrain
- murah
- bisa cite source

### Kekurangan

- butuh retrieval bagus
- kompleksitas sistem naik

---

## Fine-tuning

### Kelebihan

- model lebih adaptif
- gaya jawaban bisa spesifik

### Kekurangan

- mahal
- update lambat
- training rumit

---

Rule umum:

> **Knowledge baru → pakai RAG**  
> **Behavior baru → pakai fine-tuning**

---

# 10. Tantangan dalam RAG

## Chunking jelek

Chunk terlalu besar:

- noise

Chunk terlalu kecil:

- konteks hilang

---

## Retrieval jelek

Chunk relevan tidak ditemukan.

---

## Context terlalu banyak

Prompt penuh.

---

## Embedding mismatch

Dokumen dan query pakai model berbeda.

---

## Prompt buruk

LLM tidak fokus pada context.

---

# 11. RAG Stack untuk Golang

Contoh arsitektur sederhana:

```text
Go API
   |
   |-- Embedding API
   |
   |-- pgvector
   |
   |-- Retriever
   |
   |-- Prompt Builder
   |
   |-- LLM API
```

---

## Tools Go populer

### LangChainGo

Framework orchestration.

---

### pgvector

Simpan embedding di PostgreSQL.

---

### sqlx / gorm

Database access.

---

### OpenAI / Gemini SDK

Untuk embedding + generation.

---

# 12. Pseudo-code sederhana

```go
question := "Apa jatah cuti?"

queryVector := Embed(question)

chunks := SearchSimilar(queryVector)

prompt := BuildPrompt(question, chunks)

answer := AskLLM(prompt)

fmt.Println(answer)
```

Itulah inti RAG.

---

# 13. Inti yang harus diingat

> RAG = cari informasi dulu, baru jawab.

Ringkas:

```text
Documents
   ↓
Chunking
   ↓
Embedding
   ↓
Vector DB
   ↓
User Query
   ↓
Retrieve relevant chunks
   ↓
Add to prompt
   ↓
LLM Answer
```

Tanpa retrieval:

```text
LLM menebak
```

Dengan RAG:

```text
LLM menjawab berdasarkan data
```