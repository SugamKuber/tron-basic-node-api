# LOCAL TRON JAVA NODE SETUP

- Make sure you are using Oracle JDK 1.8 ONLY
use the below commands to verify
```
>>java -version

java version "1.8.0_401"
Java(TM) SE Runtime Environment (build 1.8.0_401-b10)
```

- Use the latest version of JAVA-TRON Build with [FullNode.jar
](https://github.com/tronprotocol/java-tron/releases/tag/GreatVoyage-v4.7.5)
- You can also copy the repo and build with ```gradlew```
- Use the config file given or create from [here](https://github.com/tronprotocol/tron-deployment)

- Run the Jar file with the below command
```
java -Xmx24g -XX:+UseConcMarkSweepGC -jar FullNode.jar -c config.conf
```
- Check the logs with
```
tail -f logs/tron.log                                                
```
- Run a sample curl to check the node 
```
curl http://localhost:8090/wallet/getnowblock
```

It should look something like this
![image](https://github.com/SugamKuber/tron-basic-node-api/assets/86460476/8ce2dcc5-55ea-45ef-aa8a-273340ba3b07)
![image](https://github.com/SugamKuber/tron-basic-node-api/assets/86460476/4288eb96-174c-4ece-9f4a-691c1b174588)
![image](https://github.com/SugamKuber/tron-basic-node-api/assets/86460476/f8665570-fe30-411f-bb65-1e53289c6983)



