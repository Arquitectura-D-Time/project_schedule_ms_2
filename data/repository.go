package data

import (
	"context"
	"project_schedule_ms/model"
)

type Horario interface {
	Fetch(ctx context.Context, num int64) ([]*model.Horario, error)
	GetByID(ctx context.Context, id int64) ([]*model.Horario, error)
	GetByIDTutor(ctx context.Context, idtutor int64) ([]*model.Horario, error)
	GetByNombre(ctx context.Context, nombre string) ([]*model.Horario, error)
	GetByFecha(ctx context.Context, fecha string) ([]*model.Horario, error)
	GetByHora(ctx context.Context, hora string) ([]*model.Horario, error)
	Create(ctx context.Context, p *model.Horario) (int64, error)
	Update(ctx context.Context, p *model.Horario) (*model.Horario, error)
	Delete(ctx context.Context, id int64, id2 int64) (bool, error)
}

type Agendadas interface {
	Fetch(ctx context.Context, num int64) ([]*model.Agendadas, error)
	GetByID(ctx context.Context, id int64) ([]*model.Agendadas, error)
	GetByID2(ctx context.Context, id2 int64) ([]*model.Agendadas, error)
	Create(ctx context.Context, p *model.Agendadas) (int64, error)
	Update(ctx context.Context, p *model.Agendadas) (*model.Agendadas, error)
	Delete(ctx context.Context, id int64, id2 int64) (bool, error)
}
