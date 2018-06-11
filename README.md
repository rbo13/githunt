# Githunt


Githunt is a simple CLI tool for searching github users. Written in Golang.


![CircleCI](https://img.shields.io/circleci/project/github/RedSparr0w/node-csgo-parser.svg)




### Installation

Install the dependencies.

```sh
$ cd githunt
$ go get github.com/ogier/pflag
$ go get github.com/olekukonko/tablewriter
```
### Libraries

Githunt is using some helpful libraries. Instructions on how to use them in your own application are linked below.

| Plugin | README |
| ------ | ------ |
| pflag | https://github.com/ogier/pflag/blob/master/README.md] 
| tablewriter | https://github.com/olekukonko/tablewriter/blob/master/README.md] 

#### How to run?

```sh
$ go build main.go
$ go install
$ githunt -u <some_github_username>
```
