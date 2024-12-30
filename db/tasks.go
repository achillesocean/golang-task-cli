package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks") //what the hell is 'tasks' in there?
																// I think its the name of the bucket
var db *bolt.DB//pay attention to how this package-level variable is gonna be used.

type Task struct {
	Key int
	Value string
}

// initialize databse
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error { //are we returning a function here, because it's getting defined, not called. also, how are we passing in a function into the function?
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
		// what gets returned if the callback function doesn't return an error
	})
}

func CreateTask(task string) (int, error) {
	// we will be using the db.update function because it and what we're doing is a write-transaction
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}


func ListTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next(){
			tasks = append(tasks, Task{
				Key: btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}