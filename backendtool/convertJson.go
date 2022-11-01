package backendtool

import (
	"net/http"
)

func thirdPartInterfaceExample() {
	response, err := http.Get("http://data.nba.net/10s/prod/v1/2022/schedule.json")
		if err != nil || response.StatusCode != http.StatusOK {
				return
		}
		// bodyValue, err := io.ReadAll(response.Body)
		defer response.Body.Close()
}