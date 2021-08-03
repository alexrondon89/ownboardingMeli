# ownboardingMeli

-API that show you current value about one crypto, or a crypto array for a specific currency like usd.

-The client used to get coin prices is coingecko API.

-This api is resilient, that mean it will never crash if one coin or currency not exist. When something get wrong with
 a coin, partial response is build with status code 206. If there are no errors then response build is successful with 
 status code 200

-To improve the time to response, this api make use of go routine over each coin price request

NOTE: always when you do a request to get bitcoin price, you will receive a partial response because there is an intentional
 panic execution to make use of Recover method() 

Endpoints:

Getting price for a specific coin with currency eur
http://localhost:8080/meli/coinprice?Coin=cardano&Currency=eur

Getting price for an array of coins with currency usd
http://localhost:8080/meli/listprice?Coins=cardano&Coins=bitcoin&Coins=ethereum&Currency=usd

