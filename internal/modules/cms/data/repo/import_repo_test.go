package repo

import (
	"context"
	"testing"

	"thmanyah/internal/modules/cms/biz"

	"github.com/google/uuid"
)

func TestImportRepo_CRUDJourney(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.Cleanup(context.Background(), t)

	ctx := context.Background()
	repo := NewImportRepository(helper.Pool)
	userID := uuid.MustParse(GetTestUserID())

	// Create a test category first
	categoryRepo := NewCategoryRepository(helper.Pool)
	testCategory := &biz.Category{
		Name:      "Test Import Category",
		Type:      biz.CategoryTypePodcast,
		CreatedBy: userID,
	}
	err := categoryRepo.Create(ctx, testCategory)
	AssertNoError(t, err, "creating test category for import tests")
	categoryID := testCategory.ID

	// Test 1: Create Import
	t.Run("Create", func(t *testing.T) {
		importData := &biz.ImportData{
			SourceType:   "rss",
			SourceURL:    "https://example.com/feed.xml",
			SourceConfig: biz.Metadata{"url": "https://example.com/feed.xml"},
			CategoryID:   categoryID,
			Status:       biz.ImportStatusPending,
			TotalItems:   100,
			CreatedBy:    userID,
			FieldMapping: biz.Metadata{"title": "title", "description": "summary"},
		}

		err := repo.Create(ctx, importData)
		AssertNoError(t, err, "creating import")

		if importData.ID == uuid.Nil {
			t.Error("Expected import ID to be set")
		}
		if importData.CreatedAt.IsZero() {
			t.Error("Expected CreatedAt to be set")
		}
	})

	// Test 2: Create Import with Invalid Category (should fail)
	t.Run("CreateInvalidCategory", func(t *testing.T) {
		invalidCategoryID := uuid.New()
		importData := &biz.ImportData{
			SourceType: "rss",
			SourceURL:  "https://example.com/invalid.xml",
			CategoryID: invalidCategoryID,
			Status:     biz.ImportStatusPending,
			CreatedBy:  userID,
		}

		err := repo.Create(ctx, importData)
		AssertError(t, err, "creating import with invalid category")
	})

	// Test 3: Get Import By ID
	t.Run("GetByID", func(t *testing.T) {
		// Create import for this test
		testImport := &biz.ImportData{
			SourceType:   "api",
			SourceURL:    "https://api.example.com/data",
			SourceConfig: biz.Metadata{"api_key": "test123"},
			CategoryID:   categoryID,
			Status:       biz.ImportStatusPending,
			TotalItems:   50,
			CreatedBy:    userID,
		}

		err := repo.Create(ctx, testImport)
		AssertNoError(t, err, "creating import for GetByID test")

		retrieved, err := repo.GetByID(ctx, testImport.ID)
		AssertNoError(t, err, "getting import by ID")

		if retrieved.ID != testImport.ID {
			t.Errorf("Expected ID %s, got %s", testImport.ID, retrieved.ID)
		}
		if retrieved.SourceType != testImport.SourceType {
			t.Errorf("Expected source type %s, got %s", testImport.SourceType, retrieved.SourceType)
		}
		if retrieved.TotalItems != testImport.TotalItems {
			t.Errorf("Expected total items %d, got %d", testImport.TotalItems, retrieved.TotalItems)
		}
	})

	// Test 4: Get Non-existent Import
	t.Run("GetByIDNotFound", func(t *testing.T) {
		nonExistentID := uuid.New()
		_, err := repo.GetByID(ctx, nonExistentID)
		AssertError(t, err, "getting non-existent import")
	})

	// Test 5: Update Import
	t.Run("Update", func(t *testing.T) {
		// Create import to update
		testImport := &biz.ImportData{
			SourceType: "csv",
			SourceURL:  "https://example.com/data.csv",
			Status:     biz.ImportStatusPending,
			TotalItems: 75,
			CategoryID: categoryID,
			CreatedBy:  userID,
		}

		err := repo.Create(ctx, testImport)
		AssertNoError(t, err, "creating import for update test")

		// Update it
		newStatus := biz.ImportStatusProcessing
		newTotalItems := int32(80)
		newSourceURL := "https://example.com/updated.csv"

		updates := &biz.UpdateImportRequest{
			Status:     &newStatus,
			TotalItems: &newTotalItems,
			SourceURL:  &newSourceURL,
		}

		updated, err := repo.Update(ctx, userID, testImport.ID, updates)
		AssertNoError(t, err, "updating import")

		if updated.Status != newStatus {
			t.Errorf("Expected status %s, got %s", newStatus, updated.Status)
		}
		if updated.TotalItems != newTotalItems {
			t.Errorf("Expected total items %d, got %d", newTotalItems, updated.TotalItems)
		}
		if updated.SourceURL != newSourceURL {
			t.Errorf("Expected source URL %s, got %s", newSourceURL, updated.SourceURL)
		}
	})

	// Test 6: List Imports with Filters
	t.Run("ListWithFilters", func(t *testing.T) {
		// Create imports with different statuses
		pendingImport := &biz.ImportData{
			SourceType: "json",
			SourceURL:  "https://example.com/pending.json",
			Status:     biz.ImportStatusPending,
			CategoryID: categoryID,
			CreatedBy:  userID,
		}

		completedImport := &biz.ImportData{
			SourceType: "xml",
			SourceURL:  "https://example.com/completed.xml",
			Status:     biz.ImportStatusCompleted,
			CategoryID: categoryID,
			CreatedBy:  userID,
		}

		err := repo.Create(ctx, pendingImport)
		AssertNoError(t, err, "creating pending import")

		err = repo.Create(ctx, completedImport)
		AssertNoError(t, err, "creating completed import")

		// Filter by status
		pendingStatus := biz.ImportStatusPending
		filter := biz.ImportFilter{Status: &pendingStatus}
		pagination := biz.PaginationRequest{Page: 1, PageSize: 10}
		sort := biz.SortRequest{SortBy: "created_at", SortOrder: "desc"}

		result, _, err := repo.List(ctx, filter, pagination, sort)
		AssertNoError(t, err, "listing pending imports")

		// All results should be pending
		for _, imp := range result {
			if imp.Status != biz.ImportStatusPending {
				t.Errorf("Expected only pending imports, got %s", imp.Status)
			}
		}
	})

	// Test 7: Progress Update Operations
	t.Run("ProgressUpdates", func(t *testing.T) {
		// Create import for progress testing
		testImport := &biz.ImportData{
			SourceType: "feed",
			SourceURL:  "https://example.com/progress.xml",
			Status:     biz.ImportStatusPending,
			TotalItems: 100,
			CategoryID: categoryID,
			CreatedBy:  userID,
		}

		err := repo.Create(ctx, testImport)
		AssertNoError(t, err, "creating import for progress test")

		// Update status to processing
		processingStatus := biz.ImportStatusProcessing
		statusUpdate := &biz.UpdateImportRequest{Status: &processingStatus}

		_, err = repo.Update(ctx, userID, testImport.ID, statusUpdate)
		AssertNoError(t, err, "updating to processing status")

		// Update progress
		processed := int32(50)
		success := int32(45)
		errors := int32(5)

		progressUpdate := &biz.UpdateImportProgressRequest{
			ProcessedItems: &processed,
			SuccessCount:   &success,
			ErrorCount:     &errors,
		}

		err = repo.UpdateProgress(ctx, userID, testImport.ID, progressUpdate)
		AssertNoError(t, err, "updating import progress")

		// Verify progress update
		updated, err := repo.GetByID(ctx, testImport.ID)
		AssertNoError(t, err, "getting updated import")

		if updated.ProcessedItems != processed {
			t.Errorf("Expected processed items %d, got %d", processed, updated.ProcessedItems)
		}
		if updated.SuccessCount != success {
			t.Errorf("Expected success count %d, got %d", success, updated.SuccessCount)
		}
		if updated.ErrorCount != errors {
			t.Errorf("Expected error count %d, got %d", errors, updated.ErrorCount)
		}
	})

	// Test 8: Error and Warning Management
	t.Run("ErrorWarningManagement", func(t *testing.T) {
		// Create import for error/warning testing
		testImport := &biz.ImportData{
			SourceType: "api",
			SourceURL:  "https://api.example.com/errors",
			Status:     biz.ImportStatusProcessing,
			CategoryID: categoryID,
			CreatedBy:  userID,
		}

		err := repo.Create(ctx, testImport)
		AssertNoError(t, err, "creating import for error/warning test")

		// Add errors
		err = repo.AddError(ctx, userID, testImport.ID, "Failed to parse item 10")
		AssertNoError(t, err, "adding first error")

		err = repo.AddError(ctx, userID, testImport.ID, "Network timeout on item 25")
		AssertNoError(t, err, "adding second error")

		// Add warnings
		err = repo.AddWarning(ctx, userID, testImport.ID, "Missing thumbnail for item 5")
		AssertNoError(t, err, "adding first warning")

		err = repo.AddWarning(ctx, userID, testImport.ID, "Invalid date format for item 15")
		AssertNoError(t, err, "adding second warning")

		// Verify errors and warnings were added
		updated, err := repo.GetByID(ctx, testImport.ID)
		AssertNoError(t, err, "getting import with errors/warnings")

		if len(updated.Errors) != 2 {
			t.Errorf("Expected 2 errors, got %d", len(updated.Errors))
		}
		if len(updated.Warnings) != 2 {
			t.Errorf("Expected 2 warnings, got %d", len(updated.Warnings))
		}

		// Check specific error messages
		expectedErrors := []string{"Failed to parse item 10", "Network timeout on item 25"}
		for i, expectedError := range expectedErrors {
			if i < len(updated.Errors) && updated.Errors[i] != expectedError {
				t.Errorf("Expected error '%s', got '%s'", expectedError, updated.Errors[i])
			}
		}

		// Check specific warning messages
		expectedWarnings := []string{"Missing thumbnail for item 5", "Invalid date format for item 15"}
		for i, expectedWarning := range expectedWarnings {
			if i < len(updated.Warnings) && updated.Warnings[i] != expectedWarning {
				t.Errorf("Expected warning '%s', got '%s'", expectedWarning, updated.Warnings[i])
			}
		}
	})

	// Test 9: Complete Import Journey
	t.Run("CompleteJourney", func(t *testing.T) {
		// Create a new import
		importData := &biz.ImportData{
			SourceType:   "rss",
			SourceURL:    "https://example.com/complete.xml",
			SourceConfig: biz.Metadata{"refresh_interval": "1h"},
			CategoryID:   categoryID,
			Status:       biz.ImportStatusPending,
			TotalItems:   200,
			CreatedBy:    userID,
			FieldMapping: biz.Metadata{"title": "title", "content": "description"},
		}

		err := repo.Create(ctx, importData)
		AssertNoError(t, err, "creating import for complete journey")

		// Start processing
		processingStatus := biz.ImportStatusProcessing
		statusUpdate := &biz.UpdateImportRequest{Status: &processingStatus}

		_, err = repo.Update(ctx, userID, importData.ID, statusUpdate)
		AssertNoError(t, err, "updating to processing status")

		// Simulate processing with progress updates
		processed := int32(50)
		success := int32(45)
		errors := int32(5)
		progressUpdate := &biz.UpdateImportProgressRequest{
			ProcessedItems: &processed,
			SuccessCount:   &success,
			ErrorCount:     &errors,
		}

		err = repo.UpdateProgress(ctx, userID, importData.ID, progressUpdate)
		AssertNoError(t, err, "updating progress")

		// Add some errors and warnings during processing
		err = repo.AddError(ctx, userID, importData.ID, "Failed to parse item 10")
		AssertNoError(t, err, "adding error during processing")

		err = repo.AddWarning(ctx, userID, importData.ID, "Missing metadata for item 20")
		AssertNoError(t, err, "adding warning during processing")

		// Complete the import
		completedStatus := biz.ImportStatusCompleted
		finalProcessed := int32(200)
		finalSuccess := int32(195)
		finalErrors := int32(5)

		finalUpdate := &biz.UpdateImportRequest{Status: &completedStatus}
		_, err = repo.Update(ctx, userID, importData.ID, finalUpdate)
		AssertNoError(t, err, "completing import")

		finalProgressUpdate := &biz.UpdateImportProgressRequest{
			ProcessedItems: &finalProcessed,
			SuccessCount:   &finalSuccess,
			ErrorCount:     &finalErrors,
		}

		err = repo.UpdateProgress(ctx, userID, importData.ID, finalProgressUpdate)
		AssertNoError(t, err, "final progress update")

		// Verify final state
		final, err := repo.GetByID(ctx, importData.ID)
		AssertNoError(t, err, "getting final import state")

		if final.Status != biz.ImportStatusCompleted {
			t.Errorf("Expected status %s, got %s", biz.ImportStatusCompleted, final.Status)
		}
		if final.ProcessedItems != finalProcessed {
			t.Errorf("Expected processed items %d, got %d", finalProcessed, final.ProcessedItems)
		}
		if final.SuccessCount != finalSuccess {
			t.Errorf("Expected success count %d, got %d", finalSuccess, final.SuccessCount)
		}
		if final.ErrorCount != finalErrors {
			t.Errorf("Expected error count %d, got %d", finalErrors, final.ErrorCount)
		}
		if len(final.Errors) == 0 {
			t.Error("Expected to have error messages")
		}
		if len(final.Warnings) == 0 {
			t.Error("Expected to have warning messages")
		}

		// Verify database consistency
		exists, err := helper.RowExists(ctx, "imports", "id = $1", final.ID)
		AssertNoError(t, err, "checking import exists in database")
		if !exists {
			t.Error("Expected import to exist in database")
		}
	})
}
