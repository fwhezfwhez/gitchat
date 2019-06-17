package main

// 用户进度
type UserInviteActivity struct {
	ActivityId  int    `json:"activity_id"`             // 活动号
	DateVersion string `json:"date_version"`            // 进度版本
	UserId      int    `json:"user_id"`                 // 用户号

	InviteWho []int `json:"invite_who"`                 // 邀请了谁
	InvitedBy []int `json:"invited_by"`                 // 被哪些人邀请了
}

