package activityModel

type ActivityConfig struct{

}
// well marshaled struct
type ActivityConfig2 struct {
	Id        int         `json:"id"`
	PropAward []PropAward `json:"prop_id"`
}

type PropAward struct {
	PropId   int `json:"prop_id"`
	PropNum  int `json:"prop_num"`
	ExpireIn int `json:"expire_in"` // 小时
}
