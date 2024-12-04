package utils

import "time"

func QueueCopD(copsD chan<- time.Duration, spent time.Duration) {
	select {
	case copsD <- spent:
	default:
		panic("you should make the copsD channel bigger")
	}
}
