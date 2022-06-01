package services

import "example/wongnai-test/model"

func FilterByAge(data model.Data) map[string]int {
	ageCount := map[string]int{}

	for i := 0; i < len(data.Data); i++ {
		if data.Data[i].NationEn == "Thailand" {
			if data.Data[i].Age == 0 {
				ageCount["N/A"]++
			}

			if data.Data[i].Age > 0 && data.Data[i].Age < 31 {
				ageCount["0-30"]++
			}

			if data.Data[i].Age > 30 && data.Data[i].Age < 61 {
				ageCount["31-60"]++
			}

			if data.Data[i].Age > 60 {
				ageCount["61+"]++

			}
		}
	}

	return ageCount
}
