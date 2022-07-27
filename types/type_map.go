/** ***********************************************
 * File Name : list_map.go
 * Author    : rootkiter
 * E-mail    : rootkiter@rootkiter.com
 * Created   : 2022-07-20 18:04:56 CST
************************************************* */

package main
import (
    "fmt"
)

func main() {
    // map
    abc := map[string] string{
        "key1": "value1",
        "key2": "value2",
    }
    abc["key3"] = "value3"

    fmt.Println(abc)
}
