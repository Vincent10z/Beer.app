package models

type Plan struct {
	ID          string `gorm:"primaryKey;type:text" json:"id"`
	Name        string `gorm:"size:255;not null" json:"name"`
	Description string `gorm:"size:1024" json:"description"`
	Price       int    `gorm:"not null" json:"price"`
}

const (
	BasicPlan      = "pln_rkflbh60iklsujcm8it7"
	ProPlan        = "pln_9bkyi2anud7zpelnlm18"
	EnthusiastPlan = "pln_0glmrsmrx4xe29k53zya"
)
