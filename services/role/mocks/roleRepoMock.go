package mocks

import (
	"architecture_template/services/role/entities"
	"context"

	"github.com/stretchr/testify/mock"
)

type RoleRepoMock struct {
	mock.Mock
}

func (mock *RoleRepoMock) GetAllRoles(c context.Context) (*[]entities.Role, error) {
	mockData := mock.Called(c)

	var res1 *[]entities.Role
	if mockFunc, ok := mockData.Get(0).(func(context.Context) *[]entities.Role); ok {
		res1 = mockFunc(c)
	} else {
		res1 = mockData.Get(0).(*[]entities.Role)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(1).(func(context.Context) error); ok {
		res2 = mockFunc(c)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (mock *RoleRepoMock) GetRolesByName(name string, c context.Context) (*[]entities.Role, error) {
	mockData := mock.Called(name, c)
	//----------------------------------
	var res1 *[]entities.Role
	if mockFunc, ok := mockData.Get(0).(func(string, context.Context) *[]entities.Role); ok {
		res1 = mockFunc(name, c)
	} else {
		res1 = mockData.Get(0).(*[]entities.Role)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(0).(func(string, context.Context) error); ok {
		res2 = mockFunc(name, c)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (mock *RoleRepoMock) GetRolesByStatus(status bool, c context.Context) (*[]entities.Role, error) {
	mockData := mock.Called(status, c)
	//----------------------------------
	var res1 *[]entities.Role
	if mockFunc, ok := mockData.Get(0).(func(bool, context.Context) *[]entities.Role); ok {
		res1 = mockFunc(status, c)
	} else {
		res1 = mockData.Get(0).(*[]entities.Role)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(0).(func(bool, context.Context) error); ok {
		res2 = mockFunc(status, c)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (mock *RoleRepoMock) GetRoleById(id string, c context.Context) (*entities.Role, error) {
	mockData := mock.Called(id, c)
	//----------------------------------
	var res1 *entities.Role
	if mockFunc, ok := mockData.Get(0).(func(string, context.Context) *entities.Role); ok {
		res1 = mockFunc(id, c)
	} else {
		res1 = mockData.Get(0).(*entities.Role)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(1).(func(string, context.Context) error); ok {
		res2 = mockFunc(id, c)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (mock *RoleRepoMock) CreateRole(r entities.Role, c context.Context) error {
	mockData := mock.Called(r, c)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(entities.Role, context.Context) error); ok {
		return mockFunc(r, c)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (mock *RoleRepoMock) UpdateRole(r entities.Role, c context.Context) error {
	mockData := mock.Called(r, c)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(entities.Role, context.Context) error); ok {
		return mockFunc(r, c)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (mock *RoleRepoMock) RemoveRole(id string, c context.Context) error {
	mockData := mock.Called(id, c)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(string, context.Context) error); ok {
		return mockFunc(id, c)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (mock *RoleRepoMock) ActivateRole(id string, c context.Context) error {
	mockData := mock.Called(id, c)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(string, context.Context) error); ok {
		return mockFunc(id, c)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}
