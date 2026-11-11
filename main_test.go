package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-10, -1, 0, 1, 7, 8, 9, 1000} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			data := generateRandomElements(size)
			wantLen := size
			if wantLen < 0 {
				wantLen = 0
			}
			if len(data) != wantLen {
				t.Fatalf("длина = %d, ожидается %d", len(data), wantLen)
			}
			for i, value := range data {
				if value <= 0 {
					t.Fatalf("data[%d] = %d, ожидается положительное число", i, value)
				}
			}
		})
	}
}

func TestMaximumFunctions(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	cases := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"zero", []int{0}, 0},
		{"negative", []int{-9, -2, -7}, -2},
		{"equal", []int{5, 5, 5, 5, 5, 5, 5, 5}, 5},
		{"ascending", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"descending", []int{8, 7, 6, 5, 4, 3, 2, 1}, 8},
		{"remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, 99},
		{"limits", []int{minInt, 0, maxInt}, maxInt},
		{"minimum", []int{minInt}, minInt},
	}
	for name, findMax := range map[string]func([]int) int{
		"maximum":   maximum,
		"maxChunks": maxChunks,
	} {
		t.Run(name, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					original := slices.Clone(tc.data)
					if got := findMax(tc.data); got != tc.want {
						t.Errorf("максимум = %d, ожидается %d", got, tc.want)
					}
					if !slices.Equal(tc.data, original) {
						t.Error("исходный слайс изменён")
					}
				})
			}
		})
	}
}

func TestMaxChunksEveryPosition(t *testing.T) {
	for size := 1; size <= 65; size++ {
		for position := 0; position < size; position++ {
			data := make([]int, size)
			for i := range data {
				data[i] = -10
			}
			data[position] = -1
			if got := maxChunks(data); got != -1 {
				t.Fatalf("размер %d, позиция %d: максимум = %d", size, position, got)
			}
		}
	}
}
