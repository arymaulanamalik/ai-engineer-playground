# Belajar AI Engineer — AI Agent (LLM yang Bisa Bertindak)

## 1. Apa itu AI Agent?

**AI Agent** adalah **LLM yang tidak hanya menjawab, tapi juga bisa mengambil tindakan (action).**

Definisi sederhana:

> **Agent = LLM + kemampuan menggunakan tools + kemampuan memutuskan langkah**

Kalau chatbot biasa:

```text
User → Tanya
LLM → Jawab
```

Kalau agent:

```text
User → Tanya
LLM → Berpikir
LLM → Pilih tool
LLM → Jalankan tool
LLM → Analisa hasil
LLM → Ambil langkah berikutnya
LLM → Jawab
```

---

# 2. Analogi paling gampang

Bayangkan dua asisten.

### Chatbot biasa

Kamu bilang:

```text
Cari cuaca Jakarta hari ini
```

Chatbot hanya bilang:

```text
Saya tidak tahu pasti.
```

---

### AI Agent

Dia akan:

```text
1. Buka weather API
2. Ambil data
3. Baca hasil
4. Ringkas
5. Jawab ke kamu
```

Itulah agent.

---

# 3. Bedanya RAG vs Agent

Ini sering membingungkan.

## RAG

Tujuan:

```text
Mencari informasi
```

Flow:

```text
Cari dokumen → kasih ke LLM → jawab
```

---

## Agent

Tujuan:

```text
Melakukan aksi
```

Flow:

```text
Pikir → pilih tool → execute → evaluasi → lanjut
```

---

Contoh:

### RAG

```text
Tanya SOP cuti
```

Cari PDF → jawab.

---

### Agent

```text
Ajukan cuti saya untuk besok
```

Agent bisa:

- cek kalender
- cek saldo cuti
- submit request
- kirim email

---

# 4. Komponen AI Agent

## A. LLM (Brain)

Otak agent.

Bertugas:

- memahami tujuan
- reasoning
- memilih langkah

---

## B. Tools

Kemampuan tambahan.

Contoh:

- database query
- REST API
- web search
- calculator
- email sender
- filesystem
- terminal
- CRM system

---

## C. Memory

Menyimpan state.

Contoh:

```text
User sedang booking hotel
sudah pilih tanggal
belum pilih kamar
```

---

## D. Planner

Membuat rencana.

Contoh:

```text
1. Cari tiket
2. Bandingkan harga
3. Booking
```

Kadang planner implicit di LLM.

---

## E. Executor

Menjalankan tool.

---

## F. Observer

Membaca hasil tool dan menentukan langkah berikutnya.

---

# 5. Flow dasar AI Agent

```text
User Goal
   ↓
LLM reasoning
   ↓
Choose Tool
   ↓
Execute Tool
   ↓
Observe Result
   ↓
Need another action?
   ├─ Yes → loop
   └─ No → final answer
```

Ini disebut:

> **Reason → Act → Observe (ReAct)**

---

# 6. Contoh nyata

Prompt:

```text
Cari restoran sushi terbaik dekat kantor saya dan booking untuk jam 7 malam.
```

Agent bisa:

### Step 1

Gunakan Maps API.

Cari restoran.

---

### Step 2

Bandingkan rating.

---

### Step 3

Gunakan reservation API.

Booking meja.

---

### Step 4

Konfirmasi hasil.

---

LLM biasa tidak bisa melakukan ini sendiri.

---

# 7. Function Calling = dasar Agent

Kebanyakan agent modern dibangun dengan:

> **Function Calling / Tool Calling**

Contoh tools:

```json
[
  {
    "name": "get_weather"
  },
  {
    "name": "send_email"
  }
]
```

LLM bisa memilih:

```text
Saya perlu panggil get_weather()
```

---

# 8. Single-step vs Multi-step Agent

## Single-step

Satu tool.

Contoh:

```text
cek cuaca
```

---

## Multi-step

Banyak tool berantai.

Contoh:

```text
Cari customer
→ ambil invoice
→ hitung denda
→ kirim email reminder
```

Lebih kompleks.

---

# 9. Agent Patterns

## Tool Use Agent

Pakai tool sesuai kebutuhan.

Paling umum.

---

## Planning Agent

Bikin rencana dulu.

---

## Multi-Agent

Beberapa agent kerja sama.

Contoh:

```text
Research Agent
Coding Agent
Testing Agent
```

---

## Supervisor Agent

Satu agent mengontrol agent lain.

---

# 10. Risiko AI Agent

## Infinite loop

Agent terus mencoba.

---

## Tool misuse

Memanggil tool salah.

---

## Hallucinated actions

Mengira tool berhasil padahal gagal.

---

## Security risk

Kalau tool punya akses:

- database
- email
- file system

Harus hati-hati.

---

# 11. Stack AI Agent di Golang

Contoh arsitektur:

```text
Go App
   |
   |-- LLM API
   |
   |-- Tool Registry
   |     |- Search()
   |     |- DBQuery()
   |     |- SendEmail()
   |
   |-- Agent Loop
   |
   |-- Memory Store
```

---

## Pseudo-code sederhana

```go
for {
    decision := AskLLM(state)

    if decision.Tool != "" {
        result := RunTool(decision.Tool)
        state.Add(result)
        continue
    }

    break
}
```

---

# 12. Kapan pakai Agent?

Gunakan agent kalau perlu:

- multi-step reasoning
- tool usage
- workflow automation
- decision making

Jangan pakai agent kalau cukup:

- Q&A biasa
- summarization
- simple RAG

Karena:

> Agent lebih kompleks dan lebih mahal.

---

# 13. Inti yang harus diingat

> Agent = LLM yang bisa bertindak.

Ringkas:

```text
LLM
 + Memory
 + Tools
 + Reasoning
 + Action Loop
 = Agent
```

Chatbot:

```text
hanya menjawab
```

Agent:

```text
bisa melakukan sesuatu
```