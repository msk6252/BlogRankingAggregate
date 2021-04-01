package analytics

import (
	"log"
	"regexp"
  "strconv"

	"github.com/msk6252/BlogRankingAggregate/tools/aws"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	analytics "google.golang.org/api/analytics/v3"
)

const REGION = "ap-northeast-1"

func GetAnalytics() *analytics.GaData {
  key, err := aws.GetSecret("BLOG_RANKING_SECRET", "SECRET_JSON", REGION)
  if err != nil {
    log.Fatalln("GET SECRET ERROR: ", err)
  }

  byteKey := ([]byte)(key)
  viewId, err := aws.GetSecret("BLOG_RANKING_SECRET", "ANALYTICS_VIEW_ID", REGION)

  jwtConf, err := google.JWTConfigFromJSON(
    byteKey,
    analytics.AnalyticsReadonlyScope,
  )

  if err != nil {
    log.Fatalln("Authenticate Google Auth:", err)
  }

  httpClient := jwtConf.Client(oauth2.NoContext)

  svc, err := analytics.New(httpClient)

  if err != nil {
    log.Fatalln(err)
  }

  result, err := svc.Data.Ga.Get("ga:" + viewId, "7daysAgo", "today", "ga:pageviews").Dimensions("ga:pagePath").Filters("ga:pagePath=~^/archives/").Sort("-ga:pageviews").MaxResults(20).Do()

  if err != nil {
    log.Fatalln(err)
  }

  return result
}

func CreatePathHashMap(aryURL *analytics.GaData) map[string]int {
  rep := regexp.MustCompile(`\?.*|/amp.*`)
  hashmap := map[string]int{}

  for _, row := range aryURL.Rows {
    mapKey := rep.ReplaceAllString(row[0], "")
    pageview, err := strconv.Atoi(row[1])
    if err != nil {
      log.Fatalln("STRCONV Error: ", err)
    }
    hashmap[mapKey] = hashmap[mapKey] + pageview
  }

  return hashmap
}

func MaxPVOrderMap(hashMap map[string]int, rank int) map[int]map[string]int {
  orderHashMap := map[int]map[string]int{}
  deleteUrl := ""
  for i:=1; i<=rank; i++ {
    max := -1
    for url, pageview := range hashMap {
      if max <= pageview {
        max = pageview
        deleteUrl = url
        orderHashMap[i] = map[string]int{url : pageview}
      }
    }
    delete(hashMap,deleteUrl)
  }
  return orderHashMap
}
