# example
**example** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/example@latest! | sudo bash
```
`username/example` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

### Start 2 Node

## config.toml
```
persistent_peers = "592f96a5c6f45413e6655a49015b21fd67ee8158@62.112.10.250:26656,4336e8b0eee818009b2b98e5a09aa43644efc369@62.112.10.186:26656"
```

## genesis.json
```json
{
  "app_name": "exampled",
  "app_version": "",
  "genesis_time": "2024-02-04T19:51:23.651868156Z",
  "chain_id": "example",
  "initial_height": 1,
  "app_hash": null,
  "app_state": {
    "06-solomachine": null,
    "07-tendermint": null,
    "auth": {
      "params": {
        "max_memo_characters": "256",
        "tx_sig_limit": "7",
        "tx_size_cost_per_byte": "10",
        "sig_verify_cost_ed25519": "590",
        "sig_verify_cost_secp256k1": "1000"
      },
      "accounts": [
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "cosmos1y6xz2ggfc0pcsmyjlekh0j9pxh6hk87yfwcjct",
          "pub_key": null,
          "account_number": "0",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "cosmos1t20dq93dy7qw70wgkgjprvt7tqjke907qyn359",
          "pub_key": null,
          "account_number": "1",
          "sequence": "0"
        }
      ]
    },
    "authz": {
      "authorization": []
    },
    "bank": {
      "params": {
        "send_enabled": [],
        "default_send_enabled": true
      },
      "balances": [
        {
          "address": "cosmos1y6xz2ggfc0pcsmyjlekh0j9pxh6hk87yfwcjct",
          "coins": [
            {
              "denom": "stake",
              "amount": "1000000000000000000000000000"
            }
          ]
        },
        {
          "address": "cosmos1t20dq93dy7qw70wgkgjprvt7tqjke907qyn359",
          "coins": [
            {
              "denom": "stake",
              "amount": "1000000000000000000000000000"
            }
          ]
        }
      ],
      "supply": [
        {
          "denom": "stake",
          "amount": "2000000000000000000000000000"
        }
      ],
      "denom_metadata": [],
      "send_enabled": []
    },
    "capability": {
      "index": "1",
      "owners": []
    },
    "circuit": {
      "account_permissions": [],
      "disabled_type_urls": []
    },
    "consensus": null,
    "crisis": {
      "constant_fee": {
        "denom": "stake",
        "amount": "1000"
      }
    },
    "distribution": {
      "params": {
        "community_tax": "0.0",
        "base_proposer_reward": "0.000000000000000000",
        "bonus_proposer_reward": "0.000000000000000000",
        "withdraw_addr_enabled": true
      },
      "fee_pool": {
        "community_pool": []
      },
      "delegator_withdraw_infos": [],
      "previous_proposer": "",
      "outstanding_rewards": [],
      "validator_accumulated_commissions": [],
      "validator_historical_rewards": [],
      "validator_current_rewards": [],
      "delegator_starting_infos": [],
      "validator_slash_events": []
    },
    "evidence": {
      "evidence": []
    },
    "example": {
      "params": {}
    },
    "feegrant": {
      "allowances": []
    },
    "feeibc": {
      "identified_fees": [],
      "fee_enabled_channels": [],
      "registered_payees": [],
      "registered_counterparty_payees": [],
      "forward_relayers": []
    },
    "genutil": {
      "gen_txs": [
        {
          "body": {
            "messages": [
              {
                "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
                "description": {
                  "moniker": "localtestnet",
                  "identity": "",
                  "website": "",
                  "security_contact": "",
                  "details": ""
                },
                "commission": {
                  "rate": "0.250000000000000000",
                  "max_rate": "0.250000000000000000",
                  "max_change_rate": "0.010000000000000000"
                },
                "min_self_delegation": "1",
                "delegator_address": "",
                "validator_address": "cosmosvaloper1y6xz2ggfc0pcsmyjlekh0j9pxh6hk87yv6v85c",
                "pubkey": {
                  "@type": "/cosmos.crypto.ed25519.PubKey",
                  "key": "IP/gOBw/gXRNiT5ItH4vayEjhXWQ/Rl9t8APeVZiqZI="
                },
                "value": {
                  "denom": "stake",
                  "amount": "1000000000000000000000"
                }
              }
            ],
            "memo": "4336e8b0eee818009b2b98e5a09aa43644efc369@62.112.10.186:26656",
            "timeout_height": "0",
            "extension_options": [],
            "non_critical_extension_options": []
          },
          "auth_info": {
            "signer_infos": [
              {
                "public_key": {
                  "@type": "/cosmos.crypto.secp256k1.PubKey",
                  "key": "A92x5YqLnNhrBrkm5qJ+EhzAPj9qiYF4qjjqh4A/MewT"
                },
                "mode_info": {
                  "single": {
                    "mode": "SIGN_MODE_DIRECT"
                  }
                },
                "sequence": "0"
              }
            ],
            "fee": {
              "amount": [],
              "gas_limit": "200000",
              "payer": "",
              "granter": ""
            },
            "tip": null
          },
          "signatures": [
            "cQN9Yyc/wbEakw/62DUxcZegJWtP3AQC0akQxrPZ58hdgtvKsxUFypkOFVLLM/d/59VqH5Byuoea113gJzyicA=="
          ]
        },
        {
          "body": {
            "messages": [
              {
                "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
                "description": {
                  "moniker": "localtestnet",
                  "identity": "",
                  "website": "",
                  "security_contact": "",
                  "details": ""
                },
                "commission": {
                  "rate": "0.250000000000000000",
                  "max_rate": "0.250000000000000000",
                  "max_change_rate": "0.010000000000000000"
                },
                "min_self_delegation": "1",
                "delegator_address": "",
                "validator_address": "cosmosvaloper1t20dq93dy7qw70wgkgjprvt7tqjke9079s8yck",
                "pubkey": {
                  "@type": "/cosmos.crypto.ed25519.PubKey",
                  "key": "0Q9M6OJR1Pyp9P3yyoWm4nKeLYikPnIgHA2eA4JJXbY="
                },
                "value": {
                  "denom": "stake",
                  "amount": "1000000000000000000000"
                }
              }
            ],
            "memo": "592f96a5c6f45413e6655a49015b21fd67ee8158@62.112.10.250:26656",
            "timeout_height": "0",
            "extension_options": [],
            "non_critical_extension_options": []
          },
          "auth_info": {
            "signer_infos": [
              {
                "public_key": {
                  "@type": "/cosmos.crypto.secp256k1.PubKey",
                  "key": "A6gd6cxhHvWXSYN7xMdNKUylLjCTwKVW/Ciq1DndOPax"
                },
                "mode_info": {
                  "single": {
                    "mode": "SIGN_MODE_DIRECT"
                  }
                },
                "sequence": "0"
              }
            ],
            "fee": {
              "amount": [],
              "gas_limit": "200000",
              "payer": "",
              "granter": ""
            },
            "tip": null
          },
          "signatures": [
            "hEkxnaqh9ZIf8RD2rVcTTXD8ubEspWnRWzm+WT/6pGUr8IkaEUf+urlRoHMcI+nDrza3Hg/cFtbmUUtMjeJJ/A=="
          ]
        }
      ]
    },
    "gov": {
      "starting_proposal_id": "1",
      "deposits": [],
      "votes": [],
      "proposals": [],
      "deposit_params": null,
      "voting_params": null,
      "tally_params": null,
      "params": {
        "min_deposit": [
          {
            "denom": "stake",
            "amount": "10000000"
          }
        ],
        "max_deposit_period": "172800s",
        "voting_period": "172800s",
        "quorum": "0.334000000000000000",
        "threshold": "0.500000000000000000",
        "veto_threshold": "0.334000000000000000",
        "min_initial_deposit_ratio": "0.000000000000000000",
        "proposal_cancel_ratio": "0.500000000000000000",
        "proposal_cancel_dest": "",
        "expedited_voting_period": "86400s",
        "expedited_threshold": "0.667000000000000000",
        "expedited_min_deposit": [
          {
            "denom": "stake",
            "amount": "50000000"
          }
        ],
        "burn_vote_quorum": false,
        "burn_proposal_deposit_prevote": false,
        "burn_vote_veto": true,
        "min_deposit_ratio": "0.010000000000000000"
      },
      "constitution": ""
    },
    "group": {
      "group_seq": "0",
      "groups": [],
      "group_members": [],
      "group_policy_seq": "0",
      "group_policies": [],
      "proposal_seq": "0",
      "proposals": [],
      "votes": []
    },
    "ibc": {
      "client_genesis": {
        "clients": [],
        "clients_consensus": [],
        "clients_metadata": [],
        "params": {
          "allowed_clients": [
            "06-solomachine",
            "07-tendermint",
            "09-localhost"
          ]
        },
        "create_localhost": false,
        "next_client_sequence": "0"
      },
      "connection_genesis": {
        "connections": [],
        "client_connection_paths": [],
        "next_connection_sequence": "0",
        "params": {
          "max_expected_time_per_block": "30000000000"
        }
      },
      "channel_genesis": {
        "channels": [],
        "acknowledgements": [],
        "commitments": [],
        "receipts": [],
        "send_sequences": [],
        "recv_sequences": [],
        "ack_sequences": [],
        "next_channel_sequence": "0"
      }
    },
    "interchainaccounts": {
      "controller_genesis_state": {
        "active_channels": [],
        "interchain_accounts": [],
        "ports": [],
        "params": {
          "controller_enabled": true
        }
      },
      "host_genesis_state": {
        "active_channels": [],
        "interchain_accounts": [],
        "port": "icahost",
        "params": {
          "host_enabled": true,
          "allow_messages": [
            "*"
          ]
        }
      }
    },
    "mint": {
      "minter": {
        "inflation": "0.130000000000000000",
        "annual_provisions": "0.000000000000000000"
      },
      "params": {
        "mint_denom": "stake",
        "inflation_rate_change": "0.130000000000000000",
        "inflation_max": "0.200000000000000000",
        "inflation_min": "0.070000000000000000",
        "goal_bonded": "0.670000000000000000",
        "blocks_per_year": "6311520"
      }
    },
    "mydist": {
      "params": {}
    },
    "params": null,
    "runtime": null,
    "slashing": {
      "params": {
        "signed_blocks_window": "100",
        "min_signed_per_window": "0.500000000000000000",
        "downtime_jail_duration": "600s",
        "slash_fraction_double_sign": "0.050000000000000000",
        "slash_fraction_downtime": "0.010000000000000000"
      },
      "signing_infos": [],
      "missed_blocks": []
    },
    "staking": {
      "params": {
        "unbonding_time": "1814400s",
        "max_validators": 100,
        "max_entries": 7,
        "historical_entries": 10000,
        "bond_denom": "stake",
        "min_commission_rate": "0.000000000000000000"
      },
      "last_total_power": "0",
      "last_validator_powers": [],
      "validators": [],
      "delegations": [],
      "unbonding_delegations": [],
      "redelegations": [],
      "exported": false
    },
    "transfer": {
      "port_id": "transfer",
      "denom_traces": [],
      "params": {
        "send_enabled": true,
        "receive_enabled": true
      },
      "total_escrowed": []
    },
    "upgrade": {},
    "vesting": {}
  },
  "consensus": {
    "params": {
      "block": {
        "max_bytes": "22020096",
        "max_gas": "-1"
      },
      "evidence": {
        "max_age_num_blocks": "100000",
        "max_age_duration": "172800000000000",
        "max_bytes": "1048576"
      },
      "validator": {
        "pub_key_types": [
          "ed25519"
        ]
      },
      "version": {
        "app": "0"
      },
      "abci": {
        "vote_extensions_enable_height": "0"
      }
    }
  }
}
```

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
