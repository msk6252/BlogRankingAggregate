package twitter

import (
 "log"
 "os"

 "github.com/joho/godotenv"
 "github.com/dghubble/go-twitter/twitter"
 "github.com/dghubble/oauth1"
)

func main() {
  consumerKey := GetEnv("CONSUMER_KEY")
  consumerSecret := GetEnv("CONSUMER_KEY_SECRET")
  accessToken := GetEnv("ACCESS_TOKEN")
  accessSecret := GetEnv("ACCESS_TOKEN_SECRET")

  config := oauth1.NewConfig(consumerKey, consumerSecret)
  token := oauth1.NewToken(accessToken, accessSecret)

  httpClient := config.Client(oauth1.NoContext, token)

  client := twitter.NewClient(httpClient)

  txt := "ツイートする本文"

  //tweet, res, err := client.Statuses.Update("ツイートする本文", nil)
  _, r, e := client.Statuses.Update(txt, nil)
  if e != nil {
   log.Println("err", e)
  }
  // ツイート情報とhttpレスポンス
  log.Println("tweet", r)
  log.Println("res", r)
}

func GetEnv(key string) string {
  err := godotenv.Load()

  if err != nil {
    log.Fatalln("Get env error:", err)
  }

  value := os.Getenv(key)

  if value == "" {
    log.Fatalln("Empty .env")
  }

  return value
}
