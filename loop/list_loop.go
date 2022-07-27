/** ***********************************************
 * File Name : list_loop.go
 * Author    : rootkiter
 * E-mail    : rootkiter@rootkiter.com
 * Created   : 2022-07-21 17:55:00 CST
************************************************* */

package main
import (
    "fmt"
)

func main() {
    abc := []int{20, 30, 40, 50}
    abc = append(abc, 60)
    tmp := abc[:]
    for k, v := range(tmp) {
        fmt.Println(k, v)
    }
}
