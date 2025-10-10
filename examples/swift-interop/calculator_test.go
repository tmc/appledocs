package main

import (
	"testing"

	"github.com/ebitengine/purego"
)

func TestCalculator(t *testing.T) {
	// Open the Swift library with calculator
	lib, err := purego.Dlopen("./libcalculator.dylib", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		t.Skipf("Calculator library not built, run: swiftc -emit-library calculator.swift -o libcalculator.dylib")
	}

	// Register functions
	var calcCreate func() uintptr
	var calcAdd func(uintptr, float64, float64) float64
	var calcMultiply func(uintptr, float64, float64) float64
	var calcStore func(uintptr, float64)
	var calcRecall func(uintptr) float64
	var calcRelease func(uintptr)

	purego.RegisterLibFunc(&calcCreate, lib, "calculator_create")
	purego.RegisterLibFunc(&calcAdd, lib, "calculator_add")
	purego.RegisterLibFunc(&calcMultiply, lib, "calculator_multiply")
	purego.RegisterLibFunc(&calcStore, lib, "calculator_store")
	purego.RegisterLibFunc(&calcRecall, lib, "calculator_recall")
	purego.RegisterLibFunc(&calcRelease, lib, "calculator_release")

	// Create calculator instance
	calc := calcCreate()
	defer calcRelease(calc)

	// Test addition
	result := calcAdd(calc, 10.5, 5.5)
	if result != 16.0 {
		t.Errorf("Expected 16.0, got %f", result)
	}

	// Test multiplication
	result = calcMultiply(calc, 4.0, 2.5)
	if result != 10.0 {
		t.Errorf("Expected 10.0, got %f", result)
	}

	// Test memory
	calcStore(calc, 42.0)
	result = calcRecall(calc)
	if result != 42.0 {
		t.Errorf("Expected 42.0, got %f", result)
	}
}
