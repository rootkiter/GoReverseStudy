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
    // list
    abc := []int{20,30, 40, 50}
    abc = append(abc, 60)

    fmt.Println(abc)
}
