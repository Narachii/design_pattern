package main

import (
	"fmt"
	"sync"
)

//The sync.Once will only perform the operation only once. See below code
var once sync.Once

type singleOnce struct {}

var singleOnceInstance *singleOnce

func getOnceInstance() *singleOnce {
	if singleOnceInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating Single Once Instance New")
				singleOnceInstance = &singleOnce{}
			},
		)
	} else {
		fmt.Println("Single Once Instance already created")
	}
	return singleOnceInstance
}
