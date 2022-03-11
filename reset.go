package main

import (
	"errors"
	"fmt"
	registry "golang.org/x/sys/windows/registry"
	"strings"
)

const PremiumPath = "Software\\PremiumSoft\\NavicatPremium"
const ClsidPath = "Software\\Classes\\CLSID"

func arrayStringContains(array []string, str string) (int, error) {
	for key, item := range array {
		if strings.Contains(item, str) {
			return key, nil
		}
	}
	return 0, errors.New("Not Found " + str)
}

func clearPremiumKey() {
	key, err := registry.OpenKey(registry.CURRENT_USER, PremiumPath, registry.ALL_ACCESS)
	if err != nil {
		return
	}
	defer key.Close()
	names, err := key.ReadSubKeyNames(0)
	if err != nil {
		return
	}
	index, err := arrayStringContains(names, "Registration")
	if err != nil {
		fmt.Println("Not Found Registration")
		return
	}
	registrationKey := names[index]
	//Delete registry key
	err = registry.DeleteKey(registry.CURRENT_USER, registrationKey)
	if err != nil {
		return
	}
}

func clearClsidKey() {
	key, err := registry.OpenKey(registry.CURRENT_USER, ClsidPath, registry.ALL_ACCESS)
	if err != nil {
		return
	}
	defer key.Close()
	names, err := key.ReadSubKeyNames(0)
	if err != nil {
		return
	}

	for _, name := range names {
		subKey, err := registry.OpenKey(registry.CURRENT_USER, ClsidPath+"\\"+name, registry.ALL_ACCESS)
		if err != nil {
			return
		}
		subKey.Close()
		subNames, err := subKey.ReadSubKeyNames(0)
		if err != nil {
			return
		}
		index, err := arrayStringContains(subNames, "Info")
		if err != nil {
			continue
		}
		//Delete registry key
		subPath := ClsidPath + "\\" + name + "\\" + subNames[index]
		err = registry.DeleteKey(registry.CURRENT_USER, subPath)
		if err != nil {
			continue
		}
	}

}

func main() {
	clearPremiumKey()
	clearClsidKey()
}
