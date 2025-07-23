# ASCII Banner Generator in Go

## Overview

This project is a command-line ASCII banner generator written in Go. It takes a string input and outputs an ASCII art representation of that string into a specified file.

The program accepts flags and arguments to customize the output file name, the string to convert, and the style of the ASCII art (banner).

---

## Features

- Output ASCII art into a file specified by the `--output=<fileName.txt>` flag.
- Supports multiple ASCII art styles (banners).
- Accepts a minimum of one argument (the string to convert) and optional arguments for style and output file.

---

## Usage

```bash
go run . --output=<fileName.txt> [STRING] [BANNER]
