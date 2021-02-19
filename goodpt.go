package goodpt

import (
	"time"
)

type Trains []struct {
	ID                     string    `json:"@id"`
	Type                   string    `json:"@type"`
	DcDate                 time.Time `json:"dc:date"`
	Context                string    `json:"@context"`
	DctValid               time.Time `json:"dct:valid"`
	OdptDelay              int       `json:"odpt:delay,omitempty"`
	OwlSameAs              string    `json:"owl:sameAs"`
	OdptRailway            string    `json:"odpt:railway"`
	OdptOperator           string    `json:"odpt:operator"`
	OdptToStation          string    `json:"odpt:toStation"`
	OdptTrainType          string    `json:"odpt:trainType,omitempty"`
	OdptFromStation        string    `json:"odpt:fromStation"`
	OdptTrainNumber        string    `json:"odpt:trainNumber"`
	OdptRailDirection      string    `json:"odpt:railDirection"`
	OdptCarComposition     int       `json:"odpt:carComposition,omitempty"`
	OdptDestinationStation []string  `json:"odpt:destinationStation"`
	OdptViaRailway         []string  `json:"odpt:viaRailway,omitempty"`
	OdptTrainOwner         string    `json:"odpt:trainOwner,omitempty"`
	OdptOriginStation      []string  `json:"odpt:originStation,omitempty"`
}

type Railways []struct {
	ID               string    `json:"@id"`
	Type             string    `json:"@type"`
	DcDate           time.Time `json:"dc:date"`
	Context          string    `json:"@context"`
	DcTitle          string    `json:"dc:title"`
	OwlSameAs        string    `json:"owl:sameAs"`
	OdptLineCode     string    `json:"odpt:lineCode"`
	OdptOperator     string    `json:"odpt:operator"`
	OdptStationOrder []struct {
		OdptIndex        int    `json:"odpt:index"`
		OdptStation      string `json:"odpt:station"`
		OdptStationTitle struct {
			En string `json:"en"`
			Ja string `json:"ja"`
		} `json:"odpt:stationTitle"`
	} `json:"odpt:stationOrder"`
	OdptAscendingRailDirection  string `json:"odpt:ascendingRailDirection"`
	OdptDescendingRailDirection string `json:"odpt:descendingRailDirection"`
	OdptRailwayTitle struct {
		En     string `json:"en"`
		Ja     string `json:"ja"`
		Ko     string `json:"ko"`
		ZhHans string `json:"zh-Hans"`
		ZhHant string `json:"zh-Hant"`
	} `json:"odpt:railwayTitle,omitempty"`
}

type Stations []struct {
	ID               string    `json:"@id"`
	Type             string    `json:"@type"`
	DcDate           time.Time `json:"dc:date"`
	GeoLat           float64   `json:"geo:lat,omitempty"`
	Context          string    `json:"@context"`
	DcTitle          string    `json:"dc:title"`
	GeoLong          float64   `json:"geo:long,omitempty"`
	OwlSameAs        string    `json:"owl:sameAs"`
	OdptRailway      string    `json:"odpt:railway"`
	OdptOperator     string    `json:"odpt:operator"`
	OdptStationCode  string    `json:"odpt:stationCode"`
	OdptPassengerSurvey   []string `json:"odpt:passengerSurvey,omitempty"`
	OdptStationTimetable  []string `json:"odpt:stationTimetable"`
	OdptConnectingRailway []string `json:"odpt:connectingRailway,omitempty"`
	OdptStationTitle      struct {
		En     string `json:"en"`
		Ja     string `json:"ja"`
		Ko     string `json:"ko"`
		ZhHans string `json:"zh-Hans"`
		ZhHant string `json:"zh-Hant"`
	} `json:"odpt:stationTitle,omitempty"`
}
	
type RailDirections []struct {
	ID                     string    `json:"@id"`
	Type                   string    `json:"@type"`
	DcDate                 time.Time `json:"dc:date"`
	Context                string    `json:"@context"`
	DcTitle                string    `json:"dc:title"`
	OwlSameAs              string    `json:"owl:sameAs"`
	OdptOperator           string    `json:"odpt:operator"`
	OdptRailDirectionTitle struct {
		En string `json:"en"`
		Ja string `json:"ja"`
	} `json:"odpt:railDirectionTitle"`
}

