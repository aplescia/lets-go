# Let's Go

A module with some useful packages for anyone developing Golang applications. 
---
Packages include:

* **util** -- a generic utility package
* **redis** -- a package that provides utilities relevant to connecting with and interacting with a 
Redis cluster.
* **sql** -- a package that provides utilities relevant to connecting to and 
interacting with a Postgres database.
* **alb** -- a package that helps format responses to be sent from an AWS Lambda function back to a
 triggering ALB, and
therefore back to a client.
* **kinesis** -- AWS Kinesis utility functions (putting payloads to a kinesis stream).

## How do I use it?

```shell script
go env -w GOPRIVATE=github.com/Chewy-Inc/lets-go
```

then add the following to your go.mod file (or directly import it):

```shell script
github.com/Chewy-Inc/lets-go
```
 
 If you see:
 ```shell script
fatal: could not read Username for 'https://github.com': terminal prompts disabled
```

try running the following:

```shell script
git config --global url."git@github.com:".insteadOf "https://github.com"
```