.PHONY:	all
all: start

.PHONY: start
start: .fabric-setup
	(cd network && ./network.sh up createChannel -c mychannel -ca)

.PHONY: stop
stop: 
	(cd network && ./network.sh down)

.fabric-setup: 
	./bootstrap.sh -s -d
	touch .fabric-setup