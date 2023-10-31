package datastore_test

import (
	"testing"

	"github.com/hojin-kr/datastore"
)

func TestPut(t *testing.T) {
	datastore := datastore.GcpDatastore{}
	datastore.Init()
	datastore.Put("test", "test")
	value := datastore.Get("test")
	if value != "test" {
		t.Errorf("Expected %s, got %s", "test", value)
	}
}
