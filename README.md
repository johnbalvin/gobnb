# Airbnb scraper in Go

## Overview
This project is an open-source tool developed in Golang for extracting product information from Airbnb. It's designed to be fast, efficient, and easy to use, making it an ideal solution for developers looking for Airbnb product data.

## Features
- Full search support
- Extracts detailed product information from Airbnb
- Implemented in Go for performance and efficiency
- Easy to integrate with existing Go projects
- The code is optimize to work on this format: ```https://www.airbnb.com/rooms/[roomID]```

## Examples

### Quick testing

```Go
    package main

    import (
        "encoding/json"
        "fmt"
        "log"
        "os"
        "time"
        "github.com/johnbalvin/gobnb"
        "github.com/johnbalvin/gobnb/search"
    )
    func main(){
        client := gobnb.NewClient("USD", nil)
        // zoom value from 1 - 20, so from the "square" like I said on the coorinates
        // This represents how much zoom there is on this square.
        zoomvalue := 2
        check := search.Check{
            In:  time.Now().AddDate(0, 0, 1),
            Out: time.Now().AddDate(0, 0, 7),
        }
        //coordinates should be 2 points one from shouth and one from north(if you think it like a square
        //this presents the two points of the diagonal from this square)
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
        results, err := client.SearchFirstPage(zoomvalue, coords, check)
        if err != nil {
            log.Println(err)
            return
        }
        rawJSON, _ := json.MarshalIndent(results, "", "  ")
        fmt.Printf("%s", rawJSON) //in case you don't have write permisions
        if err := os.WriteFile("./searchResult.json", rawJSON, 06444); err != nil {
            log.Println(err)
            return
        }
    }
```

```Go
    package main

    import (
        "encoding/json"
        "fmt"
        "log"
        "os"
	    "time"
        "github.com/johnbalvin/gobnb"
        "github.com/johnbalvin/gobnb/search"
    )
    func main(){
        client := gobnb.NewClient("EUR", nil)
        // zoom value from 1 - 20, so from the "square" like I said on the coorinates
        // This represents how much zoom there is on this square.
        zoomvalue := 15
        checkIn := search.Check{
            In:  time.Now().AddDate(0, 0, 1),
            Out: time.Now().AddDate(0, 0, 7),
        }
        //coordinates should be 2 points one from shouth and one from north(if you think it like a square
        //this presents the two points of the diagonal from this square)
        coords := search.CoordinatesInput{
            Sw: search.CoordinatesValues{
                Latitude: 0.9539058343440772,
                Longitud: -79.65750456127796,
            },
            Ne: search.CoordinatesValues{
                Latitude: 0.9747511155111473,
                Longitud: -79.64106021485907,
            },
        }
        results, err := client.SearchAll(zoomvalue, coords, checkIn)
        if err != nil {
            log.Println(err)
            return
        }
        rawJSON, _ := json.MarshalIndent(results, "", "  ")
        fmt.Printf("%s", rawJSON) //in case you don't have write permisions
        if err := os.WriteFile("./searchResultAll.json", rawJSON, 06444); err != nil {
            log.Println(err)
            return
        }
    }
```
```Go
    package main

    import (
        "encoding/json"
        "fmt"
        "log"
        "os"
	    "time"
        "github.com/johnbalvin/gobnb"
        "github.com/johnbalvin/gobnb/search"
    )
    func main(){
        client := gobnb.NewClient("MXN", nil)
        // zoom value from 1 - 20, so from the "square" like I said on the coorinates
        // This represents how much zoom there is on this square.
        zoomvalue := 2
        checkInOut := search.Check{
            In:  time.Now().AddDate(0, 0, 1),
            Out: time.Now().AddDate(0, 0, 7),
        }
        //coordinates should be 2 points one from shouth and one from north(if you think it like a square
        //this presents the two points of the diagonal from this square)
        coords := search.CoordinatesInput{
            Sw: search.CoordinatesValues{
                Latitude: -1.03866277790021,
                Longitud: -77.53091734683608,
            },
            Ne: search.CoordinatesValues{
                Latitude: -1.1225978433925647,
                Longitud: -77.59713412765507,
            },
        }
        searchResults, err := client.SearchFirstPage(zoomvalue, coords, checkInOut)
        if err != nil {
            log.Println(err)
            return
        }
        rawJSON, _ := json.MarshalIndent(results, "", "  ")
        fmt.Printf("%s", rawJSON) //in case you don't have write permisions
        if err := os.WriteFile("./searchResultAll.json", rawJSON, 06444); err != nil {
            log.Println(err)
            return
        }
        var datas []details.Data
        for i, result := range searchResults {
            data, err := client.DetailsFromRoomID(result.RoomID, checkInOut)
            if err != nil {
               log.Println(err)
			   return
		    }
		    datas = append(datas, data)
		    log.Printf("Progress: %d/%d id: %d\n", i+1, len(searchResults), result.RoomID)
        }   
        rawJSON2, _ := json.MarshalIndent(datas, "", "  ")
        fmt.Printf("%s", rawJSON2) //in case you don't have write permisions
        if err := os.WriteFile("./details.json", rawJSON2, 0644); err != nil {
		   log.Println(err)
		   return
        }
    }
```

