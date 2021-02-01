# go-callers-bench

This program tries to measure the costs of calling `runtime.Callers()` in Go.


Numbers from my machine (16" MBP 2019, 2.6 GHz 6-Core Intel Core i7):

```
# macOS
$ $ uname -a
Darwin COMP-C02DJ092MD6V 19.6.0 Darwin Kernel Version 19.6.0: Tue Nov 10 00:10:30 PST 2020; root:xnu-6153.141.10~1/RELEASE_X86_64 x86_64 i386 MacBookPro16,1 Darwin
$ for i in {1..5}; do go run . -depth 16; done
1.022µs/op
1.019µs/op
1.045µs/op
1.14µs/op
1.026µs/op
$ for i in {1..5}; do go run . -depth 160; done
5.938µs/op
6.146µs/op
5.922µs/op
6.027µs/op
6.064µs/op

# docker for mac
$ uname -a
Linux 021bbdbc5cbe 4.19.121-linuxkit #1 SMP Tue Dec 1 17:50:32 UTC 2020 x86_64 GNU/Linux
$ for i in {1..5}; do go run . -depth 16; done
1.082µs/op
1.096µs/op
1.081µs/op
1.076µs/op
1.071µs/op
$ for i in {1..5}; do go run . -depth 160; done
6.254µs/op
6.285µs/op
6.393µs/op
6.285µs/op
6.268µs/op
```
