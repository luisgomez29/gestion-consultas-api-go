package repositories

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
)

// PermissionRepository encapsula la lógica para acceder a los permisos.
type PermissionRepository interface {
	All() ([]*models.Permission, error)
	Get(codename string) (*models.Permission, error)
	Create(p *models.Permission) (*models.Permission, error)
	Update(codename string, p *models.Permission) (*models.Permission, error)
	Delete(codename string) (uint, error)
}

type permissionRepository struct {
	conn *pgxpool.Pool
}

// NewPermissionRepository crea un nuevo repositorio de permisos.
func NewPermissionRepository(db *pgxpool.Pool) PermissionRepository {
	return permissionRepository{conn: db}
}

func (r permissionRepository) All() ([]*models.Permission, error) {
	query := `SELECT id, name, codename, content_type_id FROM permissions`

	var perms []*models.Permission
	if err := pgxscan.Select(context.Background(), r.conn, &perms, query); err != nil {
		return nil, err
	}
	return perms, nil
}

func (r permissionRepository) Get(codename string) (*models.Permission, error) {
	panic("implement me")
}

func (r permissionRepository) Create(p *models.Permission) (*models.Permission, error) {
	panic("implement me")
}

func (r permissionRepository) Update(codename string, p *models.Permission) (*models.Permission, error) {
	panic("implement me")
}

func (r permissionRepository) Delete(codename string) (uint, error) {
	panic("implement me")
}
