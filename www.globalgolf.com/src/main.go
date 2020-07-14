package main

/*

# Fetching datas from https://www.globalgolf.com and export to an excel file 

## What to fetch

Summary for all products from the paths:
* https://www.globalgolf.com/golf-clubs
* https://www.globalgolf.com/golf-balls
* https://www.globalgolf.com/golf-bags
* https://www.globalgolf.com/golf-shoes

## Example

Taking "TaylorMade Milled Grind Satin Chrome Sand Wedge 56° Used Golf Club" for example, the summary should be 

* Title: TaylorMade Milled Grind Satin Chrome Sand Wedge 56° Used Golf Club
* Regular Price: ‪HK$451.81
* Current Price: HK$338.84
* Condition: Very Good
* Brand: TaylorMade
* Model: Milled Grind Satin Chrome
* Player Type: Men
* Clubs Type: Sand Wedge
* Bounce: 12°
* Dexterity: Right Hand
* Flex: Wedgeflex
* Length: Standard
* Lie Angle: Standard
* Loft: 56°
* Shaft: True Temper Dynamic Gold Steel
* Grip: Standard
* Has Headcover: No
* Has Tool: No
* Shaft Material: Steel
* SKU: 1037625-AAG-3B3-AB2
* Image https://image.globalgolf.com/dynamic/1037625/aag/sole-view/taylormade-milled-grind-satin-chrome-wedge.jpg?s=1240
* Image https://image.globalgolf.com/dynamic/1037625/aag/toe-view/taylormade-milled-grind-satin-chrome-wedge.jpg?s=1240
* Image https://image.globalgolf.com/dynamic/1037625/aag/club-face/taylormade-milled-grind-satin-chrome-wedge.jpg?s=1240

## Note

The all html files downloaded by wget.

*/

import (
  "fmt"
  "log"
  "os"
  "strings"
  "path/filepath"
  "github.com/PuerkitoBio/goquery"
  "github.com/360EntSecGroup-Skylar/excelize"
)

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func AddUniqueKey(keys []string, key string) []string {
	_, exists := Find(keys, key)
	if !exists {
		return append(keys, key)
	}
	return keys
}

func ForeachDocuments(root string, handler func(doc *goquery.Document)) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".html") {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		doc, err := goquery.NewDocumentFromReader(f)
		if err != nil {
			log.Fatal(err)
		}

		handler(doc)

		return nil
	})
}

func main() {
	excel := excelize.NewFile()

	columns := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V","W", "X", "Y", "Z",
						"AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV","W", "AX", "AY", "AZ"};

	var datas []map[string]string = make([]map[string]string, 0)
	var keys []string = make([]string, 0)

	root := "../res"

	ForeachDocuments(root, func(doc *goquery.Document) {
		data := make(map[string]string)

		title := doc.Find("h1#prodTitle").Text()
		if len(title) == 0 { return }
		title_key := "Title"
		keys = AddUniqueKey(keys, title_key)
		data[title_key] = title

		regular_price := strings.TrimSpace(doc.Find("td.txtline").Text())
		if len(regular_price) == 0 {
			regular_price = strings.TrimSpace(doc.Find("span.hg.b").Text())
		}
		regular_price_key := "Regular Price"
		keys = AddUniqueKey(keys, regular_price_key)
		data[regular_price_key] = regular_price
		
		current_price := strings.TrimSpace(doc.Find("td.grn").Text())
		current_price_key := "Current Price"
		keys = AddUniqueKey(keys, current_price_key)
		data[current_price_key] = current_price

		titles := make([]string, 0)
		doc.Find("div.s-1-2.conseg.left").Each(func(i int, s *goquery.Selection) {
			text := s.Find("span").Text()
			titles = append(titles, text)
		})

		values := make([]string, 0)
		doc.Find("div.s-1-2.conseg.right").Each(func(i int, s *goquery.Selection) {
			text := s.Find("span").Text()
			values = append(values, text)
		})

		if len(titles) == len(values) {
			for idx, _ := range titles {
				detail_key := titles[idx]
				keys = AddUniqueKey(keys, detail_key)
				data[detail_key] = values[idx]
			}
		}

		doc.Find("a.cloud-zoom-gallery").Each(func(i int, s *goquery.Selection) {
			image_url, _ := s.Attr("href")
			image_key := fmt.Sprintf("Image%d", i + 1)
			keys = AddUniqueKey(keys, image_key)
			data[image_key] = image_url
		})

		datas = append(datas, data)
	})
	
	for idx, key := range keys {
		excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[idx], 1), key)
	}

	for row, data := range datas {
		for column, key := range keys {
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row + 2), data[key])
		}
	}
	
	if err := excel.SaveAs("golf-xxx.xlsx"); err != nil {
		fmt.Println(err)
	}
}