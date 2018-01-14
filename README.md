# The HTTP-servers

## What is this?

Here you'll find a vast amount of examples of HTTP-servers written in different programming languages (using a vast amount of different frameworks).

All of the HTTP-servers behaves in the same way:
* Each HTTP-server serves an endpoint ```/headers```, and when you browse to that endpoint the headers of your request will be returned

To make it slightly easier for you i've provided what's needed to run the different examples in Docker containers. To reach them there's a proxy Docker container (nginx).

Please contribute! :-)

### Used languages and frameworks

* Python (with Flask)
* Go
* Java (with Spark)

## Getting started

_Work in progress.._

### Endpoints

* http://localhost/go
* http://localhost/flask
* http://localhost/

## Todo

* Add an easier way to run and start everything
* Provide an index of all examples and serve this via e.g. the proxy
* Write the Getting started section in this README
* Add Makefile for the nginx-proxy
* Docker tags, hello?
* Find other ways of serving clients via the HTTP-servers, to test more functionality of the  built-in libraries and third party frameworks of the languages
