package services

import (
	"example/wongnai-test/model"
	"testing"
)

func TestFilterByProvince(t *testing.T) {
	tests := []struct {
		caseName     string
		provinceName []string
		data         model.Data
		want         map[string]int
	}{
		{
			caseName: "two province",
			provinceName: []string{
				"A", "B",
			},
			data: model.Data{
				Data: []model.RawData{
					{
						Age:      0,
						NationEn: "Thailand",
						Province: "A",
					},
					{
						Age:      1,
						NationEn: "Thailand",
						Province: "B",
					},
					{
						Age:      2,
						NationEn: "Thailand",
						Province: "B",
					},
				},
			},
			want: map[string]int{
				"A": 1,
				"B": 2,
			},
		},
		{
			caseName: "unknow province",
			provinceName: []string{
				"N/A",
			},
			data: model.Data{
				Data: []model.RawData{
					{
						Age:      0,
						NationEn: "Thailand",
						Province: "",
					},
					{
						Age:      0,
						NationEn: "Thailand",
						Province: "",
					},
				},
			},
			want: map[string]int{
				"N/A": 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			result := FilterByProvince(tt.data)

			for i := 0; i < len(tt.provinceName); i++ {
				value, ok := result[tt.provinceName[i]]

				if ok {
					if value != tt.want[tt.provinceName[i]] {
						t.Error("value is not equal to tt.want")
					}
				} else {
					t.Errorf("%v not found", tt.provinceName[i])
				}
			}
		})

	}
}
