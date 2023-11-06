package datastore_test

import (
	"testing"

	"github.com/hojin-kr/datastore"
)

func TestPut(t *testing.T) {
	datastore := datastore.GcpDatastore{}
	datastore.Init()
	datastore.Put("test2", "test")
	value := datastore.Get("test2")
	if value != "test" {
		t.Errorf("Expected %s, got %s", "test", value)
	}
}

func TestPutCustomEntity(t *testing.T) {
	datastore := datastore.GcpDatastore{}
	datastore.Init()
	type TestEntity struct {
		Test  string `datastore:"test"`
		Test2 string `datastore:"test2"`
		Test3 string `datastore:"test3"`
	}
	entity := TestEntity{
		Test:  "test",
		Test2: "test2",
		Test3: "test5",
	}
	uuid := "testtesttest"
	datastore.PutEntity(uuid, &entity)
	var entity2 TestEntity
	datastore.GetEntity(uuid, &entity2)
	if entity2.Test != "test" {
		t.Errorf("Expected %s, got %s", "test", entity2.Test)
	}
	if entity2.Test2 != "test2" {
		t.Errorf("Expected %s, got %s", "test2", entity2.Test2)
	}
	if entity2.Test3 != "test5" {
		t.Errorf("Expected %s, got %s", "test5", entity2.Test3)
	}

}

// filterdlist
func TestFilteredList(t *testing.T) {
	datastore := datastore.GcpDatastore{}
	datastore.Init()
	type Friend struct {
		UUID_Send    string `json:"uuid_send"`
		UUID_Receive string `json:"uuid_receive"`
		Status       string `json:"status"`
		Friend_ID    string `json:"friend_id"`
	}
	var friends Friend
	ret := datastore.FilteredList(&friends, "Status", "=", "pending", 100)
	t.Logf("%+v\n", ret)
}
