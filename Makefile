.PHONY:	all
all: start

.PHONY: start
start: .fabric-setup
	(cd network && ./network.sh up createChannel -c mychannel -ca)

.PHONY: stop
stop: 
	(cd network && ./network.sh down)
	rm -rf application/wallet/*

.PHONY: deploy-cc
deploy-cc:
	(cd network && ./network.sh deployCC -ccn basic -ccl javascript)

.PHONY: deploy-cc-go
deploy-cc-go:
	(cd network && ./network.sh deployCC -ccn basic -ccl go)

.PHONY: invoke-cc
invoke-cc: .app-setup
	(cd application && node app.js)

.PHONY: clean
clean: stop
	rm -rf application/wallet

.fabric-setup: 
	./bootstrap.sh -d
	touch .fabric-setup

.app-setup:
	(cd application && npm install)
	touch .app-setup