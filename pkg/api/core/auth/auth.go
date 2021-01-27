package auth

import "github.com/yoneyan/oit_kouzoka/pkg/api/tool/config"

func Authentication(token string) bool {
	return token == config.Conf.Token
}
