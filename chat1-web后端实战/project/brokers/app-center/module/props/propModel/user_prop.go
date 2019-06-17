package propModel

type UserProp struct {
	Id        int    `gorm:"column:id;default:" json:"id" form:"id"`
	UserId    int    `gorm:"column:user_id;default:" json:"user_id" form:"user_id"`
	PropId    int    `gorm:"column:prop_id;default:" json:"prop_id" form:"prop_id"`
	ExpireIn  int    `gorm:"column:expire_in;default:" json:"expire_in" form:"expire_in"`
	PropNum   int    `gorm:"column:prop_num;default:" json:"prop_num" form:"prop_num"`
	PropTitle string `gorm:"column:prop_title;default:" json:"prop_title" form:"prop_title"`
}

func (o UserProp) TableName() string {
	return "user_prop"
}
