package piscine

func ReverseMenuIndex(menu []string) []string {
	menuLength := len(menu)
	reversedMenu := make([]string, menuLength)

	for i := 0; i < menuLength; i++ {
		reversedMenu[i] = menu[menuLength-i-1]
	}

	return reversedMenu
}
