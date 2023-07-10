package main

func main() {
	sli := []int{1, 2, 3}
	sli = append(sli, 4, 5, 6)
	sli = append(sli)
}

func append(args ...interface{}) []int {

	println(args)
	return []int{0}

}
