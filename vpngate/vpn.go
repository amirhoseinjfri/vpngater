package vpngate

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const (
	UNWANTED_CLASS = "vg_table_header"
	MESSAGE_LEN    = 3500
)

type Server struct {
	Country        string `json:"country"`
	Hostname       string `json:"hostname"`
	Port           string `json:"port"`
	CurrentSession string `json:"current_session"`
}

func refreshServer() []Server {
	serverList := []Server{}

	c := colly.NewCollector()
	c.OnHTML("table#vg_hosts_table_id > tbody", func(h *colly.HTMLElement) {
		if h.Index == 2 {
			h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				if el.ChildAttr("td", "class") != UNWANTED_CLASS {
					var server Server
					country := el.ChildText("td:nth-child(1)")
					sessionCount := el.ChildText("td:nth-child(3) b > span")
					hostnameAndPort := el.ChildText("td:nth-child(8) p > span > b > span")
					if hostnameAndPort != "" {
						var port, ip string
						hostname := strings.Split(hostnameAndPort, ":")
						if len(hostname) > 1 {
							port = hostname[1]
							ip = hostname[0]
						} else {
							port = "443"
							ip = hostnameAndPort
						}
						server = Server{
							Country:        country,
							CurrentSession: sessionCount,
							Hostname:       ip,
							Port:           port,
						}
						serverList = append(serverList, server)
					}
				}
			})
		}
	})
	c.Visit("https://www.vpngate.net/en/")

	return serverList
}

func GetServerList() []string {
	var semiFinalServerList string
	var finalServerList []string
	serverList := refreshServer()
	for _, v := range serverList {
		semiFinalServerList += fmt.Sprintf("✅ %s - %s\n📶 %s | %s\n\n", v.Country, v.CurrentSession, v.Hostname, v.Port)
		if len(semiFinalServerList) > MESSAGE_LEN {
			finalServerList = append(finalServerList, semiFinalServerList)
			semiFinalServerList = ""
		}
	}
	return finalServerList
}
