package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func readKeyboard(input_device_path string, keyboard_map map[string]string) {
	f, err := os.Open(input_device_path)

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer f.Close()
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			encoded_string := hex.EncodeToString(buf[:n])
			fmt.Println(buf[:n])
			fmt.Println(encoded_string)
		}
	}
}

func getKeyMap() map[string]string {
	cmd := exec.Command("dumpkeys", "-l")
	stdout, err := cmd.Output()
	key_map := make(map[string]string)
	if err != nil {
		fmt.Println(err.Error())
		return key_map
	}

	re_hexcode, err := regexp.Compile("0[xX][\\d\\w]+")
	re_hexcode_and_spaces, err := regexp.Compile("0[xX][\\d\\w]+\\s+")

	lines := strings.Split(string(stdout), "\n")

	for _, line := range lines {
		hex_code := re_hexcode.FindString(line)
		value := re_hexcode_and_spaces.ReplaceAllString(line, "")
		key_map[hex_code] = value
	}

	return key_map
}

func main() {
	key_map := getKeyMap()
	readKeyboard("/dev/input/by-path/platform-i8042-serio-0-event-kbd", key_map)
}