type TrainTimetables []struct {
	ID                       string    `json:"@id"`
	Type                     string    `json:"@type"`
	DcDate                   time.Time `json:"dc:date"`
	Context                  string    `json:"@context"`
	DctIssued                string    `json:"dct:issued"`
	OdptTrain                string    `json:"odpt:train"`
	OwlSameAs                string    `json:"owl:sameAs"`
	OdptRailway              string    `json:"odpt:railway"`
	OdptCalendar             string    `json:"odpt:calendar"`
	OdptOperator             string    `json:"odpt:operator"`
	OdptTrainType            string    `json:"odpt:trainType"`
	OdptTrainNumber          string    `json:"odpt:trainNumber"`
	OdptOriginStation        []string  `json:"odpt:originStation"`
	OdptRailDirection        string    `json:"odpt:railDirection"`
	OdptDestinationStation   []string  `json:"odpt:destinationStation"`
	OdptTrainTimetableObject []struct {
		OdptDepartureTime    string `json:"odpt:departureTime,omitempty"`
		OdptDepartureStation string `json:"odpt:departureStation,omitempty"`
		OdptArrivalTime      string `json:"odpt:arrivalTime,omitempty"`
		OdptArrivalStation   string `json:"odpt:arrivalStation,omitempty"`
	} `json:"odpt:trainTimetableObject"`
}

type StationTimetables []struct {
	ID                         string    `json:"@id"`
	Type                       string    `json:"@type"`
	DcDate                     time.Time `json:"dc:date"`
	Context                    string    `json:"@context"`
	DctIssued                  string    `json:"dct:issued"`
	OwlSameAs                  string    `json:"owl:sameAs"`
	OdptRailway                string    `json:"odpt:railway"`
	OdptStation                string    `json:"odpt:station"`
	OdptCalendar               string    `json:"odpt:calendar"`
	OdptOperator               string    `json:"odpt:operator"`
	OdptRailDirection          string    `json:"odpt:railDirection"`
	OdptStationTimetableObject []struct {
		OdptTrain              string   `json:"odpt:train"`
		OdptTrainType          string   `json:"odpt:trainType"`
		OdptTrainNumber        string   `json:"odpt:trainNumber"`
		OdptDepartureTime      string   `json:"odpt:departureTime"`
		OdptDestinationStation []string `json:"odpt:destinationStation"`
		OdptIsLast             bool     `json:"odpt:isLast,omitempty"`
	} `json:"odpt:stationTimetableObject"`
}

type Operators []struct {
	ID                string    `json:"@id"`
	Type              string    `json:"@type"`
	DcDate            time.Time `json:"dc:date"`
	Context           string    `json:"@context"`
	DcTitle           string    `json:"dc:title"`
	OwlSameAs         string    `json:"owl:sameAs"`
	OdptOperatorTitle struct {
		En string `json:"en"`
		Ja string `json:"ja"`
	} `json:"odpt:operatorTitle"`
}

type TrainTypes []struct {
	ID                 string    `json:"@id"`
	Type               string    `json:"@type"`
	DcDate             time.Time `json:"dc:date"`
	Context            string    `json:"@context"`
	DcTitle            string    `json:"dc:title"`
	OwlSameAs          string    `json:"owl:sameAs"`
	OdptOperator       string    `json:"odpt:operator"`
	OdptTrainTypeTitle struct {
		En string `json:"en"`
		Ja string `json:"ja"`
	} `json:"odpt:trainTypeTitle"`
}

type Calendar []struct {
	ID                string    `json:"@id"`
	Type              string    `json:"@type"`
	DcDate            time.Time `json:"dc:date,omitempty"`
	Context           string    `json:"@context"`
	DcTitle           string    `json:"dc:title"`
	OwlSameAs         string    `json:"owl:sameAs"`
	OdptCalendarTitle struct {
		En string `json:"en"`
		Ja string `json:"ja"`
	} `json:"odpt:calendarTitle,omitempty"`
	OdptOperator string   `json:"odpt:operator,omitempty"`
	OdptDay      []string `json:"odpt:day,omitempty"`
	OdptNote     string   `json:"odpt:note,omitempty"`
	OdptDuration string   `json:"odpt:duration,omitempty"`
}
