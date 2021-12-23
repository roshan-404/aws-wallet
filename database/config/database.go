package config

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-redis/redis"
)

var DB_client *dynamodb.Client
var S3session *s3.S3
var RedisClient *redis.Client


func ConnectionDB() (*dynamodb.Client, error) {
	creds := credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
	
	//loading configs
	cfg, err  := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		return nil, err
	}

	//creating session
	S3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})))

	//creating client 
	DB_client = dynamodb.NewFromConfig(cfg)

	var tableName = os.Getenv("TABLE_NAME")
	// check if table exists
	resp, errr := GetTableInfo(DB_client, tableName)
	if errr != nil {
		// create a new table
		fmt.Println("Creating new table : " + tableName)
		_, err := createTable(DB_client, tableName)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Table created!")
		}
	} else {
		fmt.Println("Table Size (bytes)", resp.Table.TableSizeBytes)
	}

	RedisClient = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

	return DB_client, nil
}
