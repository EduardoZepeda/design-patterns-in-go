package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) CreateSingleConnection() {
	fmt.Println("Creating Singleton Database")
	time.Sleep(1 * time.Second)
	fmt.Println("Database created")
}

var db *Database
var lock sync.Mutex

func getDatabaseInstance() *Database {
	if db == nil {
		fmt.Println("Creading database connection")
		db = &Database{}
	} else {
		fmt.Println("DB Already Created")
	}
	return db
}

func main() {
	lock.Lock()
	defer lock.Unlock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()
}
