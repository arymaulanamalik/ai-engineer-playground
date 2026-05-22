# Belajar AI Engineer — Temperature (Mengatur Kreativitas LLM)

## 1. Apa itu Temperature?

**Temperature** adalah parameter untuk mengatur **seberapa random / kreatif** jawaban LLM.

Sederhananya:

> **Temperature = tombol kreativitas model**

Biasanya nilainya:

```text
0.0 → sangat pasti / deterministic
2.0 → sangat random
```

Umumnya dipakai:

```text
0.0 sampai 1.0
```

---

# 2. Analogi paling gampang

Bayangkan model mau memilih kata berikutnya.

Contoh:

```text
Saya minum kopi dengan ...
```

Model punya kemungkinan:

```text
susu     → 60%
gula     → 25%
es       → 10%
garam    → 5%
```

---

## Kalau temperature rendah

Misalnya:

```text
temperature = 0
```

Model hampir selalu pilih:

```text
susu
```

Karena probabilitas paling tinggi.

Hasil:

- stabil
- konsisten
- predictable

---

## Kalau temperature tinggi

Misalnya:

```text
temperature = 1.2
```

Model bisa pilih:

```text
gula
```

atau:

```text
es
```

bahkan kadang:

```text
garam
```

Hasil:

- lebih kreatif
- lebih variatif
- kadang aneh

---

# 3. Contoh nyata

Prompt:

```text
Buat nama startup AI.
```

---

## Temperature rendah (0.1)

Output bisa seperti:

```text
AI Solutions
SmartAI
DataMind
```

Aman, tapi biasa.

---

## Temperature tinggi (1.0)

Output bisa seperti:

```text
NeuroSpark
QuantumWhisper
MindForge
```

Lebih kreatif.

---

# 4. Kapan pakai temperature rendah?

Biasanya untuk tugas yang butuh **akurasi dan konsistensi**.

Contoh:

- coding
- SQL generation
- JSON output
- classification
- summarization
- extraction
- debugging
- customer support FAQ

Rekomendasi:

```text
temperature = 0.0–0.3
```

---

# 5. Kapan pakai temperature tinggi?

Untuk tugas yang butuh **ide dan kreativitas**.

Contoh:

- brainstorming
- copywriting
- storytelling
- naming
- marketing text
- desain ide

Rekomendasi:

```text
temperature = 0.7–1.0
```

---

# 6. Sebagai AI Engineer, default aman berapa?

Biasanya:

```text
0.2
```

Kenapa?

Karena:

- cukup stabil
- tidak terlalu kaku
- masih sedikit fleksibel

Ini sering jadi default production.

---

# 7. Dampak Temperature pada Hallucination

Semakin tinggi temperature:

- model lebih berani
- model lebih kreatif
- risiko halusinasi naik

Semakin rendah:

- lebih hati-hati
- lebih konsisten
- biasanya lebih aman

Tapi:

> Temperature rendah **tidak menjamin benar**.

Model tetap bisa salah.

---

# 8. Temperature ≠ Intelligence

Ini penting.

Temperature **tidak membuat model lebih pintar**.

Temperature hanya mengubah:

```text
cara memilih token
```

Bukan menambah pengetahuan.

---

# 9. Contoh di API

Contoh pakai OpenAI API:

```json
{
  "model": "gpt-4",
  "temperature": 0.2
}
```

---

Contoh Golang:

```go
req := openai.ChatCompletionRequest{
    Model: "gpt-4",
    Temperature: 0.2,
}
```

---

# 10. Cheat Sheet

## Coding

```text
0.0–0.2
```

---

## JSON output

```text
0.0
```

---

## RAG / QA

```text
0.1–0.3
```

---

## Chatbot umum

```text
0.5
```

---

## Brainstorming

```text
0.7–0.9
```

---

## Creative writing

```text
1.0+
```

---

# 11. Inti yang harus diingat

> Temperature mengatur kreativitas model.

Ringkasnya:

```text
Low temperature
→ stabil
→ konsisten
→ aman

High temperature
→ kreatif
→ variatif
→ lebih berisiko
```

---

# Rule Praktis

Kalau bingung:

```text
Mulai dari 0.2
```

Itu pilihan paling aman untuk AI Engineer.
