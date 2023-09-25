package item

func getDescription(description string) (s string) {
	if description != "" {
		s = "# " + description
	}
	return
}
