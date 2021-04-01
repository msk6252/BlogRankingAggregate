package main

import (
  //"os"
  //"log"

  //"github.com/msk6252/BlogRankingAggregate/tools/analytics"
  "github.com/msk6252/BlogRankingAggregate/tools/aws"
  //"github.com/msk6252/BlogRankingAggregate/tools/twitter"
  "fmt"
)

func main() {
  fmt.Println(aws.GetSecret("GO_AWS_SECRET_MANAGER", "AWS_GO_SECRET_KEY", "ap-northeast-1"))
}
