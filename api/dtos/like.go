package dtos

// Like LikeGORMモデル
type Like struct {
	ArticleID string  `gorm:"primaryKey"`
	Article   Article `gorm:"foreignKey:ArticleID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID    string  `gorm:"primaryKey"`
	User      User    `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
