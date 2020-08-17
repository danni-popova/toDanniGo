package todo

import "github.com/stretchr/testify/mock"

type RepositoryMock struct {
	mock.Mock
}

func (r RepositoryMock) List() {

}
