package config

var InfosWprecon informationsWPRecon

type informationsWPRecon struct {
	Target        string
	TotalRequests int

	Verbose bool
	Force   bool

	TimeStart, TimeFinish string
	TimeSleepRequests     int
	TimeOutRequests       int

	OtherInformationsString    map[string]string
	OtherInformationsInt       map[string]int
	OtherInformationsSlice     map[string][]string
	OtherInformationsBool      map[string]bool
	OtherInformationsMapString map[string]map[string]string
}

func init() {
	InfosWprecon.OtherInformationsSlice = map[string][]string{}
	InfosWprecon.OtherInformationsString = make(map[string]string)
	InfosWprecon.OtherInformationsInt = make(map[string]int)
	InfosWprecon.OtherInformationsBool = make(map[string]bool)
	InfosWprecon.OtherInformationsMapString = make(map[string]map[string]string)

	InfosWprecon.OtherInformationsMapString["target.http.plugins.versions"] = make(map[string]string)
	InfosWprecon.OtherInformationsMapString["target.http.themes.versions"] = make(map[string]string)
}
