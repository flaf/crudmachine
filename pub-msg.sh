#!/bin/sh

scriptname="$0"
scriptdir=$(cd "${scriptname%/*}" && pwd)

for i in $(seq 1 20)
do
    go run nats-pub.go --token "token123456789" \
        -s "nats.ac-versailles.fr" toto '{"id":'$i', "action": "read", "data": {"toto": 12345, "tutu": {"tata": "zzz", "foo": "bar"} } }'
done

