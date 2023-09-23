package ryans

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/m4hi2/techpricematchbd/pkg"
	"github.com/m4hi2/techpricematchbd/pkg/utils"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func SearchProduct(name string) ([]pkg.Product, error) {

	sUrl := fmt.Sprintf("https://www.ryanscomputers.com/api/search?keyword=%s&returnType=searchPageHTML", url.QueryEscape(name))

	client := &http.Client{}
	req, err := http.NewRequest("GET", sUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	sp, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return nil, err
	}
	return extractNProducts(sp, 5)
}
func extractNProducts(sp *goquery.Document, n int) ([]pkg.Product, error) {

	ps := utils.TakeNSelections(sp.Find(".category-single-product"), n)

	var products []pkg.Product
	for _, p := range ps {
		anchor := p.Find(".image-box").Find("a")

		pLink, _ := anchor.Attr("href")

		pName := p.Find(".list-view-text").Text()
		pImageLink, _ := p.Find("img").Attr("src")

		pPriceText := strings.TrimSpace(p.Find(".pr-text").Text())
		pPriceText = strings.TrimSpace(strings.ReplaceAll(pPriceText, "Tk", ""))
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
