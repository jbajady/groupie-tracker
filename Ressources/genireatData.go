package Func

func GenriateData(id int) {
	er := FetchData(Artists[id-1].Relations, &Relation)
	er1 := FetchData(Artists[id-1].Locations, &Location)
	er2 := FetchData(Artists[id-1].ConcertDates, &Date)
	if er != nil || er1 != nil || er2 != nil {
		return
	}
}
