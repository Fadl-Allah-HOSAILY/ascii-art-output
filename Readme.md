<h1 align="center">🎨 ASCII Art Banner Generator</h1>
<p align="center">
  <strong>A command-line tool written in Go that generates ASCII art banners and writes them to a file.</strong><br>
  The tool allows customizable banners and enforces strict usage formatting for reliability and consistency.
</p>

---

## 📋 Features

- 🔤 Converts input strings into ASCII art using a chosen banner style.
- 💾 Supports writing output to a specified file using the `--output=<fileName.txt>` flag.
- 🧭 Provides usage instructions on incorrect input.
- 🎨 Supports multiple banner styles (e.g., `standard`, `shadow`, etc.).
- 🛠️ Accepts:
  - A **single string argument** with a default banner.
  - An **optional banner style**.
  - An `--output` flag to **redirect output to a file**.

---

## 🚀 Usage

```bash
go run . [OPTION] [STRING] [BANNER]
