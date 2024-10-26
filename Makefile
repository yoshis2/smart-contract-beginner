.PHONY: abi
abi:
	docker compose run --rm geth sh

.PHONY: start
start:
	docker compose run --rm solc --optimize --abi --overwrite ./contracts/smart_contract_start.sol -o build
	docker compose run --rm solc --optimize --bin --overwrite ./contracts/smart_contract_start.sol -o build
	docker compose run --rm geth abigen --abi=./build/SmartContractStart.abi --bin=./build/SmartContractStart.bin --pkg=contracts --out=./build/smart_contract_start.go

.PHONY: version
version:
	docker compose run --rm solidity --version
