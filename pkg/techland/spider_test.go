package techland_test

import (
	"github.com/m4hi2/techpricematchbd/pkg/techland"
	"log"
	"testing"
)

func TestSearch(t *testing.T) {

	products, err := techland.SearchProduct("blah")

	if err != nil {
		log.Fatal("TECHLAND: error while searching product: ", err)
	}

	for i, product := range products {
		log.Printf("%d \t Name: %s", i, product.Name)
		log.Printf("%d \t Link: %s", i, product.Link)
		log.Printf("%d \t Price: %d", i, product.Price)
		log.Printf("%d \t ImageUrl: %s", i, product.ImageUrl)
	}

}
