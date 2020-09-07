package models

import "testing"

// PostTest - структура для тестирования
type PostTest struct {
	Id     string
	Header string
	Text   string
}

func TestGetOne(t *testing.T) {
	testCases := []PostTest{
		{Id: "5f50f46feadcf22a6827f7d7", Header: "my first post", Text: "some text 1"},
		{Id: "5f50fb50eadcf22a6827f7d8", Header: "my second post", Text: "some text 2"},
		{Id: "5f523faefe21305c8992ec1f", Header: "мой третий пост", Text: "какой-то текст3"},
	}

	for _, testCase := range testCases {
		result, _ := GetOne(Ctx, Db, testCase.Id)
		if result.Header != testCase.Header && result.Text != testCase.Text {
			t.Errorf("Получили Header: %s и Text: %s, а ожидалось Header: %s и Text: %s", result.Header, result.Text, testCase.Header, testCase.Text)
		}
	}
}

var p *Post

func BenchmarkGetCollectionName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = p.GetCollectionName()
	}
}

func BenchmarkGetAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GetAll(Ctx, Db)
	}
}
