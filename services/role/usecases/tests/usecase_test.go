package tests

import (
	"architecture_template/constants/notis"
	"architecture_template/services/role/entities"
	"architecture_template/services/role/mocks"
	businesslogics "architecture_template/services/role/usecases/businessLogics"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllRoles(t *testing.T) {
	var mockRepo = &mocks.RoleRepoMock{}
	var mockService = businesslogics.InitializeService(mockRepo, &log.Logger{})

	mockRepo.On(repoGetAll, mock.Anything).Return(sampleRoles, nil)

	actual, err := mockService.GetAllRoles(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, len(*sampleRoles), len(*actual))
}

func TestGetRolesByName(t *testing.T) {
	var mockRepo = &mocks.RoleRepoMock{}
	var mockService = businesslogics.InitializeService(mockRepo, &log.Logger{})

	mockRepo.On(repoGetByName, mock.AnythingOfType("string"), mock.Anything).Return(sampleNameBasedRoles, nil)

	actual, err := mockService.GetRolesByName(searchKw, context.Background())

	assert.NoError(t, err)
	assert.Equal(t, len(*sampleNameBasedRoles), len(*actual))
}

func TestGetRolesByStatus(t *testing.T) {
	for i := 0; i < 2; i++ {
		var mockRepo = &mocks.RoleRepoMock{}
		var mockService = businesslogics.InitializeService(mockRepo, &log.Logger{})

		var expected = sampleActiveRoles
		var rawStatus = rawActive
		if i == 1 {
			expected = sampleInactiveRoles
			rawStatus = rawInactive
		}

		mockRepo.On(repoGetByStatus, mock.AnythingOfType("bool"), mock.Anything).Return(expected, nil)

		actual, err := mockService.GetRolesByStatus(rawStatus, context.Background())

		assert.NoError(t, err)
		assert.Equal(t, len(*expected), len(*actual))
	}
}
func TestGetRoleById(t *testing.T) {
	var mockRepo = &mocks.RoleRepoMock{}
	var mockService = businesslogics.InitializeService(mockRepo, &log.Logger{})

	mockRepo.On(repoGetById, mock.AnythingOfType("string"), mock.Anything).Return(sampleRole, nil)

	actual, err := mockService.GetRoleById(existed, context.Background())

	assert.NoError(t, err)
	assert.True(t, sampleRole.RoleId == actual.RoleId)
	assert.True(t, sampleRole.RoleName == actual.RoleName)
	assert.True(t, sampleRole.ActiveStatus == actual.ActiveStatus)
}

func TestCreateRole(t *testing.T) {
	var mockRepo = &mocks.RoleRepoMock{}
	var mockService = businesslogics.InitializeService(mockRepo, &log.Logger{})

	mockRepo.On(repoGetAll, mock.Anything).Return(sampleRoles, nil)

	mockRepo.On(repoCreate, mock.AnythingOfType(role), mock.Anything).Return(nil)

	assert.NoError(t, mockService.CreateRole(newRole, context.Background()))
}

func TestUpdateRole(t *testing.T) {
	var mockRepo = &mocks.RoleRepoMock{}
	var mockService = businesslogics.InitializeService(mockRepo, &log.Logger{})

	mockRepo.On(repoGetById, mock.AnythingOfType("string"), mock.Anything).Return(sampleRole, nil)

	mockRepo.On(repoUpdate, mock.AnythingOfType(role), mock.Anything).Return(nil)

	assert.NoError(t, mockService.UpdateRole(*sampleRole, context.Background()))
}

// Sample test for an invalid case
func TestUpdateRoleRequestNotExistRoleReturnsErrorOfWarning(t *testing.T) {
	var mockRepo = &mocks.RoleRepoMock{}
	var mockService = businesslogics.InitializeService(mockRepo, &log.Logger{})

	var emptyResponse *entities.Role = nil
	mockRepo.On(repoGetById, mock.AnythingOfType("string"), mock.Anything).Return(emptyResponse, nil)

	var errMsg error = mockService.UpdateRole(*sampleRole, context.Background())

	assert.Error(t, errMsg)

	assert.Equal(t, notis.UndefinedRoleWarnMsg, errMsg.Error())
}

func TestRemoveRole(t *testing.T) {
	var mockRepo = &mocks.RoleRepoMock{}
	var mockService = businesslogics.InitializeService(mockRepo, &log.Logger{})

	mockRepo.On(repoRemove, mock.AnythingOfType("string"), mock.Anything).Return(nil)

	assert.NoError(t, mockService.RemoveRole(existed, context.Background()))
}

func TestActivateRole(t *testing.T) {
	var mockRepo = &mocks.RoleRepoMock{}
	var mockService = businesslogics.InitializeService(mockRepo, &log.Logger{})

	mockRepo.On(repoActivate, mock.AnythingOfType("string"), mock.Anything).Return(nil)

	assert.NoError(t, mockService.ActivateRole(existed, context.Background()))
}

// Note: There is still a wide range of test cases. These are just regular cases as samples, use them as template to create further test cases
