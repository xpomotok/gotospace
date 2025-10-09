package mymaps

type Product struct {
	Name string
	Category string
	Price float64
}

func groupByCategory(products []Product) map[string][]Product {
	groups := make(map[string][]Product)

	for _, p := range products {


		groups[p.Category] = append(groups[p.Category], p)
	}

	return groups
}