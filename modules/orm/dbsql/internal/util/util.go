package util

import "log"

func Catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
