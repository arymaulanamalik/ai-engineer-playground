# Belajar AI Engineer — Memahami LLM (Large Language Model)

## 1. Apa itu LLM?

**LLM = Large Language Model**

Artinya:

- **Large** → modelnya sangat besar (miliaran parameter)
- **Language** → tugasnya memahami dan menghasilkan bahasa
- **Model** → rumus matematis / neural network yang belajar dari banyak data teks

Contoh LLM:

- GPT (OpenAI)
- Claude (Anthropic)
- Gemini (Google)
- Llama (Meta)
- Mistral (Mistral AI)

---

## 2. Analogi paling gampang

Bayangkan LLM seperti **mesin prediksi kata yang sangat pintar**.

Contoh:

Kalimat:

```text
Saya makan nasi dengan ...
```

Kamu mungkin jawab:

```text
ayam
```

LLM juga begitu.

Dia menghitung:

- peluang kata "ayam" = 40%
- peluang kata "ikan" = 20%
- peluang kata "sendok" = 10%

Lalu memilih kata yang paling cocok.

Jadi sebenarnya:

> **LLM = mesin yang menebak token berikutnya**

---

## 3. Kenapa bisa pintar?

Karena dilatih dengan **data sangat banyak**:

- buku
- artikel
- website
- dokumentasi code
- forum seperti Reddit
- repositori seperti GitHub

Proses ini disebut:

## Training

Model membaca triliunan kata dan belajar pola.

Contoh:

```python
if x > 0:
```

sering diikuti:

```python
print(x)
```

Atau:

```text
Jakarta adalah ibu kota ...
```

diikuti:

```text
Indonesia
```

---

## 4. Token itu apa?

LLM **tidak membaca kata**, tapi membaca **token**.

Token = potongan teks kecil.

Contoh:

```text
ChatGPT itu pintar
```

Bisa jadi token:

```text
["Chat", "GPT", " itu", " pintar"]
```

Atau:

```text
["Ch", "at", "GP", "T"]
```

Tergantung tokenizer.

### Kenapa penting?

Karena semua biaya dan limit dihitung dari token.

Contoh:

- 1 kata ≈ 1–3 token
- 1000 token ≈ 750 kata

---

## 5. Arsitektur utama: Transformer

Hampir semua LLM modern pakai:

> **Transformer**

Paper penting:

```text
Attention Is All You Need (2017)
```

Ini yang membuat GPT, Claude, Gemini mungkin ada.

### Bagian paling penting: Attention

Artinya model bisa "memperhatikan" kata penting.

Contoh:

```text
Budi memasukkan buku ke tas karena tas itu besar.
```

Kata **"itu"** merujuk ke apa?

LLM melihat konteks dan tahu:

```text
"itu" = "tas"
```

Itulah **attention**.

---

## 6. Parameter itu apa?

Parameter = angka internal model.

Semacam "ingatan" model.

Contoh ukuran model:

- 7B → 7 miliar parameter
- 70B → 70 miliar parameter
- 175B → 175 miliar parameter

Semakin besar:

✅ biasanya lebih pintar  
❌ lebih mahal  
❌ lebih lambat

---

## 7. Cara kerja saat kamu bertanya (Inference)

Misal kamu tanya:

```text
Apa ibu kota Jepang?
```

### Step 1: Tokenisasi

```text
"Apa", " ibu", " kota", " Jepang"
```

---

### Step 2: Embedding (ubah ke angka)

Semua token diubah jadi vector angka.

Misalnya:

```text
"Apa" -> [0.12, 0.88, ...]
```

---

### Step 3: Masuk Transformer

Model menghitung hubungan antar token.

---

### Step 4: Prediksi token berikutnya

Model memilih:

```text
Tokyo
```

---

### Step 5: Ulang terus

Token demi token:

```text
Tokyo adalah ibu kota Jepang.
```

---

## 8. Kenapa kadang LLM halusinasi?

Karena LLM **tidak benar-benar tahu fakta**.

Dia hanya memprediksi teks yang *terlihat benar*.

Jadi bisa saja:

- terdengar yakin
- tapi salah

Ini disebut:

> **Hallucination**

### Solusinya

- gunakan **RAG**
- grounding ke database
- tool calling
- verifikasi sumber

---

## 9. Sebagai AI Engineer, apa yang biasanya kamu lakukan?

Biasanya **tidak melatih LLM dari nol**.

Kamu lebih sering:

### a. Prompt Engineering

Mengatur cara bertanya.

Contoh:

```text
Jawab sebagai senior Golang engineer.
```

---

### b. RAG (Retrieval Augmented Generation)

Memberi dokumen tambahan.

```text
User bertanya
↓
Cari data di database/vector DB
↓
Masukkan ke prompt
↓
LLM menjawab
```

---

### c. Function / Tool Calling

LLM bisa panggil tool:

- database
- API
- calculator
- search engine

---

### d. Fine-tuning

Melatih ulang model agar spesifik.

Contoh:

- chatbot hukum
- chatbot medis
- customer service

---

## 10. Kalau pakai Golang, biasanya stack-nya seperti ini

```text
Go App
   |
   |-- OpenAI API / Claude API / Gemini API
   |
   |-- Embedding API
   |
   |-- Vector DB
         |- PostgreSQL + pgvector
         |- Pinecone
         |- Weaviate
   |
   |-- RAG Pipeline
```

### Library Go yang sering dipakai

- LangChainGo
- OpenAI Go SDK
- Google Gemini Go SDK

---

# Inti yang wajib kamu pahami

Kalau diringkas:

```text
Input text
   ↓
Tokenization
   ↓
Embedding (ubah ke angka)
   ↓
Transformer + Attention
   ↓
Predict next token
   ↓
Generate jawaban
```

Itu inti semua LLM.