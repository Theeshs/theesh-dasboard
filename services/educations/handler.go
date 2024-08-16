package educations

import (
	"context"
	"theedashboard/ent"
	"theedashboard/ent/education"
	"theedashboard/services/user"
	"theedashboard/utils"
	"time"
)

func FetchUserEducations(client *ent.Client, userID uint) ([]*ent.Education, error) {
	return client.Education.Query().
		Where(education.UserID(userID)).
		Order(ent.Desc(education.FieldStartDate)).All(context.Background())
}

func FetchUserEducation(client *ent.Client, userID uint, educationID uint) (*ent.Education, error) {
	return client.Education.Query().Where(education.UserID(userID), education.ID(educationID)).Only(context.Background())
}

func CreateEducation(client *ent.Client, userID uint, eduData Education) (*ent.Education, error) {
	tx, err := client.Tx(context.Background())
	if err != nil {
		return nil, err
	}

	userAvailable, err := user.FetchUserByID(client, userID)
	if err != nil {
		return nil, err
	}

	startDate, _ := utils.ConvertJsonDate(eduData.StartDate)
	endDate, _ := utils.ConvertJsonDate(eduData.EndDate)

	newEducation, err := tx.Education.Create().
		SetInstitueName(eduData.InstitueName).
		SetStartDate(startDate).SetEndDate(endDate).
		SetCurrentyStudying(eduData.CurrentlyStudying).
		SetUser(userAvailable).SetDegreeType(eduData.DegreeType).
		SetDescription(eduData.Description).
		SetCreatedAt(time.Now()).
		SetAreaOfStudy(eduData.AreaOfStudy).
		SetModeOfStudy(eduData.ModeOfStudy).Save(context.Background())

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return newEducation, nil
}

func UpdateEducation(client *ent.Client, userID uint, educationID uint, eduData Education) (*ent.Education, error) {
	eduToUpdate, err := FetchUserEducation(client, userID, educationID)

	if err != nil {
		return nil, err
	}
	startDate, _ := utils.ConvertJsonDate(eduData.StartDate)
	endDate, _ := utils.ConvertJsonDate(eduData.EndDate)

	updatedEdu, err := client.Education.UpdateOneID(eduToUpdate.ID).
		SetInstitueName(eduData.InstitueName).
		SetStartDate(startDate).SetEndDate(endDate).
		SetCurrentyStudying(eduData.CurrentlyStudying).
		SetDegreeType(eduData.DegreeType).
		SetDescription(eduData.Description).
		SetCreatedAt(time.Now()).
		SetAreaOfStudy(eduData.AreaOfStudy).
		SetModeOfStudy(eduData.ModeOfStudy).Save(context.Background())

	if err != nil {
		return nil, err
	}

	return updatedEdu, nil
}
