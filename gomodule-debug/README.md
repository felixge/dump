This repo shows a problem I have with go mod. The go.mod file has a single requirement:

```
$ cat go.mod
module github.com/felixge/dump/gomodule-debug

go 1.15

require github.com/prometheus/client_golang v0.8.1-0.20161102184821-e56081f7b9d0
```

Running `go mod graph` prints this:

```
$ go mod graph
github.com/felixge/dump/gomodule-debug github.com/prometheus/client_golang@v0.8.1-0.20161102184821-e56081f7b9d0
```

That makes sense. But what doesn't make sense is what happens if I run `go mod tidy`:

```
$ go mod tidy
go: finding module for package github.com/beorn7/perks/quantile
go: finding module for package github.com/prometheus/client_model/go
go: finding module for package github.com/prometheus/common/model
go: finding module for package github.com/golang/protobuf/proto
go: finding module for package github.com/prometheus/common/expfmt
go: finding module for package github.com/prometheus/procfs
go: found github.com/beorn7/perks/quantile in github.com/beorn7/perks v1.0.1
go: found github.com/golang/protobuf/proto in github.com/golang/protobuf v1.4.2
go: found github.com/prometheus/client_model/go in github.com/prometheus/client_model v0.2.0
go: found github.com/prometheus/common/expfmt in github.com/prometheus/common v0.14.0
go: found github.com/prometheus/common/model in github.com/prometheus/common v0.14.0
go: found github.com/prometheus/procfs in github.com/prometheus/procfs v0.2.0
$ git diff
diff --git a/gomodule-debug/go.mod b/gomodule-debug/go.mod
index 1bc2c6e..00d0d06 100644
--- a/gomodule-debug/go.mod
+++ b/gomodule-debug/go.mod
@@ -2,4 +2,8 @@ module github.com/felixge/dump/gomodule-debug
 
 go 1.15
 
-require github.com/prometheus/client_golang v0.8.1-0.20161102184821-e56081f7b9d0
+require (
+       github.com/prometheus/client_golang v1.7.1
+       github.com/prometheus/common v0.14.0 // indirect
+       github.com/prometheus/procfs v0.2.0 // indirect
+)
```

Why does `go mod tidy` upgrade my dependency from v0.8.1 to v1.7.1?
