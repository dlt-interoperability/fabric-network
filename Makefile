.PHONY:	all
all: start

.PHONY: start
start: .fabric-setup
	(cd test-network && ./network.sh up createChannel -c mychannel -ca)

.PHONY: stop
stop: 
	(cd test-network && ./network.sh down)

.fabric-setup: 
	./bootstrap.sh -s -d
	touch .fabric-setup