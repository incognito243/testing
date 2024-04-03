package main

import (
	"testing"
)

type Payment struct {
	n      int32
	m      int32
	method string
}

type TestCasePayment struct {
	name     string
	payment  Payment
	expected int32
}

type TestCaseSum struct {
	name     string
	n        int
	expected int
}

func TestCheckPayAmount(t *testing.T) {
	testCases := []TestCasePayment{
		{
			name: "id00",
			payment: Payment{
				n:      100000,
				m:      50000,
				method: "Credit",
			},
			expected: 0,
		},
		{
			name: "id01",
			payment: Payment{
				n:      100000,
				m:      50000,
				method: "Installment",
			},
			expected: 15000,
		},
		{
			name: "id02",
			payment: Payment{
				n:      100000,
				m:      50000,
				method: "Online",
			},
			expected: 50000,
		},
		{
			name: "id03",
			payment: Payment{
				n:      100000,
				m:      1,
				method: "Credit",
			},
			expected: 0,
		},
		{
			name: "id04",
			payment: Payment{
				n:      100000,
				m:      1,
				method: "Installment",
			},
			expected: 30000,
		},
		{
			name: "id05",
			payment: Payment{
				n:      100000,
				m:      1,
				method: "Online",
			},
			expected: 99999,
		},
		{
			name: "id06",
			payment: Payment{
				n:      100000,
				m:      2,
				method: "Credit",
			},
			expected: 0,
		},
		{
			name: "id07",
			payment: Payment{
				n:      100000,
				m:      2,
				method: "Installment",
			},
			expected: 30000,
		},
		{
			name: "id08",
			payment: Payment{
				n:      100000,
				m:      2,
				method: "Online",
			},
			expected: 99998,
		},
		{
			name: "id09",
			payment: Payment{
				n:      1,
				m:      2,
				method: "Credit",
			},
			expected: -1,
		},
		{
			name: "id10",
			payment: Payment{
				n:      100,
				m:      1,
				method: "Credit",
			},
			expected: 0,
		},
		{
			name: "id11",
			payment: Payment{
				n:      101,
				m:      1,
				method: "Installment",
			},
			expected: 30,
		},
		{
			name: "id12",
			payment: Payment{
				n:      100,
				m:      1,
				method: "Online",
			},
			expected: 99,
		},
		{
			name: "id13",
			payment: Payment{
				n:      100,
				m:      2,
				method: "NoMethod",
			},
			expected: -1,
		},
		{
			name: "id14",
			payment: Payment{
				n:      1,
				m:      2,
				method: "Credit",
			},
			expected: -1,
		},
		{
			name: "id15",
			payment: Payment{
				n:      101,
				m:      1,
				method: "Installment",
			},
			expected: 30,
		},
		{
			name: "id16",
			payment: Payment{
				n:      100,
				m:      1,
				method: "Credit",
			},
			expected: 0,
		},
		{
			name: "id17",
			payment: Payment{
				n:      100,
				m:      1,
				method: "Online",
			},
			expected: 99,
		},
		{
			name: "id13",
			payment: Payment{
				n:      100,
				m:      1,
				method: "NoMethod",
			},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected != needPayAmount(tc.payment.n, tc.payment.m, tc.payment.method) {
				t.Fail()
			}
		})
	}
}

func TestCalc(t *testing.T) {
	testCases := []TestCaseSum{
		{
			name:     "id0",
			n:        0,
			expected: 0,
		},
		{
			name:     "id1",
			n:        3,
			expected: 14,
		},
		{
			name:     "id2",
			n:        0,
			expected: 0,
		},
		{
			name:     "id3",
			n:        0,
			expected: 0,
		},
		{
			name:     "id4",
			n:        4,
			expected: 30,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected != calc(tc.n) {
				t.Fail()
			}
		})
	}
}
