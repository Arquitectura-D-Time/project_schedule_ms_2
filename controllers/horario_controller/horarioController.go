package horario_controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	driver "project_schedule_ms/common"
	repository "project_schedule_ms/data"
	horario "project_schedule_ms/data/horario_mysql"
	model "project_schedule_ms/model"
)

func NewHorarioHandler(db *driver.DB) *Horario {
	return &Horario{
		repo: horario.NewSQLHorario(db.SQL),
	}
}

type Horario struct {
	repo repository.Horario
}

func (p *Horario) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) Create(w http.ResponseWriter, r *http.Request) {
	horario := model.Horario{}
	json.NewDecoder(r.Body).Decode(&horario)

	newID, err := p.repo.Create(r.Context(), &horario)
	fmt.Println(newID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	fmt.Println(err)
}

func (p *Horario) Update(w http.ResponseWriter, r *http.Request) {
	idtutoria, _ := strconv.Atoi(chi.URLParam(r, "IDtutoria"))
	idtutor, _ := strconv.Atoi(chi.URLParam(r, "IDtutor"))

	data := model.Horario{
		IDtutoria: int64(idtutoria),
		IDtutor:  int64(idtutor),
	}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "IDtutoria"))
	
	payload, err := p.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) GetByIDTutor(w http.ResponseWriter, r *http.Request) {
	idtutor, _ := strconv.Atoi(chi.URLParam(r, "IDtutor"))
	
	payload, err := p.repo.GetByIDTutor(r.Context(), int64(idtutor))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) GetByNombre(w http.ResponseWriter, r *http.Request) {
	nombre := strconv.QuoteToASCII(chi.URLParam(r, "NombreMateria"))
	fmt.Println(nombre)
	payload, err := p.repo.GetByNombre(r.Context(),string(nombre))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) GetByFecha(w http.ResponseWriter, r *http.Request) {
	fecha := (chi.URLParam(r, "Fecha"))
	fmt.Println(fecha)
	payload, err := p.repo.GetByFecha(r.Context(), string(fecha))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) GetByHora(w http.ResponseWriter, r *http.Request) {
	hora := (chi.URLParam(r, "HoraInicio"))
	fmt.Println(hora)
	payload, err := p.repo.GetByHora(r.Context(), string(hora))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Horario) Delete(w http.ResponseWriter, r *http.Request) {
	idtutoria, _ := strconv.Atoi(chi.URLParam(r, "IDtutoria"))
	idtutor, _ := strconv.Atoi(chi.URLParam(r, "IDtutor"))
	_, err := p.repo.Delete(r.Context(), int64(idtutoria), int64(idtutor))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
