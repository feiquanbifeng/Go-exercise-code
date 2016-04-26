package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    _ "github.com/mattn/go-oci8"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "os"
)

type Result struct {
    Location   Location `json:"location"`
    Precise    int      `json:"precise"`
    Confidence int      `json:"confidence"`
    Level      string   `json:"level"`
}

type Location struct {
    Lng float64 `json:"lng"`
    Lat float64 `json:"lat"`
}

type LngAndLat struct {
    Status int    `json:"status"`
    Result Result `json:"result"`
}

func httpGet(address string) (lng, lat float64, err error) {
    urlStr := "http://api.map.baidu.com/geocoder/v2/?address=" + url.QueryEscape(address) + "&output=json&ak=xFzaWRyWuNOO2aWtesCwYeCB"
    resp, err := http.Get(urlStr)

    if err != nil {
        return 0, 0, err
    }

    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return parseResult(body)
}

func parseResult(jsonStream []byte) (lng, lat float64, err error) {
    var lngLat LngAndLat
    err = json.Unmarshal(jsonStream, &lngLat)
    if err != nil {
        return 0, 0, err
    }
    if lngLat.Status == 0 {
        return lngLat.Result.Location.Lng, lngLat.Result.Location.Lat, nil
    }
    return 0, 0, nil
}

func main() {

    // 字符集
    os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8")

    // 为log添加短文件名,方便查看行数
    log.SetFlags(log.Lshortfile | log.LstdFlags)

    fmt.Println("Oracle Driver example")

    // 用户名/密码@实例名  跟sqlplus的conn命令类似
    db, err := sql.Open("oci8", "bms/123456@10.0.0.203:1521/XE")
    if err != nil {
        log.Fatal(err)
    }
    rows, err := db.Query("SELECT ID, ADDRESS FROM BMS_MERCH_USER_ADDRESS WHERE ID = 449")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    for rows.Next() {
        var address string
        var id int64
        rows.Scan(&id, &address)

        //converts a  string from UTF-8 to gbk encoding.
        lng, lat, er := httpGet(address)
        if er != nil {
            continue
        }

        // update lng
        if lng != 0 && lat != 0 {
            upd := fmt.Sprintf("UPDATE BMS_MERCH_USER_ADDRESS SET LONGITUDE = %f, LATITUDE = %f WHERE ID = %d", lng, lat, id)
            fmt.Println(upd)
            db.Exec(upd)
        }
    }

    rows.Close()

}
