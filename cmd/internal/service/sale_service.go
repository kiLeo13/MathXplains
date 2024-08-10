package service

import "fmt"

var SalesKeyName = "sales_count"

func GetSalesCount() int {
	count, err := configRepo.GetInt(SalesKeyName)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return count
}

func UpdateSalesCount(value int) (int, *APIError) {
	res, err := configRepo.PatchInt(SalesKeyName, value)
	if err != nil {
		return 0, ErrorInternalServer
	}
	return res, nil
}
