package libs

import (
  "io"
  "net/http"
  "os"
  "path/filepath"
  "strings"

  "github.com/sagernet/sing-box/common/geosite"
  "github.com/sagernet/sing/common"
  "github.com/v2fly/v2ray-core/v5/app/router/routercommon"
  "google.golang.org/protobuf/proto"
)

func parseGeoSite(vGeositeData []byte) (map[string][]geosite.Item, error) {
  vGeositeList := routercommon.GeoSiteList{}
  err := proto.Unmarshal(vGeositeData, &vGeositeList)
  if err != nil {
    return nil, err
  }
  domainMap := make(map[string][]geosite.Item)
  for _, vGeositeEntry := range vGeositeList.Entry {
    code := strings.ToLower(vGeositeEntry.CountryCode)
    domains := make([]geosite.Item, 0, len(vGeositeEntry.Domain)*2)
    attributes := make(map[string][]*routercommon.Domain)
    for _, domain := range vGeositeEntry.Domain {
      if len(domain.Attribute) > 0 {
        for _, attribute := range domain.Attribute {
          attributes[attribute.Key] = append(attributes[attribute.Key], domain)
        }
      }
      switch domain.Type {
      case routercommon.Domain_Plain:
        domains = append(domains, geosite.Item{
          Type:  geosite.RuleTypeDomainKeyword,
          Value: domain.Value,
        })
      case routercommon.Domain_Regex:
        domains = append(domains, geosite.Item{
          Type:  geosite.RuleTypeDomainRegex,
          Value: domain.Value,
        })
      case routercommon.Domain_RootDomain:
        if strings.Contains(domain.Value, ".") {
          domains = append(domains, geosite.Item{
            Type:  geosite.RuleTypeDomain,
            Value: domain.Value,
          })
        }
        domains = append(domains, geosite.Item{
          Type:  geosite.RuleTypeDomainSuffix,
          Value: "." + domain.Value,
        })
      case routercommon.Domain_Full:
        domains = append(domains, geosite.Item{
          Type:  geosite.RuleTypeDomain,
          Value: domain.Value,
        })
      }
    }
    domainMap[code] = common.Uniq(domains)
    for attribute, attributeEntries := range attributes {
      attributeDomains := make([]geosite.Item, 0, len(attributeEntries)*2)
      for _, domain := range attributeEntries {
        switch domain.Type {
        case routercommon.Domain_Plain:
          attributeDomains = append(attributeDomains, geosite.Item{
            Type:  geosite.RuleTypeDomainKeyword,
            Value: domain.Value,
          })
        case routercommon.Domain_Regex:
          attributeDomains = append(attributeDomains, geosite.Item{
            Type:  geosite.RuleTypeDomainRegex,
            Value: domain.Value,
          })
        case routercommon.Domain_RootDomain:
          if strings.Contains(domain.Value, ".") {
            attributeDomains = append(attributeDomains, geosite.Item{
              Type:  geosite.RuleTypeDomain,
              Value: domain.Value,
            })
          }
          attributeDomains = append(attributeDomains, geosite.Item{
            Type:  geosite.RuleTypeDomainSuffix,
            Value: "." + domain.Value,
          })
        case routercommon.Domain_Full:
          attributeDomains = append(attributeDomains, geosite.Item{
            Type:  geosite.RuleTypeDomain,
            Value: domain.Value,
          })
        }
      }
      domainMap[code+"@"+attribute] = common.Uniq(attributeDomains)
    }
  }
  return domainMap, nil
}

func Build_GeoSite() error {
  outputFile, err := os.Create("geosite.db")
  if err != nil {
    return err
  }
  defer outputFile.Close()
  response, err := http.Get("https://github.com/Loyalsoldier/v2ray-rules-dat/raw/release/geosite.dat")
  if err != nil {
    return err
  }
  defer response.Body.Close()
  vData, err := io.ReadAll(response.Body)
  if err != nil {
    return err
  }
  domainMap, err := parseGeoSite(vData)
  if err != nil {
    return err
  }
  err = geosite.Write(outputFile, domainMap)
  if err != nil {
    return err
  } else {
    outputPath, _ := filepath.Abs("geosite.db")
    os.Stderr.WriteString("write " + outputPath + "\n")
  }
  return nil
}
