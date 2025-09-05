package database_test

import (
	"E-matBackend/pkg/database"
	"context"
	"testing"
	"time"
)

func TestMySQLConnection(t *testing.T) {
	db := database.MySQLConnection()
	defer db.Close()

	// Simple query to verify connection
	var version string
	err := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}

	t.Logf("MySQL version: %s", version)
}

func TestRedisConnection(t *testing.T) {
	rdb := database.Redisconnection()
	defer rdb.Close()

	ctx := context.Background()
	err := rdb.Set(ctx, "test_key", "hello", 10*time.Second).Err()
	if err != nil {
		t.Fatalf("Redis set failed: %v", err)
	}

	val, err := rdb.Get(ctx, "test_key").Result()
	if err != nil {
		t.Fatalf("Redis get failed: %v", err)
	}

	if val != "hello" {
		t.Errorf("Expected 'hello', got '%s'", val)
	}
}
