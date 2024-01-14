package domain

type Product struct {
	IDProduct   string `json:"id_product"`
	ProductName string `json:"product_name"`
	Price       uint64 `json:"price"`
	Count       int    `json:"count"`
	ImageURL    string `json:"image_url"`
}
