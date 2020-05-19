package utils

import "os"

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// dir, err := os.Getwd()
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println(dir)
