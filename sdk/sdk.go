package sdk

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ISDK interface {
	IComicsSDK
}

type ClientSDK struct{}

func NewClientSDK() ClientSDK {
	return ClientSDK{}
}

func (c ClientSDK) get(url string, out interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("err : %s", err.Error())
		return fmt.Errorf("ooopsss an error occurred, please try again")
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		log.Fatalf("err : %s", err.Error())
		return fmt.Errorf("ooopsss! an error occurred, please try again")
	}
	return nil
}
