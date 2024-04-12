package gobnb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/johnbalvin/gobnb/details"
	"github.com/johnbalvin/gobnb/search"
	"github.com/johnbalvin/gobnb/trace"
)

func test1() {
	client := NewClient("USD", nil)
	// zoom value from 1 - 20, so from the "square" like I said on the coorinates, this represents how much zoom on this squere there is
	zoomvalue := 2
	checkIn := search.Check{
		In:  time.Now().AddDate(0, 0, 1),
		Out: time.Now().AddDate(0, 0, 7),
	}
	//coordinates should be 2 points one from shouth and one from north(if you think it like a square, this points represent the two points of the diagonal from this square)
	coords := search.CoordinatesInput{
		Ne: search.CoordinatesValues{
			Latitude: 11.626466321336217,
			Longitud: -83.16752421667513,
		},
		Sw: search.CoordinatesValues{
			Latitude: 8.565185490351908,
			Longitud: -85.62044033549569,
		},
	}
	results, err := client.SearchFirstPage(zoomvalue, coords, checkIn)
	if err != nil {
		errData := trace.NewOrAdd(1, "main", "test2", err, "")
		log.Println(errData)
		return
	}
	rawJSON, _ := json.MarshalIndent(results, "", "  ")
	fmt.Printf("%s", rawJSON) //in case you don't have write permisions
	if err := os.WriteFile("./searchResult.json", rawJSON, 06444); err != nil {
		log.Println(err)
		return
	}
}
func test2() {
	client := NewClient("EUR", nil)
	// zoom value from 1 - 20, so from the "square" like I said on the coorinates, this represents how much zoom on this squere there is
	zoomvalue := 9
	checkIn := search.Check{
		In:  time.Now().AddDate(0, 0, 1),
		Out: time.Now().AddDate(0, 0, 7),
	}
	//coordinates should be 2 points one from shouth and one from north(if you think it like a square, this points represent the two points of the diagonal from this square)
	coords := search.CoordinatesInput{
		Ne: search.CoordinatesValues{
			Latitude: 0.9539058343440772,
			Longitud: -79.65750456127796,
		},
		Sw: search.CoordinatesValues{
			Latitude: 0.9747511155111473,
			Longitud: -79.64106021485907,
		},
	}
	results, err := client.SearchAll(zoomvalue, coords, checkIn)
	if err != nil {
		errData := trace.NewOrAdd(1, "main", "test2", err, "")
		log.Println(errData)
		return
	}
	rawJSON, _ := json.MarshalIndent(results, "", "  ")
	fmt.Printf("%s", rawJSON) //in case you don't have write permisions
	if err := os.WriteFile("./searchResultAll.json", rawJSON, 06444); err != nil {
		log.Println(err)
		return
	}
}
func test3() {
	client := NewClient("MXN", nil)
	zoomvalue := 2
	checkIn := search.Check{
		In:  time.Now().AddDate(0, 0, 1),
		Out: time.Now().AddDate(0, 0, 7),
	}
	coords := search.CoordinatesInput{
		Ne: search.CoordinatesValues{
			Latitude: -1.03866277790021,
			Longitud: -77.53091734683608,
		},
		Sw: search.CoordinatesValues{
			Latitude: -1.1225978433925647,
			Longitud: -77.59713412765507,
		},
	}
	searchResults, err := client.SearchFirstPage(zoomvalue, coords, checkIn)
	if err != nil {
		errData := trace.NewOrAdd(2, "main", "test2", err, "")
		log.Println(errData)
		return
	}
	rawJSON, _ := json.MarshalIndent(searchResults, "", "  ")
	fmt.Printf("%s", rawJSON) //in case you don't have write permisions
	if err := os.WriteFile("./searchResult.json", rawJSON, 0644); err != nil {
		log.Println(err)
		return
	}
	var datas []details.Data
	for i, result := range searchResults {
		data, err := client.DetailsFromRoomID(result.RoomID)
		if err != nil {
			errData := trace.NewOrAdd(2, "main", "test2", err, "")
			log.Println(errData)
			return
		}
		datas = append(datas, data)
		log.Printf("Progress: %d/%d id: %d\n", i+1, len(searchResults), result.RoomID)
	}
	rawJSON2, _ := json.MarshalIndent(datas, "", "  ")
	fmt.Printf("%s", rawJSON2) //in case you don't have write permisions
	if err := os.WriteFile("./details.json", rawJSON, 0644); err != nil {
		log.Println(err)
		return
	}
}
func test4() {
	//Make sure you have write permissions
	if err := os.MkdirAll("./test/0", 0644); err != nil {
		log.Println("test 1 -> err: ", err)
		return
	}
	client := NewClient("USD", nil)
	mainURL := "https://www.airbnb.com"
	ids, err := client.DetailsMainRoomIds(mainURL)
	if err != nil {
		log.Println("test1 -> err: ", err)
		return
	}
	var datas []details.Data
	for i, id := range ids {
		folderPath := fmt.Sprintf("./test/%d/images", i)
		os.MkdirAll(folderPath, 0644)
		data, err := client.DetailsFromRoomID(id)
		if err != nil {
			log.Println("test1 -> err: ", err)
			continue
		}
		if err := data.SetImages(client.ProxyURL); err != nil {
			log.Println("test1 -> err: ", err)
		}
		for j, img := range data.Images {
			fname3 := fmt.Sprintf("./test/%d/images/%d%s", i, j, img.Extension)
			os.WriteFile(fname3, img.Content, 0644)
		}
		datas = append(datas, data)
		log.Printf("Progress getting data: %d/%d\n", i+1, len(ids))
		filePath := fmt.Sprintf("./test/%d/data.json", i)
		f, err := os.Create(filePath)
		if err != nil {
			log.Println("test 5 -> err: ", err)
			continue
		}
		json.NewEncoder(f).Encode(data)
		f.Close()
	}
	f, _ := os.Create("./test/datas.json")
	json.NewEncoder(f).Encode(datas)
}
func test5() {
	roomURL := "https://www.airbnb.com/rooms/5264493"
	client := DefaultClient()
	data, err := client.DetailsFromRoomURL(roomURL)
	if err != nil {
		log.Println("test:2 -> err: ", err)
		return
	}
	rawJSON, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("%s", rawJSON) //in case you don't have write permisions
	if err := os.WriteFile("./details.json", rawJSON, 0644); err != nil {
		log.Println(err)
		return
	}
}
