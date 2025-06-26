package repo

import (
	"context"
	"testing"

	"thmanyah/internal/modules/cms/biz"

	"github.com/google/uuid"
)

func TestCategoryRepo_CRUDJourney(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.Cleanup(context.Background(), t)

	ctx := context.Background()
	repo := NewCategoryRepository(helper.Pool)
	userID := uuid.MustParse(GetTestUserID())

	// Test 1: Create Category
	t.Run("Create", func(t *testing.T) {
		category := &biz.Category{
			Name:        "Journey Test Category",
			Description: "Test category for CRUD journey",
			Type:        biz.CategoryTypePodcast,
			Metadata:    biz.Metadata{"test": "journey"},
		}

		err := repo.Create(ctx, category)
		AssertNoError(t, err, "creating category")

		if category.ID == uuid.Nil {
			t.Error("Expected category ID to be set")
		}
		if category.CreatedAt.IsZero() {
			t.Error("Expected CreatedAt to be set")
		}
	})

	// Test 2: Create Duplicate Name (should fail)
	t.Run("CreateDuplicate", func(t *testing.T) {
		duplicate := &biz.Category{
			Name:        "Journey Test Category", // Same name
			Description: "Duplicate category",
			Type:        biz.CategoryTypeEducational,
		}

		err := repo.Create(ctx, duplicate)
		AssertError(t, err, "creating category with duplicate name")
	})

	// Test 3: Get Category By ID
	t.Run("GetByID", func(t *testing.T) {
		// Create a new category for this test
		testCategory := &biz.Category{
			Name:        "Get By ID Test",
			Description: "For GetByID test",
			Type:        biz.CategoryTypeDocumentary,
		}
		err := repo.Create(ctx, testCategory)
		AssertNoError(t, err, "creating category for GetByID test")

		retrieved, err := repo.GetByID(ctx, testCategory.ID)
		AssertNoError(t, err, "getting category by ID")

		if retrieved.ID != testCategory.ID {
			t.Errorf("Expected ID %s, got %s", testCategory.ID, retrieved.ID)
		}
		if retrieved.Name != testCategory.Name {
			t.Errorf("Expected name %s, got %s", testCategory.Name, retrieved.Name)
		}
	})

	// Test 4: Get Non-existent Category
	t.Run("GetByIDNotFound", func(t *testing.T) {
		nonExistentID := uuid.New()
		_, err := repo.GetByID(ctx, nonExistentID)
		AssertError(t, err, "getting non-existent category")
	})

	// Test 5: Update Category
	t.Run("Update", func(t *testing.T) {
		// Create category to update
		testCategory := &biz.Category{
			Name:        "Update Test Category",
			Description: "Original description",
			Type:        biz.CategoryTypePodcast,
			CreatedBy:   userID,
		}
		err := repo.Create(ctx, testCategory)
		AssertNoError(t, err, "creating category for update test")

		// Update it
		newName := "Updated Category Name"
		newDescription := "Updated description"
		newType := biz.CategoryTypeNews

		updates := &biz.UpdateCategoryRequest{
			Name:        &newName,
			Description: &newDescription,
			Type:        &newType,
		}

		updated, err := repo.Update(ctx, userID, testCategory.ID, updates)
		AssertNoError(t, err, "updating category")

		if updated.Name != newName {
			t.Errorf("Expected name %s, got %s", newName, updated.Name)
		}
		if updated.Description != newDescription {
			t.Errorf("Expected description %s, got %s", newDescription, updated.Description)
		}
		if updated.Type != newType {
			t.Errorf("Expected type %s, got %s", newType, updated.Type)
		}
	})

	// Test 6: List Categories with Filters
	t.Run("ListWithFilters", func(t *testing.T) {
		// Create categories of different types
		categories := []*biz.Category{
			{
				Name: "Sports Category",
				Type: biz.CategoryTypeSportsEvent,
			},
			{
				Name: "Entertainment Category",
				Type: biz.CategoryTypeEntertainment,
			},
		}

		for _, cat := range categories {
			err := repo.Create(ctx, cat)
			AssertNoError(t, err, "creating test category for list")
		}

		// Test filter by type
		sportsType := biz.CategoryTypeSportsEvent
		filter := biz.CategoryFilter{Type: &sportsType}
		pagination := biz.PaginationRequest{Page: 1, PageSize: 10}
		sort := biz.SortRequest{SortBy: "name", SortOrder: "asc"}

		result, paginationResp, err := repo.List(ctx, filter, pagination, sort)
		AssertNoError(t, err, "listing categories by type")

		// Should find at least our sports category
		found := false
		for _, cat := range result {
			if cat.Type == biz.CategoryTypeSportsEvent {
				found = true
				break
			}
		}
		if !found {
			t.Error("Expected to find sports category in filtered results")
		}

		if paginationResp.TotalCount < 1 {
			t.Error("Expected total count to be at least 1")
		}
	})

	// Test 7: List with Search Query
	t.Run("ListWithSearch", func(t *testing.T) {
		// Create category with unique name for search
		searchCategory := &biz.Category{
			Name:        "Unique Search Term Category",
			Description: "For testing search functionality",
			Type:        biz.CategoryTypePodcast,
			CreatedBy:   userID,
		}
		err := repo.Create(ctx, searchCategory)
		AssertNoError(t, err, "creating category for search test")

		// Search for it
		searchTerm := "Unique Search"
		filter := biz.CategoryFilter{SearchQuery: &searchTerm}
		pagination := biz.PaginationRequest{Page: 1, PageSize: 10}
		sort := biz.SortRequest{SortBy: "name", SortOrder: "asc"}

		result, _, err := repo.List(ctx, filter, pagination, sort)
		AssertNoError(t, err, "searching categories")

		// Should find our category
		found := false
		for _, cat := range result {
			if cat.ID == searchCategory.ID {
				found = true
				break
			}
		}
		if !found {
			t.Error("Expected to find category in search results")
		}
	})

	// Test 8: Delete Category
	t.Run("Delete", func(t *testing.T) {
		// Create category to delete
		deleteCategory := &biz.Category{
			Name:      "Category to Delete",
			Type:      biz.CategoryTypePodcast,
			CreatedBy: userID,
		}
		err := repo.Create(ctx, deleteCategory)
		AssertNoError(t, err, "creating category for delete test")

		// Delete it
		err = repo.Delete(ctx, userID, deleteCategory.ID)
		AssertNoError(t, err, "deleting category")

		// Verify it's gone
		_, err = repo.GetByID(ctx, deleteCategory.ID)
		AssertError(t, err, "getting deleted category")
	})

	// Test 9: List All Categories (pagination test)
	t.Run("ListPagination", func(t *testing.T) {
		filter := biz.CategoryFilter{}
		pagination := biz.PaginationRequest{Page: 1, PageSize: 5}
		sort := biz.SortRequest{SortBy: "created_at", SortOrder: "desc"}

		result, paginationResp, err := repo.List(ctx, filter, pagination, sort)
		AssertNoError(t, err, "listing all categories")

		if len(result) > 5 {
			t.Errorf("Expected at most 5 results, got %d", len(result))
		}

		if paginationResp.Page != 1 {
			t.Errorf("Expected page 1, got %d", paginationResp.Page)
		}
		if paginationResp.PageSize != 5 {
			t.Errorf("Expected page size 5, got %d", paginationResp.PageSize)
		}
	})
}
