package analytics

import (
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
  "log"
  "github.com/msk6252/BlogRankingAggregate/tools/aws"
  analytics "google.golang.org/api/analytics/v3"
)

const REGION = "ap-northeast-1"

func GetAnalytics() (*analytics.GaData, error) {
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
    log.Fatalln(err)
  }

  httpClient := jwtConf.Client(oauth2.NoContext)

  svc, err := analytics.New(httpClient)

  if err != nil {
    log.Fatalln(err)
  }

  result, err := svc.Data.Ga.Get("ga:" + viewId, "7daysAgo", "today", "ga:pageviews").Dimensions("ga:pagePath").Filters("ga:pagePath=~^/archives/").Sort("-ga:pageviews").MaxResults(20).Do()

  return result, err
}
