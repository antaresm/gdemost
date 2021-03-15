package models

type Bridge struct {
	Id 		uint           `json:"id"`
	Name 	string           `json:"name"`
	NameEng string           `json:"name_eng"`
	Description 	string    `json:"description"`
	DescriptionEng 	string `json:"description_eng"`
	Divorces []Divorce       `json:"divorces" gorm:"foreignKey:BridgeId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Lat float32              `json:"lat"`
	Lng float32              `json:"lng"`
	PhotoCloseUrl 	string     `json:"photo_close_url"`
	PhotoOpenUrl 	string      `json:"photo_open_url"`
	Public 		bool       `json:"public"`
}

type Divorce struct {
	Id uint `json:"-"`
	BridgeId uint `json:"-"`
	Start string `json:"start"`
	End string `json:"end"`
}