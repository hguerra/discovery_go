package usecase

import "github.com/hguerra/discovery_go/modules/web/routerchi/internal/domain/user"

func FindUsers() ([]user.User, error) {
	users := []user.User{
		{
			ID:        1,
			FirstName: "Heitor",
			LastName:  "Carneiro",
		},
	}
	return users, nil
}
