package analytics

import (
  "io/ioutil"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
  "github.com/joho/godotenv"
  "os"
  "log"
  analytics "google.golang.org/api/analytics/v3"
)


func getAnalytics() (*analytics.GaData, error) {
  key, _ := ioutil.ReadFile("./secret.json")

  viewId := GetEnv("ViewID")

  jwtConf, err := google.JWTConfigFromJSON(
    key,
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
