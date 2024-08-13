package experiences

import (
	"context"
	"theedashboard/ent"
	"theedashboard/ent/experience"
	"theedashboard/services/user"
	"theedashboard/utils"
	"time"
)

func FetchUserExperiences(client *ent.Client, userID uint) ([]*ent.Experience, error) {
	return client.Experience.Query().
		Where(experience.UserID(userID)).
		Order(ent.Desc(experience.FieldStartDate)).All(context.Background())
}

func FetchUserExperience(client *ent.Client, userID uint, experienceID uint) (*ent.Experience, error) {
	return client.Experience.Query().Where(experience.UserID(userID), experience.ID(experienceID)).Only(context.Background())
}

func CreateExperience(client *ent.Client, userID uint, experience Experience) (*ent.Experience, error) {
	tx, err := client.Tx(context.Background())
	if err != nil {
		return nil, err
	}

	userAvailable, err := user.FetchUserByID(client, userID)

	if err != nil {
		return nil, err
	}
	startDate, _ := utils.ConvertJsonDate(experience.StartDate)
	endDate, _ := utils.ConvertJsonDate(experience.EndDate)
	newExperience, err := tx.Experience.Create().
		SetCompanyName(experience.CompanyName).
		SetStartDate(startDate).SetEndDate(endDate).
		SetPosition(experience.Position).
		SetCurrentPlace(experience.CurrentPlace).
		SetUser(userAvailable).
		SetCreatedAt(time.Now()).Save(context.Background())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return newExperience, nil

}

func UpdateExperience(client *ent.Client, userID uint, expId uint, experience Experience) (*ent.Experience, error) {
	expToUpdate, err := FetchUserExperience(client, userID, expId)
	if err != nil {
		return nil, err
	}
	startDate, _ := utils.ConvertJsonDate(experience.StartDate)
	endDate, _ := utils.ConvertJsonDate(experience.EndDate)
	updateExp, err := client.Experience.UpdateOneID(expToUpdate.ID).
		SetCompanyName(experience.CompanyName).
		SetStartDate(startDate).
		SetEndDate(endDate).
		SetCurrentPlace(experience.CurrentPlace).
		SetPosition(experience.Position).
		SetUpdatedAt(time.Now()).Save(context.Background())

	if err != nil {
		return nil, err
	}

	return updateExp, nil

}
