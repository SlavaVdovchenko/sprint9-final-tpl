package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {

	test := []struct {
		name string
		in   int
		want []int
	}{
		{
			name: "len(slice) == 0",
			in:   0,
			want: []int{},
		},
		{
			name: "len(slice) == 1",
			in:   1,
			want: []int{0},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, generateRandomElements(tc.in))
		})
	}
}
func TestMaximum(t *testing.T) {

	test := []struct {
		name string
		in   []int
		want int
	}{
		{
			name: "true maximum",
			in:   []int{1, 5, 6, 10, 23},
			want: 23,
		},
		{
			name: "len(slice) == 1",
			in:   []int{23},
			want: 23,
		},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, maximum(tc.in))
		})
	}
}

/*func TestGenerateRandomElements(t *testing.T) {
	size := 100
	res := generateRandomElements(size)
	if len(res) != size {
		t.Errorf("expected length %d, got %d", size, len(res))
	}

	result := generateRandomElements(0)
	if len(result) != 0 {
		t.Errorf("expected empty slice for size 0, got %v", result)
	}

	result = generateRandomElements(1)
	if len(result) != 1 || result[0] != 0 {
		t.Errorf("expected [0] for size 1, got %v", result)
	}
}

func TestMaximum_Max(t *testing.T) {
	data := []int{1, 3, 2, 5, 4}
	expected := 5
	result := maximum(data)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}

}

func TestMaximum_OneValue(t *testing.T) {
	data := []int{36}
	expected := 36
	result := maximum(data)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
} */
