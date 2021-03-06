package triggers

import (
	"github.com/olivia-ai/olivia/language"
	"github.com/olivia-ai/olivia/util"
	"strings"
)

type Capital struct{}

// Replace the content of the response by the country and his capital
func (capital Capital) ReplaceContent() string {
	// Escape if it isn't a message which contains a Country
	if !strings.Contains(Response, "${CAPITAL}") {
		return Response
	}

	country := language.FindCountry(Entry)

	// If there isn't a country respond with a data.Message
	if country.Code == "" {
		return util.GetMessage("no country")
	}

	response := strings.Replace(Response, "${CAPITAL}", country.Capital, 1)
	response = strings.Replace(response, "${COUNTRY}", country.OfficialName+" "+country.Flag, 1)

	return response
}
