package mock

import (
	"errors"
	"exchange_DB/domain"
	"sync"
)

type mockDynamoDBRepo struct {
	mockDB sync.Map
}

func NewMockDynamoDBRepo() *mockDynamoDBRepo {
	return &mockDynamoDBRepo{}
}

func (d *mockDynamoDBRepo) Store(data domain.Data) error {
	d.mockDB.Store(data.Name, data.Text)
	return nil
}

func (d *mockDynamoDBRepo) GetByName(name string) (domain.Data, error) {
	text, ok := d.mockDB.Load(name)
	if !ok {
		return domain.Data{}, errors.New("No Item")
	}
	return domain.Data{Name: name, Text: text.(string)}, nil
}
