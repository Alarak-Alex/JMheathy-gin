package utils

import "github.com/duke-git/lancet/strutil"

func ReplaceString(s string) string {

	replaces := map[string]string{
		"**": "***",
		"党":  "",
		"国家": "",
		"台湾": "",
	}

	return strutil.ReplaceWithMap(s, replaces)
}
