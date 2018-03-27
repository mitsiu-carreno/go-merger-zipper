# Go Merger Zipper

Helper to retreive a list of csv from mongo and merge then into a single file to later compress it into a zip file

## Deployment

For deployment, build the code into the target architecture 
```
$ GOOS=linux GOARCH=amd64 go build main.go
```

For running the build, first set the variables required (MAIN_DB_HOST, MAIN_DB_DB, MAIN_DB_USER, MAIN_DB_PASSWORD, MAIN_DB_COLLECTION...)
```
$ source .env.prod.unmillon
$ ./main
```