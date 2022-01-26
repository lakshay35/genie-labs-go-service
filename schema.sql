CREATE TABLE IF NOT EXISTS assets (
  id TEXT PRIMARY KEY,
  token_id TEXT NOT NULL,
  collection_address TEXT NOT NULL,
  collection_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS asset_price (
  id TEXT PRIMARY KEY,
  status TEXT NOT NULL,
  current_offer_id TEXT,
  curr_price_usd FLOAT,
  curr_price_jpy FLOAT,
  curr_price_eth FLOAT
);



/**

API Secret: y1tnfvktFn3Quv0KvWvg0xPqk/GExJ1u9roGoUUmHw3nQdtWEYl5zLtfpdjZ2Oxm8irxj9JwbT3HtkceJpbRVw==
Passphrase: ujt4bls0up
**/