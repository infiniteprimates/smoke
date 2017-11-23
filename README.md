# smoke [![License](http://img.shields.io/badge/license-mit-blue.svg)](LICENSE) [![wercker status](https://app.wercker.com/status/cf2b022fe1d667a6e16709a43d7e38ac/s/master "wercker status")](https://app.wercker.com/project/byKey/cf2b022fe1d667a6e16709a43d7e38ac) [![codecov](https://codecov.io/gh/infiniteprimates/smoke/branch/master/graph/badge.svg)](https://codecov.io/gh/infiniteprimates/smoke) [![Go Report Card](https://goreportcard.com/badge/github.com/infiniteprimates/smoke)](https://goreportcard.com/report/github.com/infiniteprimates/smoke) [![Dependency Status](https://dependencyci.com/github/infiniteprimates/smoke/badge)](https://dependencyci.com/github/infiniteprimates/smoke) [![godoc](https://godoc.org/github.com/infiniteprimates/smoke?status.svg)](https://godoc.org/github.com/infiniteprimates/smoke)

A small web based messaging service especially designed for containers.

## Build

It is assumed that you have a fully functional Go environment configured (Go runtime, $GOPATH, etc.).

### Requirements

You will need the following software in order to build Smoke.

* [Go >= 1.8](https://golang.org)
* [Go dep >= 0.3.2](https://github.com/golang/dep) 
* [Wercker CLI](https://www.wercker.com/cli) - This is only required if you are interested in testing/troubleshooting the Wercker build.
* [NodeJS >= 6.11](https://nodejs.org) - Consider using [NVM](https://github.com/creationix/nvm) to install and manage NodeJS.
* [Angular CLI >= 1.4.9](https://cli.angular.io/)

### Steps

To build the smoke binary:

1. Clone this repository to $GOPATH/src/github.com/infiniteprimates/smoke
1. Change directories into $GOPATH/src/github.com/infiniteprimates/smoke
1. Install go dependencies using the dep tool. Run 'dep ensure'.
1. Build the smoke binary. Run 'go build'.

To build the Angular UI:

1. Change directories into $GOPATH/src/github.com/infiniteprimates/smoke/ui-src
1. Install javascript dependencies. Run 'npm install'.
1. Build the Angular UI. Run 'ng build'. The built UI is located in $GOPATH/src/github.com/infiniteprimates/smoke/ui.
