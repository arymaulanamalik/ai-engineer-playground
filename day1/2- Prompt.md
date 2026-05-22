# Belajar AI Engineer — Prompt (Dasar Prompt Engineering)

## 1. Apa itu Prompt?

**Prompt** adalah **instruksi atau input** yang kita berikan ke LLM agar model menghasilkan jawaban yang kita inginkan.

Contoh sederhana:

```text
Apa ibu kota Jepang?
```

Itu juga prompt.

Contoh lain:

```text
Jelaskan Golang seperti saya anak SMA.
```

Itu juga prompt.

Jadi:

> **Prompt = cara kita berbicara dengan LLM**

---

# 2. Kenapa Prompt Penting?

LLM sangat bergantung pada prompt.

Prompt bagus → jawaban bagus  
Prompt jelek → jawaban ngawur

Contoh:

### Prompt buruk

```text
jelasin docker
```

Hasil:

- terlalu umum
- bisa terlalu panjang
- bisa tidak sesuai kebutuhan

---

### Prompt lebih baik

```text
Jelaskan Docker untuk backend engineer Golang.
Berikan analogi sederhana dan contoh dockerfile.
```

Hasil:

- lebih relevan
- lebih terarah
- lebih berguna

---

# 3. Cara pikir LLM terhadap Prompt

LLM sebenarnya membaca prompt seperti:

```text
Oh, user ingin:
- topik: Docker
- level: backend engineer
- bahasa: Indonesia
- butuh: analogi
- butuh: contoh code
```

Semakin jelas instruksi, semakin bagus hasilnya.

---

# 4. Struktur Prompt yang Bagus

Template sederhana:

```text
Role + Task + Context + Format + Constraint
```

Mari kita pecah.

---

## A. Role (suruh model jadi siapa)

Contoh:

```text
Kamu adalah senior Golang engineer.
```

atau:

```text
Jawab sebagai AI architect.
```

Kenapa penting?

Karena model akan menyesuaikan gaya jawabannya.

---

## B. Task (apa yang harus dilakukan)

Contoh:

```text
Jelaskan apa itu Redis.
```

atau:

```text
Buatkan code Golang untuk upload file.
```

Task harus jelas.

---

## C. Context (konteks tambahan)

Contoh:

```text
Saya pemula.
Saya baru belajar backend.
```

atau:

```text
Saya memakai Go tanpa framework.
```

Context membantu model menjawab lebih tepat.

---

## D. Format (bentuk output)

Contoh:

```text
Jawab dalam bullet point.
```

atau:

```text
Buatkan markdown.
```

atau:

```text
Berikan tabel perbandingan.
```

Ini sangat penting.

---

## E. Constraint (batasan)

Contoh:

```text
Jangan lebih dari 200 kata.
```

atau:

```text
Gunakan bahasa sederhana.
```

atau:

```text
Gunakan Golang, jangan Python.
```

Ini menjaga output tetap sesuai kebutuhan.

---

# 5. Contoh Prompt Lengkap

Contoh prompt jelek:

```text
buat api login
```

---

Versi lebih bagus:

```text
Kamu adalah senior Golang engineer.

Buatkan REST API login menggunakan Go standard library tanpa framework.

Gunakan:
- clean architecture
- mysql
- jwt authentication

Berikan:
- struktur folder
- contoh code
- penjelasan singkat

Gunakan markdown.
```

Hasil biasanya jauh lebih bagus.

---

# 6. Teknik Prompt Engineering Dasar

## Zero-shot Prompting

Langsung tanya tanpa contoh.

```text
Jelaskan apa itu vector database.
```

Paling umum dipakai.

---

## One-shot Prompting

Kasih satu contoh.

```text
Contoh:

Input: apel
Output: fruit

Input: mobil
Output:
```

Model akan mengikuti pola.

---

## Few-shot Prompting

Kasih beberapa contoh.

```text
Input: apel
Output: fruit

Input: wortel
Output: vegetable

Input: ayam
Output:
```

Bagus untuk klasifikasi.

---

## Chain-of-Thought Prompting

Minta model berpikir langkah demi langkah.

```text
Jelaskan langkah demi langkah.
```

atau:

```text
Think step by step.
```

Bagus untuk:

- logika
- debugging
- analisis

---

# 7. Prompt Patterns yang Sering Dipakai AI Engineer

## Explain

```text
Jelaskan RAG dengan bahasa sederhana.
```

---

## Generate Code

```text
Buatkan code Golang untuk membaca CSV.
```

---

## Refactor

```text
Refactor code ini agar clean.
```

---

## Debug

```text
Kenapa code ini panic?
```

---

## Summarize

```text
Ringkas dokumen ini.
```

---

## Transform

```text
Ubah JSON ini jadi struct Go.
```

---

## Extract

```text
Ambil semua email dari text ini.
```

---

# 8. Common Mistakes

## Terlalu pendek

```text
jelasin api
```

Terlalu ambigu.

---

## Tidak kasih konteks

```text
buat login
```

Login apa?

---

## Tidak minta format

Model bisa jawab berantakan.

---

## Tidak kasih constraint

Model bisa terlalu panjang.

---

# 9. Formula Prompt yang Aman

Kalau bingung, pakai template ini:

```text
Kamu adalah [ROLE].

Tugas:
[TASK]

Konteks:
[CONTEXT]

Output:
[FORMAT]

Batasan:
[CONSTRAINT]
```

Contoh:

```text
Kamu adalah senior AI engineer.

Tugas:
Jelaskan embedding.

Konteks:
Saya backend engineer Golang yang baru belajar AI.

Output:
Markdown dengan bullet point.

Batasan:
Gunakan bahasa sederhana.
```

---

# 10. Inti yang Harus Diingat

> Prompt yang baik = instruksi yang jelas.

Checklist:

- siapa model harus bertindak?
- apa tugasnya?
- konteks apa?
- output seperti apa?
- batasannya apa?

Semakin jelas prompt → semakin bagus jawaban.