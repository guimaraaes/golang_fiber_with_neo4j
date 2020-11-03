package utils

func Properties(propertiesmap map[string]interface{}) string {
	properties := ""
	i := 0
	for prop, value := range propertiesmap {
		i = i + 1
		properties = properties + prop + ": " + TransToString(value)
		if i <= len(propertiesmap)-1 {
			properties = properties + ", "
		}
	}
	return properties
}
