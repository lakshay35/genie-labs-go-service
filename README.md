# How would you design the system with elasticity and scalability in mind?

I would use a autoscaler system. Something like Kubernetes has support for this out of the box. Using something like HPA (Horizontal Pod Autoscaler) allows us to define maximum and minimum number of applications pods allowed and what threshold per pod to start scaling up or down at. AWS Autoscaling is another solution that can be used with AWS native services.

# Quick start

Setup .env file with the following
```
PORT=8000

DATABASE_URL=postgres://<username>:<password>@localhost:5432/genie?sslmode=disable

ETHEREUM_NODE_WSS="<ALCHEMY_OR_SOME_OTHER_NODE_WSS_URL"
```

run `go install` to install dependencies

run `make run` to start the service and watch logs for any transfer events recorded from the BAYC, MAYC contracts

# Approach

I set up the service to connect to an Ethereum Node Provider using web sockets over RPC. I listen for events that are emitted from the contract and parse them to see if an event is a "Transfer" event and call rarible APIs to update the quotes. I thought of using polling but wasn't sure if we were prioritizing absolute real time updates. Under the assumption that we were I used the event trigger approach. 

Additionally, I couldnt' find out how the continuation header value was supposed to be parsed and sent back so I left the code to make an api calls once per contract. However, in the typescript solution I submitted, I had the docs for it and made chained API calls to process each token in a collection
