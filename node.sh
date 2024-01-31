#!/bin/bash

KEYS=(
    valid0
    valid1
)

APP="exampled"
CHAINID="example"
MONIKER="localtestnet"
KEYRING="test"
LOGLEVEL="info"
HOMEDIR="$HOME/.example"
TRACE=""
IP="62.112.10.186"

CONFIG=$HOMEDIR/config/config.toml
APP_TOML=$HOMEDIR/config/app.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

# ignite chain build

rm -rf "$HOMEDIR"

VAL0_MNEMONIC="copper push brief egg scan entry inform record adjust fossil boss egg comic alien upon aspect dry avoid interest fury window hint race symptom"

VAL1_MNEMONIC="reform spell firm okay devote wife identify honey glass filter jump upset off situate cool ensure better square swallow canoe garlic pumpkin shock virus"

echo "$VAL0_MNEMONIC" | $APP keys add "${KEYS[0]}" --recover --keyring-backend "$KEYRING" --home "$HOMEDIR"
echo "$VAL1_MNEMONIC" | $APP keys add "${KEYS[1]}" --recover --keyring-backend "$KEYRING" --home "$HOMEDIR"

$APP init $MONIKER -o --chain-id $CHAINID --home "$HOMEDIR"

jq '.app_state["distribution"]["params"]["community_tax"]="0.0"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "2000000000stake"/g' "$APP_TOML"


for KEY in "${KEYS[@]}"; do
    $APP genesis add-genesis-account "$KEY" 1000000000000000000000000000stake --keyring-backend $KEYRING --home "$HOMEDIR"
done

$APP genesis gentx "${KEYS[0]}" 1000000000000000000000stake --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR" --commission-rate "0.25" --commission-max-rate "0.25"

$APP genesis collect-gentxs --home "$HOMEDIR"
$APP genesis validate --home "$HOMEDIR"

# $APP start --home "$HOMEDIR"