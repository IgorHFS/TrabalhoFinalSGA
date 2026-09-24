package services

import "testing"

func TestParseHorario(t *testing.T) {
	validos := [][2]string{{"08:00", "10:00"}, {"13:30", "14:15"}}
	for _, horario := range validos {
		if _, _, err := parseHorario(horario[0], horario[1]); err != nil {
			t.Fatalf("%s-%s deveria ser válido: %v", horario[0], horario[1], err)
		}
	}
	invalidos := [][2]string{{"10:00", "10:00"}, {"10:00", "09:00"}, {"abc", "10:00"}}
	for _, horario := range invalidos {
		if _, _, err := parseHorario(horario[0], horario[1]); err == nil {
			t.Fatalf("%s-%s deveria ser inválido", horario[0], horario[1])
		}
	}
}
