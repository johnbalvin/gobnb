package gobnb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/johnbalvin/gobnb/trace"
)

func TestSaveOnDisk() {
	test1()
}

func TestNoImages() ([]Data, error) {
	return test2(false)
}

func TestImages() ([]Data, error) {
	return test2(true)
}

func test1() {
	//Make sure you have write permissions
	if err := os.MkdirAll("./test/0", 0644); err != nil {
		log.Println("test 1 -> err: ", err)
		return
	}
	client := NewClient("USD", nil)
	mainURL := "https://www.airbnb.com"
	ids, err := client.GetMainRoomIds(mainURL)
	if err != nil {
		log.Println("test1 -> err: ", err)
		return
	}
	var datas []Data
	for i, id := range ids {
		folderPath := fmt.Sprintf("./test/%d/images", i)
		os.MkdirAll(folderPath, 0644)
		data, err := client.GetFromRoomID(id)
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

func test2(images bool) ([]Data, error) {
	client := NewClient("USD", nil)
	mainURL := "https://www.airbnb.com"
	ids, err := client.GetMainRoomIds(mainURL)
	if err != nil {
		return nil, trace.NewOrAdd(2, "main", "test2", err, "")
	}
	var datas []Data
	for i, id := range ids {
		folderPath := fmt.Sprintf("./test/%d/images", i)
		os.MkdirAll(folderPath, 0644)
		data, err := client.GetFromRoomID(id)
		if err != nil {
			errData := trace.NewOrAdd(3, "main", "test2", err, "")
			log.Println(errData)
			continue
		}
		if images {
			if err := data.SetImages(client.ProxyURL); err != nil {
				errData := trace.NewOrAdd(4, "main", "test2", err, "")
				log.Println(errData)
			}
		}
		datas = append(datas, data)
		log.Printf("Progress getting data: %d/%d\n", i+1, len(ids))
	}
	return datas, nil
}
