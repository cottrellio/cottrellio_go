package validating

// ValidFiltersMap valid filters by model.
var ValidFiltersMap = map[string]interface{}{
	"user": map[string]map[string]string{
		"search": map[string]string{
			"validator": "max=100",
		},
		"created_from": map[string]string{
			"validator": "isodate",
		},
		"created_to": map[string]string{
			"validator": "isodate",
		},
	},
}