### Basic data

```Go
    package main

    import (
        "encoding/json"
        "fmt"
        "log"
        "os"
        "github.com/johnbalvin/gobnb"
    )
    func main(){
        roomURL:="https://www.airbnb.com/rooms/[roomID]"
        client := gobnb.DefaultClient()
        checkInOut := search.Check{
            In:  time.Now().AddDate(0, 0, 1),
            Out: time.Now().AddDate(0, 0, 7),
        }
        data, err := client.DetailsFromRoomURL(roomURL, checkInOut)
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
```

```Go
    package main

    import (
        "encoding/json"
        "fmt"
        "log"
        "os"
        "github.com/johnbalvin/gobnb"
    )
    func main(){
        romID:=[roomID]
        checkInOut := search.Check{
            In:  time.Now().AddDate(0, 0, 1),
            Out: time.Now().AddDate(0, 0, 7),
        }
        client := gobnb.DefaultClient()
        data, err := client.DetailsFromRoomID(romID, checkInOut)
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
```
### Basic data and images
```Go
    package main

    import (
        "encoding/json"
        "log"
        "os"
        "github.com/johnbalvin/gobnb"
    )
    func main(){
        //you need to have write permissions, the result will be save inside folder "test"
        if err := os.MkdirAll("./test/images", 0644); err != nil {
            log.Println("test 1 -> err: ", err)
            return
        }
        roomURL:="https://www.airbnb.com/rooms/[roomID]"
        client := gobnb.DefaultClient()
        checkInOut := search.Check{
          //you can ommit the checkin, checkout date if you dont want the price
          //  In:  time.Now().AddDate(0, 0, 1),
          //  Out: time.Now().AddDate(0, 0, 7),
        }
        data,  err := client.DetailsFromRoomURL(roomURL, checkInOut)
        if err != nil {
            log.Println("test:2 -> err: ", err)
            return
        }
        if err := data.SetImages(client.ProxyURL); err != nil {
            log.Println("test:3 -> err: ", err)
            return
        }
		for j, img := range data.Images {
			fname3 := fmt.Sprintf("./test/images/%d%s", j, img.Extension)
			os.WriteFile(fname3, img.Content, 0644)
		}
        f, err := os.Create("./test/data.json")
        if err != nil {
            log.Println("test:4 -> err: ", err)
            return
        }
        json.NewEncoder(f).Encode(data)
    }
```

### With proxy

```Go
    package main

    import (
        "encoding/json"
        "log"
        "os"
        "github.com/johnbalvin/gobnb"
    )
    func main(){
        //you need to have write permissions, the result will be save inside folder "test"
        if err := os.MkdirAll("./test/images", 0644); err != nil {
            log.Println("test 1 -> err: ", err)
            return
        }
        proxyURL, err := gobnb.ParseProxy("http://[IP | domain]:[port]", "username", "password")
        if err != nil {
            log.Println("test:1 -> err: ", err)
            return
        }
        client := gobnb.NewClient("MXN", proxyURL)
        roomURL:="https://www.airbnb.com/rooms/[roomID]"
        checkInOut := search.Check{
            In:  time.Now().AddDate(0, 0, 1),
            Out: time.Now().AddDate(0, 0, 7),
        }
        data,  err := client.DetailsFromRoomURL(roomURL, checkInOut)
        if err != nil {
            log.Println("test:2 -> err: ", err)
            continue
        }
        if err := data.SetImages(client.ProxyURL); err != nil {
            log.Println("test:3 -> err: ", err)
            return
        }
		for j, img := range data.Images {
			fname3 := fmt.Sprintf("./test/images/%d%s", j, img.Extension)
			os.WriteFile(fname3, img.Content, 0644)
		}
        f, err := os.Create("./test/data.json")
        if err != nil {
            log.Println("test:4 -> err: ", err)
            return
        }
        json.NewEncoder(f).Encode(data)
    }
```