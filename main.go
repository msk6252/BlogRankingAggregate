package main

import (
	//"os"
	//"log"
	"fmt"

	"github.com/msk6252/BlogRankingAggregate/tools/analytics"
	//"github.com/msk6252/BlogRankingAggregate/tools/twitter"
)

func main() {
  result := analytics.GetAnalytics()
  hashMap := analytics.CreatePathHashMap(result)
  orderHashMap := analytics.MaxPVOrderMap(hashMap, 3)
  fmt.Println(orderHashMap)
}
