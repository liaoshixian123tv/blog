package model

// Model 基礎資料
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"createdBy"`
	ModifiedBy string `json:"modifiedBy"`
	CreatedOn  int64  `json:"createdOn"`
	ModifiedOn int64  `json:"modifiedOn"`
	DeleteOn   int64  `json:"deleteOn"`
	IsDelete   bool   `json:"isDelete"`
}

// Tag 標籤
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"stqate"`
}

// Article 文章
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"coverImageUrl"`
	State         uint8  `json:"state"`
}

// ArtcleTag 文章標籤
type ArtcleTag struct {
	*Model
	TagID  int64 `json:"tagID"`
	Artcle int64 `json:"artcle"`
}

// TableName 標籤資料表表名
func (t Tag) TableName() string {
	return "blog_tag"
}

// TableName 標籤資料表表名
func (t Article) TableName() string {
	return "blog_artcle"
}

// TableName 標籤資料表表名
func (t ArtcleTag) TableName() string {
	return "blog_artcle_tag"
}
