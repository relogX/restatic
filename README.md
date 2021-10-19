<div align="center">
<img  width="320"  src="https://user-images.githubusercontent.com/4745789/135118232-51cf3223-288f-4bd5-8c08-138c7266b4aa.png" align="center"  alt="restatic logo" />

---------------------------------------

[![Release](https://img.shields.io/github/release/relogX/restatic/all.svg)](https://github.com/relogX/restatic/releases)
[![Twitter Follow](https://img.shields.io/twitter/follow/relog_x.svg?label=Follow&style=social)](https://twitter.com/intent/follow?screen_name=relog_x)
[![License](https://img.shields.io/github/license/relogX/restatic.svg)](LICENSE)

</div>

#  What is Restatic?

Restatic is a simple HTTP server that serves a local directory over HTTP. It is written in [Go](https://golang.org/), and its usage is inspired by [Python's](https://www.python.org/)  [http.server](https://docs.python.org/3/library/http.server.html) module.

##  Using Restatic

To use Restatic, download the latest release for your platform and fire the following commands.

 1. Download the latest [restatic/releases](https://github.com/relogX/restatic/releases) for your platform.
 2. Extract the binary from the just downloaded compressed artifact
 3. Execute the binary as shown in the following steps

```
$ ./restatic -p 5030 -d .

        ██████  ███████ ███████ ████████  █████  ████████ ██  ██████ 
        ██   ██ ██      ██         ██    ██   ██    ██    ██ ██      
        ██████  █████   ███████    ██    ███████    ██    ██ ██      
        ██   ██ ██           ██    ██    ██   ██    ██    ██ ██      
        ██   ██ ███████ ███████    ██    ██   ██    ██    ██  ██████ 

        by Relog - https://relog.in

INFO[0000] server listening on :5030
INFO[0000] =========================
```

 4. Browsing the URL on your favourite browser will load a webpage like this

![screen-01](https://user-images.githubusercontent.com/4745789/135251623-f8ea8024-75b7-4150-a869-26135212822d.PNG)

##  Developing Restatic

If you are a developer and want to modify restatic, you will have first to set up a dev environment, and it has the following pre-requisites

- [Go 1.17.1](https://golang.org/)

Once you have set up all the pre-requisites, following the steps to start your development server.

- Clone the repository https://github.com/relogX/restatic
- Start the server `go run cmd/restatic/main.go`

Once you start the server, it will download all the necessary packages and listen to the configured port. The default port is 5030.

##  Linting

Maintaining coding standards is extremely critical, and restatic follows the standard [Gofmt](https://pkg.go.dev/cmd/gofmt) to reformat the code. It is customary to fire the following command before you commit.

```
make lint
```

##  Contribution Guidelines

The Code Contribution Guidelines are published at [CONTRIBUTING.md](https://github.com/relogX/restatic/blob/master/CONTRIBUTING.md); please read them before you start making any changes. This would allow us to have a consistent standard of coding practices and developer experience.

##  The RelogX Umbrella
<div align="center">
<br />
<img  width="240"  src="https://user-images.githubusercontent.com/4745789/133601178-711aa4eb-f836-4e93-a554-22006648f75f.png" align="center"  alt="relog logo" />
<br />
<br />
</div>

[Relog](https://relog.in) is an initiative that aims to transform engineering education and provide high-quality engineering courses, projects, and resources to the community. To better understand all the common systems, we aim to build our own replicated utilities, for example, a load balancer, static file server, API rate limiter, etc. All the projects fall under [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), and you can find their source code at [github.com/relogX](https://github.com/relogX).

##  License
Restatic is under [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)
