package main

import (
    "time"
    // DISABLE at Heroku
    // "io"
    "strings"
    "encoding/csv"
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/patrickmn/go-cache"
)

var _cacheHandler *cache.Cache = nil

// https://data.gov.tw/dataset/152408
// https://data.nhi.gov.tw/resource/Nhi_Fst/Fstdata.csv
func apiGetFstdata() (output gin.H) {
    output = gin.H{
        "status": false,
        "errorCode": 0,
        "cache": false,
        "header": map[string]int{},
        "data": nil,
    }
    // https://github.com/patrickmn/go-cache

    if _cacheHandler == nil {
        _cacheHandler = cache.New(5*time.Minute, 10*time.Minute)
    }
    cacheKey := "Fstdata_csv"
    raw, found := _cacheHandler.Get(cacheKey)
    if !found {
        response, err := http.Get("https://data.nhi.gov.tw/resource/Nhi_Fst/Fstdata.csv")
        if err != nil || response.StatusCode != http.StatusOK {
            output["errorCode"] = 1
            return
        }
        reader := response.Body
        defer reader.Close()
        //
        // DISABLE at Heroku
        //
        rawApi := ""
        // rawApi, err := io.ReadAll(reader)
        // if err != nil {
        //     output["errorCode"] = 2
        //     return
        // }
        stringReader := strings.NewReader(string(rawApi))
        lines, err := csv.NewReader(stringReader).ReadAll()
        if err != nil {
            output["errorCode"] = 3
            return
        }
        _data := make([][]string, 0)
        _header_length := 0
        for index, line := range lines {
            if index == 0 {
                _header := map[string]int{}
                for i, name := range line {
                    _header[ string(name) ] = i
                }
                _header_length = len(_header)
                if _header_length == 0 {
                    output["errorCode"] = 4
                    return
                }
                output["header"] = _header
            } else {
                _field := make([]string, 0, _header_length)
                for _, value := range line {
                    _field = append(_field, value)
                }
                _data = append(_data, _field)
            }
        }
        output["data"] = _data
        output["status"] = len(_data) > 0
        output["update"] = string(time.Now().Format(time.RFC3339))
        cacheString, err := json.Marshal(output)
        if err != nil {
            output["errorCode"] = 5
            return
        }
        _cacheHandler.Set(cacheKey, string(cacheString), cache.DefaultExpiration)
    } else {
        jsonString, test:= raw.(string)
        if test {
            if err := json.Unmarshal([]byte(jsonString), &output); err != nil {
                output["errorCode"] = 6
            }
        }
        output["cache"] = true
    }
    return
}

