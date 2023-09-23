package techland

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/m4hi2/techpricematchbd/pkg"
	"github.com/m4hi2/techpricematchbd/pkg/utils"
	"net/http"
	"strconv"
	"strings"
)

const SearchEndpoint = "https://www.techlandbd.com/index.php?route=product/search&search="

func SearchProduct(name string) ([]pkg.Product, error) {

	// &fq=1 is a filter for TECH LAND for in stock items.
	res, err := http.Get(fmt.Sprintf("%s%s&fq=1", SearchEndpoint, name))
	if err != nil {
		return nil, err
	}

	sp, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return nil, err
	}

	return extractNProducts(sp, 5)
}

func extractNProducts(sp *goquery.Document, n int) ([]pkg.Product, error) {

	ps := utils.TakeNSelections(sp.Find(".product-thumb"), n)

	var products []pkg.Product
	for _, p := range ps {
		anchor := p.Find("a.product-img")

		pLink, _ := anchor.Attr("href")
		// TechLand loves to add the search string as qParam to the product link
		pLink = utils.CleanLinkQParams(pLink)

		pName := p.Find(".name").Text()
		pImageLink, _ := p.Find("img").Attr("src")

		pPriceText := strings.TrimSpace(p.Find(".price").Text())
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
