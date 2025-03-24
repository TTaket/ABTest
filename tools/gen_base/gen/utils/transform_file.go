package utils

import (
	"fmt"
	"strings"
)

type Token struct {
	Before string
	After  string
}

var (
	config Config
)

var (
	gen_dir    = "./gen/"
	config_dir = gen_dir + "config/"
)
var (
	yaml_file_path = config_dir + "config.yml"
)

func init() {
	err := ReadYamlFile(yaml_file_path, &config)
	if err != nil {
		fmt.Printf("ReadYamlFile failed: %v\n", err)
		return
	}
}

// TransformFile reads a yaml file and transforms it into a string
// 会默认加入所有的%num
func TransformFile(context string, tokens []Token) (string, error) {
	// read yaml file
	var ConfigList Config
	err := ReadYamlFile(yaml_file_path, &ConfigList)
	if err != nil {
		return "", err
	}

	// to token
	var token []Token
	token = append(token, Token{Before: "%1", After: ConfigList.BaseInfo.ModuleName})
	token = append(token, Token{Before: "%2", After: ConfigList.BaseInfo.ModuleNameBig})
	token = append(token, Token{Before: "%3", After: ConfigList.BaseInfo.ModuleNameSmall})
	token = append(token, Token{Before: "%4", After: ConfigList.BaseInfo.RPCServiceName})
	token = append(token, Token{Before: "%5", After: ConfigList.BaseInfo.RPCServiceClient})
	token = append(token, Token{Before: "%6", After: ConfigList.BaseInfo.Port})

	for _, v := range ConfigList.FuncInfo {
		token = append(token, Token{Before: v.Key, After: v.Value})
	}

	for _, v := range tokens {
		token = append(token, Token{Before: v.Before, After: v.After})
	}

	//按照token替换
	for _, v := range token {
		context = strings.ReplaceAll(context, v.Before, v.After)
	}
	//两次保证不漏
	for _, v := range token {
		context = strings.ReplaceAll(context, v.Before, v.After)
	}
	return context, nil
}
