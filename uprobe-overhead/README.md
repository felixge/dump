# uprobe overhead

Seems to be about 3usec per uprobe firing.

```
vagrant:uprobe-overhead$ time sudo bpftrace -c main ./bpf.bt 
Attaching 3 probes...
sum: 1783293664


@calls: 1000000
@switches: 0

real	0m3.472s
user	0m0.048s
sys	0m3.009s
vagrant:uprobe-overhead$ time sudo ./main 
sum: 1783293664

real	0m0.010s
user	0m0.009s
sys	0m0.000s
```
