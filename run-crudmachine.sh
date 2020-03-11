#!/bin/sh

scriptname="$0"
scriptdir=$(cd "${scriptname%/*}" && pwd)

(
    cd crudmachine
    go build ./pkg/* && GOBIN="$(pwd)/bin" go install ./cmd/crudmachine && ./bin/crudmachine -config test.yml
)


