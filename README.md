# cmd-go-js

cmd-go-js is a "feature branch" of cmd/go command with experimental changes. The goal was to explore adding support for additional Go compilers and environments by making use of existing -compiler and -exec flags:

```
-compiler name
	name of compiler to use, as in runtime.Compiler (gccgo or gc).

-exec xprog
    Run the test binary using xprog. The behavior is the same as
    in 'go run'. See 'go help run' for details.
```

Specifically, I wanted to try adding support for [GopherJS compiler](https://github.com/gopherjs/gopherjs) which targets GOARCH=js architecture. It compiles Go code to JavaScript which can be executed by a JavaScript engine (such as V8 JavaScript Engine, Node.js, or any browser with JavaScript support).

#### What works

```bash
$ GOARCH=amd64 go build -compiler=gc .../hello-world
$ ./hello-world
Hello brave new world! It is working on go1.5.3 darwin/amd64!

$ GOARCH=js go build -compiler=gopherjs github.com/gophergala2016/cmd-go-js/cmd/samples/hello-world
$ node ./hello-world
Hello brave new world! It is working on go1.5.3 darwin/js!
```



### Installation

**Note:** Due to time limits, this project works on OS X only. Adding support for other other systems is easy, but it hasn't been done yet.

Go 1.5 is required.

```
go get -u github.com/gopherjs/gopherjs
GO15VENDOREXPERIMENT=1 go get -u github.com/gophergala2016/cmd-go-js/cmd/go
```

Since the binary is also called `go`, make sure you run the right one.

```bash
$ $GOPATH/bin/go version
go version go1.5.3 darwin/amd64 (with experimental changes)
```











### Requirements

- Go 1.5.3
	- Install at https://golang.org/doc/install.
	- Go 1.4.x is not supported anymore, 1.5.x is required.
	- Go 1.6.x is not supported yet, see [gopherjs/gopherjs#355](https://github.com/gopherjs/gopherjs/issues/355) tracking progress.
- GopherJS
	- Install at https://github.com/gopherjs/gopherjs#installation-and-usage.
	- Only latest `master` version is supported.
