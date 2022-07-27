/** ***********************************************
 * File Name : routine.go
 * Author    : rootkiter
 * E-mail    : rootkiter@rootkiter.com
 * Created   : 2022-07-21 19:02:25 CST
************************************************* */

package main
import (
    "fmt"
    "time"
)

func routine1 () {
    for {
        fmt.Println("routine 1")
        time.Sleep(1*1000*1000*1000)
    }
}

func main() {
    go routine1 ()

    go func (arg string) {
        for {
            fmt.Println("routine 2 ", arg)
            time.Sleep(1*500*1000*1000)
        }
    } ("HELLO-Routine")
}
