Repro bug in gopls/VSCode, https://github.com/golang/go/issues/75096

Navigate to `pkg2/bar.go` and right click `BarFunc` and select `Find All References`. It only find itself, not the usage in `pkg1`. If you make the `BarFunc` non-generic, e.g. `BarFunc func() string`, it works fine.

