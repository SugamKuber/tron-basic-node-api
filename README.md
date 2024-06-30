# Tron Java Node


### Implementation
- Refer this [doc](https://github.com/SugamKuber/tron-basic-node-api/blob/main/tron-java-node/README.md) to setup Java Tron on your linux
- Refer this [doc](https://github.com/SugamKuber/tron-basic-node-api/blob/main/tron-api-server/README.md) to setup the API server and interact with Tron Node

### API's

After the above implementation, Use the basic APIs below

This server uses LOCAL TRON NODE to check the validity of address, To get the balance & transcation tronscan API's are used, You can implement the local indexer on golang by querying with LOCAL TRON NODE directly

- Get User Balance API
```
curl http://localhost:8080/b/:address
```
- Get User Transcations
```
curl http://localhost:8080/txn/:address
```
  - Add limit & start if needed (default limit=10 & start = 0)
```
curl http://localhost:8080/txn/:address?limit=:number&start=:number
```