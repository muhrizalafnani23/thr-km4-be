package main

func removeUnrelatedData(mapData map[string]interface{}, key string) map[string]interface{} {
	delete(mapData, key)
	return mapData
}

func main() {}
