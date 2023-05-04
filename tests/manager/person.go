package manager

import "github.com/LiveAlone/GoLibSourceAnalyse/tests/model"

type PersonManager struct {
}

func SelectOne(id string) string {
	return "YQJ"
}

func (m *PersonManager) QueryPerson(id int64) (*model.Person, error) {
	return &model.Person{
		Id:   id,
		Name: "test",
		Age:  18,
	}, nil
}
