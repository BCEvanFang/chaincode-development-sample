# Chaincode Development Guide

## Create project in `GOPATH`

若要建立 `Simple Contract` ，可在 `GOPATH` 下建立對應的資料夾

```sh
cd $GOPATH/src/

# Chaincode project
mkdir simplecontract

cd simplecontract

# Main entry point
touch main.go

# Chaincode package folder
mkdir simpleasset

# Put smart contract logic here
touch ./simpleasset/simpleasset.go

# Create test files
touch main_test.go
touch ./simpleasset/simpleasset_test.go
```

資料夾結構

- simplecontract /
  - main.go
  - main_test.go
  - simpleasset /
    - simpleasset.go
    - simpleasset_test.go

---

## 單元測試指令
```sh
# 執行全部單元測試
go test ./...

# 執行 simpleasset 的單元測試
go test ./simpleasset/

# 執行單元測試，並印出詳細資訊
go test ./... -v

# 執行單元測試，並計算覆蓋率
go test ./... -cover

# 將覆蓋率輸出為檔案，並以html做視覺化呈現
go test ./... -coverprofile=cover.out
go tool cover -html=cover.out
```