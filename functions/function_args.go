/** ***********************************************
 * File Name : function_args.go
 * Author    : rootkiter
 * E-mail    : rootkiter@rootkiter.com
 * Created   : 2022-07-21 19:04:21 CST
************************************************* */

package main
import (
    "fmt"
)

type Rect struct {
    Length int
    Width  int
}

func RectPrint(rect *Rect) {
    fmt.Println("Rect.Length = ", rect.Length)
    fmt.Println("Rect.Width  = ", rect.Width)
    fmt.Println("Rect.Area   = ", rect.Length * rect.Width)
}

func StringPrint(str string) {
    fmt.Println("HELLO ", str)
}

func BytesPrint(bts []byte) {
    fmt.Println(bts)
}

func main() {
    teststring := "Rootkiter"
    StringPrint(teststring)
    bts := []byte(teststring)
    BytesPrint(bts)
    BytesPrint(bts[3:])

    rect := &Rect{40, 50}
    RectPrint(rect)

    fmt.Println(rect)
}
