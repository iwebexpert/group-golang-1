package repository

import (
	"testing"

	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/blog"
	"github.com/Toringol/group-golang-1/tree/master/s_shepelev/blog/app/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSelectPostByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := blog.NewMockRepository(ctrl)

	id := "skjbsdlhjsd;hjsdlh;sd"
	returnPost := &model.Post{
		ID:          id,
		Title:       "New post",
		Description: "hello",
		Author:      "Toringol",
	}

	db.EXPECT().SelectPostByID(id).Return(returnPost, nil)

	post, err := db.SelectPostByID(id)

	assert.Equal(t, post, returnPost)
	assert.NoError(t, err)

}

func BenchmarkSelectPostByID(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	db := blog.NewMockRepository(ctrl)

	id := "skjbsdlhjsd;hjsdlh;sd"
	returnPost := &model.Post{
		ID:          id,
		Title:       "New post",
		Description: "hello",
		Author:      "Toringol",
	}

	db.EXPECT().SelectPostByID(id).Return(returnPost, nil).AnyTimes()

	for i := 0; i < b.N; i++ {
		post, err := db.SelectPostByID(id)

		assert.Equal(b, post, returnPost)
		assert.NoError(b, err)
	}
}
