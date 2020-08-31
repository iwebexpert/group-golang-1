package tests

import (
	"dbo"
	"fmt"
	"model"
	"testing"
)

func TestConnectionDb(t *testing.T) {
	var Posts []model.Post

	dbo.GetConnectionDb()
	dbo.DB.Find(&Posts)

	if len(Posts) > 0 {
		fmt.Println("There is posts")
	} else {
		fmt.Println("There is no posts, new database")
	}

	_ = dbo.DB.Close()
}
