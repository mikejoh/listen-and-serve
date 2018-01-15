# Listen and serve

## Explain!

Listen and serve is a repository where you'll find ~~a vast amount of~~ examples of HTTP-servers written in different programming languages (using ~~a vast amount of~~ different web frameworks). All of the examples provides a Dockerfile for you so that you easily can run them seperatly as containers or via the nginx proxy container. 

The current examples all does the same thing:
* Listen on a port
* Serve content over HTTP
* Expose an endpoint for serving content, ```/headers``` to start with

Hopefully you'll learn more about Docker and serving content over HTTP in different programming languages.

## Getting started

_Work in progress.._

## Used languages and frameworks

* Python (using Flask)
* Go
* Java (using Spark)
* Node

## Endpoints

* /go 
* /flask
* /spark
* /node

## Todo

* Add an easier way to run and start everything
* Provide an index of all examples and serve this via e.g. the proxy
* Write the Getting started section in this README
* Add Makefile for the nginx-proxy
* Docker tags, hello?
* Find other ways of serving clients via the HTTP-servers, to test more functionality of the  built-in libraries and third party frameworks of the languages
