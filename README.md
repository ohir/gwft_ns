## Go workfile tests - Bootstrap repo

This emulates a take-off in hobby settings, where protocol, server, client, and app modules
are kept together by the go.work file in the project repo.

Set of modules to test [go workfile mode](https://github.com/golang/go/issues/45713) useability.
Workfile mode is developed by the Go Team and it has a natural skew toward monorepo workflows.

To install gotip and workspace aware branch:
```
go install golang.org/dl/gotip@latest
gotip download dev.cmdgo
```

To install this (gwft_ns) development tree:

```
git clone https://github.com/ohir/gwft_ns.git
```

Now you can try to build the `app`:

```
cd gwft_ns/app
gotip build -v               # use directories as set in gwft_ns/go.work file
```


### Tags:

- compiles_015 | test broken server/client module. Should start to compile w/ `-workfile=off`.
- does_not_compile_016 | test whether tool resolved [this reported bug](https://github.com/golang/go/issues/45713#issuecomment-901475788).
- does_not_compile_021 | test whether tool resolved [this reported bug](https://github.com/golang/go/issues/45713#issuecomment-902996813).

### Notes

The modules code was updated to the v0.2.1 from the "[public](https://github.com/ohir/gwft)" subrepos.
This "bootstrap" repo uses the simplest `workspace` form, with `go.work` file being:
```
go 1.17

directory (
	./server	// example.com/mS  - uses protocol
	./client        // example.com/mC  - uses protocol
	./protocol      // example.com/mP  - no dependencies
	./app           // example.com/app - command, uses /mS, /mC, /mP
)

```

