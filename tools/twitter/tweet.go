package twitter

import (
	"log"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/msk6252/BlogRankingAggregate/tools/aws"
	"github.com/msk6252/BlogRankingAggregate/tools/wordpress"
)

const REGION = "ap-northeast-1"
const SECRET_NAME = "BLOG_RANKING_SECRET"

func Tweet(hashMap map[int]map[string]int, rank int) {
  consumerKey, err := aws.GetSecret(SECRET_NAME, "TWITTER_CONSUMER_KEY", REGION)
  consumerSecret, err := aws.GetSecret(SECRET_NAME, "TWITTER_CONSUMER_KEY_SECRET", REGION)
  accessToken, err := aws.GetSecret(SECRET_NAME, "TWITTER_ACCESS_TOKEN", REGION)
  accessSecret, err := aws.GetSecret(SECRET_NAME, "TWITTER_ACCESS_TOKEN_SECRET", REGION)

  if err != nil {
    log.Fatalln(err)
  }

  config := oauth1.NewConfig(consumerKey, consumerSecret)
  token := oauth1.NewToken(accessToken, accessSecret)

  httpClient := config.Client(oauth1.NoContext, token)

  client := twitter.NewClient(httpClient)
  for i:=rank; i>=1; i-- {
    tweet := CreateTweet(hashMap[i], i)

    //tweet, res, err := client.Statuses.Update("ツイートする本文", nil)
    _, r, e := client.Statuses.Update(tweet, nil)
    if e != nil {
     log.Println("err", e)
    }
    // ツイート情報とhttpレスポンス
    log.Println("tweet", r)
    log.Println("res", r)
  }
}

func CreateTweet(hashMap map[string]int, rank int) string {
  BLOG_BASE_URL, err := aws.GetSecret(SECRET_NAME, "BLOG_BASE_URL", REGION)

  if err != nil {
    log.Fatalln(err)
  }

  strRank := strconv.Itoa(rank)
  tweet := ""
  for key := range hashMap {
    tweet =
      "今週１週間で見られた記事ランキング!\n\n" +
      "【👑" + strRank + "位】\n\n" +
      wordpress.GetBlogTitle(key) + "\n" +
      BLOG_BASE_URL + key + "\n\n"
  }
  return tweet
}
