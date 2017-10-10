#!/usr/bin/env bash
export TLSCERT=/etc/letsencrypt/live/shaheensharifian.me/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/shaheensharifian.me/privkey.pem
docker rm -f zipsvr

docker run -d \
-p 443:443 \
--name zipsvr \
-v /Users/shaheensharifian/Documents/Info344/go/src/github.com/ssharif6/info344-in-class/zipsvr/tls:/tls:ro \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
ssharif6/zipsvr
