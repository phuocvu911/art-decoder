# art-decoder

A command-line tool that converts encoded art notation into text-based ASCII art — and back again.

## Project Overview

`art-decoder` lets artists like Chris describe repetitive ASCII art compactly using a simple bracket notation:

```
[N pattern]
```

This expands `pattern` exactly `N` times inline. Mixed freely with literal characters, it makes large art easy to author and store.

## Setup & Installation

**Requirements:** Go 1.18+ (tested on Go 1.22)

```bash
# Clone or copy the project, then:
cd art-decoder
go build -o art-decoder .
```

Or run directly without building:

```bash
go run . "[5 #][5 -_]-[5 #]"
```

## Usage

### Decode (default)

Decode a single encoded string:

```bash
./art-decoder "[5 #][5 -_]-[5 #]"
# Output: #####-_-_-_-_-_-#####

./art-decoder "ABC[10 D]EFG"
# Output: ABCDDDDDDDDDDEFG
```

### Encode

Convert plain text into compact notation with `--encode` (or `-e`):

```bash
./art-decoder --encode "#####-_-_-_-_-_-#####"
# Output: #####[5 -_]-#####

./art-decoder --encode "ABCDDDDDDDDDDEFG"
# Output: ABC[10 D]EFG
```

### Multi-line mode

Process multiple lines read from stdin with `--multi` (or `-m`):

```bash
# Decode a multi-line encoded file:
./art-decoder --multi < plane.encoded

# Encode a multi-line art file:
./art-decoder --encode --multi < plane.art > plane.encoded

# Round-trip check:
./art-decoder --encode --multi < plane.art | ./art-decoder --multi
```

## Encoding Format

| Notation       | Meaning                                |
|----------------|----------------------------------------|
| `[5 #]`        | `#####` — repeat `#` five times        |
| `[3 -_]`       | `-_-_-_` — repeat `-_` three times    |
| `ABC[2 !]`     | `ABC!!` — literal chars + repetition  |
| `[10 D]`       | `DDDDDDDDDD`                           |

Rules:
- The number and pattern are separated by **one space**.
- Everything after that first space (including extra spaces) is the pattern.
- Square brackets are **not** printable in output.

## Error Handling

Any malformed input prints `Error` followed by a newline. Malformed cases:

| Case | Example |
|------|---------|
| First argument not a number | `[abc #]` |
| No space separator | `[5#]` |
| Empty pattern | `[5 ]` |
| Unbalanced open bracket | `[5 #` |
| Unbalanced close bracket | `5 #]` |

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--encode` | `-e` | Encode mode: plain text → notation |
| `--multi`  | `-m` | Multi-line mode: read lines from stdin |
| `--help`   | `-h` | Show usage |

## Additional Features

- **Greedy encoder**: The encoder finds the best (longest-saving) repeating unit at each position, so `[5 -_]` beats five separate `-_` pairs.
- **Large buffer support**: Multi-line mode handles lines up to 1 MB for very large art files.
- **Composable flags**: `--encode --multi` works together for full pipeline use.