package repo

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	testDBName = "test_db"
	testUser   = "test_user"
	testPass   = "test_pass"
)

// TestHelper provides test utilities for repository tests
type TestHelper struct {
	Container testcontainers.Container
	Pool      *pgxpool.Pool
	DSN       string
}

// SetupTestDB creates a PostgreSQL testcontainer and returns a connection pool
func SetupTestDB(t *testing.T) *TestHelper {
	ctx := context.Background()

	// Create PostgreSQL container
	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:17-alpine"),
		postgres.WithDatabase(testDBName),
		postgres.WithUsername(testUser),
		postgres.WithPassword(testPass),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second)),
	)
	if err != nil {
		t.Fatalf("Failed to start PostgreSQL container: %v", err)
	}

	// Get connection string
	dsn, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to get connection string: %v", err)
	}

	// Create connection pool
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Fatalf("Failed to create connection pool: %v", err)
	}

	// Test connection
	if err := pool.Ping(ctx); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	helper := &TestHelper{
		Container: container,
		Pool:      pool,
		DSN:       dsn,
	}

	// Initialize schema
	if err := helper.InitSchema(ctx, t); err != nil {
		t.Fatalf("Failed to initialize schema: %v", err)
	}

	return helper
}

// InitSchema loads the test schema and initial data
func (h *TestHelper) InitSchema(ctx context.Context, t *testing.T) error {
	t.Helper()

	initSQLPath := filepath.Join("platform", "sql", "init.sql")
	sqlBytes, err := os.ReadFile("../../../../../" + initSQLPath)
	if err != nil {
		return fmt.Errorf("failed to read platform/sql/init.sql: %w", err)
	}

	// Execute the main schema SQL
	if _, err := h.Pool.Exec(ctx, string(sqlBytes)); err != nil {
		return fmt.Errorf("failed to execute platform/sql/init.sql: %w", err)
	}

	// Add test seed data
	seedSQL := `
-- Insert test users
INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES
('550e8400-e29b-41d4-a716-446655440000', 'Test User 1', 'test1@example.com', 'hashed_password_1', NOW(), NOW()),
('550e8400-e29b-41d4-a716-446655440001', 'Test User 2', 'test2@example.com', 'hashed_password_2', NOW(), NOW());

-- Insert test categories
INSERT INTO categories (id, name, description, type, created_at, updated_at, created_by, metadata) VALUES
('660e8400-e29b-41d4-a716-446655440000', 'Test Category 1', 'Test category description', 'CATEGORY_TYPE_PODCAST', NOW(), NOW(), '550e8400-e29b-41d4-a716-446655440000', '{"test": "data"}'::jsonb),
('660e8400-e29b-41d4-a716-446655440001', 'Test Category 2', 'Another test category', 'CATEGORY_TYPE_EDUCATIONAL', NOW(), NOW(), '550e8400-e29b-41d4-a716-446655440001', '{"test": "data2"}'::jsonb);
`

	// Execute the seed data SQL
	if _, err := h.Pool.Exec(ctx, seedSQL); err != nil {
		return fmt.Errorf("failed to execute seed data: %w", err)
	}

	return nil
}

// Cleanup terminates the test container and closes connections
func (h *TestHelper) Cleanup(ctx context.Context, t *testing.T) {
	if h.Pool != nil {
		h.Pool.Close()
	}
	if h.Container != nil {
		if err := h.Container.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	}
}

// TruncateTables cleans up data between tests
func (h *TestHelper) TruncateTables(ctx context.Context, tables ...string) error {
	for _, table := range tables {
		if _, err := h.Pool.Exec(ctx, fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table)); err != nil {
			return fmt.Errorf("failed to truncate table %s: %w", table, err)
		}
	}
	return nil
}

// ResetSequences resets auto-increment sequences
func (h *TestHelper) ResetSequences(ctx context.Context, sequences ...string) error {
	for _, seq := range sequences {
		if _, err := h.Pool.Exec(ctx, fmt.Sprintf("ALTER SEQUENCE %s RESTART WITH 1", seq)); err != nil {
			return fmt.Errorf("failed to reset sequence %s: %w", seq, err)
		}
	}
	return nil
}

// GetTestUserID returns a test user ID
func GetTestUserID() string {
	return "550e8400-e29b-41d4-a716-446655440000"
}

// GetTestUserID2 returns a second test user ID
func GetTestUserID2() string {
	return "550e8400-e29b-41d4-a716-446655440001"
}

// GetTestCategoryID returns a test category ID
func GetTestCategoryID() string {
	return "660e8400-e29b-41d4-a716-446655440000"
}

// GetTestCategoryID2 returns a second test category ID
func GetTestCategoryID2() string {
	return "660e8400-e29b-41d4-a716-446655440001"
}

// CompareTime compares two times with a tolerance for database precision
func CompareTime(t1, t2 time.Time, tolerance time.Duration) bool {
	diff := t1.Sub(t2)
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}

// AssertTimeClose asserts that two times are close within tolerance
func AssertTimeClose(t *testing.T, expected, actual time.Time, tolerance time.Duration, msg string) {
	if !CompareTime(expected, actual, tolerance) {
		t.Errorf("%s: expected time %v, got %v (diff: %v)", msg, expected, actual, expected.Sub(actual))
	}
}

// AssertNoError is a helper to assert no error occurred
func AssertNoError(t *testing.T, err error, msg string) {
	if err != nil {
		t.Fatalf("%s: unexpected error: %v", msg, err)
	}
}

// AssertError is a helper to assert an error occurred
func AssertError(t *testing.T, err error, msg string) {
	if err == nil {
		t.Fatalf("%s: expected error but got none", msg)
	}
}

// CountRows counts rows in a table with optional WHERE clause
func (h *TestHelper) CountRows(ctx context.Context, table, whereClause string, args ...interface{}) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)
	if whereClause != "" {
		query += " WHERE " + whereClause
	}

	var count int
	err := h.Pool.QueryRow(ctx, query, args...).Scan(&count)
	return count, err
}

// RowExists checks if a row exists with given conditions
func (h *TestHelper) RowExists(ctx context.Context, table, whereClause string, args ...interface{}) (bool, error) {
	count, err := h.CountRows(ctx, table, whereClause, args...)
	return count > 0, err
}
