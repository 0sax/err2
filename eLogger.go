package err2

import "log"

func LogErr1(parent, child string, err error) {
	errMsg := "no error provided"
	if err != nil {
		errMsg = err.Error()
	}
	log.Printf("%v failed at '%v', because: %v",
		parent, child, errMsg)
}
