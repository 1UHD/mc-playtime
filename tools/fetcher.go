package tools

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func get_first_and_last_line(file_data string) (string, string) {

	log_file := strings.Split(string(file_data), "\n")

	if len(log_file) < 2 {
		return "[00:00:00]", "[00:00:00]"
	}

	var first_line string = log_file[0]

	//backtracking to find the last line that has a timestamp (this can happen when the game crashes)
	i := 2
	for !(strings.Contains(strings.Split(log_file[len(log_file)-i], " ")[0], "[")) {
		i++
	}

	var last_line string = log_file[len(log_file)-i]

	return first_line, last_line
}

func get_seconds_timestamp(line string) int32 {
	timestamp := strings.ReplaceAll(strings.ReplaceAll(strings.Split(line, " ")[0], "]", ""), "[", "")

	time := strings.Split(timestamp, ":")
	hours, err := strconv.Atoi(time[0])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	minutes, err := strconv.Atoi(time[1])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	seconds, err := strconv.Atoi(time[2])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	var elapsed int32 = int32(hours*3600 + minutes*60 + seconds)

	return elapsed
}

func Get_time_spent_in_session(file_data string) int32 {
	first_line, last_line := get_first_and_last_line(file_data)

	first_timestamp := get_seconds_timestamp(first_line)
	last_timestamp := get_seconds_timestamp(last_line)

	time_spent := last_timestamp - first_timestamp

	return time_spent
}

func read_gz_file(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer reader.Close()

	content, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(content)
}

func Get_time_for_directory(directory string) int32 {
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	var time_spent int32 = 0

	fmt.Println(len(files))
	for _, file := range files {
		if strings.Contains(file.Name(), "log.gz") {
			file_path := directory + file.Name()
			file_data := read_gz_file(file_path)
			time_spent_in_session := Get_time_spent_in_session(file_data)
			fmt.Println("Time spent in", file_path, ":", strconv.Itoa(int(time_spent_in_session)))
			time_spent += time_spent_in_session
		}
	}

	return time_spent
}
