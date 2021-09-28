<div align="center">
<img  width="320"  src="https://user-images.githubusercontent.com/4745789/135118232-51cf3223-288f-4bd5-8c08-138c7266b4aa.png" align="center"  alt="restatic logo" />

---------------------------------------

[![Release](https://img.shields.io/github/release/relogHQ/restatic/all.svg)](https://github.com/relogHQ/restatic/releases)
[![Twitter Follow](https://img.shields.io/twitter/follow/relogHQ.svg?label=Follow&style=social)](https://twitter.com/intent/follow?screen_name=relogHQ)

[![License](https://img.shields.io/github/license/apache/pinot.svg)](LICENSE)

</div>

#  What is Restatic?

Restatic is a simple HTTP server that serves a local directory over HTTP. It is written in [Go](https://golang.org/) and its usage is inspired by [Python's](https://www.python.org/)  [http.server](https://docs.python.org/3/library/http.server.html) module.

##  Using Restatic

To use Restatic download the latest release for your platform from  and fire the following commands

 1. Download the latest [restatic/releases](https://github.com/relogHQ/restatic/releases) for your platform.
 2. Extract the binary from the compressed artifact
 3. Execute the binary

```
$ ./restatic -p 4001 -d .

        ██████  ███████ ███████ ████████  █████  ████████ ██  ██████ 
        ██   ██ ██      ██         ██    ██   ██    ██    ██ ██      
        ██████  █████   ███████    ██    ███████    ██    ██ ██      
        ██   ██ ██           ██    ██    ██   ██    ██    ██ ██      
        ██   ██ ███████ ███████    ██    ██   ██    ██    ██  ██████ 

        by Relog - https://relog.in

INFO[0000] server listening on :4001                    
INFO[0000] ========================
```
  
##  Developing Restatic

Here are the pre-requisites for setting up the development environment.

- Go 1.17.1

The steps:
- Clone the repository
- Start the server `go run cmd/restatic/main.go`

##  Linting
Before you commit, ensure that the code is formatted using `gofmt` and to do that execute
```
make lint
```
##  Guidelines

If you are a developer and want to contribute to the project, please follow the [code contribution guidelines](https://github.com/relogHQ/restatic/blob/master/CONTRIBUTING.md).

##  Relog Umbrella
<div align="center">
<br />
<img  width="240"  src="https://user-images.githubusercontent.com/4745789/133601178-711aa4eb-f836-4e93-a554-22006648f75f.png" align="center"  alt="relog logo" />
<br />
<br />
</div>

Relog consists of a lot of handy utilities, systems, and projects that are aimed at making people realize how easy it is to build seemingly complex components. All the projects are open sourced and fall under [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0). You can check out all the open source projects at [github.com/relogHQ](https://github.com/relogHQ).

##  License
Restatic is under [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)
