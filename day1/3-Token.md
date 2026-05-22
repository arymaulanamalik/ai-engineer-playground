# Belajar AI Engineer — Tokens (Dasar yang Wajib Dipahami)

## 1. Apa itu Token?

**Token** adalah **unit kecil teks** yang dibaca oleh LLM.

Penting:

> LLM tidak membaca kalimat langsung.  
> LLM membaca **token demi token**.

---

## 2. Analogi paling gampang

Bayangkan token seperti **potongan puzzle**.

Kalimat:

```text
Saya belajar AI Engineer
```

Bisa dipecah menjadi:

```text
["Saya", " belajar", " AI", " Engineer"]
```

Itulah token.

---

## 3. Token ≠ Kata

Ini yang sering bikin bingung.

**1 token tidak selalu 1 kata.**

Contoh:

```text
ChatGPT
```

Bisa jadi:

```text
["Chat", "GPT"]
```

Atau:

```text
["Ch", "at", "GP", "T"]
```

Tergantung tokenizer model.

---

## 4. Token bisa berupa apa saja

Token bisa berupa:

### Kata

```text
Indonesia
```

---

### Potongan kata

```text
un + believable
```

---

### Spasi

Kadang:

```text
" hello"
```

(spasi ikut token)

---

### Simbol

```text
{
}
(
)
=
+
```

---

### Angka

```text
2026
```

---

### Emoji

```text
🚀
```

Semua dihitung token.

---

## 5. Kenapa Token Penting?

Karena semua hal di LLM dihitung berdasarkan token:

### Biaya

API provider biasanya charge berdasarkan token.

Contoh:

```text
$ per 1M input tokens
$ per 1M output tokens
```

---

### Context Window

Model punya batas token.

Contoh:

- 8k token
- 32k token
- 128k token
- 1M token

Kalau melebihi limit:

- prompt dipotong
- informasi hilang

---

### Kecepatan

Semakin banyak token:

- makin lambat
- makin mahal

---

## 6. Input Token vs Output Token

### Input Token

Semua yang kamu kirim.

Contoh:

```text
Jelaskan Docker
```

---

### Output Token

Semua jawaban model.

Contoh:

```text
Docker adalah platform...
```

Biasanya provider menghitung dua-duanya.

---

## 7. Contoh Hitung Token

Kalimat:

```text
Apa ibu kota Jepang?
```

Mungkin jadi:

```text
["Apa", " ibu", " kota", " Jepang", "?"]
```

Total:

```text
5 tokens
```

---

Kalimat:

```text
func main() {}
```

Mungkin jadi:

```text
["func", " main", "(", ")", " {", "}"]
```

Total:

```text
6 tokens
```

Code juga dihitung token.

---

## 8. Tokenizer itu apa?

**Tokenizer** adalah alat untuk memecah teks menjadi token.

Flow:

```text
Text
 ↓
Tokenizer
 ↓
Tokens
 ↓
Model
```

Contoh tokenizer:

- BPE (Byte Pair Encoding)
- SentencePiece
- WordPiece

Kamu tidak harus hafal algoritmanya dulu.

Cukup tahu:

> Tokenizer mengubah teks menjadi token.

---

## 9. Cara lihat jumlah token

Beberapa tools populer:

### OpenAI Tokenizer

Cari:

```text
OpenAI Tokenizer
```

---

### tiktoken (Python)

Library resmi untuk menghitung token.

---

### Go libraries

Kalau pakai Golang, biasanya:

```bash
go get github.com/pkoukk/tiktoken-go
```

Contoh:

```go
enc, _ := tiktoken.EncodingForModel("gpt-4")
tokens := enc.Encode("Halo dunia", nil, nil)

fmt.Println(len(tokens))
```

---

## 10. Kenapa prompt panjang bisa bermasalah?

Misalnya:

```text
- system prompt besar
- 20 dokumen RAG
- history chat panjang
- pertanyaan user
```

Semua masuk context.

Bisa jadi:

```text
120k tokens
```

Masalah:

- mahal
- lambat
- bisa melebihi limit

Makanya AI Engineer harus hemat token.

---

## 11. Cara menghemat token

### Ringkas prompt

Jangan:

```text
Tolong bantu saya jika Anda berkenan...
```

Lebih baik:

```text
Jelaskan...
```

---

### Ringkas chat history

Jangan kirim semua history.

Summarize dulu.

---

### Chunk dokumen

Untuk RAG:

jangan kirim seluruh PDF.

Potong jadi chunk kecil.

---

### Batasi output

Misalnya:

```text
Jawab maksimal 100 kata.
```

---

## 12. Approximation kasar

Rule of thumb:

```text
1 token ≈ 4 karakter bahasa Inggris
```

atau:

```text
100 tokens ≈ 75 kata
```

Tidak selalu akurat, tapi cukup membantu estimasi.

---

## 13. Inti yang harus diingat

> Token adalah mata uang LLM.

Semua dihitung dengan token:

- biaya
- limit
- kecepatan
- kualitas context

Flow sederhananya:

```text
Text
 ↓
Tokenizer
 ↓
Tokens
 ↓
LLM
 ↓
Output Tokens
```