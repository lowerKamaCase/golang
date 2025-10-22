package product

type ProductCreateRequest struct {
	Name string `json:"name" validate:"required,max=50"`
}

type ProductUpdateRequest struct {
	Name string `json:"name" validate:"required,max=50"`
	ID   uint
}
