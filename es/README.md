# es library
AWS ElasticSearch wrapper to create new client using signer v4

This repository is experimental and used by internal team only.

## Generating clients
Client constructor requires 2 arguments:
1. ES URL
2. Region

This is how you use it

`client, err := es.New(os.Getenv("ES_URL"), os.Getenv("ES_REGION"))`