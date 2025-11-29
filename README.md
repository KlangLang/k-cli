<div align="center">
  <img src="assets/logo.png" width="160">
</div>

# loom

Official command-line interface for the Klang programming language.

_Loom is the command-line interface that shapes every Klang project.  
Designed for consistency, built for clarity._

---

## Installation

Install loom in a single step. The environment adjusts around it.

### Linux

Run the installer:

```sh
chmod +x install.sh
./install.sh
```

This will:

- detect your platform
- download the correct release
- extract loom
- install it to `/usr/local/bin` (or `~/.local/bin` when sudo is unavailable)

After installation:

```sh
loom --version
```

### Windows

Download the release for your architecture from:

```
https://github.com/KlangLang/loom/releases
```

Extract the `.zip` and place `loom.exe` somewhere in your PATH,
such as:

```
%USERPROFILE%\bin
```

Or use `install.bat` if provided.

---

## Commands

Loom provides a concise, predictable workflow.

```text
loom new <project>     Create a new Klang project
loom lex <file.k>      Lexicalize a Klang source file

-V, --version          Show versions
-h, --help             Show help
```

---

## Usage

Create a new project:

```sh
loom new my_project
```

Analyze a Klang source file:

```sh
loom lex main.k
```

---

## Releases

All versions of loom are available at:

```
https://github.com/KlangLang/loom/releases
```

Each release contains builds for Linux, Windows, and macOS.

---

## About

Loom is part of the Klang toolchain.
Its purpose is to provide a precise, stable interface for project initialization,
lexical analysis, and future language tooling.
