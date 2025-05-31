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

func TestMaxChanks(t *testing.T) {
	test := []struct {
		name string
		in   []int
		want int
	}{
		{name: "review test",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 9,
		},
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
			assert.Equal(t, tc.want, maxChunks(tc.in))
		})
	}

}
