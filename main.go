package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	driver "project_schedule_ms/common"
	ac "project_schedule_ms/controllers/agendadas_controller"
	hc "project_schedule_ms/controllers/horario_controller"
)

func main() {
	/*
		dbName := os.Getenv("DB_NAME")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
	*/

	//connection, err := driver.ConnectSQL(dbHost, dbPort, "Fernando", dbPass, dbName)

	connection, err := driver.ConnectSQL("146.148.107.218", "3003", "Fernando", "2123", "agendamiento")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	horarioHandler := hc.NewHorarioHandler(connection)
	agendadasHandler := ac.NewAgendadasHandler(connection)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/horario", horarioRouter(horarioHandler))
		rt.Mount("/agendadas", agendadasRouter(agendadasHandler))
	})

	fmt.Println("Server listen at :5003")
	http.ListenAndServe(":5003", r)
}

// A completely separate router for posts routes
func horarioRouter(horarioHandler *hc.Horario) http.Handler {
	r := chi.NewRouter()
	r.Get("/", horarioHandler.Fetch)
	r.Get("/id/{IDtutoria:[0-9]+}", horarioHandler.GetByID)
	r.Get("/idtutor/{IDtutor:[0-9]+}", horarioHandler.GetByIDTutor)
	r.Get("/nombre/{NombreMateria}", horarioHandler.GetByNombre)
	r.Get("/fecha/{Fecha}", horarioHandler.GetByFecha)
	r.Get("/hora/{HoraInicio}", horarioHandler.GetByHora)
	r.Post("/", horarioHandler.Create)
	r.Put("/{IDtutoria:[0-9]+}/{IDtutor:[0-9]+}", horarioHandler.Update)
	r.Delete("/{IDtutoria:[0-9]+}/{IDtutor:[0-9]+}", horarioHandler.Delete)

	return r
}

func agendadasRouter(agendadasHandler *ac.Agendadas) http.Handler {
	r := chi.NewRouter()
	r.Get("/", agendadasHandler.Fetch)
	r.Get("/tutoria/{IDtutoria:[0-9]+}", agendadasHandler.GetByID)
	r.Get("/alumno/{IDalumno:[0-9]+}", agendadasHandler.GetByID2)
	r.Post("/", agendadasHandler.Create)
	r.Put("/{IDtutoria:[0-9]+}/{IDalumno:[0-9]+}", agendadasHandler.Update)
	r.Delete("/{IDtutoria:[0-9]+}/{IDalumno:[0-9]+}", agendadasHandler.Delete)

	return r
}
