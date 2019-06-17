package propModel

type Prop struct {
	Id         int    `gorm:"column:id;default:" json:"id" form:"id"`
	PropName   string `gorm:"column:prop_name;default:" json:"prop_name" form:"prop_name"`
	CanPresent int    `gorm:"column:can_present;default:" json:"can_present" form:"can_present"`
	CanDestroy int    `gorm:"column:can_destroy;default:" json:"can_destroy" form:"can_destroy"`
	AutoUsed   int    `gorm:"column:auto_used;default:" json:"auto_used" form:"auto_used"`
	ImageUrl   string `gorm:"column:image_url;default:" json:"image_url" form:"image_url"`
}

func (o Prop) TableName() string {
	return "prop"
}
