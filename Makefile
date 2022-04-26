export GO111MODULE=on

makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.DEFAULT_GOAL := install

install:
	${GO_MOD} go install ./

get-lint:
	if [ ! -f ./bin/golangci-lint ]; then \
		wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s latest; \
	fi;

lint: get-lint
	./bin/golangci-lint run ./... --timeout 5m0s

check:
	gosec ./...

start: install
	$(GOPATH)/bin/chainbridge-fee-oracle server -c $(makeFileDir)config.yaml -k $(makeFileDir)keyfile.priv -t secp256k1

genmocks:
	mockgen -destination=./store/mock/store.go -source=./store/store.go
	mockgen -destination=./oracle/mock/oracle.go github.com/ChainSafe/chainbridge-fee-oracle/oracle GasPriceOracle,ConversionRateOracle

test:
	go clean -testcache
	go test --race ./...

## license: Adds license header to missing files.
license:
	@echo "  >  \033[32mAdding license headers...\033[0m "
	GO111MODULE=off go get -u github.com/google/addlicense
	addlicense -c "ChainSafe Systems" -f ./script/header.txt -y 2021 .

## license-check: Checks for missing license headers
license-check:
	@echo "  >  \033[Checking for license headers...\033[0m "
	GO111MODULE=off go get -u github.com/google/addlicense
	addlicense -check -c "ChainSafe Systems" -f ./script/header.txt -y 2021 .


