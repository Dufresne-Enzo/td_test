package main

import "testing"

func testAddition(t *testing.T) {
	var param1 int
	var param2 int
	var funcReturn int

	param1 = 3
	param2 = 4
	funcReturn = monAddition(param1, param2)
	if funcReturn <= 0 {
		t.Error("monAddition return is <= 0")
	}
	if reflect.TypeOf(funcReturn) != int {
		t.Error("La valeur de retour n'est pas de type int")
	}
}

func testSoustraction(t *testing.T) {
	var param1 int
	var param2 int
	var funcReturn int

	param1 = 1
	param2 = 9
	funcReturn = maSoustraction(param1, param2)
	if funcReturn >= 0 {
		t.Error("maSoustraction return is >= 0")
	}
	if reflect.TypeOf(funcReturn) != int {
		t.Error("La valeur de retour n'est pas de type int")
	}
}

func testMultiplication(t *testing.T) {
	var param1 int
	var param2 int
	var funcReturn int

	param1 = 2
	param2 = 9
	funcReturn = maMultiplication(param1, param2)
	if funcReturn >= 0 {
		t.Error("maMultiplication return is != 18")
	}
	if reflect.TypeOf(funcReturn) != int {
		t.Error("La valeur de retour n'est pas de type int")
	}
}

func testDivision(t *testing.T) {
	var param1 int
	var param2 int
	var funcReturn int

	param1 = 10
	param2 = 2
	funcReturn = maDivision(param1, param2)
	if funcReturn == 0 {
		t.Error("maDivision return is = 0")
	}
	if reflect.TypeOf(funcReturn) != int {
		t.Error("La valeur de retour n'est pas de type int")
	}
}

