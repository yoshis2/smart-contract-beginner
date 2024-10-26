# 簡単なスマートコントラクトをGo言語で実装するプロジェクト

## 設定方法は以下のURLに記載しています

<a href="https://crypto-currency-academy.com/ethereum-smart-contract/" target="_blank">Go言語でEthereum Smart Contract環境構築</a>

## gethの使い方

docker run -v /local/path:/sources ethereum/solc:stable -o /sources/output --abi --bin /sources/Contract.sol



solc --optimize --bin ./contracts/smart_contract_start.sol -o build 
