# Faucetchain

## Faucet module

- cosmos17s95c5jpc6x2l3edwh4dm8yhac68yru7cre47d

## Init accounts

```
rm -rf ~/.faucetchaind
./bin/faucetchaind init faucetchain --chain-id faucetchain
./bin/faucetchaind keys add validator --keyring-backend test
./bin/faucetchaind keys add faucet-account --keyring-backend test

./bin/faucetchaind genesis add-genesis-account $(./bin/faucetchaind keys show validator -a --keyring-backend test) 1000000000stake
./bin/faucetchaind genesis add-genesis-account $(./bin/faucetchaind keys show faucet-account -a --keyring-backend test) 1000000000stake

./bin/faucetchaind genesis gentx validator 500000000stake --chain-id faucetchain --keyring-backend test
./bin/faucetchaind genesis collect-gentxs
./bin/faucetchaind genesis validate-genesis

# Edit app.toml explicitly to set minimum-gas-prices
# [base]
# minimum-gas-prices = "0.01stake"

./bin/faucetchaind start

```