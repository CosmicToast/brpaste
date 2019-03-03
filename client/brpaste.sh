#!/bin/sh
brpaste() {
    host='https://brpaste.xyz'
    out=$(curl -sF 'data=<-' $host || printf fail)
    [ "$out" = fail ] && echo fail || printf '%s/%s\n' "$host" "$out"
}

brpaste
