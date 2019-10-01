package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	pb "github.com/syuni/pact-sample/pb/addressbook"
)

const filename = "addressbook.db"

// DB data
var addressbook pb.AddressBook // nolint

func loadDB(filePath string) (*[]byte, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

// Load is read DB file and store data
func Load() error {
	cwd, _ := os.Getwd()
	db, err := loadDB(filepath.Join(cwd, "provider", "store", filename))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(*db, &addressbook); err != nil {
		return err
	}

	log.Println("DB file loaded successfully")

	return nil
}

// GetPeople returns all people
func GetPeople() []*pb.Person {
	return addressbook.GetPeople()
}

// FetchByName returns a person who has specified name
func FetchByID(id int32) (*pb.Person, error) {
	for _, people := range addressbook.GetPeople() {
		if people.GetId() == id {
			return people, nil
		}
	}
	return nil, fmt.Errorf("can not find person. Given id: %d", id)
}
