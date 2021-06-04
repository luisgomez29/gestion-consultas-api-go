package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/luisgomez29/gestion-consultas-api/api/models"
)

// GroupRepository encapsula la lógica para acceder a los permisos.
type GroupRepository interface {
	All() ([]*models.Permission, error)
	Get(codename string) (*models.Permission, error)
	Create(p *models.Permission) (*models.Permission, error)
	Update(codename string, p *models.Permission) (*models.Permission, error)
	Delete(codename string) (uint, error)
	SetUser(user, group uint) error
}

type groupRepository struct {
	conn *pgxpool.Pool
}

// NewGroupRepository crea un nuevo repositorio de grupos.
func NewGroupRepository(db *pgxpool.Pool) GroupRepository {
	return groupRepository{conn: db}
}

func (r groupRepository) All() ([]*models.Permission, error) {
	panic("implement me")
}

func (r groupRepository) Get(codename string) (*models.Permission, error) {
	panic("implement me")
}

func (r groupRepository) Create(p *models.Permission) (*models.Permission, error) {
	panic("implement me")
}

func (r groupRepository) Update(codename string, p *models.Permission) (*models.Permission, error) {
	panic("implement me")
}

func (r groupRepository) Delete(codename string) (uint, error) {
	panic("implement me")
}

func (r groupRepository) SetUser(user, group uint) error {
	query := `INSERT INTO user_groups(user_id, group_id) VALUES ($1, $2)`

	perm := new(models.Permission)
	if _, err := r.conn.Exec(context.Background(), query, user, group); err != nil {
		return perm.ValidatePgError(err)
	}
	return nil
}
