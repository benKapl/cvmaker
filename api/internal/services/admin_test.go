package services

import (
	"context"
	"errors"
	"testing"

	"github.com/benKapl/cvmaker-api/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAdminService_ResetDatabase(t *testing.T) {
	ctx := context.Background()

	t.Run("successful database reset", func(t *testing.T) {
		mockDb := new(mocks.MockQuerier)
		adminService := NewAdminService(mockDb, "test_platform")

		mockDb.On("DeleteUsers", ctx).Return(nil).Once()

		err := adminService.ResetDatabase(ctx)

		assert.NoError(t, err)
		mockDb.AssertExpectations(t)
	})

	t.Run("database error during reset", func(t *testing.T) {
		mockDb := new(mocks.MockQuerier)
		adminService := NewAdminService(mockDb, "test_platform")
		dbErr := errors.New("failed to delete users")

		mockDb.On("DeleteUsers", ctx).Return(dbErr).Once()

		err := adminService.ResetDatabase(ctx)

		assert.ErrorIs(t, err, dbErr)
		mockDb.AssertExpectations(t)
	})
}
