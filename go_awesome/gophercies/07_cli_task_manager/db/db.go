package db

import (
	"fmt"
	"strconv"
	"time"

	"go.etcd.io/bbolt"
)

var bucketName = []byte("tasks")
var DBPath = "tasks.db"

func getDB() (*bbolt.DB, error) {
	db, err := bbolt.Open(DBPath, 0600, &bbolt.Options{
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

func AddTask(task string) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer func(db *bbolt.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	err = db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		nextID, err := b.NextSequence()
		if err != nil {
			fmt.Println("error closing db:", err)
			return err
		}
		return b.Put([]byte(strconv.FormatUint(nextID, 10)), []byte(task))
	})
	return err
}

func ListTasks() ([]string, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	defer func(db *bbolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("error closing db:", err)
			panic(err)
		}
	}(db)

	var tasks []string
	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.ForEach(func(k, v []byte) error {
			tasks = append(tasks, string(k)+": "+string(v))
			return nil
		})
	})
	return tasks, err
}

func CompleteTask(id int) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer func(db *bbolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("error closing db:", err)
			panic(err)
		}
	}(db)
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.Delete([]byte(strconv.Itoa(id)))
	})
}
