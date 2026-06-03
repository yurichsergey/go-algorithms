package db

import (
	"time"

	"go.etcd.io/bbolt"
)

var bucketName = []byte("tasks")

func getDB() (*bbolt.DB, error) {
	db, err := bbolt.Open("tasks.db", 0600, &bbolt.Options{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})
	return db, err
}
