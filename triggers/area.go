package triggers

import (
	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/util"
	"strconv"
	"strings"
)

type Area struct{}

func (area Area) ReplaceContent() string {
	// Escape if it isn't a message which contains a Country
	if !strings.Contains(Response, "${AREA}") {
		return Response
	}

	country := language.FindCountry(Entry)

	// If there isn't a country respond with a data.Message
	if country.Code == "" {
		return util.GetMessage("no country")
	}

	response := strings.Replace(Response, "${AREA}", strconv.FormatFloat(country.Area, 'f', 2, 64), 1)
	response = strings.Replace(response, "${COUNTRY}", country.OfficialName+" "+country.Flag, 1)

	return response
}
