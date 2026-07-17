package main

import (
	"encoding/xml"
	"fmt"
)

type RaceReport struct {
	XMLName             xml.Name     `xml:"report"`
	CompetitionLocation string       `xml:"competition>location"`
	CompetitionClass    string       `xml:"competition>class"`
	Results             []RaceResult `xml:"racer"`
}

type RaceResult struct {
	XMLName   xml.Name `xml:"racer"`
	GlobalId  int      `xml:"global_id,attr,omitempty"`
	Nick      string   `xml:"nick"`
	BestLapMs int      `xml:"best_lam_ms"`
	Laps      float32  `xml:"laps"`
}

func FilterXML(input string, laps float32) (output string, err error) {
	var rp RaceReport
	err = xml.Unmarshal([]byte(input), &rp)
	if err != nil {
		return
	}
	filter := make([]RaceResult, 0, len(rp.Results))
	for _, racer := range rp.Results {
		if racer.Laps > laps {
			filter = append(filter, racer)
		}
	}
	rp.Results = filter
	var data []byte
	data, err = xml.MarshalIndent(rp, "", "  ")
	if err != nil {
		return
	}
	output = string(data)
	return
}

func main() {
	input := `<report>
  <competition>
    <location>РФ, Санкт-Петербург, Дворец творчества юных техников</location>
    <class>ТА-24</class>
  </competition>
  <racer global_id="100">
    <nick>RacerX</nick>
    <best_lap_ms>61012</best_lap_ms>
    <laps>52.3</laps>
  </racer>
  <racer global_id="127">
    <nick>Иван The Шумахер</nick>
    <best_lap_ms>61023</best_lap_ms>
    <laps>51</laps>
  </racer>
  <racer global_id="203">
    <nick>Петя Иванов</nick>
    <best_lap_ms>63000</best_lap_ms>
    <laps>49.9</laps>
    <!--Болид не соответствует техническому регламенту, 
    результат не учитывается в общем рейтинге-->
  </racer>
  <racer>
    <nick>Гость 1</nick>
    <best_lap_ms>123001</best_lap_ms>
    <laps>25.8</laps>
  </racer>
</report>`

	output, err := FilterXML(input, 50)
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}
	fmt.Println(output)
}
