curl --request POST \
    --url http://localhost:1323/wallets \
    --header 'content-type: application/json' \
    --data '{"name": "Wallet-1"}'

sleep 1

curl --request POST \
    --url http://localhost:1323/wallets \
    --header 'content-type: application/json' \
    --data '{"name": "Wallet-2"}'

curl --request POST \
    --url http://localhost:1323/wallets \
    --header 'content-type: application/json' \
    --data '{"name": "Wallet-3"}'

# create an existing wallet
curl --request POST \
    --url http://localhost:1323/wallets \
    --header 'content-type: application/json' \
    --data '{"name": "Wallet-1"}'

# create a wallet with unsupported fields
curl --request POST \
    --url http://localhost:1323/wallets \
    --header 'content-type: application/json' \
    --data '{"name": "Wallet-x", "capacity" : 122}'

curl --request GET \
    --url http://localhost:1323/wallets \
    --header 'content-type: application/json'

sleep 1

curl --request PUT \
    --url http://localhost:1323/wallets/Wallet-1\
    --header 'content-type: application/json'\
    --data '{"name": "Wallet-1-new"}'

# edit an unknown wallet
curl --request PUT \
    --url http://localhost:1323/wallets/Wallet-unknown\
    --header 'content-type: application/json'\
    --data '{"name": "new-name"}'

curl --request DELETE \
    --url http://localhost:1323/wallets/Wallet-3\
    --header 'content-type: application/json'

curl --request GET \
    --url http://localhost:1323/wallets \
    --header 'content-type: application/json'

#----------------------------------------------
curl --request POST \
    --url http://localhost:1323/Wallet-2/coins \
    --header 'content-type: application/json' \
    --data '{"name": "Go Coin", "symbol": "GOC", "amount": 7.1, "rate": 21.4}'

curl --request POST \
    --url http://localhost:1323/Wallet-2/coins \
    --header 'content-type: application/json' \
    --data '{"name": "Bitcoin", "symbol": "BIT", "amount": 4.1, "rate": 15.4}'

curl --request POST \
    --url http://localhost:1323/Wallet-2/coins \
    --header 'content-type: application/json' \
    --data '{"name": "Mammad Coin", "symbol": "MMD", "amount": 12.1, "rate": 30}'

# Unknown fields in coin JSON
curl --request POST \
    --url http://localhost:1323/Wallet-2/coins \
    --header 'content-type: application/json' \
    --data '{"salam": "khoobi", "symbol": "ssssymbol"}'

curl --request GET \
    --url http://localhost:1323/wallets \
    --header 'content-type: application/json'


curl --request GET \
    --url http://localhost:1323/Wallet-2 \
    --header 'content-type: application/json'

# get an unknown wallet info
curl --request GET \
    --url http://localhost:1323/Wallet-unknown \
    --header 'content-type: application/json'

sleep 1
curl --request PUT \
    --url http://localhost:1323/Wallet-2/GOC \
    --header 'content-type: application/json' \
    --data '{"name": "Go Coin", "symbol": "GOCCC", "amount": 1, "rate": 20}'

curl --request GET \
    --url http://localhost:1323/Wallet-2 \
    --header 'content-type: application/json'

curl --request DELETE \
    --url http://localhost:1323/Wallet-2/MMD \
    --header 'content-type: application/json'

curl --request GET \
    --url http://localhost:1323/Wallet-2 \
    --header 'content-type: application/json'

curl --request GET \
    --url http://localhost:1323/wallets \
    --header 'content-type: application/json'

curl --request DELETE \
    --url http://localhost:1323/wallets/Wallet-1-new\
    --header 'content-type: application/json'

curl --request GET \
    --url http://localhost:1323/wallets \
    --header 'content-type: application/json'

# Bad endpoint example
curl --request GET \
    --url http://localhost:1323/salam/khoobi \
    --header 'content-type: application/json'