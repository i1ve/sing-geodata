package libs

import (
  "io"
  "net"
  "net/http"
  "os"
  "sort"
  "path/filepath"
  "strings"

  "github.com/maxmind/mmdbwriter"
  "github.com/maxmind/mmdbwriter/inserter"
  "github.com/maxmind/mmdbwriter/mmdbtype"
  "github.com/oschwald/geoip2-golang"
  "github.com/oschwald/maxminddb-golang"
)

func parseGeoIP(binary []byte) (metadata maxminddb.Metadata, countryMap map[string][]*net.IPNet, err error) {
  database, err := maxminddb.FromBytes(binary)
  if err != nil {
    return
  }
  metadata = database.Metadata
  networks := database.Networks(maxminddb.SkipAliasedNetworks)
  countryMap = make(map[string][]*net.IPNet)
  var country geoip2.Enterprise
  var ipNet *net.IPNet
  for networks.Next() {
    ipNet, err = networks.Network(&country)
    if err != nil {
      return
    }
    var code string
    if country.Country.IsoCode != "" {
      code = strings.ToLower(country.Country.IsoCode)
    } else if country.RegisteredCountry.IsoCode != "" {
      code = strings.ToLower(country.RegisteredCountry.IsoCode)
    } else if country.RepresentedCountry.IsoCode != "" {
      code = strings.ToLower(country.RepresentedCountry.IsoCode)
    } else if country.Continent.Code != "" {
      code = strings.ToLower(country.Continent.Code)
    } else {
      continue
    }
    countryMap[code] = append(countryMap[code], ipNet)
  }
  err = networks.Err()
  return
}

func newWriter(metadata maxminddb.Metadata, codes []string) (*mmdbwriter.Tree, error) {
  return mmdbwriter.New(mmdbwriter.Options{
    DatabaseType:            "sing-geoip",
    Languages:               codes,
    IPVersion:               int(metadata.IPVersion),
    RecordSize:              int(metadata.RecordSize),
    Inserter:                inserter.ReplaceWith,
    DisableIPv4Aliasing:     true,
    IncludeReservedNetworks: true,
  })
}

func write(writer *mmdbwriter.Tree, dataMap map[string][]*net.IPNet, output string, codes []string) error {
  if len(codes) == 0 {
    codes = make([]string, 0, len(dataMap))
    for code := range dataMap {
      codes = append(codes, code)
    }
  }
  sort.Strings(codes)
  codeMap := make(map[string]bool)
  for _, code := range codes {
    codeMap[code] = true
  }
  for code, data := range dataMap {
    if !codeMap[code] {
      continue
    }
    for _, item := range data {
      err := writer.Insert(item, mmdbtype.String(code))
      if err != nil {
        return err
      }
    }
  }
  outputFile, err := os.Create(output)
  if err != nil {
    return err
  }
  defer outputFile.Close()
  _, err = writer.WriteTo(outputFile)
  return err
}

func Build_GeoIP() error {
  response, err := http.Get("https://github.com/d2184/geoip/raw/release/Country.mmdb")
  if err != nil {
    return err
  }
  defer response.Body.Close()
  binary, err := io.ReadAll(response.Body)
  if err != nil {
    return err
  }
  metadata, countryMap, err := parseGeoIP(binary)
  if err != nil {
    return err
  }
  allCodes := make([]string, 0, len(countryMap))
  for code := range countryMap {
    allCodes = append(allCodes, code)
  }
  writer, err := newWriter(metadata, allCodes)
  if err != nil {
    return err
  }
  err = write(writer, countryMap, "geoip.db", nil)
  if err != nil {
    return err
  } else {
    outputPath, _ := filepath.Abs("geoip.db")
    os.Stderr.WriteString("write " + outputPath + "\n")
  }
  writer, err = newWriter(metadata, []string{"cn", "private"})
  if err != nil {
    return err
  }
  err = write(writer, countryMap, "geoip-lite.db", []string{"cn", "private"})
  if err != nil {
    return err
  } else {
    outputPath, _ := filepath.Abs("geoip-lite.db")
    os.Stderr.WriteString("write " + outputPath + "\n")
  }
  return nil
}
