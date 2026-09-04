package model

import "encoding/json"

// ForceJSONMarshal Принудительная маршализация в формат JSON.
func ForceJSONMarshal(v any) []byte {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return j
}

// ForceJSONUnmarshal Принудительная распаковка из формата JSON.
func ForceJSONUnmarshal(data []byte, v any) {
	if err := json.Unmarshal(data, v); err != nil {
		panic(err)
	}
}
