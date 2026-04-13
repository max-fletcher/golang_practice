package formatters

import (
	"dummy/internal/database"
	"time"

	"github.com/google/uuid"
)

// NOTE: Exporting type fields
// Turns out any type that has fields that are not exported as pascal-case(or at least starts with a capital case) is
// not exported. So if you used id instead of ID, the returned struct(and consequently the JSON) will be missing that field.
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NOTE: Exporting functions from a folder(a.k.a a package as different folders are each considered a package in go)
// Turns out any exported function from a folder/package needs to be pascal-case(or at least starts with a capital case)
// else, the function is not made available when imported and used. This is why the convention for writing golang often
// is to use pascal-case.

// This is a function that converts struct keys from pascal-case to use camel-case. This is so we get camel-cased JSON.
// Remember that database.go has the type for the user being fetched from database and that is being passed here.
func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}
