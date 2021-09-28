package main

func isBalanced(s string) string {
	// Write your code here
	if s == "" {
		return "YES"
	}

	if len(s)%2 == 1 {
		// fmt.Println("mdjdidi")
		return "NO"
	}

	store := []string{}

	for _, val := range s {
		str := string(val)

		if str == "(" || str == "{" || str == "[" {
			store = append(store, str)
		} else {
			lastString := store[len(store)-1]
			if str == ")" {
				if lastString == "(" {
					store = store[:len(store)-1]
				}
			} else if str == "}" {
				if lastString == "{" {
					store = store[:len(store)-1]
				}
			} else if str == "]" {
				if lastString == "[" {
					store = store[:len(store)-1]
				}
			} else {
				return "NO"
			}
		}
	}

	if len(store) == 0 && len(s)%2 == 0 {
		return "YES"
	}
	return "NO"
}

func main() {

}
