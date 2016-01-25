package notam

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetNotams(designator string) Notam {
	var notamList Notam
    
	designator = strings.ToUpper(designator)
	if len(designator) != 4 {
		log.Fatalln("Designator invalid")
	}

	v := url.Values{}
	v.Set("searchType", "0")
	v.Add("designatorsForLocation", designator)
    log.Println("start fetching NOTAMs....")
	resp, err := http.PostForm("https://notams.aim.faa.gov/notamSearch/search", v)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &notamList)
	if err != nil {
		log.Fatalln(err.Error())
	}
    log.Println("Feteched", len(notamList.Notamlist), "NOTAMs")
  
	return notamList
}
