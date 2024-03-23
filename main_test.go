package main

import (
	"fmt"
	"testing"
)

type Triangle struct {
	a float64
	b float64
	c float64
}

type TestCase struct {
	name     string
	triangle Triangle
	expected string
}

func TestCheckTriangleType(t *testing.T) {
	testCases := []TestCase{
		{
			name: "tam-giac-deu",
			triangle: Triangle{
				a: 5,
				b: 5,
				c: 5,
			},
			expected: "Tam giác đều",
		},
		{
			name: "tam-giac-can",
			triangle: Triangle{
				a: 5,
				b: 5,
				c: 6,
			},
			expected: "Tam giác cân",
		},
		{
			name: "tam-giac-thuong",
			triangle: Triangle{
				a: 4,
				b: 5,
				c: 6,
			},
			expected: "Tam giác thường",
		},
		{
			name: "khong-phai-tam-giac",
			triangle: Triangle{
				a: 100,
				b: 5,
				c: 6,
			},
			expected: "Không phải tam giác",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected == checkTriangleType(tc.triangle.a, tc.triangle.b, tc.triangle.c) {
				fmt.Printf("pass")
			} else {
				fmt.Printf("failed")
				t.Fail()
			}
		})
	}
}
