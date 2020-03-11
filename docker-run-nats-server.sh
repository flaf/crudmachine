#!/bin/sh

#docker run --name nats --rm --detach --publish=4222:4222 --publish=8222:8222 nats-streaming --auth "12345"

scriptname="$0"
scriptdir=$(cd "${scriptname%/*}" && pwd)

docker stop $(docker ps -qa) 2>/dev/null || true

docker run -v "${scriptdir}:/ssl" --name nats --rm --detach \
    --publish=4222:4222 --publish=8222:8222 nats \
    --auth 'token123456789' \
    --tls=true --tlscert=/ssl/wildcard.ac-versailles.fr.crt --tlskey=/ssl/wildcard.ac-versailles.fr.key \
    #--tlscacert=/ssl/wildcard.ac-versailles.fr.intermediate.crt
    #--tls=true --tlscert=/ssl/ssl-cert-snakeoil.pem --tlskey=/ssl/ssl-cert-snakeoil.key


