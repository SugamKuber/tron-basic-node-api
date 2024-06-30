# API server setup

- Git clone the repo
```
git clone https://github.com/SugamKuber/tron-basic-node-api
cd tron-api-server
```

- Make sure golang is installed, If not install golang
```
go version        
```

- Add the ENV variables needed

```
TRON_API_KEY=
LOCAL_NODE_API_URL=
```

- To get LOCAL_NODE_API_URL after you start the local node (follow [this doc](https://github.com/SugamKuber/tron-basic-node-api/blob/main/tron-java-node/README.md)) use your localhost here

- To get TRON_API_KEY on [tronscan dev api](https://tronscan.org/#/developer/api)
I am using this to get txns and full details on balance of user, You can also build a indexer from scratch insted of using this api by querying with the local node setup directly  


- Install the pkg & start the server
```
go mod tidy
go run cmd/main.go
```

- Check if the server is running
```
curl http://localhost:8080/h
```