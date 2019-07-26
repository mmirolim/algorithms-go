package hk

import (
	"sort"
)

// https://www.hackerrank.com/challenges/jim-and-the-orders/problem
func JimOrders(input [][]int32) []int32 {
	// last is customer id
	orders := make([][3]int32, len(input))
	var customerID int32 = 1
	for i := range input {
		orders[i] = [3]int32{input[i][0], input[i][1], customerID}
		customerID++
	}

	// sort orders by order + prep_time and customer id in ascending order
	sort.Slice(orders, func(i, j int) bool {
		serveTimeI := orders[i][0] + orders[i][1]
		serveTimeJ := orders[j][0] + orders[j][1]
		if serveTimeI == serveTimeJ {
			// sort by customer id
			if orders[i][2] < orders[j][2] {
				return true
			}
			return false
		} else if serveTimeI < serveTimeJ {
			return true
		} else {
			return false
		}
	})
	customersOrder := make([]int32, len(input))
	for i := range orders {
		customersOrder[i] = orders[i][2]
	}
	return customersOrder
}
