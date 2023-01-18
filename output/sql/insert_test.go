package sql

import (
	"context"
	"fmt"
	"testing"
)

type User struct {
	Name string `db:"name"`
	ID   int64  `db:"id"`
	// Role []uint `db:"role"`
	Hide bool  `db:"hide"`
	HHH  *bool `db:"hhh"`
	// NNN  []*int `db:"nnnn"`
}

func TestSQLInsert(t *testing.T) {
	for _, a := range []struct {
		source interface{}
	}{
		{&User{Name: "sonico"}},
		{[]User{{Name: "bbq"}, {Name: "ccv"}}},
	} {
		result, err := SQLInsert(context.Background(), a.source, "users", "user", "db")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(result))
	}
}
