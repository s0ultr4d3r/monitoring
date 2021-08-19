package main

func sortPrblm(slc []Prblm, tag string) []Prblm {
	prSlc := make([]Prblm, 0, 50)
	for i := 0; i < len(slc); i++ {
		if slc[i].Tag == tag {
			prSlc = append(prSlc, slc[i])
		}
	}
	return prSlc
}

func sortSnap() []Prblm {
	snapSlc := sortPrblm(ramSrvSlc(), "snapshot is old")
	return snapSlc
}

func sortMemFree() []Prblm {
	memFreeSlc := sortPrblm(ramSrvSlc(), "mem_free")
	return memFreeSlc
}

/* func sortOldConf() []Prblm {
	oldConfSlc := sortPrblm(ramSrvSlc(), "old config")
	return oldConfSlc
} */
