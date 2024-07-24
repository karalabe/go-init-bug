The `good` and `bad` package contains the exact same code, but the build order of the files is different (a.go renamed to c.go in one vs the other). You can flip the faulty/good behavior in those packages by renaming the files.

```go
% go run main.go

Good order: 220
Bad order:  0
```

