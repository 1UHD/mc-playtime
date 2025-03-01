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

const curseforge_picture string = "assets/curseforge.png"
const curseforge_mac string = "/Library/Application Support/curseforge/minecraft/Instances/"
const curseforge_linux string = "/.curseforge/minecraft/Instances/"
const curseforge_windows string = "\\curseforge\\minecraft\\Instances\\"

const multimc_picture string = "assets/multimc.png"
const multimc_mac string = "/Library/Application Support/MultiMC/instances/"
const multimc_linux string = "/.local/share/MultiMC/instances/"
const multimc_windows string = "\\AppData\\MultiMC\\instances\\"

const prismlauncher_picture string = "assets/prismlauncher.png"
const prismlauncher_mac string = "/Library/Application Support/PrismLauncher/instances/"
const prismlauncher_linux string = "/.local/share/PrismLauncher/instances/"
const prismlauncher_windows string = "\\AppData\\PrismLauncher\\instances\\"

func Get_path(path string) string {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return home_dir + path
}

func Get_os_path() (string, string, string, string, string, string) {
	operating_system := runtime.GOOS

	switch operating_system {
	case "darwin":
		return Get_path(vanilla_mac), Get_path(lunar_mac), Get_path(badlion_mac), Get_path(curseforge_mac), Get_path(multimc_mac), Get_path(prismlauncher_mac)
	case "linux":
		return Get_path(vanilla_linux), Get_path(lunar_linux), Get_path(badlion_linux), Get_path(curseforge_linux), Get_path(multimc_linux), Get_path(prismlauncher_linux)
	case "windows":
		return Get_path(vanilla_windows), Get_path(lunar_windows), Get_path(badlion_windows), Get_path(curseforge_windows), Get_path(multimc_windows), Get_path(prismlauncher_windows)
	}

	return "", "", "", "", "", ""
}

type Version struct {
	Path    string
	Picture string
	Title   string
}

func get_all_instance_log_path(application_path string, pic_path string) ([]Version, error) {
	instances, err := os.ReadDir(application_path)
	if err != nil {
		return nil, err
	}

	var _versions []Version

	for _, instance := range instances {
		var ext string

		if instance.IsDir() {

			if runtime.GOOS == "windows" {
				ext = "\\.minecraft\\logs\\"
			} else {
				ext = "/.minecraft/logs/"
			}

			_, err := os.Stat(application_path + instance.Name() + ext)

			if err == nil {
				newVersion := Version{
					Path:    application_path + instance.Name() + ext,
					Picture: pic_path,
					Title:   instance.Name(),
				}

				_versions = append(_versions, newVersion)
			}
		}
	}

	return _versions[:], nil
}

func Get_versions() []Version {
	vanilla_path, lunar_path, badlion_path, curseforge_path, multimc_path, prismlauncher_path := Get_os_path()

	var versions = []Version{
		{
			Path:    vanilla_path,
			Picture: vanilla_picture,
			Title:   "Vanilla",
		},
		{
			Path:    lunar_path,
			Picture: lunar_picture,
			Title:   "Lunar",
		},
		{
			Path:    badlion_path,
			Picture: badlion_picture,
			Title:   "Badlion",
		},
	}

	forge_paths, err := get_all_instance_log_path(curseforge_path, curseforge_picture)
	if err != nil {
		fmt.Println("no curseforge instances found")
	} else {
		versions = append(versions, forge_paths...)
	}

	multimc_paths, err := get_all_instance_log_path(multimc_path, multimc_picture)
	if err != nil {
		fmt.Println("no multimc instances found")
	} else {
		versions = append(versions, multimc_paths...)
	}

	prismlauncher_paths, err := get_all_instance_log_path(prismlauncher_path, prismlauncher_picture)
	if err != nil {
		fmt.Println("no prismlauncher instances found")
	} else {
		versions = append(versions, prismlauncher_paths...)
	}

	return versions[:]
}
