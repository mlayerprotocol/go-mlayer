nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run . daemon --network-private-key $1
