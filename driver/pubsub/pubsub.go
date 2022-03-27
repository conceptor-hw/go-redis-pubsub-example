package pubsub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var (
	R_ctx    = context.Background()
	R_client *redis.Client
)

var pub_channel string = "go_channel"
var pub_prover_channel string = "prover_sub_channel"
var pub *redis.IntCmd = nil

type ProverMessage struct {
	Previous_block_hash string
	Block_height        uint32
	Block_timestamp     int64
	Difficulty_target   uint64
}

func init() {
	R_client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := R_client.Ping(R_ctx).Result()
	fmt.Println(pong, err)

}

// func PubBinaryData(channel string, data interface{}) error {

// 	fmt.Println("======== binnary pub info is ==========", data)
// 	buffer1 := new(bytes.Buffer)
// 	encoder1 := gob.NewEncoder(buffer1)
// 	err := encoder1.Encode(data)
// 	fmt.Println("binary data is ", buffer1.Bytes())

// 	err = R_client.Publish(R_ctx, channel, buffer1.Bytes()).Err()
// 	if err != nil {
// 		return errors.New("publish Data wrong... " + err.Error())
// 	}
// 	return nil
// }

func PubNormalMsg(channel string, data interface{}) error {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	messageString := string(jsonBytes)
	fmt.Println("messagei is", messageString)
	err = R_client.Publish(R_ctx, channel, messageString).Err()
	if err != nil {
		return errors.New("publish normal message wrong... " + err.Error())
	}
	return nil
}
