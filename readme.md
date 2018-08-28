Yama
======== Yet Another Yaml pArser

This is a small repo for parsing and merging yaml files

Maybe use yq instead for parsing and merging or appending yaml

https://github.com/mikefarah/yq

If you need schema validation use yamale

https://github.com/23andMe/Yamale

If you need linter use yamllint

https://github.com/adrienverge/yamllint


Helm
=========

Use lintconf.yml for yamllint

https://github.com/kubernetes/charts/blob/master/test/circle/lintconf.yml

Use yamale schema for Chart

https://github.com/kubernetes/charts/blob/master/test/circle/yaml-schemas/Chart.yaml

Install all these plus the semver tools

https://github.com/kubernetes/charts/blob/master/test/circle/install.sh

Get dependencies
====================

Use go mod to get dependencies instead of dep this is only available on 1.11
```
go mod init
```

Make sure to get rid `Gopkg.lock` and `Gopkg.toml` and `vendor/`

run `go run main.go` should save all your repos under `$GOPATH/src/mod`

the only difference is if your repo is under `$GOPATH` it will not work
(recursive reference and all[?])

`make build` should also fetch all the dependencies before building the
binary
To build
=============

`make build`
