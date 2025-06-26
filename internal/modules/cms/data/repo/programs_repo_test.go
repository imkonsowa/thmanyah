package repo

import (
	"context"
	"testing"

	"thmanyah/internal/modules/cms/biz"

	"github.com/google/uuid"
)

func TestProgramRepo_CRUDJourney(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.Cleanup(context.Background(), t)

	ctx := context.Background()
	repo := NewProgramRepository(helper.Pool)
	userID := uuid.MustParse(GetTestUserID())
	categoryID := uuid.MustParse(GetTestCategoryID())

	// Test 1: Create Program
	t.Run("Create", func(t *testing.T) {
		program := &biz.Program{
			Title:        "Test Program Journey",
			Description:  "Test program for CRUD journey",
			CategoryID:   categoryID,
			Status:       biz.ProgramStatusDraft,
			CreatedBy:    userID,
			UpdatedBy:    userID,
			ThumbnailURL: "https://example.com/thumb.jpg",
			Tags:         []string{"test", "journey"},
			Metadata:     biz.Metadata{"test": "program"},
			IsFeatured:   true,
			ViewCount:    100,
			Rating:       4.5,
		}

		err := repo.Create(ctx, program)
		AssertNoError(t, err, "creating program")

		if program.ID == uuid.Nil {
			t.Error("Expected program ID to be set")
		}
		if program.CreatedAt.IsZero() {
			t.Error("Expected CreatedAt to be set")
		}
	})

	// Test 2: Create with Invalid Category (should fail)
	t.Run("CreateInvalidCategory", func(t *testing.T) {
		invalidCategoryID := uuid.New()
		program := &biz.Program{
			Title:      "Invalid Category Program",
			CategoryID: invalidCategoryID,
			Status:     biz.ProgramStatusDraft,
			CreatedBy:  userID,
			UpdatedBy:  userID,
		}

		err := repo.Create(ctx, program)
		AssertError(t, err, "creating program with invalid category")
	})

	// Test 3: Get Program By ID
	t.Run("GetByID", func(t *testing.T) {
		// Create program for this test
		testProgram := &biz.Program{
			Title:      "Get By ID Test Program",
			CategoryID: categoryID,
			Status:     biz.ProgramStatusDraft,
			CreatedBy:  userID,
			UpdatedBy:  userID,
			Tags:       []string{"get", "test"},
		}
		err := repo.Create(ctx, testProgram)
		AssertNoError(t, err, "creating program for GetByID test")

		// Get without category
		retrieved, err := repo.GetByID(ctx, testProgram.ID)
		AssertNoError(t, err, "getting program by ID")

		if retrieved.ID != testProgram.ID {
			t.Errorf("Expected ID %s, got %s", testProgram.ID, retrieved.ID)
		}
		if retrieved.Title != testProgram.Title {
			t.Errorf("Expected title %s, got %s", testProgram.Title, retrieved.Title)
		}
		if len(retrieved.Tags) != len(testProgram.Tags) {
			t.Errorf("Expected %d tags, got %d", len(testProgram.Tags), len(retrieved.Tags))
		}

	})

	// Test 4: Update Program
	t.Run("Update", func(t *testing.T) {
		// Create program to update
		testProgram := &biz.Program{
			Title:       "Original Title",
			Description: "Original Description",
			CategoryID:  categoryID,
			Status:      biz.ProgramStatusDraft,
			CreatedBy:   userID,
			UpdatedBy:   userID,
			Tags:        []string{"original"},
		}
		err := repo.Create(ctx, testProgram)
		AssertNoError(t, err, "creating program for update test")

		// Update it
		newTitle := "Updated Title"
		newDescription := "Updated Description"
		newStatus := biz.ProgramStatusPublished
		newTags := []string{"updated", "tags"}
		sourceURL := "https://example.com/source"

		updates := &biz.UpdateProgramRequest{
			Title:       &newTitle,
			Description: &newDescription,
			Status:      &newStatus,
			Tags:        &newTags,
			SourceURL:   &sourceURL,
		}

		updated, err := repo.Update(ctx, userID, testProgram.ID, updates)
		AssertNoError(t, err, "updating program")

		if updated.Title != newTitle {
			t.Errorf("Expected title %s, got %s", newTitle, updated.Title)
		}
		if updated.Description != newDescription {
			t.Errorf("Expected description %s, got %s", newDescription, updated.Description)
		}
		if updated.Status != newStatus {
			t.Errorf("Expected status %s, got %s", newStatus, updated.Status)
		}
		if len(updated.Tags) != len(newTags) {
			t.Errorf("Expected %d tags, got %d", len(newTags), len(updated.Tags))
		}
	})

	// Test 5: List Programs with Filters
	t.Run("ListWithFilters", func(t *testing.T) {
		// Create programs with different statuses
		draftProgram := &biz.Program{
			Title:      "Draft Program",
			CategoryID: categoryID,
			Status:     biz.ProgramStatusDraft,
			CreatedBy:  userID,
			UpdatedBy:  userID,
			Tags:       []string{"draft"},
		}

		publishedProgram := &biz.Program{
			Title:      "Published Program",
			CategoryID: categoryID,
			Status:     biz.ProgramStatusPublished,
			CreatedBy:  userID,
			UpdatedBy:  userID,
			Tags:       []string{"published"},
			IsFeatured: true,
		}

		err := repo.Create(ctx, draftProgram)
		AssertNoError(t, err, "creating draft program")

		err = repo.Create(ctx, publishedProgram)
		AssertNoError(t, err, "creating published program")

		// Filter by status
		publishedStatus := biz.ProgramStatusPublished
		filter := biz.ProgramFilter{Status: &publishedStatus}
		pagination := biz.PaginationRequest{Page: 1, PageSize: 10}
		sort := biz.SortRequest{SortBy: "created_at", SortOrder: "asc"}

		result, _, err := repo.List(ctx, filter, pagination, sort)
		AssertNoError(t, err, "listing published programs")

		// All results should be published
		for _, prog := range result {
			if prog.Status != biz.ProgramStatusPublished {
				t.Errorf("Expected only published programs, got %s", prog.Status)
			}
		}

		// Filter by featured
		featuredOnly := true
		filter = biz.ProgramFilter{FeaturedOnly: &featuredOnly}

		result, _, err = repo.List(ctx, filter, pagination, sort)
		AssertNoError(t, err, "listing featured programs")

		// All results should be featured
		for _, prog := range result {
			if !prog.IsFeatured {
				t.Error("Expected only featured programs")
			}
		}
	})

	// Test 7: Bulk Operations
	t.Run("BulkOperations", func(t *testing.T) {
		// Create multiple programs for bulk operations
		var programIDs []uuid.UUID
		for i := 0; i < 3; i++ {
			program := &biz.Program{
				Title:      "Bulk Test Program",
				CategoryID: categoryID,
				Status:     biz.ProgramStatusDraft,
				CreatedBy:  userID,
				UpdatedBy:  userID,
			}

			err := repo.Create(ctx, program)
			AssertNoError(t, err, "creating program for bulk test")
			programIDs = append(programIDs, program.ID)
		}

		// Bulk update
		newStatus := biz.ProgramStatusPublished
		featured := true
		updates := &biz.BulkUpdateProgramsRequest{
			Status:     &newStatus,
			IsFeatured: &featured,
		}

		updatedCount, err := repo.BulkUpdate(ctx, userID, programIDs, updates)
		AssertNoError(t, err, "bulk updating programs")

		if updatedCount != int32(len(programIDs)) {
			t.Errorf("Expected %d updated programs, got %d", len(programIDs), updatedCount)
		}

		// Verify updates
		for _, id := range programIDs {
			program, err := repo.GetByID(ctx, id)
			AssertNoError(t, err, "getting updated program")

			if program.Status != newStatus {
				t.Errorf("Expected status %s, got %s", newStatus, program.Status)
			}
			if !program.IsFeatured {
				t.Error("Expected program to be featured")
			}
		}

		// Bulk delete
		err = repo.BulkDelete(ctx, userID, programIDs)
		AssertNoError(t, err, "bulk deleting programs")

		// Verify deletion
		for _, id := range programIDs {
			_, err := repo.GetByID(ctx, id)
			AssertError(t, err, "getting deleted program")
		}
	})

	// Test 8: View Count Operations
	t.Run("ViewCount", func(t *testing.T) {
		// Create program with initial view count
		program := &biz.Program{
			Title:      "View Count Test",
			CategoryID: categoryID,
			Status:     biz.ProgramStatusDraft,
			CreatedBy:  userID,
			UpdatedBy:  userID,
			ViewCount:  10,
		}

		err := repo.Create(ctx, program)
		AssertNoError(t, err, "creating program for view count test")

		initialCount := program.ViewCount

		// Increment view count
		err = repo.IncrementViewCount(ctx, program.ID)
		AssertNoError(t, err, "incrementing view count")

		// Verify increment
		updatedProgram, err := repo.GetByID(ctx, program.ID)
		AssertNoError(t, err, "getting updated program")

		if updatedProgram.ViewCount != initialCount+1 {
			t.Errorf("Expected view count %d, got %d", initialCount+1, updatedProgram.ViewCount)
		}
	})

	// Test 9: Episodes Count Update
	t.Run("EpisodesCount", func(t *testing.T) {
		// Create program for episodes count test
		program := &biz.Program{
			Title:      "Episodes Count Test",
			CategoryID: categoryID,
			Status:     biz.ProgramStatusDraft,
			CreatedBy:  userID,
			UpdatedBy:  userID,
		}

		err := repo.Create(ctx, program)
		AssertNoError(t, err, "creating program for episodes count test")

		// Update episodes count (this method calculates from episodes table)
		err = repo.UpdateEpisodesCount(ctx, program.ID)
		AssertNoError(t, err, "updating episodes count")

		// Since we haven't created any episodes, count should be 0
		updatedProgram, err := repo.GetByID(ctx, program.ID)
		AssertNoError(t, err, "getting updated program")

		if updatedProgram.EpisodesCount != 0 {
			t.Errorf("Expected episodes count 0, got %d", updatedProgram.EpisodesCount)
		}
	})
}
