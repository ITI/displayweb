package main

import (
    "io/ioutil"
    "os"
    "encoding/xml"
)

type Output struct {
    Name string `xml:"name,attr"`
    Port int `xml:"port,attr"`
    WallID int `xml:"wallid,attr"`
    MacAddr string `xml:"mac,attr"`
}
func (v Output) GetName() (string) {
    return v.Name
}
func (v Output) GetPort() (int) {
    return v.Port
}

type Input struct {
    Name string `xml:"name,attr"`
    Port int `xml:"port,attr"`
}
func (v Input) GetName() (string) {
    return v.Name
}
func (v Input) GetPort() (int) {
    return v.Port
}

type Config struct {
    Inputs []Input `xml:"inputs>input"`
    Outputs []Output `xml:"outputs>output"`
}

type Panel struct {
    Name string `xml:"name,attr"`
    Source string `xml:"source"`
}

type WInput struct {
    Name string `xml:"name,attr"`
    Outputs []string `xml:"output"`
}

type WPanel struct {
    Position int `xml:"pos,attr"`
    Name string `xml:"name,attr"`
}

type WSize struct {
    X int `xml:"x,attr"`
    Y int `xml:"y,attr"`
}

type Wall struct {
    Size WSize `xml:"size"`
    Panels []WPanel `xml:"panel"`
}

type Layout struct {
    Name string `xml:"name,attr"`
    PanelConfigs []Panel `xml:"panelconfig>panel"`
    VideoSwitch []WInput `xml:"videoswitch>input"`
    VideoWalls []Wall `xml:"videowall>wall"`
}

type Configuration struct {
    Config Config `xml:"config"`
    Layouts []Layout `xml:"layouts>layout"`
}

func parseConfig(cfgfile string) (*Configuration, error) {
    cfile, err := os.Open(cfgfile)
    if err != nil {
        return nil, err
    }
    defer cfile.Close()
    cfgbuf, _ := ioutil.ReadAll(cfile)

    var config Configuration
    err = xml.Unmarshal(cfgbuf, &config)
    if err != nil {
        return nil, err
    }
    return &config, nil
}
