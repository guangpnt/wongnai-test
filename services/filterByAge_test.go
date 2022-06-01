package services

import (
	"example/wongnai-test/model"
	"testing"
)

func TestFilterByAge(t *testing.T) {
	tests := []struct {
		caseName string
		data     model.Data
		want     map[string]int
	}{
		{
			caseName: "N/A",
			data: model.Data{
				Data: []model.RawData{
					{
						Age:      0,
						NationEn: "Thailand",
						Province: "Nan",
					},
					{
						Age:      1,
						NationEn: "Thailand",
						Province: "Nan",
					},
				},
			},
			want: map[string]int{
				"N/A": 1,
			},
		},
		{
			caseName: "0-30",
			data: model.Data{
				Data: []model.RawData{
					{
						Age:      10,
						NationEn: "Thailand",
						Province: "Nan",
					},
					{
						Age:      20,
						NationEn: "Thailand",
						Province: "Nan",
					},
				},
			},
			want: map[string]int{
				"0-30": 2,
			},
		},
		{
			caseName: "31-60",
			data: model.Data{
				Data: []model.RawData{
					{
						Age:      35,
						NationEn: "Thailand",
						Province: "Nan",
					},
					{
						Age:      35,
						NationEn: "Thailand",
						Province: "Nan",
					},
					{
						Age:      100,
						NationEn: "Thailand",
						Province: "Nan",
					},
				},
			},
			want: map[string]int{
				"31-60": 2,
			},
		},
		{
			caseName: "61+",
			data: model.Data{
				Data: []model.RawData{
					{
						Age:      100,
						NationEn: "Thailand",
						Province: "Nan",
					},
					{
						Age:      20,
						NationEn: "Thailand",
						Province: "Nan",
					},
				},
			},
			want: map[string]int{
				"61+": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			result := FilterByAge(tt.data)

			value, ok := result[tt.caseName]

			if ok {
				if value != tt.want[tt.caseName] {
					t.Error("value is not equal to tt.want")
				}
			} else {
				t.Errorf("%v not found", tt.caseName)
			}
		})

	}
}
