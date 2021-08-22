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

To install this (gwft-nosub) development tree:

```
git clone --tags https://github.com/ohir/gwft_ns.git
```

Now you can build app in either mode:

```
cd gwft_ns/app
gotip -workfile=off build    # use modules
gotip build                  # use directories as set in gwft_ns/go.work file
```

Note that this repo uses `example.com/` prefixed paths.
You can change go.mod and git config to point to your github clone(s) issuing (in the gwft_ns directory):

```
find . -name config -exec sed -ire "s#ohir/#yoursGH/#" {} \;
find . -name go.mod -exec sed -ire "s#ohir/#yoursGH/#" {} \;
```

### Tags:

- compiles_015 | test broken server/client module. 
- does_not_compile_016 | test whether tool resolved [this reported bug](https://github.com/golang/go/issues/45713#issuecomment-901475788).
- does_not_compile_021 | test whether tool resolved [this reported bug](https://github.com/golang/go/issues/45713#issuecomment-902996813).
- does_not_compile_021 

