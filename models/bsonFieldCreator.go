package models

import "fmt"

func BsonFieldCreator(name string) string {
	return fmt.Sprintf(`bson:"%s" json:"%s,omitempty"`, name, name)
}