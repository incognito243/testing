package main

import "fmt"

const (
	Installment = "Installment"
	Credit      = "Credit"
	Online      = "Online"
)

func main() {
	fmt.Print(needPayAmount(10000, 1000, Credit))
}

func needPayAmount(realAmount int32, paidAmount int32, method string) int32 {
	if realAmount < paidAmount {
		return -1
	}
	amount := realAmount - paidAmount

	var res int32
	switch method {
	case Installment:
		res = int32(float32(amount+2) * 0.3)
		break
	case Credit:
		res = 0
		break
	case Online:
		res = amount
		break
	default:
		res = -1
		break
	}
	return res
}

func calc(n int) int {
	ans := 0
	if n < 0 {
		ans = 0
	} else {
		ans = int(n * (n + 1) * (2*n + 1) / 6)
	}
	return ans
}
