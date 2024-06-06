## Overview

This project implements various compression/decompression algorithms and string search algorithms. The implemented algorithms are:
- Compression/Decompression: Huffman, LZW, Shannon-Fano
- String Search: Karp-Rabin, Knuth-Morris-Pratt (KMP)

Each algorithm has its own subproject within the main project directory. The structure is organized to facilitate ease of use and modularity.
## Directory Structure

```bash
./
.idea/
huffman/
huffman/bin/
huffman/cmd/
huffman/cmd/decode/
huffman/cmd/encode/
huffman/pkg/
huffman/pkg/huffman/
karprabin/
karprabin/bin/
karprabin/bin/search/
karprabin/cmd/
karprabin/cmd/search/
karprabin/pkg/
karprabin/pkg/karprabin/
kmp/
kmp/bin/
kmp/bin/search/
kmp/cmd/
kmp/cmd/search/
kmp/pkg/
kmp/pkg/kmp/
lzw/
lzw/bin/
lzw/cmd/
lzw/cmd/decode/
lzw/cmd/encode/
lzw/pkg/
lzw/pkg/lzw/
sfano/
sfano/bin/
sfano/cmd/
sfano/cmd/decode/
sfano/cmd/encode/
sfano/pkg/
sfano/pkg/sfano/
check_positions.sh
huffman/cmd/decode/main.go
huffman/cmd/encode/main.go
huffman/go.mod
huffman/pkg/huffman/decode.go
huffman/pkg/huffman/encode.go
huffman/pkg/huffman/huffman.go
huffman/pkg/huffman/stats.go
huffman/pkg/huffman/util.go
karprabin/cmd/search/main.go
karprabin/go.mod
karprabin/pkg/karprabin/search.go
karprabin/pkg/karprabin/stats.go
karprabin/pkg/karprabin/util.go
kmp/cmd/search/main.go
kmp/go.mod
kmp/pkg/kmp/search.go
kmp/pkg/kmp/stats.go
kmp/pkg/kmp/util.go
lzw/cmd/decode/main.go
lzw/cmd/encode/main.go
lzw/go.mod
lzw/pkg/lzw/decode.go
lzw/pkg/lzw/encode.go
lzw/pkg/lzw/stats.go
lzw/pkg/lzw/util.go
sfano/cmd/decode/main.go
sfano/cmd/encode/main.go
sfano/go.mod
sfano/pkg/sfano/decode.go
sfano/pkg/sfano/encode.go
sfano/pkg/sfano/stats.go
sfano/pkg/sfano/util.go
```


## General Compilation and Running Instructions
### Prerequisites
- Go (version 1.16 or higher)
### Building the Projects
1. **Huffman Compression/Decompression**
- Navigate to the `huffman` directory:

```bash
cd huffman
``` 
- Build the encoder and decoder:

```bash
go build -o bin/encode cmd/encode/main.go
go build -o bin/decode cmd/decode/main.go
``` 
2. **LZW Compression/Decompression**
- Navigate to the `lzw` directory:

```bash
cd lzw
``` 
- Build the encoder and decoder:

```bash
go build -o bin/encode cmd/encode/main.go
go build -o bin/decode cmd/decode/main.go
``` 
3. **Shannon-Fano Compression/Decompression**
- Navigate to the `sfano` directory:

```bash
cd sfano
``` 
- Build the encoder and decoder:

```bash
go build -o bin/encode cmd/encode/main.go
go build -o bin/decode cmd/decode/main.go
``` 
4. **Karp-Rabin String Search**
- Navigate to the `karprabin` directory:

```bash
cd karprabin
``` 
- Build the search program:

```bash
go build -o bin/search cmd/search/main.go
``` 
5. **Knuth-Morris-Pratt (KMP) String Search**
- Navigate to the `kmp` directory:

```bash
cd kmp
``` 
- Build the search program:

```bash
go build -o bin/search cmd/search/main.go
```
### Running the Programs

Each program requires specific command-line arguments to run. Below are the examples of how to run each program.
1. **Huffman Compression**

