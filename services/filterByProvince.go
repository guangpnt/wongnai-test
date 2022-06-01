package services

import "example/wongnai-test/model"

func FilterByProvince(data model.Data) map[string]int {
	provinceCount := map[string]int{}

	for i := 0; i < len(data.Data); i++ {
		if data.Data[i].NationEn == "Thailand" {
			if data.Data[i].Province == "" {
				provinceCount["N/A"]++
			} else {
				provinceCount[data.Data[i].Province]++
			}
		}
	}

	return provinceCount
}
