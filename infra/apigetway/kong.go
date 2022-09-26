package apigetway

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Kong struct {
	url string
}

func NewKong(url string) *Kong {
	return &Kong{url: url}
}

const KongCommunicationFailure = "communication failure"

type ConsumerInfoResponse struct {
	Key string `json:"key"`
}

func (k *Kong) GetConsumerInfo(in ConsumerInfoInput) (out ConsumerInfoOutput) {
	resp, err := http.Get(fmt.Sprintf("%s/consumers/%s/%s", k.url, in.UserName, in.AuthType))
	if err != nil || resp.StatusCode > 299 {
		out.Error = errors.New(KongCommunicationFailure)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		out.Error = err
		return
	}

	respKong := ConsumerInfoResponse{}
	json.Unmarshal(body, &respKong)

	out.Key = respKong.Key
	return
}
