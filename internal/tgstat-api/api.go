package tgstat_api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/ProVitSer/tg-vrode-rabotaet-no/config"
	"github.com/ProVitSer/tg-vrode-rabotaet-no/internal/logger"
	tgstat "github.com/helios-ag/tgstat-go"
	"github.com/helios-ag/tgstat-go/callback"
	"github.com/helios-ag/tgstat-go/posts"
)

const MAX_LIMIT_POSTS = "20"

func GetCallbackInfo() (*tgstat.GetCallbackResponse, error) {
	resp, _, errInfo := callback.GetCallbackInfo(context.Background())

	if errInfo != nil {
		return nil, errInfo
	}
	return resp, nil
}

func SetCallbackSubscribeWord() {

	cbReq, res, setErr := callback.SetCallback(context.Background(), config.GlobalConfig.SubscribeWordCallbackUrl)

	if setErr != nil {
		fmt.Printf("error SetCallbackSubscribeWord: %v\n", setErr)
		fmt.Printf("status: %v\n", res.Status)
		fmt.Printf("status: %d\n", res.StatusCode)
		fmt.Printf("status: %v\n", res.Body)
		os.Exit(1)
	}

	fmt.Println(cbReq.VerifyCode)
}

func PostSearch(searchQ string) ([]tgstat.PostSearchResultItem, error) {

	logger.GetLogger().Printf("Строка поиска: %v", searchQ)

	req := posts.PostSearchRequest{
		Q:     searchQ,
		Limit: string2int(MAX_LIMIT_POSTS),
	}

	info, _, err := posts.PostSearch(context.Background(), req)

	if err != nil {
		logger.GetLogger().Printf("error PostSearch %v\n", err)
		return nil, err

	}

	postJSON, err := json.MarshalIndent(info.Response.Items, "", "  ")
	if err != nil {
		logger.GetLogger().Printf("Ошибка при сериализации JSON: %v\n", err)
	}

	logger.GetLogger().Printf("Post Info:\n%s\n", postJSON)

	return info.Response.Items, nil
}

func SubscriptionsList() ([]tgstat.Subscription, error) {

	req := callback.SubscriptionsListRequest{
		SubscriptionId:   new(string),
		SubscriptionType: new(string),
	}

	resp, _, err := callback.SubscriptionsList(context.Background(), req)

	if err != nil {
		logger.GetLogger().Printf("error SubscriptionsList: %v\n", err)
		return nil, err

	}

	sub, err := json.MarshalIndent(resp.Response, "", "  ")
	if err != nil {
		logger.GetLogger().Printf("Ошибка при сериализации JSON: %v\n", err)
	}

	logger.GetLogger().Printf("SubscriptionsList Info:\n%s\n", sub)

	return resp.Response.Subscriptions, nil
}

func SubscribeWord(q string, peerTypes string) (*int, error) {

	req := callback.SubscribeWordRequest{
		Q:          q,
		EventTypes: "new_post",
		PeerTypes:  &peerTypes,
	}

	resp, _, err := callback.SubscribeWord(context.Background(), req)

	if err != nil {
		logger.GetLogger().Printf("error SubscribeWord: %v\n", err)
		return nil, err
	}

	subInfo, err := json.MarshalIndent(resp.Response, "", "  ")
	if err != nil {
		logger.GetLogger().Printf("Ошибка при сериализации JSON: %v\n", err)
	}

	logger.GetLogger().Printf("SubscribeWord Info:\n%s\n", subInfo)

	return &resp.Response.SubscriptionId, nil

}

func Unsubscribe(subscriptionId string) error {
	_, _, err := callback.Unsubscribe(context.Background(), subscriptionId)

	if err != nil {
		logger.GetLogger().Printf("error Unsubscribe: %v\n", err)
		return err
	}

	return nil
}

func string2int(v string) *int {
	res, _ := strconv.Atoi(v)
	return &res
}
