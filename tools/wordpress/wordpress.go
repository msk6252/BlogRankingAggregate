package wordpress

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/msk6252/BlogRankingAggregate/tools/aws"
)

type Article struct {
  Id   int    `json:"id"`
  Link  string `json:"link"`
  Title struct {
    Rendered string `json:"rendered"`
  }
}

func GetBlogTitle(blogId string) string {
  blog_url, err := aws.GetSecret("BLOG_RANKING_SECRET", "BLOG_JSON_BASE_URL", "ap-northeast-1")
  if err != nil {
    log.Fatalln(err)
  }

  resp, _ := http.Get(blog_url + blogId)
  defer resp.Body.Close()
  byteArray, _ := ioutil.ReadAll(resp.Body)
  jsonBytes := ([]byte)(byteArray)
  data := make(map[string]interface{})
  if err := json.Unmarshal(jsonBytes, &data);  err != nil {
    log.Fatalln("JSON Unmarshal error:", err)
  }
  log.Println(data)

  return data["title"].(map[string]interface{})["rendered"].(string)
}

