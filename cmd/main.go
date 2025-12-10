package main

import (
    "os"

    "iptv_sc/pkg"

)

func main() {
    sc_parse.LoadConfig("resource/sichuanunicom.json")
    sc_parse.ShowConfig()
    os.Exit(0)
}
