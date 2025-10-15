package repo

type Product struct {
	ID          int     `json:"_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{
	}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productId int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productId {
			return product, nil
		}
	}

	return nil, nil
}


func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}


func (r *productRepo) Delete(productId int) error {
	var temp []*Product
	for _, product :=range r.productList{
		if product.ID != productId {
			temp = append(temp, product)
		}
	}

	r.productList = temp
	return nil
}


func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList{
		if p.ID == product.ID {
			r.productList[idx] = &product
			return p, nil
		}
	}

	return nil, nil
}



func generateInitialProducts(r *productRepo) {
	p1 := Product{
		ID:          1,
		Title:       "Mango",
		Description: "A delicious tropical fruit",
		Price:       100,
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRf0J-W_xQ8nJ2T7SeBHdkUc68NZIE0Zb4woQ&s",
	}

	p2 := Product{
		ID:          2,
		Title:       "Apple",
		Price:       50,
		Description: "A sweet red fruit",
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRf0J-W_xQ8nJ2T7SeBHdkUc68NZIE0Zb4woQ&s",
	}

	p3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "A long yellow fruit",
		Price:       30,
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQG7ElBNPs-HbYJJOMHRu7lEmphTn8-52FYKw&s",
	}

	p4 := Product{
		ID:          4,
		Title:       "Pineapple",
		Description: "A tropical fruit with a spiky exterior",
		Price:       150,
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRf0J-W_xQ8nJ2T7SeBHdkUc68NZIE0Zb4woQ&s",
	}

	p5 := Product{
		ID:          5,
		Title:       "Grapes",
		Description: "A bunch of small round fruits",
		Price:       200,
		ImageUrl:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRf0J-W_xQ8nJ2T7SeBHdkUc68NZIE0Zb4woQ&s",
	}

	r.productList = append(r.productList, &p1, &p2, &p3, &p4, &p5)
}