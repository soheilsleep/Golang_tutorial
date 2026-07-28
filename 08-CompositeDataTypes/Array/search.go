package main

func main() {
	names := [8]string{"soheil", "ali", "ahmad", "zahra", "asghar", "soghra", "taha", "saman"}
	var searchKey = "zahra"
	for i, name := range names {
		if name == searchKey {
			println("name found , index:", i)
			break
		}

	}

	for j := 0; j < len(names); j++ {
		if names[j] == searchKey {
			println("name found , index:", j)
			break
		}

	}
}
