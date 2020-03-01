# iservice-demo
iservice demo - iservice daemon for service provider

## INSTALL IRIS
```bash
git clone https://github.com/irisnet/irishub.git
cd irishub
git checkout -b develop origin/develop
make install
```

## BUILD
```bash
make install
```

## RUNNING

### generate a new key
```bash
iservice keys add iservice
```
* copy the mnemonic and address
```bash
provider_addr=<your_address>
```


### start node
```bash
# init node
iris testnet --v 1 --chain-id test

# modify token amount
sed -i '_bak' 's/150000000000000000000iris-atto/150000000000000000000000000iris-atto/' mytestnet/node0/iris/config/genesis.json

# start node
iris start --home mytestnet/node0/iris
```

### define % bind service
```bash
# set service name
service_name=price_service

# recover your address (iservice key)
iriscli keys add iservice --recover --home mytestnet/node0/iriscli

# send 1000000iris to iservice address
iriscli bank send --from node0 --to $provider_addr --amount 1000000iris --chain-id test --fee 0.3iris --commit --home mytestnet/node0/iriscli/

# define service 
iriscli service define --chain-id test --from iservice --fee 0.3iris --name $service_name --description="provide token price" --tags=price --schemas=iservice/service/service_definition.json --commit --home mytestnet/node0/iriscli/

# bind service
iriscli service bind --chain-id test --from iservice --fee 0.3iris --service-name $service_name --deposit=10000iris --pricing iservice/service/service_pricing.json --commit --home mytestnet/node0/iriscli/

# qury binding
iriscli service binding $service_name $provider_addr

# create feed
feed_name=price
iriscli oracle create --chain-id test --from node0 --fee 0.3iris  --feed-name $feed_name --latest-history 10 --service-name $service_name --input "{\"denom\":\"iris\"}" --providers $provider_addr --threshold 1 --service-fee-cap 1iris --timeout 2 --frequency 5 --total -1  --aggregate-func "avg" --value-json-path "price" --commit --home mytestnet/node0/iriscli/

# query feed
iriscli oracle query-feed $feed_name

# start feed  
iriscli oracle start $feed_name --chain-id test --from node0 --fee 0.3iris --commit --home mytestnet/node0/iriscli/

# query feed (state to running)
iriscli oracle query-feed $feed_name
```

### start iservice demo
```bash
iservice start iservice
```

### query feed value
```bash
# query feed value
iriscli oracle query-value $feed_name
```
