// Ping multiple hosts at a time, show which ones are reachable, which ones
// did not respond within a second, and which ones could not be identified
// by name (DNS).
package main
import (
    "fmt"
    "os"
    "os/exec"
    "strconv"

    "github.com/mgutz/ansi"
)
const (
    HOST_UP         = 0     // Host has responded within a second
    HOST_DOWN       = 1     // Host has not responded within a second
    HOST_UNKNOWN    = 2     // Host name cannot be resolved by DNS
    PING            = "ping"
)

type Host struct {
    name    string
    state   int
}

// Ping a host and return results into channel
func Ping(name string, ch chan Host) {
    cmd := exec.Command(PING, "-c", "1", "-w", "1", name)
    err := cmd.Start()
    if err!=nil {
        panic(err)
    }
    err = cmd.Wait()
    if err==nil {
        // No error means host responded
        ch <- Host{name, HOST_UP}
        return
    }
    // The observed returned value from the above ping line is
    // 1 for host that did not respond
    // 2 for unresolved host name.
    msg:=err.Error()
    state, _ := strconv.Atoi(string(msg[len(msg)-1]))
    ch <- Host{name, state}
    return
}

func PingHosts(hostnames []string) map[string]int {
    hosts:=make(map[string]int)

    c:=make(chan Host)
    for i:=0 ; i<len(hostnames) ; i++ {
        go Ping(hostnames[i], c)
    }
    for i:=0 ; i<len(hostnames) ; i++ {
        h := <-c
        hosts[h.name] = h.state
    }
    close(c)

    return hosts
}

func DisplayHost(host string, state int) string {
    msg:="--"
    switch state {
        case HOST_UP:
        msg = fmt.Sprintf("[+] %s", host)
        msg = ansi.Color(msg, "green")
        case HOST_DOWN:
        msg = fmt.Sprintf("[-] %s", host)
        msg = ansi.Color(msg, "red")
        case HOST_UNKNOWN:
        msg = fmt.Sprintf("[?] %s", host)
        msg = ansi.Color(msg, "yellow")
    }
    return msg
}


func main() {
    if len(os.Args)<2 {
        fmt.Println("use: "+os.Args[0]+" host ... host")
        return
    }

    ups := PingHosts(os.Args[1:])
    for k,v := range ups {
        fmt.Printf(DisplayHost(k, v)+"\n")
    }
}
