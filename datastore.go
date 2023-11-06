package datastore

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

type GcpDatastore struct {
	ProjectId string
	Kind      string
}

type GcpDatastoreEntity struct {
	Key   *datastore.Key `datastore:"__key__"`
	Value string         `datastore:"value"`
}

func (gcpDatastore *GcpDatastore) Init() {
	gcpDatastore.ProjectId = os.Getenv("GCP_PROJECT_ID")
	gcpDatastore.Kind = os.Getenv("GCP_DATASTORE_KIND")
}

func (gcpDatastore *GcpDatastore) GetClient() *datastore.Client {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, gcpDatastore.ProjectId)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

func (gcpDatastore *GcpDatastore) GetKey(client *datastore.Client, key string) *datastore.Key {
	return datastore.NameKey(gcpDatastore.Kind, key, nil)
}

func (gcpDatastore *GcpDatastore) GetIncompleteKey(client *datastore.Client) *datastore.Key {
	return datastore.IncompleteKey(gcpDatastore.Kind, nil)
}

func (gcpDatastore *GcpDatastore) Get(key string) string {

	client := gcpDatastore.GetClient()
	datastoreKey := gcpDatastore.GetKey(client, key)

	var entity GcpDatastoreEntity
	err := client.Get(context.Background(), datastoreKey, &entity)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return entity.Value
}

func (gcpDatastore *GcpDatastore) Put(key string, value string) {

	client := gcpDatastore.GetClient()
	datastoreKey := gcpDatastore.GetKey(client, key)

	entity := GcpDatastoreEntity{
		Key:   datastoreKey,
		Value: value,
	}

	_, err := client.Put(context.Background(), datastoreKey, &entity)
	if err != nil {
		fmt.Println(err)
	}
}

// put custom entity
func (gcpDatastore *GcpDatastore) PutEntity(key string, entity interface{}) {
	client := gcpDatastore.GetClient()

	datastoreKey := gcpDatastore.GetIncompleteKey(client)
	if key != "" {
		datastoreKey = gcpDatastore.GetKey(client, key)
	}
	_, err := client.Put(context.Background(), datastoreKey, entity)
	if err != nil {
		fmt.Println(err)
	}
}

// get custom entity
func (gcpDatastore *GcpDatastore) GetEntity(key string, entity interface{}) {
	client := gcpDatastore.GetClient()
	datastoreKey := gcpDatastore.GetKey(client, key)

	err := client.Get(context.Background(), datastoreKey, entity)
	if err != nil {
		fmt.Println(err)
	}
}

func (gcpDatastore *GcpDatastore) Delete(key string) {

	client := gcpDatastore.GetClient()
	datastoreKey := gcpDatastore.GetKey(client, key)

	err := client.Delete(context.Background(), datastoreKey)
	if err != nil {
		fmt.Println(err)
	}
}

// get filterd list
func (gcpDatastore *GcpDatastore) FilteredList(entity interface{}, colume string, operation string, value string, limit int) (ret interface{}) {

	client := gcpDatastore.GetClient()

	query := datastore.NewQuery(gcpDatastore.Kind).
		FilterField(colume, operation, value).
		Limit(limit)

	it := client.Run(context.Background(), query)

	// ret is list of entity
	ret = []interface{}{}
	for {
		_, err := it.Next(entity)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(entity)
		ret = append(ret.([]interface{}), entity)

	}
	return ret
}

func (gcpDatastore *GcpDatastore) List() {

	client := gcpDatastore.GetClient()

	query := datastore.NewQuery(gcpDatastore.Kind)
	it := client.Run(context.Background(), query)

	for {
		var entity GcpDatastoreEntity
		_, err := it.Next(&entity)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(entity)
	}
}
