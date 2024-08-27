package campaign

import "time"

type Campaign struct {
	ID               uint      `gorm:"primary_key" json:"id"`
	UserID           uint      `gorm:"column:user_id" json:"user_id"`
	Name             string    `gorm:"column:name" json:"name"`
	ShortDescription string    `gorm:"column:short_description" json:"short_description"`
	Description      string    `gorm:"column:long_description" json:"long_description"`
	Perks            string    `gorm:"column:benefit" json:"benefit"`
	BackerCount      int       `gorm:"column:backer_count" json:"backer_count"`
	GoalAmount       int       `gorm:"column:goal_amount" json:"goal_amount"`
	CurrentAmount    int       `gorm:"column:current_amount" json:"current_amount"`
	Slug             string    `gorm:"column:slug" json:"slug"`
	CreatedAt        time.Time `gorm:"column:created_time;autoCreateTime" json:"created_time"`
	UpdatedAt        time.Time `gorm:"column:updated_time;autoUpdateTime" json:"updated_time"`
}

type CampaignImage struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	CampaignID uint      `gorm:"column:campaign_id" json:"campaign_id"`
	FileName   string    `gorm:"column:filename" json:"filename"`
	IsPrimary  int       `gorm:"column:is_primary" json:"is_primary"`
	CreatedAt  time.Time `gorm:"column:created_time;autoCreateTime" json:"created_time"`
	UpdatedAt  time.Time `gorm:"column:updated_time;autoUpdateTime" json:"updated_time"`
}
