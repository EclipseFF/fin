package group

func getGroup(input string) string {
	var name string

	for _, t := range input {
		if t >= 48 && t <= 54 {
			name += string(t)
		}
	}

	return name
}
