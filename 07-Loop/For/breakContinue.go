package main

func main() {
	// continue
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			continue
		}
		println("the even number:", i)

	}
	//break
	for j := 1; j < 100; j++ {
		if j == 50 {
			break
		}
		println(j)

	}
}
