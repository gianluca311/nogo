package notam

//Notam Struct for Notamlist
type Notam struct {
	Notamlist []struct {
		Facilitydesignator                   string `json:"facilityDesignator"`
		Notamnumber                          string `json:"notamNumber"`
		Featurename                          string `json:"featureName"`
		Issuedate                            string `json:"issueDate"`
		Startdate                            string `json:"startDate"`
		Enddate                              string `json:"endDate"`
		Source                               string `json:"source"`
		Sourcetype                           string `json:"sourceType"`
		Icaomessage                          string `json:"icaoMessage"`
		Traditionalmessage                   string `json:"traditionalMessage"`
		Plainlanguagemessage                 string `json:"plainLanguageMessage"`
		Traditionalmessagefrom4Thword        string `json:"traditionalMessageFrom4thWord"`
		Icaoid                               string `json:"icaoId"`
		Accountid                            string `json:"accountId"`
		Airportname                          string `json:"airportName"`
		Procedure                            bool   `json:"procedure"`
		Userid                               int    `json:"userID"`
		Transactionid                        int    `json:"transactionID"`
		Cancelledorexpired                   bool   `json:"cancelledOrExpired"`
		Digitaltpplink                       bool   `json:"digitalTppLink"`
		Status                               string `json:"status"`
		Contractionsexpandedforplainlanguage bool   `json:"contractionsExpandedForPlainLanguage"`
		Keyword                              string `json:"keyword"`
		Snowtam                              bool   `json:"snowtam"`
		Geometry                             string `json:"geometry"`
		Digitallytransformed                 bool   `json:"digitallyTransformed"`
		Messagedisplayed                     string `json:"messageDisplayed"`
		Morethan300Chars                     bool   `json:"moreThan300Chars"`
		Showingfulltext                      bool   `json:"showingFullText"`
		Mappointer                           string `json:"mapPointer"`
		Locid                                int    `json:"locID"`
		Defaulticao                          bool   `json:"defaultIcao"`
		Crossovertransactionid               int    `json:"crossoverTransactionID"`
		Crossoveraccountid                   string `json:"crossoverAccountID"`
		Requestid                            int    `json:"requestID"`
	} `json:"notamList"`
	Startrecordcount    int    `json:"startRecordCount"`
	Endrecordcount      int    `json:"endRecordCount"`
	Totalnotamcount     int    `json:"totalNotamCount"`
	Filteredresultcount int    `json:"filteredResultCount"`
	Criteriacaption     string `json:"criteriaCaption"`
	Searchdatetime      string `json:"searchDateTime"`
	Error               string `json:"error"`
	Countsbytype        []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"countsByType"`
	Requestid int `json:"requestID"`
}