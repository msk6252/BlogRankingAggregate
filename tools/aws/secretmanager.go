package aws

import (
  "github.com/aws/aws-sdk-go/service/secretsmanager"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "encoding/json"
)

func GetSecret(secretName string, key string, REGION string) (string, error){
  svc := secretsmanager.New(session.New(),
    aws.NewConfig().WithRegion(REGION))

  input := &secretsmanager.GetSecretValueInput{
    SecretId:     aws.String(secretName),
    VersionStage: aws.String("AWSCURRENT"),
  }

  result, err := svc.GetSecretValue(input)
  if err != nil {
    return "", err
  }

  secretString := aws.StringValue(result.SecretString)
  res := make(map[string]interface{})
  if err := json.Unmarshal([]byte(secretString), &res); err != nil {
    return "", err
  }
  return res[key].(string), nil
}
