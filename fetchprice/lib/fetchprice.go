package lib

func FetchPrice(name, platform, condition, edition, currency string) (price float32, curr string) {
	curr = currency
	price = 1.23

	return price, curr
}
