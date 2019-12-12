package models

import (
	"time"

	log "github.com/sirupsen/logrus"

	"gin-example-structure/db"
)

type MatchinfoSummary struct {
	MatchID   string    `json:"match_id" gorm:"primary_key"`
	StartTime time.Time `json:"start_time"`
}

func (MatchinfoSummary) TableName() string {
	return "matchinfo.summary"
}
func (h MatchinfoSummary) GetByID(id string) (*MatchinfoSummary, error) {
	db := db.GetDB()
	var info MatchinfoSummary
	err := db.Where("match_id = ?", id).Find(&info).Error
	if err != nil {
		log.WithField("match_id", id).WithError(err).Error("MatchinfoSummary GetByID")
		return nil, err
	}
	return &info, nil
}
