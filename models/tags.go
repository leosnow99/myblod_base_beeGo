package models

func HandleTagsListData(tags []string) map[string]int {
	var tagsMap = make(map[string]int)
	for _, tag := range tags {
		tagsMap[tag]++
	}
	return tagsMap
}
