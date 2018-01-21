FROM alpine:latest
RUN apk update && apk add --no-cache \
    alpine-sdk \
    cabal \
    curl \
    ghc
RUN curl -sSL https://get.haskellstack.org/ | sh
RUN cabal update && cabal install Spock
RUN stack new httpserver
WORKDIR /httpserver/app
COPY Main.hs .
RUN stack build --fast --pedantic
#
#  https://github.com/commercialhaskell/stack/issues/2387
#
