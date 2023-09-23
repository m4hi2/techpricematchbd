package utils

import "github.com/PuerkitoBio/goquery"

/*
TakeNSelections - Takes N number of "Selection"s from goquery selection and makes it possible to
range over them since it returns a slice of selections.
*/
func TakeNSelections(s *goquery.Selection, n int) []*goquery.Selection {
	var rs []*goquery.Selection
	cs := 1

	s.EachWithBreak(func(i int, selection *goquery.Selection) bool {
		rs = append(rs, selection)
		cs += 1
		if cs > n {
			return false
		}

		return true
	})

	return rs
}
