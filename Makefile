.PHONY:	all
all: start

.PHONY: start
start: .fabric-setup
	./network/network.sh up createChannel -c mychannel -ca

.PHONY: stop
stop: 
	./network/network.sh down

.fabric-setup: 
	curl -sSL https://bit.ly/2ysbOFE | bash -s
	touch .fabric-setup
