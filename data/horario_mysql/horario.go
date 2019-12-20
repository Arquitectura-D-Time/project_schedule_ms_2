package horario_mysql

import (
	"context"
	"database/sql"
	"fmt"
	repo "project_schedule_ms/data"
	model "project_schedule_ms/model"
)

func NewSQLHorario(Conn *sql.DB) repo.Horario {
	return &mysqlHorario{
		Conn: Conn,
	}
}

type mysqlHorario struct {
	Conn *sql.DB
}

func (m *mysqlHorario) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.Horario, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*model.Horario, 0)
	for rows.Next() {
		data := new(model.Horario)

		err := rows.Scan(
			&data.IDtutoria,
			&data.IDtutor,
			&data.NombreMateria,
			&data.Fecha,
			&data.HoraInicio,
			&data.HoraFinal,
			&data.Cupos,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlHorario) Fetch(ctx context.Context, num int64) ([]*model.Horario, error) {
	query := "Select IDtutoria, IDtutor, NombreMateria, Fecha, HoraInicio, HoraFinal, Cupos From Horario"

	return m.fetch(ctx, query)
}

func (m *mysqlHorario) GetByID(ctx context.Context, IDtutoria int64) ([]*model.Horario, error) {
	query := "Select IDtutoria, IDtutor, NombreMateria, Fecha, HoraInicio, HoraFinal, Cupos From Horario where IDtutoria=?"
	rows, err := m.fetch(ctx, query, IDtutoria)
	if err != nil {
		return nil, err
	} else {
		return rows, nil
	}
}

func (m *mysqlHorario) GetByIDTutor(ctx context.Context, IDtutor int64) ([]*model.Horario, error) {
	query := "Select IDtutoria, IDtutor, NombreMateria, Fecha, HoraInicio, HoraFinal, Cupos From Horario where IDtutor=?"
	rows, err := m.fetch(ctx, query, IDtutor)
	if err != nil {
		return nil, err
	} else {
		return rows, nil
	}
}

func (m *mysqlHorario) GetByNombre(ctx context.Context, NombreMateria string) ([]*model.Horario, error) {
	fmt.Println(NombreMateria)
	query := "Select IDtutoria, IDtutor, NombreMateria, Fecha, HoraInicio, HoraFinal, Cupos From Horario where NombreMateria like '%?%'"
	fmt.Println(query)
	rows, err := m.fetch(ctx, query, NombreMateria)
	if err != nil {
		return nil, err
	} else {
		return rows, nil
	}
}

func (m *mysqlHorario) GetByFecha(ctx context.Context, Fecha string) ([]*model.Horario, error) {
	query := "Select IDtutoria, IDtutor, NombreMateria, Fecha, HoraInicio, HoraFinal, Cupos From Horario where Fecha like '%?%'"
	rows, err := m.fetch(ctx, query, Fecha)
	if err != nil {
		return nil, err
	} else {
		return rows, nil
	}
}

func (m *mysqlHorario) GetByHora(ctx context.Context, HoraInicio string) ([]*model.Horario, error) {
	query := "Select IDtutoria, IDtutor, NombreMateria, Fecha, HoraInicio, HoraFinal, Cupos From Horario where HoraInicio like '%?%'"
	rows, err := m.fetch(ctx, query, HoraInicio)
	if err != nil {
		return nil, err
	} else {
		return rows, nil
	}
}

func (m *mysqlHorario) Create(ctx context.Context, p *model.Horario) (int64, error) {
	query := "Insert Horario SET IDtutoria=?, IDtutor=?, NombreMateria=?, Fecha=?, HoraInicio=?, HoraFinal=?, Cupos=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.IDtutoria, p.IDtutor, p.NombreMateria, p.Fecha, p.HoraInicio, p.HoraFinal, p.Cupos)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlHorario) Update(ctx context.Context, p *model.Horario) (*model.Horario, error) {
	query := "Update Horario set NombreMateria=?, Fecha=?, HoraInicio=?, HoraFinal=?, Cupos=? where IDtutoria=? AND IDtutor=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.NombreMateria,
		p.Fecha,
		p.HoraInicio,
		p.HoraFinal,
		p.Cupos,
		p.IDtutoria,
		p.IDtutor,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *mysqlHorario) Delete(ctx context.Context, IDtutoria int64, IDtutor int64) (bool, error) {
	query := "Delete From Horario Where IDtutoria=? AND IDtutor=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, IDtutoria, IDtutor)
	if err != nil {
		return false, err
	}
	return true, nil
}
