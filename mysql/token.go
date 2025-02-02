package mysql

import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

// InsertToken insert token
func InsertToken(token *table.Token) error {
	if err := db.DB.Create(token).Error; err != nil {
		return err
	}
	return nil
}

// GetTokenByUserID get token by user id
func GetTokenByUserID(userID int64) (*table.Token, error) {
	var token table.Token
	if err := db.DB.Where("user_id = ?", userID).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

// GetToken get token
func GetToken(token string) (*table.Token, error) {
	var tokenInfo table.Token
	if err := db.DB.Where("token = ?", token).First(&tokenInfo).Error; err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}

// UpdateToken update token
func UpdateToken(token *table.Token) error {
	if err := db.DB.Save(token).Error; err != nil {
		return err
	}
	return nil
}
