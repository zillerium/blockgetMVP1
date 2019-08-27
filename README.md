Design
======


Client
======

Remote client has no BTFS node. The client uses gRPC to access the server network and it gets a file. Streaming is used.


BTFS Node
=========

The client runs a node and hence everything is done locally. This is much quicker but needs a node. We can test later with
send data to remote BTFS nodes directly. This might need a CORS change. Problem is security also nodes might get overloaded.
A load balancing is needed at the node level which means building a server architecture and hence the advantage should be tested.

Blockchain
==========

A local node is better for speed. If that cannot be done then the server system needs to update again using gRPC.

Sidechains
==========

IPFS uses chunks and DAGs but we also use chunks for security. Hence we have a lot of txns. This should be stored on sidechains. We
need to test on alternative forms of storage. 

Elasticsearch
=============

This needs to be tested too. We can store key JSON objects reflecting BTFS content.

Encryption
==========

Keys have to be stored somewhere and that really means an app (web or phone). A user will need some kind of private storage for their
private keys. 


