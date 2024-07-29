#!/bin/bash

go run . daemon \
--private-key a12b76e5ab05842dd110b9cf709a3ca4e965ec149a1e6d40f09f55a4fc6279d2  \
--rpc-port 9090 \
--ws-address localhost:9091 \
--rest-address localhost:9092   \
--data-dir "data/dev2/" \
-l "/ip4/127.0.0.1/tcp/6000/ws" \
-l "/ip4/127.0.0.1/tcp/6001"
    

 