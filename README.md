# Airbnb Information Extractor in Go

## Overview
This project is an open-source tool developed in Golang for extracting product information from Airbnb. It's designed to be fast, efficient, and easy to use, making it an ideal solution for developers looking for Airbnb product data.

## Features
- Extracts detailed product information from Airbnb
- Implemented in Go for performance and efficiency
- Easy to integrate with existing Go projects
- The code is optimize to work on this format: ```https://www.airbnb.com/rooms/[roomID]```

## Examples

### Quick testing

```Go
    package main

    import (
        "github.com/johnbalvin/gobnb"
    )
    func main(){
        //you need to have write permissions, the result will be save inside folder "test"
        gobnb.TestSaveOnDisk()
    }
```

```Go
    package main

    import (
        "github.com/johnbalvin/gobnb"
    )
    func main(){
        //If you have write permissions errors with the project, try printing the data at least
        datas,err:=gobnb.TestNoImages()
        if err!=nil{
            log.Println("err",err)
            return
        }
        log.Printf("allDatas: %+v\n",datas)
    }
```


### Basic data

```Go
    package main

    import (
        "encoding/json"
        "log"
        "os"
        "github.com/johnbalvin/gobnb"
    )
    func main(){
        roomURL:="https://www.airbnb.com/rooms/[roomID]"
        client := gobnb.DefaulClient()
        data, err := client.GetFromRoomURL(roomURL)
        if err != nil {
            log.Println("test:2 -> err: ", err)
            return
        }
        log.Printf("data: %+v\n",data)
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
        client := gobnb.DefaulClient()
        data,  err := client.GetFromRoomURL(roomURL)
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
        data,  err := client.GetFromRoomURL(roomURL)
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