package main

import (

	"github.com/msk6252/BlogRankingAggregate/tools/analytics"
	"github.com/msk6252/BlogRankingAggregate/tools/twitter"
)

const RANK = 3

func main() {
  result := analytics.GetAnalytics()

  hashMap := analytics.CreatePathHashMap(result)

  orderHashMap := analytics.MaxPVOrderMap(hashMap, RANK)

  twitter.Tweet(orderHashMap, RANK)
}
