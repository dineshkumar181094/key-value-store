# key-value-store
Two parts of this projects

*`apiserver`* - this server will store the key-value in ram no persistent storage is used.
[Swagger Specification](swagger.yaml)

*`kvstore (cli)`* - this is cli to interact with this server.



## How to run locally
## Using docker-compose
```sh
docker-compose up
#in new terminal
docker exec -it key-value-store_server_1 sh
# now you can run
# to set the key
./kvstore set home true
./kvstore get home
./kvstore update home false
./kvstore get home
```

## How to build locally
### Using Makefile
```sh
#Prerequisite
#golang version 1.15.x
make build
# the above file will create binary in bin directory
cd bin
./apiserver &
# this command will run api server in backend

# some testing commands
./kvstore set home true
./kvstore get home
./kvstore update home false
./kvstore get home
```

## CI/CD Intergration

- lint check on master merge.
- Test flow trigger on master merge.
- New relase is created on any Tag push.