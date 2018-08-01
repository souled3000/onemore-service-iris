package models

import (
	"time"
)

type Version struct {
	Type        int       `json:"type" xorm:"int pk"`
	AppType     int       `json:"app_type"  xorm:"int pk"`
	HtmlVersion string    `json:"html_version" xorm:"varchar(10)"`
	HtmlUrl     string    `json:"html_url" xorm:"varchar(255)"`
	CreatedAt   time.Time `json:"created_at" xorm:"created"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"updated"`
}
