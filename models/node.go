package models

type Node struct {
	Model
	Name     string                 `gorm:"column=name" json:"name"`
	OutBound map[string]interface{} `gorm:"column=outbound" json:"outBound"`
	Enable   bool                   `gorm:"column=enable" json:"enable"`
}
