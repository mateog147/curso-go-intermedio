package main

import "testing"

// comados para correr los test
// go test -v -> corre todos los test con salida detallada
// go test -run NombreDelTest -v -> corre un test especifico
// go test -cover -> muestra el porcentaje de cobertura de codigo por los test
// go test -coverprofile=coverage.out -> genera un archivo con el reporte de cobertura
// go tool cover -html=coverage.out -o coverage.html -> abre un reporte en html con el detalle de la cobertura
/* go tool cover -func=coverage.out -> muestra la cobertura por funcion
go test -cpuprofile=cpu.out  -> genera un perfil de CPU
top
list Suma  -> muestra el codigo mas costoso en tiempo de CPU en la funcion Suma
quit
web
pdf
*/
// go tool pprof cpu.out -> abre el analizador de perfiles de CPU
func TestSuma(t *testing.T) {
	resultado := Suma(2, 3)
	valorEsperado := 5

	if resultado != valorEsperado {
		t.Errorf("Suma(2, 3) = %d; se esperaba %d", resultado, valorEsperado)
	}
}

func TestMultiSuma(t *testing.T) {
	tableTests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{-2, -3, -5},
		{100, 200, 300},
	}
	for _, item := range tableTests {
		result := Suma(item.a, item.b)
		if result != item.expected {
			t.Errorf("Suma(%d, %d) = %d; se esperaba %d", item.a, item.b, result, item.expected)
		}
	}

}

func TestFibonacci(t *testing.T) {
	tableTests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{50, 12586269025},
	}
	for _, item := range tableTests {
		result := Fibonacci(item.n)
		if result != item.expected {
			t.Errorf("Fibonacci(%d) = %d; se esperaba %d", item.n, result, item.expected)
		}
	}
}
