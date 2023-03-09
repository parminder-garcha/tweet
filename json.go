package json

import (
  "os"
  "io/ioutil" 
  "encoding/json"
)

type TweetInfo struct {
        Tweet []struct {
                Id   string `json:"id"`
                Text string `json:"text"`
        } `json:"data"`
}


func readJSON(filename string) TweetInfo {

   file, err := os.Open(filename)

   if err != nil {
      panic(err)
   }

  defer file.Close()

  byteResult, _ := ioutil.ReadAll(file)

  data := TweetInfo{}

  json.Unmarshal([]byte(byteResult), data)

  return data
}
