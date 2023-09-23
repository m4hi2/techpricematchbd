package startech_test

import (
	"github.com/m4hi2/techpricematchbd/pkg/startech"
	"log"
	"testing"
)

func TestSearch(t *testing.T) {

	products, err := startech.SearchProduct("edifier")

	if err != nil {
		log.Fatal("RYANS: error while searching product: ", err)
	}

	for i, product := range products {
		log.Printf("%d \t Name: %s", i, product.Name)
		log.Printf("%d \t Link: %s", i, product.Link)
		log.Printf("%d \t Price: %d", i, product.Price)
		log.Printf("%d \t ImageUrl: %s", i, product.ImageUrl)
	}

}
