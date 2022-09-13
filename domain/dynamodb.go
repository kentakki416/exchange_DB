//  必要なデータ型を定義
package domain

type Data struct {
	Name string `dynamo:"name,hash"`
	Text string `dynamo:"text"`
}

type DynamoDBRepo interface {
	Store(data Data) error
	GetByName(name string) (Data, error)
}
