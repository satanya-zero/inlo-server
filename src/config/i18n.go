package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Lang struct {
	RouteEventMsg       string `json:"route_event_msg"`
	SetOptMsg           string `json:"setting_option_msg"`
	UnspecOptMsg        string `json:"unspecified_option_msg"`
	IllegalOptMsg       string `json:"illegal_option_msg"`
	LoadMsg             string `json:"load_msg"`
	CancelEvent         string `json:"cancel_event"`
	FoundEvent          string `json:"found_event"`
	ConnMsg             string `json:"connect_msg"`
	OpenFileFailMsg     string `json:"open_file_failed_msg"`
	ReadFileFailMsg     string `json:"read_file_failed_msg"`
	CloseFileFailMsg    string `json:"close_file_failed_msg"`
	ParseFileFailMsg    string `json:"parse_file_failed_msg"`
	FoundRouteFailMsg   string `json:"found_route_failed_msg"`
	LivingRouteExistMsg string `json:"living_route_exist_msg"`
	EnemyRouteExistMsg  string `json:"enemy_route_exist_msg"`
}

var defaultLang, _ = NewLang(defaultLangCode)

func NewLang(langCode string) (*Lang, error) {
	var lang Lang
	completePath := i18nPath + langCode + ".json"
	if jsonFile, err := os.Open(completePath); err != nil {
		return nil, errors.New(Msg.OpenFileFailMsg)
	} else if byteValue, err := ioutil.ReadAll(jsonFile); err != nil {
		return nil, errors.New(Msg.ReadFileFailMsg)
	} else if err = jsonFile.Close(); err != nil {
		return nil, errors.New(Msg.CloseFileFailMsg)
	} else if err := json.Unmarshal(byteValue, &lang); err != nil {
		return nil, errors.New(Msg.ParseFileFailMsg)
	}
	return &lang, nil
}
