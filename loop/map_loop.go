/** ***********************************************
 * File Name : map_loop.go
 * Author    : rootkiter
 * E-mail    : rootkiter@rootkiter.com
 * Created   : 2022-07-21 17:51:49 CST
************************************************* */

package main
import (
        "fmt"
)

func main() {
    def := map[string] string {
        "key1": "value1",
        "key2": "value2",
    }
    for k, v := range(def) {
        fmt.Println(k, v)
    }
}
