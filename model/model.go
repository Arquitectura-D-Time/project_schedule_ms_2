package model

type Horario struct {
	IDtutoria     int64  `json':"IDtutoria"`
	IDtutor       int64  `json':"IDtutor"`
	NombreMateria string `json':"NombreMateria"`
	Fecha         string `json':"Fecha"`
	HoraInicio    string `json':"HoraInicio"`
	HoraFinal     string `json':"HoraFinal"`
	Cupos         int64  `json':"Cupos"`
}

type Agendadas struct {
	IDtutoria    int64  `json':"IDtutoria"`
	IDalumno     int64  `json':"IDalumno"`
	NombreAlumno string `json':"NombreAlumno"`
}
