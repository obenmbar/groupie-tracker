package groupino

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func FetchData(url string, datass any) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data from url %s: %v", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("errrur en status code", err)
	}

	err1 := json.NewDecoder(resp.Body).Decode(datass)
	if err1 != nil {
		return fmt.Errorf("error en decodage de json")
	}
	return nil
}
