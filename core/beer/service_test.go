package beer_test

import (
	"database/sql"
	"testing"

	"github.com/gasorey/go-api-minetto/core/beer"
	_ "github.com/mattn/go-sqlite3"
)

func TestStore(t *testing.T) {
	b := &beer.Beer{
		ID:    1,
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	if err != nil {
		t.Fatalf("Connection with database failed %s", err.Error())
	}

	err = clearDB(db)
	if err != nil {
		t.Fatalf("Could not clear database: %s", err.Error())
	}
	defer db.Close()
	service := beer.NewService(db)
	err = service.Store(b)
	if err != nil {
		t.Fatalf("Could not save beer in database: %s", err.Error())
	}
	saved, err := service.Get(1)
	if err != nil {
		t.Fatalf("Could not get beer in database %s", err.Error())
	}
	if saved.ID != 1 {
		t.Fatalf("Invalid data, expected: %d, received:  %d", 1, saved.ID)
	}
}

func clearDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from beer")
	tx.Commit()
	return err
}
