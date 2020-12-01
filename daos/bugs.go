package daos

import (
	"bug-tracker/backend/models"
	"log"
)

//FindAllBugs find all in db
func FindAllBugs() *[]models.Bug {
	bugs := []models.Bug{}
	if err := GetConn().Find(&bugs).Error; err != nil {
		log.Println("Error when finding bugs")
	}
	return &bugs
}

//CreateBug Create a bug
func CreateBug(b *models.Bug) *models.Bug {
	b.Status = "New"
	if err := GetConn().Create(b).Error; err != nil {
		log.Println("Error when inserting new bug")
	}
	return b
}

//UpdateBug Update a bug
func UpdateBug(id int32, b *models.Bug) *models.Bug {
	uB := &models.Bug{ID: id}
	if err := GetConn().Model(uB).Updates(b).Error; err != nil {
		log.Println("Error when updating new bug")
	}
	return uB
}

//DeleteBug Delete a bug
func DeleteBug(id int32) bool {
	if err := GetConn().Delete(&models.Bug{}, id).Error; err != nil {
		log.Println("Error when inserting new bug")
		return false
	}
	return true
}

//FindBug Get a bug
func FindBug(id int32) *models.Bug {
	b := &models.Bug{}
	if err := GetConn().Find(&b, id).Error; err != nil {
		log.Println("Error when getting bug by id")
		return nil
	}
	return b
}

//UpdateBugStatus Update the bug's status in db
func UpdateBugStatus(id int32, status string) bool {
	if err := GetConn().Raw("UPDATE bugs SET status = ? where id = ?", status, id).Error; err != nil {
		log.Println("Error when updating bug status")
		return false
	}
	return true
}
