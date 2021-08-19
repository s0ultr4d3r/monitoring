package main

import (
	log "github.com/sirupsen/logrus"
)

//var srvSlc []Prblm

//parse convert apiAnsw slice to work slice
func parse() []Prblm {
	apiAnsw := apiRq().States
	srvSlc := make([]Prblm, 0, 100)
	for i := 0; i < len(apiAnsw); i++ {
		tmpPrblm := Prblm{
			Hostname:  apiAnsw[i].Alias,
			Tag:       apiAnsw[i].SmallMsg,
			Problem:   apiAnsw[i].MsgDecoded,
			MetaSrv:   apiAnsw[i].ExternalLink,
			StartTime: apiAnsw[i].TotalTime.Timestamp,
			Id:        uint64(apiAnsw[i].Pk),
		}
		srvSlc = append(srvSlc, tmpPrblm)
	}
	log.Error()

	return srvSlc
}

func ramSrvSlc() []Prblm {
	SrvSlc := parse()
	return SrvSlc
}

func prblmSlc(srvSlc []Prblm) []string {
	var prblmSlc []string
	prblmSlc = make([]string, 0, len(srvSlc))
	for _, v := range srvSlc {
		prblmSlc = append(prblmSlc, v.Tag)
	}
	return prblmSlc
}

func uniqPrblms(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}
	return us
}
