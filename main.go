package main

import (
	"github.com/m4hi2/techpricematchbd/pkg"
	"github.com/m4hi2/techpricematchbd/pkg/ryans"
	"github.com/m4hi2/techpricematchbd/pkg/startech"
	"github.com/m4hi2/techpricematchbd/pkg/techland"
	"log"
	"os"
	"strings"
	"sync"
)

//type Result struct {
//	Ryans    []pkg.Product
//	StarTech []pkg.Product
//	TechLand []pkg.Product
//}

func main() {
	searchItemArg := os.Args[1:]

	pname := strings.Join(searchItemArg, " ")

	//res := &Result{}

	ch := make(chan pkg.Product)

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(3)
		go func(wg *sync.WaitGroup) {
			ps, err := ryans.SearchProduct(pname)
			if err != nil {
				log.Println("ERROR:RYANS ", err)
			}

			if ps != nil {
				ch <- ps[0]
			}

			wg.Done()
		}(wg)

		go func(wg *sync.WaitGroup) {
			ps, err := startech.SearchProduct(pname)
			if err != nil {
				log.Println("ERROR:STARTECH ", err)
			}

			if ps != nil {
				ch <- ps[0]
			}
			wg.Done()
		}(wg)

		go func(wg *sync.WaitGroup) {
			ps, err := techland.SearchProduct(pname)
			if err != nil {
				log.Println("ERROR:TECHLAND ", err)
			}

			if ps != nil {
				ch <- ps[0]
			}
			wg.Done()
		}(wg)

		wg.Wait()
		close(ch)
	}()

	for p := range ch {
		logProduct(p)

	}
}

func logProduct(p pkg.Product) {
	log.Println("Name:  ", p.Name)
	log.Println("Price: ", p.Price)
	log.Println("Link:  ", p.Link)

}
