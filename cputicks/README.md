# cpucycles

This tool tries to measure the costs of calling the `asm_amd64.s` version of
the `cpucycles()` function in Go.

Numbers from my machine (16" MBP 2019, 2.6 GHz 6-Core Intel Core i7):

```go
# macOS
$ uname -a
Darwin COMP-C02DJ092MD6V 19.6.0 Darwin Kernel Version 19.6.0: Tue Nov 10 00:10:30 PST 2020; root:xnu-6153.141.10~1/RELEASE_X86_64 x86_64 i386 MacBookPro16,1 Darwin
$ go run .
10000000 ops in 104.234628ms
10.4 ns/op

# docker for mac
$ uname -a
Linux 21ae85ccf189 4.19.121-linuxkit #1 SMP Tue Dec 1 17:50:32 UTC 2020 x86_64 GNU/Linux
$ go run .
10000000 ops in 113.8714ms
11.4 ns/op
```
