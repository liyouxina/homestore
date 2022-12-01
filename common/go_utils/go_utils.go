package go_utils

func GetStringFromMap(key string, data map[string]interface{}) (*string, bool, error) {
	valueInterface, exists := data[key]
	if !exists {
		return nil, exists, nil
	}
	value := valueInterface.(string)
	return &value, exists, nil
}
