package table

type Version struct {
	Id      int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Version string `gorm:"column:version;type:varchar(45);NOT NULL" json:"version"`
	Url     string `gorm:"column:url;type:varchar(45);NOT NULL" json:"url"`
}

func (m *Version) TableName() string {
	return "version"
}