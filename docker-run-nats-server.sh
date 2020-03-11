#!/bin/sh

#docker run --name nats --rm --detach --publish=4222:4222 --publish=8222:8222 nats-streaming --auth "12345"

scriptname="$0"
scriptdir=$(cd "${scriptname%/*}" && pwd)

docker run -v "${scriptdir}:/ssl" --name nats --rm --detach \
    --publish=4222:4222 --publish=8222:8222 nats \
    --auth 'token123456789' --tls=true --tlscert=/ssl/ssl-cert-snakeoil.pem --tlskey=/ssl/ssl-cert-snakeoil.key

