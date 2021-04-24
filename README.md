# Let's Go

[![Go Report Card](https://goreportcard.com/badge/github.com/aplescia-chwy/lets-go)](https://goreportcard.com/report/github.com/aplescia-chwy/lets-go)
[![codecov](https://codecov.io/gh/aplescia-chwy/lets-go/branch/master/graph/badge.svg)](https://codecov.io/gh/aplescia-chwy/lets-go)

[GoDoc](https://pkg.go.dev/github.com/aplescia-chwy/lets-go)

A module with some useful packages for anyone developing Golang applications on the cloud. 
---
Packages include:

* **util** -- a generic utility package
* **rest** -- useful REST utilities. Primary deals with unmarshalling and marshalling JSON, as well as 
support for serializing/deserializing structs according to the JSON:API schema.
* **redis** -- a package that provides utilities relevant to connecting with and interacting with a 
Redis cluster. This is a common use case for AWS Elasticache.
* **sql** -- a package that provides utilities relevant to connecting to and 
interacting with a Postgres database. Makes use of the GORM library found [here](https://github.com/jinzhu/gorm).
* **alb** -- a package that helps format responses to be sent from an AWS Lambda function back to a
 triggering ALB, and
therefore back to a client.
* **kinesis** -- AWS Kinesis utility functions (putting payloads to a kinesis stream).

## How do I use it?

Add the following to your go.mod file (or directly import it):

```shell script
github.com/aplescia/lets-go
```
