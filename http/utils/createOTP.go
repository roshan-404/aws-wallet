package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func main(id string) {
    num := rand.Intn(1000000)

    json, err := json.Marshal(int(num))
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(json)

    err = client.Set("id1234", json, 0).Err()
    if err != nil {
        fmt.Println(err)
    }
    val, err := client.Get("id1234").Result()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(val)
}