```bash
./bin/encode <input file> <output file> <stats file>
```



Example:

```bash
./bin/encode example_plain.txt example_huffman_encoded.bin huffman_stats.json
``` 
2. **Huffman Decompression**

```bash
./bin/decode <input file> <output file> <stats file>
```



Example:

```bash
./bin/decode example_huffman_encoded.bin example_huffman_decoded.txt huffman_stats.json
``` 
3. **LZW Compression**

```bash
./bin/encode <input file> <output file> <stats file>
```



Example:

```bash
./bin/encode example_plain.txt example_lzw_encoded.bin lzw_stats.json
``` 
4. **LZW Decompression**

```bash
./bin/decode <input file> <output file> <stats file>
```



Example:

```bash
./bin/decode example_lzw_encoded.bin example_lzw_decoded.txt lzw_stats.json
``` 
5. **Shannon-Fano Compression**

```bash
./bin/encode <input file> <output file> <stats file>
```



Example:

```bash
./bin/encode example_plain.txt example_sfano_encoded.bin sfano_stats.json
``` 
6. **Shannon-Fano Decompression**

```bash
./bin/decode <input file> <output file> <stats file>
```



Example:

```bash
./bin/decode example_sfano_encoded.bin example_sfano_decoded.txt sfano_stats.json
``` 
7. **Karp-Rabin String Search**

```bash
./bin/search <pattern> <input file> <stats file>
```



Example:

```bash
./bin/search "search_pattern" example_plain.txt karprabin_stats.json
``` 
8. **KMP String Search**

```bash
./bin/search <pattern> <input file> <stats file>
```



Example:

```bash
./bin/search "search_pattern" example_plain.txt kmp_stats.json
```
## Statistics

Each run of the programs generates a statistics file in JSON format, containing various metrics related to the operation. The statistics include:
- Operation type (encode/decode/search)
- Input and output file names
- Running time
- Input and output sizes (in characters and bytes)
- Compression rate (for encoding operations)
- Number of different characters (for compression operations)
- Character frequencies

These statistics can be used for analyzing the performance and effectiveness of the implemented algorithms.
## Example Usage
### Example Files

The project includes two example files that end with `_plain.txt`, which can be used to test the algorithms.
### Running an Example
1. **Huffman Compression and Decompression**

Compress `example_plain.txt` and then decompress the result:

```bash
./huffman/bin/encode example_plain.txt example_huffman_encoded.bin huffman_stats.json
./huffman/bin/decode example_huffman_encoded.bin example_huffman_decoded.txt huffman_stats.json
``` 
2. **LZW Compression and Decompression**

Compress `example_plain.txt` and then decompress the result:

```bash
./lzw/bin/encode example_plain.txt example_lzw_encoded.bin lzw_stats.json
./lzw/bin/decode example_lzw_encoded.bin example_lzw_decoded.txt lzw_stats.json
``` 
3. **Shannon-Fano Compression and Decompression**

Compress `example_plain.txt` and then decompress the result:

```bash
./sfano/bin/encode example_plain.txt example_sfano_encoded.bin sfano_stats.json
./sfano/bin/decode example_sfano_encoded.bin example_sfano_decoded.txt sfano_stats.json
``` 
4. **Karp-Rabin String Search**

Search for a pattern in `example_plain.txt`:

```bash
./karprabin/bin/search "pattern" example_plain.txt karprabin_stats.json
``` 
5. **KMP String Search**

Search for a pattern in `example_plain.txt`:

```bash
./kmp/bin/search "pattern" example_plain.txt kmp_stats.json
```
## Shell Script

The `check_positions.sh` script can be used to check specific byte positions in a file and print the context around them.
### Usage

```bash
./check_positions.sh <file> <positions> [num_chars]
```


- `<file>`: The file to check.
- `<positions>`: Space-separated list of byte positions to check.
- `[num_chars]`: (Optional) Number of characters to display around the position (default is 20).

Example:

```bash
./check_positions.sh example_plain.txt 10 50
```



This will display the context around byte position 10 in `example_plain.txt`.---
