package completeword

// import (
//    "fmt"
//    "os/exec"
//    "os"
//    "bytes"
//    "io"
// )

// func Choose() {
    // app := "/bin/ls"
    // cmd, err := exec.Run(app, []string{app, "-l"}, nil, "", exec.DevNull, exec.Pipe, exec.Pipe)
    //
    // if (err != nil) {
    //    fmt.Fprintln(os.Stderr, err.String())
    //    return
    // }
    //
    // var b bytes.Buffer
    // io.Copy(&b, cmd.Stdout)
    // fmt.Println(b.String())
    //
    // cmd.Close()
// }