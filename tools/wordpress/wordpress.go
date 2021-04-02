package wordpress

import (
	"encoding/json"
	"fmt"
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

func GetBlogTitle() map[string]interface {
  BLOG_BASE_URL, err := aws.GetSecret("BLOG_RANKING_SECRET", "BLOG_BASE_URL", "ap-northeast-1")
  if err != nil {
    log.Fatalln(err)
  }
  resp, _ := http.Get(BLOG_BASE_URL + "1813")
  defer resp.Body.Close()
  byteArray, _ := ioutil.ReadAll(resp.Body)
  jsonBytes := ([]byte)(byteArray)
  data := make(map[string]interface{})
  if err := json.Unmarshal(jsonBytes, &data);  err != nil {
    fmt.Println("JSON Unmarshal error:", err)
    return
  }
  return data["title"].(map[string]interface{}))
}

