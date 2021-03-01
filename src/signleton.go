/*
 * 单例模式
 */
package main

// Signleton struct
type Signleton struct {
	count int
}

var instance *Signleton = nil

// NewSignleton func
func NewSignleton() *Signleton {
	if instance != nil {
		instance = new(Signleton)
	}
	return instance
}

func main() {
	var signleton = NewSignleton()
	
}
