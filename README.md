<img src="https://docs.entangle.fi/~gitbook/image?url=https%3A%2F%2F4040807501-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F5AajewgFWO9EkufRORqL%252Fuploads%252FDfRGGJJASR0PFitX6rbx%252FTwitter%2520%281%29.png%3Falt%3Dmedia%26token%3D09e49fe6-1b92-4bed-82e6-730ba785afaf&width=1248&dpr=2&quality=100&sign=5fbbb9f4&sv=1" alt="Entangle" style="width:100%;"/><div align="center">

# pull-update-publisher


## Run pull update publisher

```
export ETNANGLE_PUBLISHER=1TYXaYBYNuij5eHv4hrovXyMLxy7ybKPqjEYy8J8TWjAByGyxTiAwa9cvCweLVGUGbBUbnucYaE6MXYJDELdCk4
export ENT_SOLANA_PUBLISHER_CONFIG=$(pwd)/config_solana.yml
export RUST_LOG=info,price_publisher=debug,solana_tools=debug
export LD_LIBRARY_PATH="."
go run cmd/pull-update-publisher/main.go
```

## Example of getting last update and getting the accounts data (solana)

When latest update is applied one can see the following logs:


```
[2024-09-03T17:54:28Z DEBUG price_publisher::data_feed_processor] data_feed_msg: { merkle_root: bc7e102e02da9f4ad7dd7adedb42169f7229c21169938f666bb50d5f522bbc34, data_feeds: [{ timestamp: 1725386002, data_key: GAS-ETH, data: 00000000000000000000000000000000000000000000000000000322a94f8d8c, merkle_proof: [67a4bef56395552e5b16c08e9b92218f16e21c16156d5a7b931ef6fc6b26dbe8] }], signatures: [{ 1b1db1743d3a85580a84bb2849a51bbb333660fadd945e7e744e40987b8db6e55154e8b9fa2d3b941cf6a761e7aee160fccc829c24f33f5ce25d7151d8d585b667 }, { 1cc2044555c510f990a90b159ff7e58cd6cbf3e2269b5505c9a7096e3dc6b353507c4fd246ce26690d5eff5a1925c5998d02328d1fb228aa8387aad0d35ca3333c }, { 1c614c4a82d4c6406e422af9567e5b3bcfbf5340d4b7133322aeb3af257496f70c5b84d406b915b3419ec093c9945f17d722447d0aefede604ad22617e3178702e }] }
[2024-09-03T17:54:29Z DEBUG price_publisher::data_feed_processor] pda: [GAS-ETH: 7xKpRLTnN7nhtbg5bPRLhi3d73KnsiJRkoRUwbKdogFV]
[2024-09-03T17:54:29Z DEBUG solana_tools::solana_transactor::ix_compiler] Instructions: 0 Tx len: 667 CU: 400000
[2024-09-03T17:54:37Z DEBUG solana_tools::solana_transactor::transactor] Sent bundle 6646c7e0-9250-4921-b751-5dfddfda74c4 with sig 3CeRixVfcvetiGfAZqnWW71Ee7YCLJyNzsweQKkXY4LCR9PMHcQL6PMCSQMQViGizpQMDFScPYQNnr5a34Ji1EQD, awaiting 1
...
```

It's possible to use the asset related Program Derived Address (PDA) to check if it's really updated on the devnet for testing purposes

```
anchor account --provider.cluster devnet --idl udf_solana.json udf_solana.LatestUpdate 7xKpRLTnN7nhtbg5bPRLhi3d73KnsiJRkoRUwbKdogFV | jq -c
{
    "data":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,3,34,169,79,141,140],
    "dataKey":[71,65,83,45,69,84,72,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],
    "dataTimestamp":1725386002}
}
```
