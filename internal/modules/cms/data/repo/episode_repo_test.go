package repo

import (
	"context"
	"testing"

	"thmanyah/internal/modules/cms/biz"

	"github.com/google/uuid"
)

func TestEpisodeRepo_CRUDJourney(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.Cleanup(context.Background(), t)

	ctx := context.Background()
	repo := NewEpisodeRepository(helper.Pool)
	userID := uuid.MustParse(GetTestUserID())
	categoryID := uuid.MustParse(GetTestCategoryID())

	// First create a program for episodes
	programRepo := NewProgramRepository(helper.Pool)
	program := &biz.Program{
		Title:      "Test Program for Episodes",
		CategoryID: categoryID,
		Status:     biz.ProgramStatusDraft,
		CreatedBy:  userID,
		UpdatedBy:  userID,
	}
	err := programRepo.Create(ctx, program)
	AssertNoError(t, err, "creating test program")

	// Test 1: Create Episode
	t.Run("Create", func(t *testing.T) {
		episode := &biz.Episode{
			ProgramID:     program.ID,
			Title:         "Test Episode Journey",
			Description:   "Test episode for CRUD journey",
			DurationSecs:  1800, // 30 minutes
			EpisodeNumber: 1,
			SeasonNumber:  1,
			Status:        biz.EpisodeStatusDraft,
			CreatedBy:     userID,
			UpdatedBy:     userID,
			MediaURL:      "https://example.com/episode.mp3",
			ThumbnailURL:  "https://example.com/thumb.jpg",
			Tags:          []string{"episode", "test"},
			Metadata:      biz.Metadata{"key1": "value1"},
			ViewCount:     50,
			Rating:        4.2,
		}

		err := repo.Create(ctx, episode)
		AssertNoError(t, err, "creating episode")

		if episode.ID == uuid.Nil {
			t.Error("Expected episode ID to be set")
		}
		if episode.CreatedAt.IsZero() {
			t.Error("Expected CreatedAt to be set")
		}
	})

	// Test 2: Create Episode with Duplicate Number (should fail)
	t.Run("CreateDuplicate", func(t *testing.T) {
		duplicate := &biz.Episode{
			ProgramID:     program.ID,
			Title:         "Duplicate Episode",
			EpisodeNumber: 1, // Same as above
			SeasonNumber:  1, // Same as above
			Status:        biz.EpisodeStatusDraft,
			CreatedBy:     userID,
			UpdatedBy:     userID,
		}

		err := repo.Create(ctx, duplicate)
		AssertError(t, err, "creating episode with duplicate number")
	})

	// Test 3: Get Episode By ID
	t.Run("GetByID", func(t *testing.T) {
		// Create episode for this test
		testEpisode := &biz.Episode{
			ProgramID:     program.ID,
			Title:         "Get By ID Test",
			Description:   "Test Description",
			EpisodeNumber: 2,
			SeasonNumber:  1,
			Status:        biz.EpisodeStatusDraft,
			CreatedBy:     userID,
			UpdatedBy:     userID,
			Tags:          []string{"tag1", "tag2"},
		}

		err := repo.Create(ctx, testEpisode)
		AssertNoError(t, err, "creating episode for GetByID test")

		// Get episode without program
		retrieved, err := repo.GetByID(ctx, testEpisode.ID)
		AssertNoError(t, err, "getting episode by ID")

		if retrieved.ID != testEpisode.ID {
			t.Errorf("Expected ID %s, got %s", testEpisode.ID, retrieved.ID)
		}
		if retrieved.Title != testEpisode.Title {
			t.Errorf("Expected title %s, got %s", testEpisode.Title, retrieved.Title)
		}
		if len(retrieved.Tags) != len(testEpisode.Tags) {
			t.Errorf("Expected %d tags, got %d", len(testEpisode.Tags), len(retrieved.Tags))
		}
	})

	// Test 4: Update Episode
	t.Run("Update", func(t *testing.T) {
		// Create episode to update
		testEpisode := &biz.Episode{
			ProgramID:     program.ID,
			Title:         "Original Title",
			Description:   "Original Description",
			EpisodeNumber: 4,
			SeasonNumber:  1,
			Status:        biz.EpisodeStatusDraft,
			CreatedBy:     userID,
			UpdatedBy:     userID,
			Tags:          []string{"original"},
		}

		err := repo.Create(ctx, testEpisode)
		AssertNoError(t, err, "creating episode for update test")

		// Update episode
		newTitle := "Updated Title"
		newDescription := "Updated Description"
		newStatus := biz.EpisodeStatusPublished
		newTags := []string{"updated", "tags"}
		newDuration := int32(3600) // 1 hour

		updates := &biz.UpdateEpisodeRequest{
			Title:        &newTitle,
			Description:  &newDescription,
			Status:       &newStatus,
			Tags:         &newTags,
			DurationSecs: &newDuration,
		}

		updated, err := repo.Update(ctx, userID, testEpisode.ID, updates)
		AssertNoError(t, err, "updating episode")

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
		if updated.DurationSecs != newDuration {
			t.Errorf("Expected duration %d, got %d", newDuration, updated.DurationSecs)
		}
	})

	// Test 5: List Episodes with Filters
	t.Run("ListWithFilters", func(t *testing.T) {
		// Create episodes with different statuses
		draftEpisode := &biz.Episode{
			ProgramID:     program.ID,
			Title:         "Draft Episode",
			EpisodeNumber: 5,
			SeasonNumber:  1,
			Status:        biz.EpisodeStatusDraft,
			CreatedBy:     userID,
			UpdatedBy:     userID,
		}

		publishedEpisode := &biz.Episode{
			ProgramID:     program.ID,
			Title:         "Published Episode",
			EpisodeNumber: 6,
			SeasonNumber:  1,
			Status:        biz.EpisodeStatusPublished,
			CreatedBy:     userID,
			UpdatedBy:     userID,
		}

		err := repo.Create(ctx, draftEpisode)
		AssertNoError(t, err, "creating draft episode")

		err = repo.Create(ctx, publishedEpisode)
		AssertNoError(t, err, "creating published episode")

		// Filter by program ID
		filter := biz.EpisodeFilter{ProgramID: &program.ID}
		pagination := biz.PaginationRequest{Page: 1, PageSize: 10}
		sort := biz.SortRequest{SortBy: "episode_number", SortOrder: "asc"}

		result, _, err := repo.List(ctx, filter, pagination, sort)
		AssertNoError(t, err, "listing episodes by program")

		// All results should belong to the program
		for _, ep := range result {
			if ep.ProgramID != program.ID {
				t.Errorf("Expected program ID %s, got %s", program.ID, ep.ProgramID)
			}
		}

		// Filter by status
		publishedStatus := biz.EpisodeStatusPublished
		filter = biz.EpisodeFilter{Status: &publishedStatus}

		result, _, err = repo.List(ctx, filter, pagination, sort)
		AssertNoError(t, err, "listing published episodes")

		// All results should be published
		for _, ep := range result {
			if ep.Status != biz.EpisodeStatusPublished {
				t.Errorf("Expected only published episodes, got %s", ep.Status)
			}
		}
	})

	// Test 7: Soft and Hard Delete
	t.Run("DeleteOperations", func(t *testing.T) {
		// Create episode for soft delete
		softDeleteEpisode := &biz.Episode{
			ProgramID:     program.ID,
			Title:         "Soft Delete Test",
			EpisodeNumber: 7,
			SeasonNumber:  1,
			Status:        biz.EpisodeStatusDraft,
			CreatedBy:     userID,
			UpdatedBy:     userID,
		}

		err := repo.Create(ctx, softDeleteEpisode)
		AssertNoError(t, err, "creating episode for soft delete")

		// Soft delete episode
		err = repo.Delete(ctx, userID, softDeleteEpisode.ID)
		AssertNoError(t, err, "soft deleting episode")

		// Episode should still exist but be archived
		archivedEpisode, err := repo.GetByID(ctx, softDeleteEpisode.ID)
		AssertNoError(t, err, "getting archived episode")

		if archivedEpisode.Status != biz.EpisodeStatusArchived {
			t.Errorf("Expected status %s, got %s", biz.EpisodeStatusArchived, archivedEpisode.Status)
		}
	})

	// Test 8: View Count Operations
	t.Run("ViewCount", func(t *testing.T) {
		// Create episode with initial view count
		episode := &biz.Episode{
			ProgramID:     program.ID,
			Title:         "View Count Test",
			EpisodeNumber: 9,
			SeasonNumber:  1,
			Status:        biz.EpisodeStatusDraft,
			CreatedBy:     userID,
			UpdatedBy:     userID,
			ViewCount:     5,
		}

		err := repo.Create(ctx, episode)
		AssertNoError(t, err, "creating episode for view count test")

		initialCount := episode.ViewCount

		// Increment view count
		err = repo.IncrementViewCount(ctx, episode.ID)
		AssertNoError(t, err, "incrementing view count")

		// Verify increment
		updatedEpisode, err := repo.GetByID(ctx, episode.ID)
		AssertNoError(t, err, "getting updated episode")

		if updatedEpisode.ViewCount != initialCount+1 {
			t.Errorf("Expected view count %d, got %d", initialCount+1, updatedEpisode.ViewCount)
		}
	})

	// Test 9: List By Program
	t.Run("ListByProgram", func(t *testing.T) {
		// Create second program for comparison
		program2 := &biz.Program{
			Title:      "Second Program",
			CategoryID: categoryID,
			Status:     biz.ProgramStatusDraft,
			CreatedBy:  userID,
			UpdatedBy:  userID,
		}
		err := programRepo.Create(ctx, program2)
		AssertNoError(t, err, "creating second program")

		// Create episode for second program
		episode2 := &biz.Episode{
			ProgramID:     program2.ID,
			Title:         "Episode for Program 2",
			EpisodeNumber: 1,
			SeasonNumber:  1,
			Status:        biz.EpisodeStatusDraft,
			CreatedBy:     userID,
			UpdatedBy:     userID,
		}

		err = repo.Create(ctx, episode2)
		AssertNoError(t, err, "creating episode for second program")

		// List episodes by first program
		pagination := biz.PaginationRequest{Page: 1, PageSize: 10}
		sort := biz.SortRequest{SortBy: "episode_number", SortOrder: "asc"}

		result, _, err := repo.ListByProgram(ctx, program.ID, pagination, sort)
		AssertNoError(t, err, "listing episodes by program 1")

		// All results should belong to program 1
		for _, ep := range result {
			if ep.ProgramID != program.ID {
				t.Errorf("Expected program ID %s, got %s", program.ID, ep.ProgramID)
			}
		}

		// Should not find episodes from program 2
		foundProgram2Episode := false
		for _, ep := range result {
			if ep.ID == episode2.ID {
				foundProgram2Episode = true
				break
			}
		}
		if foundProgram2Episode {
			t.Error("Should not find episodes from program 2 in program 1 results")
		}
	})
}
