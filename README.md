# natspool
A simple, automatically expanding/cleaning thread-safe connection pool for [nats.io](http://nats.io/).

## Installation

    go get github.com/akaumov/natspool
    
## Documentation
    
   [Documentation](https://godoc.org/github.com/akaumov/natspool)
    
## Testing
    
    go test github.com/akaumov/natspool
    
The test action assumes you have the following running:

* A nats server listening on port 4222


This code is rewriting of "pool" packet from [radix](https://github.com/mediocregopher/radix.v2) for [nats.io](http://nats.io/).


## Copyright and licensing

Unless otherwise noted, the source files are distributed under the *MIT License*
found in the LICENSE.txt file.

[nats.io]: http://nats.io