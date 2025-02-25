package tools

import (
	"fmt"
	"os"
	"runtime"
)

const vanilla_picture string = "assets/vanilla.png"
const vanilla_mac string = "/Library/Application Support/minecraft/logs/"
const vanilla_linux string = "/.minecraft/logs/"
const vanilla_windows string = "\\AppData\\Roaming\\.minecraft\\logs\\"

const lunar_picture string = "assets/lunar.png"
const lunar_mac string = "/.lunarclient/offline/multiver/logs/"
const lunar_linux string = "/.lunarclient/offline/multiver/logs/"
const lunar_windows string = "\\.lunarclient\\offline\\multiver\\logs\\"

const badlion_picture string = "assets/badlion.png"
const badlion_mac string = "/Library/Application Support/minecraft/logs/blclient/minecraft/"
const badlion_linux string = "/.minecraft/logs/blclient/minecraft/"
const badlion_windows string = "\\AppData\\Roaming\\.minecraft\\logs\\blclient\\minecraft\\"

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

func Get_versions() []Version {
	vanilla_path, lunar_path, badlion_path := Get_os_path()

	var versions = [...]Version{
		Version{
			Path:    vanilla_path,
			Picture: vanilla_picture,
			Title:   "Vanilla",
		},
		Version{
			Path:    lunar_path,
			Picture: lunar_picture,
			Title:   "Lunar",
		},
		Version{
			Path:    badlion_path,
			Picture: badlion_picture,
			Title:   "Badlion",
		},
	}

	return versions[:]
}
