# art-decoder

A command-line tool for encoding and decoding text-based art using a compact run-length notation.

---

## Project Overview

`art-decoder` reads an encoded string and expands repeated sequences into their full text form. The notation lets artists describe ASCII images with heavy repetition in a much shorter form — instead of writing `@@@@@@@@@`, write `[9 @]`.

### Notation Format

Repeated characters are written as `[N pattern]`, where `N` is the repeat count and `pattern` is everything (character or string) after the single delimiter space.

| Encoded | Decoded |
|---|---|
| `[5 #]` | `#####` |
| `[3 -_]` | `-_-_-_` |
| `ABC[10 D]EFG` | `ABCDDDDDDDDDDEFG` |
| `[5 #][5 -_]-[5 #]` | `#####-_-_-_-_-_-#####` |

---

## Setup & Installation

**Requirements:** Go 1.18 or later

```bash
git clone ...
cd art-decoder
go build -o art-decoder .
```

---

## Usage

### Basic Decode (default)

```bash
./art-decoder "<encoded string>"
```
Example:
```
$ ./art-decoder "[5 #][5 -_]-[5 #]"
#####-_-_-_-_-_-#####

$ ./art-decoder "ABC[10 D]EFG"
ABCDDDDDDDDDDEFG
```

### Error Handling

If the input is malformed, the program prints `Error` and exits with code 1.

| Case | Example |
|---|---|
| First argument not a number | `[abc X]` |
| No space separator | `[5m]` |
| Empty pattern | `[5 ]` |
| Unbalanced opening bracket | `[10 A` |
| Unbalanced closing bracket | `hello]world` |

---

## Extras

### Encode Mode (`--encode` / `-e`)

Compresses plain text art back into the notation. The encoder scans each position and tries every possible pattern length, picking whichever gives the greatest character saving.

```
$ ./art-decoder --encode "#####-_-_-_-_-_-#####"
#####[5 -_]-#####

$ ./art-decoder -e "ABCDDDDDDDDDDEFG"
ABC[10 D]EFG
```

### Multiline Mode (`--multi` / `-m`)

Reads multiple lines from `stdin`, processing each line independently. Works with both decode (default) and encode mode. Noting that operator `<` is used to feed the content of the text file as a `stdin` to a command.

```bash
# Decode a multi-line encoded file
./art-decoder --multi < plane.encoded.txt

# Encode a plain art file
./art-decoder -e -m < plane.art.txt
```

---

## Bonus Feature: Paint Mode (`--paint` / `-p`)

Colorizes decoded art in the terminal. Each unique character is assigned a distinct, stable ANSI 256-color. Flag `-p` does not paint the ouput when `--encode` mode is on.

```bash
./art-decoder --paint "[5 #][5 -_]-[5 #]"
./art-decoder -p -m < plane.encoded.txt
```

Paint mode has the best effect when used combining with `--multi`.

---

## All Flags

```
Flags:
  --encode, -e    Encode plain text into art-decoder notation
  --multi,  -m    Read multiple lines from stdin
  --paint,  -p    Colorize output with ANSI 256 colors per unique character
  --help,   -h    Show usage
```

If unknown flag got passed in command line, the program will throw errors and print supported flags to the terminal.

---

## Design Notes

- The **decoder** is a single-pass O(n) character scanner — no regex, no splitting.
- The **encoder** uses a greedy best-saving search: at each position it tries every pattern length up to `remaining/2` and picks whichever compresses most. It only encodes when doing so actually reduces the output length.
- The **paint** feature builds a `rune → color index` map that persists across lines, so character coloring is consistent throughout a multi-line art piece.
- `bufio.Scanner` with a 1 MB line buffer handles very long lines in large art files without issues.
- Square brackets `[]` are **not printable** in this format.