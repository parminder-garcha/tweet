package main

import (
	  analyse "tweets/analyser"
	  db      "tweets/db"
)

func main() {
	
   dbConn, err := db.ConnectDB("localhost", "twitter", "5432")
   
   if err != nil {
   	  panic(err)
   }

   retVal := analyse.Sentiment(dbConn)

   if retVal != true {
   	  panic("")
   }
}
