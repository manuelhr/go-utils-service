package matrixops

import "errors"

func EsMatrizValida(input [][]int) (bool, error) {
	if input == nil {
		return false, errors.New("solicitud no válida")
	}
	if len(input) == 0 {
		return false, errors.New("solicitud no válida. Se requiere una matriz con elementos")
	}
	if len(input[0]) == 0 {
		return false, errors.New("solicitud no válida. Se requiere una matriz con elementos")
	}
	longitudLadoMatriz := len(input)
	for _, currentRow := range input {
		if longitudLadoMatriz != len(currentRow) {
			return false, errors.New("solicitud no válida. Se requiere una matriz simétrica")
		}
	}
	return true, nil
}

func Rotar90GradosAntihorario(input [][]int) ([][]int, error) {

	inputValido := false
	inputValido, errorValidacionMatriz := EsMatrizValida(input)
	if !inputValido {
		return nil, errorValidacionMatriz
	}

	//2)Preparación de variables de trabajo
	var i, j, x int
	longitudLadoMatriz := len(input)

	serializedInput := []int{}
	for i = 0; i < len(input); i++ {
		currSlice := input[i][:]
		serializedInput = append(serializedInput, currSlice...)
	}

	output := make([][]int, longitudLadoMatriz)
	for i := range input {
		output[i] = make([]int, longitudLadoMatriz)
	}

	//3) Operación de "rotación"
	for i = 0; i < longitudLadoMatriz; i++ {
		for j = longitudLadoMatriz - 1; j > -1; j-- {
			output[j][i] = serializedInput[x]
			x++
		}
	}

	return output, nil
}
