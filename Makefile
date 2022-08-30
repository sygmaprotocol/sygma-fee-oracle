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
	golangci-lint run ./... --skip-dirs e2e --timeout 5m0s

check:
	gosec -exclude-dir=e2e ./...

start: install
	$(GOPATH)/bin/sygma-fee-oracle server -c $(makeFileDir)config.yaml -d $(makeFileDir)domain.json -r $(makeFileDir)resource.json -k $(makeFileDir)keyfile.priv -t secp256k1

genmocks:
	mockgen -destination=./store/mock/store.go -source=./store/store.go
	mockgen -destination=./oracle/mock/oracle.go github.com/ChainSafe/sygma-fee-oracle/oracle GasPriceOracle,ConversionRateOracle
	mockgen -destination=./config/mock/remoteParamOperator.go -source=./remoteParam/base.go

test:
	go clean -testcache
	./scripts/test.sh

start-ganache:
	./scripts/start_ganache.sh false

e2e-test: install
	./scripts/prepare_test_data.sh
	./scripts/start.sh
	./scripts/start_ganache.sh true
	./scripts/e2e_testing.sh master

## license: Adds license header to missing files.
license:
	@echo "  >  \033[32mAdding license headers...\033[0m "
	GO111MODULE=off go get -u github.com/google/addlicense
	addlicense -c "Sygma" -f ./scripts/header.txt -y 2022 .

## license-check: Checks for missing license headers
license-check:
	@echo "  >  \033[Checking for license headers...\033[0m "
	GO111MODULE=off go get -u github.com/google/addlicense
	addlicense -check -c "Sygma" -f ./scripts/header.txt -y 2022 -skip sh .


