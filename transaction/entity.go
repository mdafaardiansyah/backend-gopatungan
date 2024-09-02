package transaction

import "time"

type Transaction struct {
	ID         int    `gorm:"primary_key" json:"id"`
	CampaignID int    `gorm:"column:campaign_id" json:"campaign_id"`
	UserID     int    `gorm:"column:user_id" json:"user_id"`
	Amount     int    `gorm:"column:amount" json:"amount"`
	Status     string `gorm:"column:status" json:"status"`
	Code       string `gorm:"column:code" json:"code"`
	//PaymentURL string    `gorm:"column:payment_url" json:"payment_url"`
	CreatedAt time.Time `gorm:"column:created_time;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_time;autoUpdateTime"`
}
