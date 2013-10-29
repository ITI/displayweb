package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "os/exec"
)

const basedir string = "/Users/jmjone/src/go"
//const basedir string = "/home/tcipg/src/go"
const webdir string = "/src/iti/displayweb/html"
const wall_cmd = basedir + "/bin/wallctl"


var (
    layoutnames []string
    cmd_opts = []string{"-f", basedir+"/wall.conf", "-i", "192.168.80.6"}
)

func index(writer http.ResponseWriter, request *http.Request) {
    http.ServeFile(writer, request, basedir + webdir + "/index.html")
}

func getNames(writer http.ResponseWriter, request *http.Request) {
    j, _ := json.Marshal(layoutnames)
    fmt.Fprintf(writer, "%s", j)
}

func setLayout(writer http.ResponseWriter, request *http.Request) {
    var j string
    var opt []string

    d := json.NewDecoder(request.Body)
    d.Decode(&j)
    if j == "power_on" {
        opt = []string{"-on"}
    } else if j == "power_off" {
        opt = []string{"-off"}
    } else {
        opt = []string{"-c", j}
    }

    cmd := exec.Command(wall_cmd)
    for _,a := range cmd_opts {
        cmd.Args = append(cmd.Args, a)
    }
    for _,a := range opt {
        cmd.Args = append(cmd.Args, a)
    }
    fmt.Printf("%s\n", cmd.Args)
    //cmd.Run(wall_cmd + cmd_opts + opt)
    op, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Error:::: %s\n", op)
    }
}

func staticRsrc(writer http.ResponseWriter, request *http.Request) {
    fmt.Printf("%s\n", request.URL.Path[10:])
    http.ServeFile(writer, request, basedir + webdir + "/resource/" + request.URL.Path[10:])

}

func main() {


    http.HandleFunc("/", index)
    http.HandleFunc("/conf/names", getNames)
    http.HandleFunc("/set", setLayout)
    http.HandleFunc("/resource/", staticRsrc)
    conf, err := parseConfig(basedir + "/wall.conf")
    if err != nil {
        fmt.Print("Bad things")
    }

    for _, name := range conf.Layouts {
        layoutnames = append(layoutnames, name.Name)
    }

    fmt.Print(layoutnames)
    http.ListenAndServe(":8080", nil)

}
