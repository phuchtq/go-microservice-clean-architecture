package mocks

import (
	"architecture_template/services/role/entities"
	"context"

	"github.com/stretchr/testify/mock"
)

type RoleServiceMock struct {
	mock.Mock
}

func (mock *RoleServiceMock) GetAllRoles(c context.Context) (*[]entities.Role, error) {
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

func (mock *RoleServiceMock) GetRolesByName(name string, c context.Context) (*[]entities.Role, error) {
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

func (mock *RoleServiceMock) GetRolesByStatus(rawStatus string, c context.Context) (*[]entities.Role, error) {
	mockData := mock.Called(rawStatus, c)
	//----------------------------------
	var res1 *[]entities.Role
	if mockFunc, ok := mockData.Get(0).(func(string, context.Context) *[]entities.Role); ok {
		res1 = mockFunc(rawStatus, c)
	} else {
		res1 = mockData.Get(0).(*[]entities.Role)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(0).(func(string, context.Context) error); ok {
		res2 = mockFunc(rawStatus, c)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (mock *RoleServiceMock) GetRoleById(id string, c context.Context) (*entities.Role, error) {
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

func (mock *RoleServiceMock) CreateRole(name string, c context.Context) error {
	mockData := mock.Called(name, c)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(string, context.Context) error); ok {
		return mockFunc(name, c)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (mock *RoleServiceMock) UpdateRole(r entities.Role, c context.Context) error {
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

func (mock *RoleServiceMock) RemoveRole(id string, c context.Context) error {
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

func (mock *RoleServiceMock) ActivateRole(id string, c context.Context) error {
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
