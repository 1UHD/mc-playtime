package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

const vanilla_mac string = "/Library/Application Support/minecraft/logs/"
const vanilla_linux string = "/"
const vanilla_windows string = "\\"

const lunar_mac string = "/Library/Application Support/minecraft/logs/"
const lunar_linux string = "/"
const lunar_windows string = "\\"

const badlion_mac string = "/Library/Application Support/minecraft/logs/"
const badlion_linux string = "/"
const badlion_windows string = "\\"

func Get_path(path string) string {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return home_dir + path
}

func Get_os_path() (string, string, string) {
	operating_system := runtime.GOOS
	fmt.Println(operating_system)

	switch operating_system {
	case "darwin":
		return Get_path(vanilla_mac), Get_path(lunar_mac), Get_path(badlion_mac)
	case "linux":
		return Get_path(vanilla_linux), Get_path(lunar_linux), Get_path(badlion_linux)
	case "windows":
		return Get_path(vanilla_windows), Get_path(lunar_windows), Get_path(badlion_windows)
	}

	return "", "", ""
}

type Version struct {
	Path    string
	Picture string
	Title   string
}

func Get_versions(filename string) ([]Version, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var versions []Version
	if err := json.Unmarshal(data, &versions); err != nil {
		return nil, err
	}
	return versions, nil
}
