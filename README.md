# iservice-demo
iservice demo - iservice daemon for service provider

## INSTALL IRIS
```bash
git clone https://github.com/irisnet/irishub.git
cd irishub
git checkout -b develop origin/develop
source scripts/setTestEnv.sh
make install
```

## BUILD
```bash
make install
```

## RUNNING

### start node
```bash
# init node
iriscli keys add acct0
iris init --chain-id test --moniker acct0
iris gentx --amount 100iris --name acct0
iris add-genesis-account $(iriscli keys show acct0 -o json | jq -r '.address') 2000000000iris
iris collect-gentxs

# add acct0 as profiler
sed -i '' "s/faa108w3x8/$(iriscli keys show acct0 -o json | jq -r '.address')/" ~/.iris/config/genesis.json

# start node
iris start
```

### Create key pair for service providers
```bash
# generate provider1 address
iservice keys add provider1

# add provider1 accout in iriscli (use the above mnemonic)
iriscli keys add provider1 --recover

# generate provider2 address
iservice keys add provider2

# add provider2 accout in iriscli (use the above mnemonic)
iriscli keys add provider2 --recover
```

### create service definition
```bash
# set service name
service_name=price_service

# define service 
iriscli service define --chain-id test --from acct0 --fee 0.3iris --name $service_name --description="provide token price" --tags=price --schemas=iservice/service/service_definition.json --commit
```

### create service binding
```bash
# send 1000000iris to provider address
iriscli bank send --from acct0 --to $(iriscli keys show provider1 -o json | jq -r '.address') --amount 1000000iris --chain-id test --fee 0.3iris --commit
iriscli bank send --from acct0 --to $(iriscli keys show provider2 -o json | jq -r '.address') --amount 1000000iris --chain-id test --fee 0.3iris --commit

# bind service
iriscli service bind --chain-id test --from provider1 --fee 0.3iris --service-name $service_name --deposit=10000iris --pricing iservice/service/service_pricing.json --commit
iriscli service bind --chain-id test --from provider2 --fee 0.3iris --service-name $service_name --deposit=10000iris --pricing iservice/service/service_pricing.json --commit

# qury bindings
iriscli service bindings $service_name
```

### start iservice daemon
```bash
iservice start provider1 huobi &
iservice start provider2 binance &
```

# create & start oracle feed
```bash
feed_name=link_usdt
iriscli oracle create --chain-id test --from acct0 --fee 0.3iris  --feed-name $feed_name --latest-history 10 --service-name $service_name --input "{\"base\":\"link\",\"quote\":\"usdt\"}" --providers $(iriscli keys show provider1 -o json | jq -r '.address'),$(iriscli keys show provider2 -o json | jq -r '.address') --threshold 2 --service-fee-cap 1iris --timeout 2 --frequency 5 --aggregate-func "avg" --value-json-path "price" --commit

# query feed
iriscli oracle query-feed $feed_name

# start feed  
iriscli oracle start $feed_name --chain-id test --from acct0 --fee 0.3iris --commit

# query feed (state to running)
iriscli oracle query-feed $feed_name
```

### query feed value
```bash
# query feed value
iriscli oracle query-value $feed_name
```
