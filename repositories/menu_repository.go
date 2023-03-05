package repositories

import "fiber-muscles/models"

func CreateMenu(
	name string,
	description string,
	target string,
	weight float64,
	isJoint bool,
	link string,
	userID string,
) (m models.Menu, err error) {
	m = models.Menu{
		Name:        name,
		Description: description,
		Target:      target,
		Weight:      weight,
		IsJoint:     isJoint,
		Link:        link,
		UserID:      userID,
	}

	if err = models.Psql.Create(&m).Error; err != nil {
		return
	}

	if err = models.Psql.Preload("User").First(&m).Error; err != nil {
		return
	}

	return
}

func GetAllMenuByID(userID string) (m []models.Menu, err error) {
	if err = models.Psql.Where("user_id = ?", userID).Preload("User").Find(&m).Error; err != nil {
		return
	}

	return
}
