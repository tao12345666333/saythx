package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"os"
)

type ThxData struct {
	From    string `json:"from"`
	Content string `json:"content"`
	To      string `json:"to"`
}

type HonorData struct {
	Name  string `json:"name"`
	Score string `json:"score"`
}

type HonorDataslice struct {
	HonorDatas []HonorData
}

func SayThx(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data ThxData
		json.NewDecoder(r.Body).Decode(&data)
		key := uuid.NewV4()

		c, err := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":6379")
		if err != nil {
			log.Panicln(err)
		}
		defer c.Close()

		if _, err := c.Do("HMSET", redis.Args{}.Add(key).AddFlat(&data)...); err != nil {
			log.Panicln(err)
		}
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func GetHonorList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c, err := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":6379")
		if err != nil {
			log.Panicln(err)
		}
		defer c.Close()

		data, err := redis.Strings(c.Do("ZREVRANGEByScore", "honorlist", "+inf", "-inf", "WITHSCORES", "limit", 0, 20))
		if err != nil {
			log.Panicln(err)
		}

		var d HonorDataslice
		for i := 0; i < len(data); i += 2 {
			d.HonorDatas = append(d.HonorDatas, HonorData{Name: data[i], Score: data[i+1]})
		}

		j, err := json.Marshal(d)
		if err != nil {
			log.Panicln(err)
		}

		fmt.Fprint(w, string(j))
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}

func main() {
	http.HandleFunc("/api/v1/list", GetHonorList)
	http.HandleFunc("/api/v1/saythx", SayThx)
	log.Println("server starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
