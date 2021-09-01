package err2

import "log"

func LogErr1(parent, child string, err error) {
	log.Printf("%v failed at '%v', because: %v",
		parent, child, err.Error())
}

func BilboLogginbs(parent, child string, err error) {
	log.Printf("%v failed at '%v', because: %v",
		parent, child, err.Error())
}
