package helpers

import "os"

func CheckDocker() bool {
	return os.Getenv("FE_URL") != ""
}
