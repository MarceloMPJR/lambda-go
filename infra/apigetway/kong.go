package apigetway

import (
	"encoding/json"
	"errors"
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

func (k *Kong) GetConsumerInfo(ConsumerInfoInput) (out ConsumerInfoOutput) {
	resp, err := http.Get(k.url)
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
