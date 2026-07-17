package cache

import "fmt"

var cache map[string]string = make(map[string]string)

func SET(key string, val string){

	cache[key] = val
	fmt.Printf("Set Value of %s: %s\n", key, val)
}

func GET(key string){
	val, ok := cache[key]
	if ok {
		fmt.Println(val)
	}else{
		fmt.Println("not initialized")
	}

}

func DEL(key string){

	delete(cache, key)
}