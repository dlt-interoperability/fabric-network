.PHONY:	all
all: start

.PHONY: start
start: .fabric-setup
	(cd network && ./network.sh up createChannel -c mychannel -ca)

.PHONY: stop
stop: 
	(cd network && ./network.sh down)
	rm -rf asset-transfer-basic/application-javascript/wallet/*

.PHONY: deploy-cc
deploy-cc:
	(cd network && ./network.sh deployCC -ccn basic -ccl javascript)

.PHONY: invoke-cc
invoke-cc: .app-setup
	(cd asset-transfer-basic/application-javascript && node app.js)

.PHONY: clean
clean: stop
	rm -rf asset-transfer-basic/application-javascript/wallet

.fabric-setup: 
	./bootstrap.sh -d
	touch .fabric-setup

.app-setup:
	(cd asset-transfer-basic/application-javascript && npm install)
	touch .app-setup