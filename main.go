package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// percobaan GET
// type Student struct{
// 	NIM string `json:"nim"`
// 	FirstName string
// 	LastName string
// 	Class string
// 	Age int
// }

// func main()  {

// 	e := echo.New()

// 	marc := Student{
// 		NIM: "123",
// 		FirstName: "Marc",
// 		LastName: "Bq",
// 		Class: "B",
// 		Age: 20,
// 	}

// 	e.GET("/awal/", func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, marc)
// 	})

// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Halaman awal!")
// 	})

// 	e.GET("/hello", hello)

// 	e.GET("/bye", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "GoodBye!")
// 	})

// 	e.Logger.Fatal(e.Start("localhost:999"))

// }

// func hello(c echo.Context) error{
// 	return c.String(http.StatusOK, "Hello World!")
// }

type Product struct {
	ID string
	Name string
	Stock int
}

var products = make(map[string]Product)

func getAllProducts(c echo.Context) error {
	productList := make ([]Product, 0, len(products))
    for _, v := range products {
		productList = append(productList, v)
	}
    return c.JSON(http.StatusOK, productList)
}

func getProduct(c echo.Context) error {
	ID := c.Param("ID")
	product, found := products[ID]
	if !found {
		return c.String(http.StatusNotFound, "Product Not Found!")
	}
    return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
	newProduct := new(Product)
	if err := c.Bind(newProduct); err != nil {
		return err
	}

	if _, found := products[newProduct.ID]; found {
		return c.String(http.StatusBadRequest, "Product Already Exsist1")
	}

	products[newProduct.ID] = *newProduct
	return c.JSON(http.StatusCreated, newProduct)
}

func updateProducts(c echo.Context) error {
	ID := c.Param("ID")
	if _, found := products[ID]; !found {
		return c.String(http.StatusNotFound, "Product Not Found!")
	}

	updateProduct := new(Product)
	if err := c.Bind(updateProduct); err != nil {
		return err
	}

	products[ID] = *updateProduct
	return c.String(http.StatusOK, "Update Succes!")
}

func deleteProducts(c echo.Context) error {
	ID := c.Param("ID")
	if _, found := products[ID]; !found {
		return c.String(http.StatusNotFound, "Product Not Found!")
	}

	delete(products, ID)
	return c.String(http.StatusOK, "Delete Succes!")
}

func main(){
	e := echo.New()

	e.GET("/products", getAllProducts)
	e.GET("/products/:ID", getProduct)
	e.POST("/products", createProduct)
	e.PUT("/products/:ID", updateProducts)
	e.DELETE("/products/:ID", deleteProducts)

	e.Logger.Fatal(e.Start("localhost:999"))
}

// catatan
// parameter: Param
// e. = dari echo
// type data diisi disini return c.String(http.StatusOK, "Delete Succes!")