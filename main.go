package main

import (
  "github.com/sirupsen/logrus"
  "sing-geodata/libs"
)

func main() {
  var err error
  err = libs.Build_GeoIP()
  if err != nil {
    logrus.Fatal(err)
  }
  err = libs.Build_GeoSite()
  if err != nil {
    logrus.Fatal(err)
  }
}