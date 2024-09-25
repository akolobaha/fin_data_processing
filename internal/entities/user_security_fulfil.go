package entities

type UserSecurityFulfil struct {
	ID     uint    `gorm:"primaryKey;column:id" json:"id"`
	Ticker string  `gorm:"not null;column:ticker" json:"ticker"`
	UserId int     `gorm:"column:user_id" json:"user_id"`
	PE     float32 `gorm:"column:p_e_msfo_fulfil" json:"pe,omitempty"`
	PBv    float32 `gorm:"column:p_bv_msfo_fulfil" json:"pbv,omitempty"`
}
