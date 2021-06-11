package responses

import "github.com/luisgomez29/gestion-consultas-api/api/models"

// UserResponse lists the fields to return for the user types.
func UserResponse(role string, user *models.User) *models.User {
	if role != models.UserAdmin.String() {
		user.IsActive = false
		user.IsSuperuser = false
		user.IsStaff = false
	}
	user.Password = ""
	return user
}

// UserManyResponse lists the fields to return for the slice user types.
func UserManyResponse(role string, users []*models.User) []*models.User {
	for i, user := range users {
		users[i] = UserResponse(role, user)
	}
	return users
}
