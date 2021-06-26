package user_test

import (
	"fmt"
	"testing"
	"github.com/mixnote/mixnote-api-go/src/core/models/user"
)

func TestUser(t *testing.T) {
	u := user.New()
	u.FName = "jh"
	fmt.Println(u.FName)
}