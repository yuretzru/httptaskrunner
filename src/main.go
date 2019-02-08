package main

import (
    "flag"
    "log"
    "encoding/base64"
    "net/http"
    "os/exec"
    "bytes"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

var addr = flag.String("addr", "127.0.0.1:56565", "http service address")
var configfile = flag.String("conf", "/etc/httptaskrunner.yml", "config file")
var favicon, _ = base64.StdEncoding.DecodeString("AAABAAEAEBAQAAEABAAoAQAAFgAAACgAAAAQAAAAIAAAAAEABAAAAAAAgAAAAAAAAAAAAAAAEAAAAAAAAAAlzAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEREREREREREREREREREREREBEREAAAAREQAREREREREREAERERERERERABEREREREREQAREREREREREAEREREREREAEREREREREAEREREREREAEREREREREAEREREREREQEREREREREREREREREREREREREREREREREREREREREAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
var VERSION = "UNKNOWN"

type Conf struct {
    Version int `yaml:"version"`
    Settings struct {
        Listen string `yaml:"listen"`
    }`yaml:"settings"`
    Commands map[string]string `yaml:"commands"`
}   

var cfg Conf


func serveFavicon(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "image/x-icon")
    w.Write(favicon)
}


func serveHome(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    cmd := r.FormValue("cmd")
    if cmd == "" {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("parameter cmd not found"))
        return
    }
    command := cfg.Commands[cmd]
    if command == "" {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("command: " + cmd + " not found"))
        return
    }

    res, err := Exec(command)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()+"\n"))
        w.Write([]byte(res))
    } else {
        w.Write([]byte(res))
    }
}

func serveVer(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(VERSION))
}


func Exec(command string) (out string, err error) {
    cmd := exec.Command("/bin/sh", "-c", command)
    var outb, errb bytes.Buffer
    cmd.Stdout = &outb
    cmd.Stderr = &errb
    err = cmd.Run()
    if err != nil {
        out = errb.String()
        return
    }
    out = outb.String()
    return
}


func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    flag.Parse()

    file, err := ioutil.ReadFile(*configfile)
    if err != nil {
        panic(err)
    }

    err = yaml.Unmarshal([]byte(file), &cfg)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    port := cfg.Settings.Listen
    if *addr != "" {
        port = *addr
    }

    http.HandleFunc("/", serveHome)
    http.HandleFunc("/favicon.ico", serveFavicon)
    http.HandleFunc("/v", serveVer)

    err = http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
