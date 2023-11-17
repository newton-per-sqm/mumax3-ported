#! /bin/bash

set -e

bin/mumax3 -vet test/*.mx3

bin/mumax3 -paranoid=false -failfast -cache /tmp -f -http "" test/*.go test/*.mx3
