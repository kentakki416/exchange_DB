package repository

import (
	"exchange_DB/domain"
	"exchange_DB/repository/mock"

	"github.com/guregu/dynamo"
)

type DynamoDBRepo struct {
	dynamoDB   *dynamo.DB
	tableName  string
	primaryKey string
}

func NewDynamoDBRepo(d *dynamo.DB, isMock bool) domain.DynamoDBRepo {
	if isMock {
		return mock.NewMockDynamoDBRepo()
	}
	return &DynamoDBRepo{
		dynamoDB:   d,
		tableName:  "test_list",
		primaryKey: "name",
	}
}

func (d *DynamoDBRepo) Store(data domain.Data) error {
	if err := d.dynamoDB.Table(d.tableName).Put(data).Run(); err != nil {
		return err
	}
	return nil
}

func (d *DynamoDBRepo) GetByName(name string) (domain.Data, error) {
	var data domain.Data
	if err := d.dynamoDB.Table(d.tableName).Get(d.primaryKey, name).One(&data); err != nil {
		return domain.Data{}, err
	}
	return data, nil
}
