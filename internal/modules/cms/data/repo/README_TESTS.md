# Repository Tests

This directory contains comprehensive integration tests for all repository operations using PostgreSQL testcontainers.

## Overview

The test suite provides:
- **PostgreSQL testcontainers** for isolated database testing
- **Complete CRUD operation tests** for all repositories
- **Edge case and error condition testing**
- **Integration tests** that verify end-to-end functionality
- **Test helpers** for common operations and assertions

## Test Structure

### Test Files
- `user_repo_test.go` - Tests for user repository operations
- `categories_repo_test.go` - Tests for category repository operations  
- `programs_repo_test.go` - Tests for program repository operations
- `episode_repo_test.go` - Tests for episode repository operations
- `import_repo_test.go` - Tests for import repository operations

### Support Files
- `test_helper.go` - Test utilities and helper functions
- `testdata/init.sql` - Database schema and seed data for tests

## Running Tests

### Prerequisites
- Docker (for testcontainers)
- Go 1.24+

### Run All Repository Tests
```bash
# From project root
go test ./internal/modules/cms/repo/... -v

# Run with coverage
go test ./internal/modules/cms/repo/... -v -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Run Specific Test Files
```bash
# Test specific repository
go test ./internal/modules/cms/repo/ -run TestUserRepo -v
go test ./internal/modules/cms/repo/ -run TestCategoryRepo -v
go test ./internal/modules/cms/repo/ -run TestProgramRepo -v
go test ./internal/modules/cms/repo/ -run TestEpisodeRepo -v
go test ./internal/modules/cms/repo/ -run TestImportRepo -v
```

### Run Individual Tests
```bash
# Test specific functionality
go test ./internal/modules/cms/repo/ -run TestProgramRepo_Create -v
go test ./internal/modules/cms/repo/ -run TestEpisodeRepo_List_WithFilters -v
```

## Test Coverage

### User Repository Tests
- ✅ Create user with validation
- ✅ Handle duplicate email errors
- ✅ Get user with password
- ✅ Get user by email/ID identifier
- ✅ Update user (full and partial)
- ✅ Error handling for non-existent users
- ✅ Integration workflow

### Category Repository Tests
- ✅ Create category with metadata
- ✅ Handle duplicate name errors
- ✅ Get category by ID
- ✅ Update category (full and partial)
- ✅ Delete category
- ✅ List with filtering (type, search query)
- ✅ Pagination and sorting
- ✅ Error handling

### Program Repository Tests
- ✅ Create program with tags and metadata
- ✅ Handle foreign key constraints (category, user)
- ✅ Get program with/without category join
- ✅ Update program (including empty tags)
- ✅ Delete program
- ✅ List with complex filters (status, category, featured, tags)
- ✅ Search functionality
- ✅ Bulk operations (update, delete)
- ✅ View count increment
- ✅ Episodes count calculation

### Episode Repository Tests
- ✅ Create episode with program association
- ✅ Handle unique constraints (episode/season numbers)
- ✅ Get episode with/without program join
- ✅ Update episode properties
- ✅ Soft delete (archive) vs hard delete
- ✅ List with filters (program, status)
- ✅ List by specific program
- ✅ View count increment
- ✅ Integration workflow

### Import Repository Tests
- ✅ Create import with configuration
- ✅ Handle foreign key constraints
- ✅ Get import by ID
- ✅ Update import status and properties
- ✅ List with pagination
- ✅ Progress tracking
- ✅ Error and warning accumulation
- ✅ Complete import workflow simulation

## Test Data

### Seed Data
The tests use predefined test data:
- **Test Users**: `550e8400-e29b-41d4-a716-446655440000`, `550e8400-e29b-41d4-a716-446655440001`
- **Test Categories**: `660e8400-e29b-41d4-a716-446655440000`, `660e8400-e29b-41d4-a716-446655440001`

### Helper Functions
- `GetTestUserID()` - Returns test user UUID
- `GetTestCategoryID()` - Returns test category UUID
- `AssertNoError()` - Helper for error assertions
- `AssertTimeClose()` - Helper for time comparisons

## Test Features

### Database Isolation
- Each test gets a fresh PostgreSQL container
- Tests run in parallel without interference
- Automatic cleanup after test completion

### Comprehensive Coverage
- **Happy path testing** - Normal operations
- **Error condition testing** - Invalid inputs, constraints
- **Edge case testing** - Empty arrays, null values
- **Integration testing** - Multi-step workflows

### Array Handling
- Tests verify proper handling of PostgreSQL arrays (tags, errors, warnings)
- Validates `pq.Array()` usage for array parameters
- Tests empty array scenarios

### Constraint Testing
- Foreign key constraint violations
- Unique constraint violations  
- NOT NULL constraint violations

### Performance Patterns
- Pagination testing with multiple pages
- Bulk operation testing
- Filtering and search testing

## Debugging Tests

### Verbose Output
```bash
go test ./internal/modules/cms/repo/ -v -run TestProgramRepo_Create
```

### Test with Race Detection
```bash
go test ./internal/modules/cms/repo/ -race -v
```

### Test Timeout
```bash
go test ./internal/modules/cms/repo/ -timeout 300s -v
```

### Container Logs
If tests fail, check Docker logs:
```bash
docker logs $(docker ps -q --filter ancestor=postgres:17-alpine)
```

## Continuous Integration

These tests are designed to run in CI environments:
- Self-contained with testcontainers
- No external database dependencies
- Deterministic test data
- Parallel execution safe

## Best Practices

1. **Test Isolation** - Each test creates its own data
2. **Cleanup** - Tests clean up containers automatically
3. **Descriptive Names** - Test names clearly indicate what's being tested
4. **Error Testing** - Both success and failure scenarios covered
5. **Integration Testing** - End-to-end workflows validated