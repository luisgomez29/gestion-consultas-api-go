package repositories

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
)

type AuthRepository interface {
	// UserLoggedIn get the user who is logged in from UserRepository.Get.
	UserLoggedIn(username string) *models.User

	// UserPermissions get the permissions that the user has in `user_permissions`.
	UserPermissions(username string) ([]*models.Permission, error)

	// UserGroupPermissions get the permissions that this user has through their groups.
	UserGroupPermissions(username string) ([]*models.Permission, error)

	// AllPermissions get all the permissions defined in the `permissions` table.
	AllPermissions() ([]*models.Permission, error)
}

type authRepository struct {
	conn       *pgxpool.Pool
	permission PermissionRepository
	user       UserRepository
}

// NewAuthRepository creates a new auth repository.
func NewAuthRepository(db *pgxpool.Pool, p PermissionRepository, u UserRepository) AuthRepository {
	return authRepository{conn: db, permission: p, user: u}
}

func (r authRepository) UserLoggedIn(username string) *models.User {
	user, _ := r.user.Get(username)
	return user
}

func (r authRepository) UserPermissions(username string) ([]*models.Permission, error) {
	query := `
			SELECT DISTINCT p.id, p.name, p.content_type_id, p.codename
			FROM permissions as p
				 INNER JOIN user_permissions up on p.id = up.permission_id
				 INNER JOIN users u on up.user_id = u.id
			WHERE username = $1
			ORDER BY p.id`

	var perms []*models.Permission
	if err := pgxscan.Select(context.Background(), r.conn, &perms, query, username); err != nil {
		return nil, err
	}
	return perms, nil
}

func (r authRepository) UserGroupPermissions(username string) ([]*models.Permission, error) {
	query := `
			SELECT p.id, p.name, p.codename, p.content_type_id
			FROM permissions AS p
				 INNER JOIN group_permissions gp on p.id = gp.permission_id
				 INNER JOIN groups g on g.id = gp.group_id
				 INNER JOIN user_groups ug on g.id = ug.group_id
				 INNER JOIN users u on u.id = ug.user_id
			WHERE u.username = $1
			ORDER BY p.id`

	var perms []*models.Permission
	if err := pgxscan.Select(context.Background(), r.conn, &perms, query, username); err != nil {
		return nil, err
	}
	return perms, nil
}

func (r authRepository) AllPermissions() ([]*models.Permission, error) {
	return r.permission.All()
}
