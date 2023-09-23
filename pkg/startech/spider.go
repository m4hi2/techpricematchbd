package startech

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/m4hi2/techpricematchbd/pkg"
	"github.com/m4hi2/techpricematchbd/pkg/utils"
	"net/http"
	"strconv"
	"strings"
)

const SearchEndpoint = "https://www.startech.com.bd/product/search?search="

func SearchProduct(name string) ([]pkg.Product, error) {

	res, err := http.Get(fmt.Sprintf("%s%s", SearchEndpoint, name))
	if err != nil {
		return nil, err
	}

	sp, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return nil, err
	}

	return extractNProducts(sp, 5)
}

/*
extractNProducts - Extracts product information into pkg.Product{} structs.
Basically the ugly scraping logic.
*/
func extractNProducts(sp *goquery.Document, n int) ([]pkg.Product, error) {
	ps := utils.TakeNSelections(sp.Find(".p-item"), n)

	var products []pkg.Product
	for _, p := range ps {
		anchor := p.Find("a")

		pLink, _ := anchor.Attr("href")
		pName := anchor.Text()
		pPriceText := strings.TrimSpace(p.Find(".p-item-price").Text())
		pImageLink, _ := p.Find("img").Attr("src")

		if pPriceText == "TBA" {
			continue
		}

		pPriceText = strings.TrimSpace(strings.ReplaceAll(pPriceText, "৳", ""))
		pPriceText = strings.TrimSpace(strings.ReplaceAll(pPriceText, ",", ""))

		if strings.Contains(pPriceText, " ") {
			ss := strings.Split(pPriceText, " ")
			pPriceText = ss[0]
		}

		pPrice, err := strconv.Atoi(pPriceText)
		if err != nil {
			return nil, err
		}

		ep := pkg.Product{
			Name:     pName,
			Link:     pLink,
			Price:    pPrice,
			ImageUrl: pImageLink,
		}

		products = append(products, ep)

	}

	return products, nil
}
