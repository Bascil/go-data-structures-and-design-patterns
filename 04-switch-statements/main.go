package main

func main(){
 println(price("tomatoes"))
}

func price(item string) int {
	switch item {
	case "apples", "tomatoes":
		return 10
	case "oranges":
		return 20
    default:
		return 0
	}
}