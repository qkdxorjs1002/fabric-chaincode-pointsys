# fabric-chaincode-pointsys

![image](https://user-images.githubusercontent.com/3241479/135601688-b313ce96-03c4-489a-99ea-c7e8b771e50d.png)

## Description
`golang`으로 작성된 체인코드이며 `Express.js`와 `Angular.js`로 작성된 
간단한 웹 페이지를 통해 CRUD 작업을 실행 해볼 수 있습니다.

1. (func AddMember)         멤버 추가
2. (func QueryMember)       멤버 정보 조회
3. (func UpdateMemberPoint) 멤버 지갑 포인트 갱신
5. (func DeleteMember)      멤버 제거


## Installation
```
$ cd ./chaincode/pointsys/go
$ go mod init
$ go mod tidy
$ go build

## Run on cli container ($ docker exec -it cli bash)
## Go to repository dir

~# peer lifecycle chaincode package pointsys.tar.gz \
    --path ./chaincode/pointsys/go/ \
    --lang golang \
    --label pointsys_1

~# peer lifecycle chaincode install pointsys.tar.gz

##############################################
## Approve chaincode with your own commands ##
##############################################

~# exit
```

## Demo
```
## Client Initialization
$ npm install

$ cd ./application
$ cp -r <Path of network dir>/organizations/peerOrganizations/org1.example.com/connection-org1.json .

$ cd ./sdk
$ node enrollAdmin.js
$ node registerUser.js

$ cd ../rest
$ node server.js

##############################################
## Connect localhost:8001 using web browser ##
##############################################
```
