package main

/*

# Fetching datas from https://www.golfbidder.co.uk and export to an excel file 

## What to fetch
Summary for all products. 

## Example

Taking Mizuno MP-18 for example, the summary should be 

* Reference: 1751502 
* Model Year: 2017 
* Golfer: Mens Right Hand 
* Brand: Mizuno 
* Club: 5-PW Iron Set 
* Shaft: NS pro Modus Tour 105 
* Shaft Material: Steel 
* Flex: Stiff 
* Flex Rating: 
* Grip: Golf Pride MCC plus 4 
* Head Condition: 6 - Fair 
* Shaft Condition: 7 - Good 
* Grip Condition: 6 - Fair 
* Price: £408.00 
* RRP £810.00
* Save £402.00 
* Image https://www.golfbidder.co.uk/images/productimages/066/callaway-md5-jaws-raw-x-grind-lw-s1829236-01.jpg 
* Image https://www.golfbidder.co.uk/images/productimages/066/callaway-md5-jaws-raw-x-grind-lw-s1829236-02.jpg 
* Image https://www.golfbidder.co.uk/images/productimages/066/callaway-md5-jaws-raw-x-grind-lw-s1829236-03.jpg 
* Image https://www.golfbidder.co.uk/images/productimages/066/callaway-md5-jaws-raw-x-grind-lw-s1829236-04.jpg 
* Image https://www.golfbidder.co.uk/images/productimages/066/callaway-md5-jaws-raw-x-grind-lw-s1829236-05.jpg

## Note

The all html files downloaded by SiteSucker.

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

func main() {
	excel := excelize.NewFile()

	excel.SetCellValue("Sheet1", "A1", "Title")
	excel.SetCellValue("Sheet1", "B1", "Reference")
	excel.SetCellValue("Sheet1", "C1", "Model Year")
	excel.SetCellValue("Sheet1", "D1", "Golfer")
	excel.SetCellValue("Sheet1", "E1", "Brand")
	excel.SetCellValue("Sheet1", "F1", "Club")
	excel.SetCellValue("Sheet1", "G1", "Shaft")
	excel.SetCellValue("Sheet1", "H1", "Material")
	excel.SetCellValue("Sheet1", "I1", "Grip")
	excel.SetCellValue("Sheet1", "J1", "Head Condition")
	excel.SetCellValue("Sheet1", "K1", "Shaft Condition")
	excel.SetCellValue("Sheet1", "L1", "Grip Condition")
	excel.SetCellValue("Sheet1", "M1", "Price")
	excel.SetCellValue("Sheet1", "N1", "RRP")
	excel.SetCellValue("Sheet1", "O1", "Saving")
	excel.SetCellValue("Sheet1", "P1", "Image")

	columns := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V","W", "X", "Y", "Z"};

	root := "../res"
	row := 2
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if (strings.HasSuffix(path, ".html")) {
			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			doc, err := goquery.NewDocumentFromReader(f)
			
			if err != nil {
				log.Fatal(err)
			}

			column := 0

			doc.Find("div#bd").Each(func(i int, s *goquery.Selection) {
				title_value := s.Find("h1").Text()
				excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), title_value)
			})
			column++

			// reference_title := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblProductReference").Text()
			reference_value := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblProductReferenceText").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), reference_value)
			column ++

			// year_title := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblModelYear").Text()
			year_value := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblModelYearText").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), year_value)
			column ++

			// golfer_title := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblGendageCode").Text()
			golfer_value := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblGendageCodeText").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), golfer_value)
			column ++

			// brand_title := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblBrand").Text()
			brand_value := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblBrandText").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), brand_value)
			column ++

			// club_title := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblModel").Text()
			club_value := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblModelText").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), club_value)
			column ++

			// shaft_title := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblShaftName").Text()
			shaft_value := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblShaftNameText").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), shaft_value)
			column ++

			// material_title := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblShaftMarterial").Text()
			material_value := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblShaftMarterialText").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), material_value)
			column ++

			// grip_title := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblGrip").Text()
			grip_value := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblGripText").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), grip_value)
			column ++

			// head_condition_title := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblHeadCondition").Text()
			head_condition_value := doc.Find("label#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblHeadConditionText_LabelWithLightBox_LabelText").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), head_condition_value)
			column ++

			// shaft_condition_title := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblShaftCondition").Text()
			shaft_condition_value := doc.Find("label#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblShaftConditionText_LabelWithLightBox_LabelText").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), shaft_condition_value)
			column ++

			// grip_condition_title := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblGripCondition").Text()
			grip_condition_value := doc.Find("label#ctl00_ContentProductNarrow_ProductItemInfo_ProductVerbosity_uiItemData_lblGripConditionText_LabelWithLightBox_LabelText").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), grip_condition_value)
			column ++

			// price_title := doc.Find("span#price-lbl").Text()

			price_integer_value := doc.Find("span.integer-part").Text()
			price_decimal_value := doc.Find("span.decimal-part").Text()
			price_value := fmt.Sprintf("%s%s", price_integer_value, price_decimal_value)
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), price_value)
			column ++

			// RRP_title := doc.Find("span#rrp-lbl").Text()
			RRP_value := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductActionsBelow_litRRP").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), RRP_value)
			column ++

			// saving_title := doc.Find("span#saving-lbl").Text()
			saving_value := doc.Find("span#ctl00_ContentProductNarrow_ProductItemInfo_ProductActionsBelow_litSaving").Text()
			excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), saving_value)
			column ++

			doc.Find(".li-product").Each(func(i int, s *goquery.Selection) {
				url, _ := s.Find("a").Attr("href")
				excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), url)
				column ++
			})

			// parts := strings.Split(path, "/")
			// excel.SetCellValue("Sheet1", fmt.Sprintf("%s%d", columns[column], row), parts[len(parts) - 2])
			// column++

			row += 1
		}
		return err
	})
	
	if err := excel.SaveAs("golfbidder.xlsx"); err != nil {
		fmt.Println(err)
	}
}