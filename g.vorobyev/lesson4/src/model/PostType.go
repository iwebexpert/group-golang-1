package model

type Post struct {
	PostID   uint   `gorm:"AUTO_INCREMENT;primary_key"`
	Title    string `sql:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	PostData string `sql:"type:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
}
