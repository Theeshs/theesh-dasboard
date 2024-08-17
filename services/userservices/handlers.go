package userservices

import (
	"context"
	"theedashboard/ent"
	"theedashboard/services/user"
	"time"
)
import "theedashboard/ent/userservice"

func FetchUsersServices(client *ent.Client, userID uint) ([]*ent.UserService, error) {
	return client.UserService.Query().
		Where(userservice.UserID(userID)).
		All(context.Background())
}

func FetchUserService(client *ent.Client, userID uint, serviceID uint) (*ent.UserService, error) {
	return client.UserService.Query().
		Where(userservice.UserID(userID), userservice.ID(serviceID)).
		Only(context.Background())
}

func SaveUserService(client *ent.Client, userID uint, service UserService) (*ent.UserService, error) {
	tx, err := client.Tx(context.Background())
	if err != nil {
		return nil, err
	}
	userAvailable, err := user.FetchUserByID(client, userID)
	if err != nil {
		return nil, err
	}

	newService, err := tx.UserService.Create().
		SetUser(userAvailable).
		SetCreatedAt(time.Now()).
		SetServiceName(service.ServiceName).
		SetServiceDescription(service.ServiceDescription).
		SetServiceIcon(service.ServiceIcon).
		Save(context.Background())

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return newService, nil
}

func EditUserService(client *ent.Client, userID uint, serviceID uint, service UserService) (*ent.UserService, error) {
	serviceForUpdate, err := FetchUserService(client, userID, serviceID)
	if err != nil {
		return nil, err
	}

	updatedService, err := client.UserService.UpdateOneID(serviceForUpdate.ID).
		SetServiceName(service.ServiceName).
		SetServiceDescription(service.ServiceDescription).
		SetServiceIcon(service.ServiceIcon).
		SetUpdatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return updatedService, nil

}
