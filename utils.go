/*
Create: 2022/7/20
Project: octopus-meta
Github: https://github.com/landers1037
Copyright Renj
*/

// Package octopus_meta
package octopus_meta

import (
	"os"
	"path"
)

func isDirExist(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}
	return true
}

func isEnvExist(env string) (string, bool) {
	if os.Getenv(env) == "" {
		return "", false
	}
	return os.Getenv(env), true
}

// load all app from dir
func load(dir string) (map[string]App, error) {
	cfs, err := getAPPCfs(dir)
	if err != nil {
		return nil, err
	}
	res, status := loadAllCfs(cfs)
	if !status {
		return res, ErrLoadPart
	}
	return res, nil
}

func chooseDir() string {
	autoLoadDir := path.Join(os.Getenv("APP_ROOT"), MetaAutoLoadDir)
	autoLoadDirNoHide := path.Join(os.Getenv("APP_ROOT"), MetaAutoLoadDirNoHide)
	if OctopusMetaDir == "" &&
		!isDirExist(autoLoadDir) &&
		!isDirExist(autoLoadDirNoHide) {
		return ""
	}
	var p string
	if OctopusMetaDir != "" {
		p = OctopusMetaDir
	} else if autoLoadDir != "" {
		p = autoLoadDir
	} else if autoLoadDirNoHide != "" {
		p = autoLoadDirNoHide
	}
	return p
}
