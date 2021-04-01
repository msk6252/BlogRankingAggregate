package main

import (
  //"os"
  "log"
  "fmt"

  "github.com/msk6252/BlogRankingAggregate/tools/analytics"
  //"github.com/msk6252/BlogRankingAggregate/tools/twitter"
)

func main() {
  result, err := analytics.GetAnalytics()
  if err != nil {
    log.Fatalln("Get Secret Error")
  }
  fmt.Println(result)
}
