package datastore

import (
	"context"
	"fmt"
	"log"
	"os"
)

import (
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

func (gcpDatastore *GcpDatastore) Delete(key string) {

	client := gcpDatastore.GetClient()
	datastoreKey := gcpDatastore.GetKey(client, key)

	err := client.Delete(context.Background(), datastoreKey)
	if err != nil {
		fmt.Println(err)
	}
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
