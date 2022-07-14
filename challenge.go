package main
import "encoding/base64"
import "fmt"

func main() {
    secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatIsIt := reverseStr(string(sd))
	fmt.Printf(whatIsIt)
}

func reverseStr(str string) (out string) {
    for _, s := range str {
        out = string(s) + out
    }

    return
}
