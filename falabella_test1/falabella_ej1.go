package tlvToMap

import (
	"regexp"
	"strconv"
)

const TYPEOFFSET = 3
const LENGHTOFFSET = 5

var validValueTypes = map[string]string{
	"A": "^[0-9A-Za-z]+$",
	"N": "^[0-9]+$",
}

type InputError struct{ S string }

func (e InputError) Error() string {
	return "The input provided is not valid: " + e.S
}

//checks if the input is empty and it contains all the fields
func checkInputErrors(fields []byte) error {
	if len(fields) == 0 {
		return InputError{"The input cant be empty"}
	}
	if len(fields) < LENGHTOFFSET {
		return InputError{"The input must provide at least the first 2 fields even if the value is empty"}
	}
	return nil
}

//returns true and an error if going out of range
func checkIndexInRange(fields []byte, offset int) (bool, error) {
	if len(fields) < offset {
		return true, InputError{"Index out of range"}
	}
	return false, nil
}

//it returns the field type checking the first letter is correct
func getType(fields []byte, index int) (string, error) {
	if outOfRange, err := checkIndexInRange(fields, index+TYPEOFFSET); outOfRange {
		return "", err
	}
	myType := string(fields[index : index+TYPEOFFSET])
	if myType[0] != 'A' && myType[0] != 'N' {
		return "", InputError{"The first character of the type field should be A or N"}
	}
	return myType, nil
}

//returns the field lenght checking that is a positive integer
func getLenght(fields []byte, index int) (int, error) {
	var lenght int
	if outOfRange, err := checkIndexInRange(fields, index+LENGHTOFFSET); outOfRange {
		return lenght, err
	}
	lenght, err := strconv.Atoi(string(fields[index+TYPEOFFSET : index+LENGHTOFFSET]))
	if lenght < 0 {
		err = InputError{"The lenght of the value must be positive"}
	}
	return lenght, err
}

//returns the field value checking that matches the type defined in the "type" field
func getValue(fields []byte, index int, lenght int, fieldType string) (string, error) {
	if outOfRange, err := checkIndexInRange(fields, index+LENGHTOFFSET+lenght); outOfRange {
		return "", err
	}
	value := string(fields[index+LENGHTOFFSET : index+LENGHTOFFSET+lenght])
	match, _ := regexp.MatchString(validValueTypes[fieldType], value)
	if !match {
		return value, InputError{"The value must match the value type defined on the first field"}
	}
	return value, nil
}

func ToMap(fields []byte) (map[string]string, error) {
	resultMap := map[string]string{}
	index := 0
	var key, value string
	var lenght int
	err := checkInputErrors(fields)

	if err != nil {
		return resultMap, err
	}
	for index < len(fields) {
		if key, err = getType(fields, index); err != nil {
			return resultMap, err
		}
		if lenght, err = getLenght(fields, index); err != nil {
			return resultMap, err
		}
		if value, err = getValue(fields, index, lenght, string(key[0])); err != nil {
			return resultMap, err
		}
		index += LENGHTOFFSET + lenght
		resultMap[key] = value
	}
	return resultMap, err
}
