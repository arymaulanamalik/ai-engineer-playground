# Belajar AI Engineer — Embedding (Cara AI Memahami Makna)

## 1. Apa itu Embedding?

**Embedding** adalah cara mengubah teks menjadi **angka (vector)** supaya komputer bisa memahami **makna**, bukan hanya kata.

Definisi sederhana:

> **Embedding = representasi numerik dari makna sebuah teks**

Contoh:

```text
"Kucing"
```

diubah menjadi:

```text
[0.12, -0.44, 0.89, 0.21, ...]
```

Itu disebut **vector embedding**.

---

# 2. Kenapa perlu Embedding?

LLM dan komputer **tidak paham kata** seperti manusia.

Komputer hanya paham angka.

Jadi teks harus diubah:

```text
Text
 ↓
Embedding Model
 ↓
Vector (angka)
```

Baru bisa dihitung.

---

# 3. Analogi paling gampang

Bayangkan setiap kata ditempatkan di **peta koordinat**.

Contoh:

```text
"Kucing" → (2, 3)
"Anjing" → (2.2, 3.1)
"Mobil"  → (10, 8)
```

Kucing dan anjing dekat.

Mobil jauh.

Artinya:

> kata yang maknanya mirip → posisi vector-nya dekat

Itulah inti embedding.

---

# 4. Similarity (kemiripan)

Embedding dipakai untuk mengukur:

> Seberapa mirip dua teks?

Contoh:

```text
"Saya lapar"
```

dan

```text
"Saya ingin makan"
```

Meski katanya berbeda, embedding-nya bisa dekat.

Karena maknanya mirip.

---

# 5. Contoh nyata vector

Contoh embedding:

```text
"Kucing" ->
[0.22, 0.55, -0.12, 0.91]
```

```text
"Anjing" ->
[0.20, 0.51, -0.10, 0.88]
```

Mirip.

---

```text
"Database" ->
[-0.88, 0.12, 0.44, -0.90]
```

Jauh.

---

# 6. Embedding Model ≠ Chat Model

Ini penting.

Ada dua model berbeda:

### Chat Model

Untuk menjawab pertanyaan.

Contoh:

```text
GPT
Claude
Gemini
```

---

### Embedding Model

Untuk mengubah teks menjadi vector.

Contoh:

```text
text-embedding-3-small
text-embedding-3-large
```

---

# 7. Embedding dipakai untuk apa?

## Semantic Search

Cari berdasarkan **makna**, bukan keyword.

Contoh:

Query:

```text
cara bikin login golang
```

Dokumen:

```text
tutorial autentikasi JWT di Go
```

Keyword beda.

Tapi embedding bisa tahu:

> ini mirip

---

## RAG (Retrieval Augmented Generation)

Flow:

```text
Dokumen
 ↓
Embedding
 ↓
Simpan di Vector DB
```

Saat user bertanya:

```text
Pertanyaan
 ↓
Embedding
 ↓
Cari vector terdekat
 ↓
Kirim ke LLM
```

Ini penggunaan paling penting.

---

## Recommendation

Contoh:

```text
User suka buku AI
```

Cari embedding buku lain yang mirip.

---

## Clustering

Mengelompokkan teks mirip.

---

## Deduplication

Cari dokumen duplikat.

---

# 8. Cara kerja embedding

Misalnya dokumen:

```text
Golang adalah bahasa backend.
```

Step:

### Step 1

Kirim ke embedding model.

---

### Step 2

Dapat vector:

```text
[0.22, -0.33, 0.98, ...]
```

---

### Step 3

Simpan ke vector database.

---

### Step 4

Saat user tanya:

```text
Backend pakai Go bagus?
```

Pertanyaan juga di-embed.

---

### Step 5

Hitung similarity.

Cari vector terdekat.

---

# 9. Similarity Metrics

Biasanya pakai:

## Cosine Similarity

Paling populer.

Nilai:

```text
1.0 → sangat mirip
0.0 → tidak mirip
-1.0 → berlawanan
```

Kamu tidak perlu hitung manual.

Vector DB yang menghitung.

---

# 10. Embedding + Vector Database

Biasanya pasangan ini selalu bersama.

Flow:

```text
Text
 ↓
Embedding Model
 ↓
Vector
 ↓
Vector Database
```

Contoh vector DB:

- pgvector (PostgreSQL)
- Pinecone
- Weaviate
- Milvus
- Qdrant

---

# 11. Contoh pakai API

Contoh request:

```json
{
  "model": "text-embedding-3-small",
  "input": "Apa itu Golang?"
}
```

Output:

```json
{
  "embedding": [0.12, -0.88, ...]
}
```

---

# 12. Contoh Golang

Pseudo code:

```go
resp, _ := client.Embeddings.Create(
    "text-embedding-3-small",
    "Apa itu Golang?",
)

vector := resp.Data[0].Embedding
```

Lalu simpan ke pgvector.

---

# 13. Hal penting yang harus diingat

## Embedding bukan jawaban

Embedding hanya representasi angka.

---

## Model embedding harus konsisten

Jangan campur:

```text
dokumen → model A
query → model B
```

Bisa jelek hasilnya.

---

## Chunking penting

Biasanya embedding dokumen dipotong:

```text
500–1000 tokens
```

Bukan seluruh PDF sekaligus.

---

# 14. Inti yang harus diingat

> Embedding mengubah makna menjadi angka.

Ringkas:

```text
Text
 ↓
Embedding Model
 ↓
Vector
 ↓
Similarity Search
 ↓
Relevant Context
 ↓
LLM
```

Tanpa embedding, **RAG tidak bisa bekerja**.