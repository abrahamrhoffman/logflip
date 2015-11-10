// Get owning Service
// Move .log -> .log.1
// Reload owning service

package main

import (
    "fmt"
    "strings"
    "gopkg.in/pipe.v2"
)

func main() {

    log0     := string("/var/log/apache2/access.log")
    log1     := string("/var/log/apache2/access.log.1")

    // Owning PID Command

    p0 := pipe.Line(
        pipe.Exec("lsof", log0),
        pipe.Exec("grep", "-v", "COMMAND"),
        pipe.Exec("awk", "{print $1}"),
        pipe.Exec("sort", "-u"),
    )
    output0, err := pipe.CombinedOutput(p0)
    if err != nil {
        fmt.Printf("%v\n", err)
    }

    process := string(output0)
    t := strings.TrimSpace(process)

    // Move Command

    p1 := pipe.Line(
        pipe.Exec("mv", log0, log1),
    )

    output1, err := pipe.CombinedOutput(p1)
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    fmt.Printf("%s", output1)

    // Service Reload Command

    p2 := pipe.Line(
        pipe.Exec("service", t, "reload"),
    )
    output2, err := pipe.CombinedOutput(p2)
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    fmt.Printf("%s", output2)
}
