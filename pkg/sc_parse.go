package sc_parse

import (
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

type ChannelItem struct {
    Caption string `json:"Caption"`
    LiveUrl string `json:"LiveUrl"`
}

type ChannelList struct {
    ChannelList  []ChannelItem `json:"ChannelList"`
}

type IPTVResp struct {
    Comment            string `json:"comment"`
    RetValues           []ChannelList `json:"retValues"`
}

var IPTVConfig *IPTVResp = nil

func ShowConfig() error {
//    if IPTVConfig != nil {
//        fmt.Printf("Comment : '%s'\n", IPTVConfig.Comment)
//    }

//	for idxChannel, itemChannel := range IPTVConfig.RetValues {
//		fmt.Printf("Index: %d, Value: %v\n", idxChannel, itemChannel)
//	}

    // only need first group of channel list, which contains all channel without by group
    fmt.Printf("#EXTM3U\n")
    itemChannel := IPTVConfig.RetValues[0].ChannelList
    //fmt.Printf("debug %v\n", itemChannel);
    chIdx := 1
    for _, item := range itemChannel {
        if !strings.Contains(item.Caption, "高清") {
            continue
        }

        groupName := ""
//        epgLogo := ""
        if strings.Contains(item.Caption, "卫视高清") {
            groupName = "省级卫视"
        } else if strings.Contains(item.Caption, "CCTV") {
            groupName = "中央电视台"
            itemVal := ""
            if idx := strings.Index(item.Caption, "高清"); idx != -1 {
                itemVal = item.Caption[:idx]
            }

            if idx := strings.Index(itemVal, "CCTV-"); idx != -1 {
                pos := idx + len("CCTV-")
                //itemPrefix := itemVal[:pos]
                itemIdx := itemVal[pos:]
                if iidx := strings.Index(itemIdx, "+"); iidx != -1 {
                    itemIdx = itemIdx[:iidx]
                }
                //epgLogo = fmt.Sprintf("resource/CCTV%s", itemIdx)
                //fmt.Printf("logo %s\n", epgLogo)
                //fmt.Printf("item %s, itemPrefix %s, itemIdx %s\n", itemVal, itemPrefix, itemIdx);
            }
        } else if strings.Contains(item.Caption, "SCTV") {
            groupName = "四川电视台"
        } else if strings.Contains(item.Caption, "CETV") {
            groupName = "中国教育电视台"
        } else if strings.Contains(item.Caption, "CHC") {
            groupName = "CHC 家庭影院"
        } else {
            groupName = "其它"
        }

        itemName := item.Caption
        if idx := strings.Index(itemName, "高清"); idx != -1 {
            itemName = itemName[:idx]
        }

        fmt.Printf("#EXTINF:-1 tvg-id=\"%d\" tvg-name=\"%s\" group-title=\"%s\",%s \n",
                    chIdx, itemName, groupName, itemName)
        fmt.Printf("%s\n", item.LiveUrl)

        chIdx++
    }

    return  nil
}

func LoadConfig(confFile string) error {
    bytes, err := os.ReadFile(confFile)
    if err != nil {
        fmt.Printf("Failed to read %s with error '%s'\n", confFile, err.Error())
        return err
    }

   // fmt.Printf("get conf\n%s\n", string(bytes))

    err = json.Unmarshal(bytes, &IPTVConfig)
    if err != nil {
        fmt.Printf("Failed to parse %s with error '%s'\n", err.Error())
        return err
    }

    return nil
}
