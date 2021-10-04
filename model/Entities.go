package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

func Parser(entity string) (Result, error) {

	entity = strings.TrimSpace(entity) //Elimino los " "

	eType := string(entity[0]) + string(entity[1]) //Tipo

	eLength := checkLength(string(entity[2]), string(entity[3])) //Tamaño (Si empieza con 0, devuelvo el sig)

	eValue := entity[4:] // Valores apartir de la pos 4

	eValueLength := len(entity[4:]) //Cant de valores desde la pos 4

	errType := verifyType(eType, eValue) //Identifico el tipo y si concuerdan los valores

	eLengthInteger, errLength := strconv.Atoi(eLength) //Parseo el tamaño de la entidad a int

	if errType == nil {
		if errLength != nil { //Errores
			return newEntity(eType, eLengthInteger, eValue), errors.New("error: el tamaño declarado tiene caracter incluido")
		} else if eLengthInteger == eValueLength {
			return newEntity(eType, eLengthInteger, eValue), nil //OK
		} else {
			return newEntity(eType, eLengthInteger, eValue), errors.New("error: el tamaño declarado no coincide con la cant de valores")
		}
	}
	return newEntity(eType, eLengthInteger, eValue), errType //OK
}

//---------------------------------------------------------------------------------

func newEntity(t string, l int, v string) Result {
	return Result{t, l, v}
}

func checkLength(e2 string, e3 string) string {
	if string(e2) == "0" {
		return e3
	} else {
		return e2 + e3
	}
}

func verifyType(eType string, eValue string) error {
	switch eType {
	case "TX":
		fmt.Println("el tipo es TX")
		return hasNumber(eValue)
	case "NN":
		fmt.Println("el tipo es NN")
		return hasChar(eValue)
	default:
		return errors.New("error: el tipo no es ni TX ni NN")
	}

}

func hasNumber(eValue string) error {
	for _, c := range eValue {
		if unicode.IsDigit(c) {
			return errors.New("error: el valor tiene numero/s")
		}
	}
	return nil //OK
}

func hasChar(e string) error {
	_, err := strconv.Atoi(e)
	if err != nil {
		return errors.New("error: el valor tiene caracter/es")
	}
	return nil //OK
}
