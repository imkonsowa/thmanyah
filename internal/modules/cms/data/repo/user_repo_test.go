package repo

import (
	"context"
	"testing"

	"thmanyah/internal/modules/cms/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

func TestUsersRepo_CRUDJourney(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.Cleanup(context.Background(), t)

	ctx := context.Background()
	logger := log.DefaultLogger
	repo, err := NewUsersRepo(helper.Pool, logger)
	AssertNoError(t, err, "creating users repo")

	// Test 1: Create User
	t.Run("Create", func(t *testing.T) {
		user := &biz.User{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "hashed_password",
		}

		createdUser, err := repo.CreateUser(ctx, user)
		AssertNoError(t, err, "creating user")

		if createdUser.ID == uuid.Nil {
			t.Error("Expected user ID to be set")
		}
		if createdUser.Name != user.Name {
			t.Errorf("Expected name %s, got %s", user.Name, createdUser.Name)
		}
		if createdUser.Email != user.Email {
			t.Errorf("Expected email %s, got %s", user.Email, createdUser.Email)
		}
		if createdUser.CreatedAt.IsZero() {
			t.Error("Expected CreatedAt to be set")
		}
	})

	// Test 2: Create Duplicate Email (should fail)
	t.Run("CreateDuplicate", func(t *testing.T) {
		duplicateUser := &biz.User{
			Name:     "Jane Doe",
			Email:    "john.doe@example.com", // Same email
			Password: "another_password",
		}

		_, err := repo.CreateUser(ctx, duplicateUser)
		AssertError(t, err, "creating user with duplicate email")
	})

	// Test 3: Get User With Password
	t.Run("GetUserWithPassword", func(t *testing.T) {
		retrievedUser, err := repo.GetUserWithPassword(ctx, "john.doe@example.com")
		AssertNoError(t, err, "getting user with password")

		if retrievedUser.Email != "john.doe@example.com" {
			t.Errorf("Expected email john.doe@example.com, got %s", retrievedUser.Email)
		}
		if retrievedUser.Password != "hashed_password" {
			t.Errorf("Expected password to be returned")
		}
	})

	// Test 4: Get User With Password (not found)
	t.Run("GetUserWithPasswordNotFound", func(t *testing.T) {
		_, err := repo.GetUserWithPassword(ctx, "nonexistent@example.com")
		AssertError(t, err, "getting non-existent user")
	})

	// Test 5: Get User By Identifier (email)
	t.Run("GetUserByEmail", func(t *testing.T) {
		retrievedUser, err := repo.GetUserByIdentifier(ctx, "john.doe@example.com")
		AssertNoError(t, err, "getting user by email")

		if retrievedUser.Email != "john.doe@example.com" {
			t.Errorf("Expected email john.doe@example.com, got %s", retrievedUser.Email)
		}
	})

	// Test 6: Get User By Identifier (ID)
	t.Run("GetUserByID", func(t *testing.T) {
		// First get the user to find their ID
		existingUser, err := repo.GetUserByIdentifier(ctx, "john.doe@example.com")
		AssertNoError(t, err, "getting user for ID test")

		retrievedUser, err := repo.GetUserByIdentifier(ctx, existingUser.ID.String())
		AssertNoError(t, err, "getting user by ID")

		if retrievedUser.ID != existingUser.ID {
			t.Errorf("Expected ID %s, got %s", existingUser.ID, retrievedUser.ID)
		}
	})

	// Test 7: Update User
	t.Run("Update", func(t *testing.T) {
		// First get the user to update
		existingUser, err := repo.GetUserByIdentifier(ctx, "john.doe@example.com")
		AssertNoError(t, err, "getting user for update")

		newName := "John Updated"
		newEmail := "john.updated@example.com"

		updates := &biz.UpdateUserRequest{
			Name:  newName,
			Email: newEmail,
		}

		updatedUser, err := repo.UpdateUser(ctx, existingUser.ID, updates)
		AssertNoError(t, err, "updating user")

		if updatedUser.Name != newName {
			t.Errorf("Expected name %s, got %s", newName, updatedUser.Name)
		}
		if updatedUser.Email != newEmail {
			t.Errorf("Expected email %s, got %s", newEmail, updatedUser.Email)
		}
	})

	// Test 8: Update User (partial update)
	t.Run("PartialUpdate", func(t *testing.T) {
		// Get the updated user
		existingUser, err := repo.GetUserByIdentifier(ctx, "john.updated@example.com")
		AssertNoError(t, err, "getting user for partial update")

		// Only update name
		newName := "John Partially Updated"
		updates := &biz.UpdateUserRequest{
			Name: newName,
		}

		updatedUser, err := repo.UpdateUser(ctx, existingUser.ID, updates)
		AssertNoError(t, err, "partially updating user")

		if updatedUser.Name != newName {
			t.Errorf("Expected name %s, got %s", newName, updatedUser.Name)
		}
		// Email should remain unchanged
		if updatedUser.Email != "john.updated@example.com" {
			t.Errorf("Expected email to remain john.updated@example.com, got %s", updatedUser.Email)
		}
	})

	// Test 9: Update Non-existent User
	t.Run("UpdateNotFound", func(t *testing.T) {
		nonExistentID := uuid.New()
		newName := "Should Fail"

		updates := &biz.UpdateUserRequest{
			Name: newName,
		}

		_, err := repo.UpdateUser(ctx, nonExistentID, updates)
		AssertError(t, err, "updating non-existent user")
	})
}
