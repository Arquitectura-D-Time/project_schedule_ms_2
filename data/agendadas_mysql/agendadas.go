package agendadas_mysql

import (
	"context"
	"database/sql"
	repo "project_schedule_ms/data"
	model "project_schedule_ms/model"
)

func NewSQLAgendadas(Conn *sql.DB) repo.Agendadas {
	return &mysqlAgendadas{
		Conn: Conn,
	}
}

type mysqlAgendadas struct {
	Conn *sql.DB
}

func (m *mysqlAgendadas) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.Agendadas, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*model.Agendadas, 0)
	for rows.Next() {
		data := new(model.Agendadas)

		err := rows.Scan(
			&data.IDtutoria,
			&data.IDalumno,
			&data.NombreAlumno,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlAgendadas) Fetch(ctx context.Context, num int64) ([]*model.Agendadas, error) {
	query := "Select IDtutoria, IDalumno, NombreAlumno From Agendadas"

	return m.fetch(ctx, query)
}

func (m *mysqlAgendadas) GetByID(ctx context.Context, IDtutoria int64) ([]*model.Agendadas, error) {
	query := "Select IDtutoria, IDalumno, NombreAlumno From Agendadas where IDtutoria=?"

	rows, err := m.fetch(ctx, query, IDtutoria)
	if err != nil {
		return nil, err
	} else {
		return rows, nil
	}
}

func (m *mysqlAgendadas) GetByID2(ctx context.Context, IDalumno int64) ([]*model.Agendadas, error) {
	query := "Select IDtutoria, IDalumno, NombreAlumno From Agendadas where IDalumno=?"

	rows, err := m.fetch(ctx, query, IDalumno)
	if err != nil {
		return nil, err
	} else {
		return rows, nil
	}
}

func (m *mysqlAgendadas) Create(ctx context.Context, p *model.Agendadas) (int64, error) {
	query := "Insert Agendadas SET IDtutoria=?, IDalumno=?, NombreAlumno=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.IDtutoria, p.IDalumno, p.NombreAlumno)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlAgendadas) Update(ctx context.Context, p *model.Agendadas) (*model.Agendadas, error) {
	query := "Update Agendadas set NombreAlumno=? where IDtutoria=? AND IDalumno=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.NombreAlumno,
		p.IDtutoria,
		p.IDalumno,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *mysqlAgendadas) Delete(ctx context.Context, IDtutoria int64, IDalumno int64) (bool, error) {
	query := "Delete From Agendadas Where IDtutoria=? AND IDalumno=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, IDtutoria, IDalumno)
	if err != nil {
		return false, err
	}
	return true, nil
}