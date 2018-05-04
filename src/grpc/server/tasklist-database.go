package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"grpc/api"
	"log"

	"github.com/boltdb/bolt"
)

type Database struct {
	db         *bolt.DB
	bucketName string
}

func (d *Database) Open() error {

	d.bucketName = "TASKS"
	db, err := bolt.Open("tasklist.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	bk, err := tx.CreateBucketIfNotExists([]byte(d.bucketName))
	if err != nil {
		return fmt.Errorf("failed to create bucket: %v", err)
	}

	_ = bk
	if err := tx.Commit(); err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) AddTask(t *taskservice.Task) error {
	tx, err := d.db.Begin(true)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	b := tx.Bucket([]byte(d.bucketName))

	t.Id, err = b.NextSequence()

	key, value, err := taskToKeyValuePair(t)
	if err != nil {
		return err
	}

	err = b.Put(key, value)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (d *Database) GetTaskList() (*taskservice.TaskList, error) {

	tl := taskservice.TaskList{}

	tx, err := d.db.Begin(true)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()
	b := tx.Bucket([]byte(d.bucketName))

	c := b.Cursor()

	for key, value := c.First(); key != nil; key, value = c.Next() {

		t, err := valueToTask(value)
		if err != nil {
			return nil, err
		}
		tl.Task = append(tl.Task, &t)
	}

	return &tl, err
}
func (d *Database) UpdateTask(taskservice.Task) error {
	return nil

}
func (d *Database) DeleteTask(t *taskservice.Task) error {

	tx, err := d.db.Begin(true)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	b := tx.Bucket([]byte(d.bucketName))

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, t.Id)

	err = b.Delete(key)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func taskToKeyValuePair(t *taskservice.Task) ([]byte, []byte, error) {

	//var key bytes.Buffer
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, t.Id)

	var value bytes.Buffer
	enc := gob.NewEncoder(&value)

	err := enc.Encode(t)

	if err != nil {
		return nil, nil, err
	}

	return key, value.Bytes(), nil
}

func valueToTask(value []byte) (taskservice.Task, error) {

	dec := gob.NewDecoder(bytes.NewBuffer(value))

	var t taskservice.Task

	err := dec.Decode(&t)
	if err != nil {
		return t, err
	}

	return t, nil
}
