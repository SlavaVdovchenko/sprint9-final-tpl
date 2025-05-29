package main

import (
	"testing"
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	size := 100
	res := GenerateRandomElements(size)
	if len(res) != size {
		t.Errorf("expected length %d, got %d", size, len(res))
	}

	result := GenerateRandomElements(0)
	if len(result) != 0 {
		t.Errorf("expected empty slice for size 0, got %v", result)
	}

	result = GenerateRandomElements(1)
	if len(result) != 1 || result[0] != 0 {
		t.Errorf("expected [0] for size 1, got %v", result)
	}
}

func TestMaximum_Max(t *testing.T) {
	data := []int{1, 3, 2, 5, 4}
	expected := 5
	result := Maximum(data)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}

}

func TestMaximum_OneValue(t *testing.T) {
	data := []int{36}
	expected := 36
	result := Maximum(data)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
