package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type PointSys struct {
	contractapi.Contract
}

/**
 * Member
 */
type Member struct {
	Name   string `json:"name"`
	Wallet Wallet `json:"wallet"`
}

func (Member) Create(Name string) *Member {
	return &Member{
		Name:   Name,
		Wallet: *Wallet{}.Create("wallet", 0),
	}
}

func (member *Member) ToJsonBytes() []byte {
	jsonBytes, err := json.Marshal(member)
	if err != nil {
		panic(err)
	}

	return jsonBytes
}

/**
 * Wallet
 */
type Wallet struct {
	Name  string `json:"name"`
	Point uint   `json:"point"`
}

func (Wallet) Create(Name string, Point uint) *Wallet {
	return &Wallet{
		Name:  Name,
		Point: Point,
	}
}

func (wallet *Wallet) ToJsonBytes() []byte {
	jsonBytes, err := json.Marshal(wallet)
	if err != nil {
		panic(err)
	}

	return jsonBytes
}

func (wallet *Wallet) IncPoint(PointToAdd uint) {
	wallet.Point += PointToAdd
}

func (wallet *Wallet) DecPoint(PointToSub uint) error {
	if wallet.Point-PointToSub < 0 {
		return fmt.Errorf("Failed to subtract point: Point < PointToSub")
	}
	wallet.Point -= PointToSub

	return nil
}

/**
 * Member 추가
 *
 * @param ctx contract transaction interface
 * @param name Member 이름
 * @return error, nil
 */
func (t *PointSys) AddMember(ctx contractapi.TransactionContextInterface, name string) error {

	// name 이름을 가진 멤버 추가 (기존 멤버 여부 확인 X)
	member := Member{}.Create(name)

	err := ctx.GetStub().PutState(name, member.ToJsonBytes())
	if err != nil {
		return err
	}

	return nil
}

/**
 * Member 제거
 *
 * @param ctx contract transaction interface
 * @param name Member 이름
 * @return error, nil
 */
func (t *PointSys) DeleteMember(ctx contractapi.TransactionContextInterface, name string) error {

	err := ctx.GetStub().DelState(name)
	if err != nil {
		return fmt.Errorf("Failed to delete state")
	}

	return nil
}

/**
 * Member 조회
 *
 * @param ctx contract transaction interface
 * @param name Member 이름
 * @return error, nil
 */
func (t *PointSys) QueryMember(ctx contractapi.TransactionContextInterface, name string) (string, error) {

	memberBytes, err := ctx.GetStub().GetState(name)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + name + "\"}"
		return "", errors.New(jsonResp)
	}

	if memberBytes == nil {
		jsonResp := "{\"Error\":\"Nil for " + name + "\"}"
		return "", errors.New(jsonResp)
	}

	return string(memberBytes), nil
}

/**
 * Member 포인트 업데이트
 *
 * @param ctx contract transaction interface
 * @param name Member 이름
 * @return error, nil
 */
func (t *PointSys) UpdateMemberPoint(ctx contractapi.TransactionContextInterface, name string, point uint) (string, error) {

	memberBytes, err := ctx.GetStub().GetState(name)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + name + "\"}"
		return "", errors.New(jsonResp)
	}

	if memberBytes == nil {
		jsonResp := "{\"Error\":\"Nil for " + name + "\"}"
		return "", errors.New(jsonResp)
	}

	var member Member
	err = json.Unmarshal(memberBytes, &member)
	if err != nil {
		panic(err)
	}

	member.Wallet.Point = point
	memberJsonBytes := member.ToJsonBytes()

	err = ctx.GetStub().PutState(name, memberJsonBytes)
	if err != nil {
		return "Failed to update state", err
	}

	return string(memberJsonBytes), nil
}

func main() {
	cc, err := contractapi.NewChaincode(new(PointSys))
	if err != nil {
		panic(err.Error())
	}
	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting PointSys chaincode: %s", err)
	}
}
