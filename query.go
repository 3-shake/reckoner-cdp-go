package reckonercdp

import (
	"encoding/json"
	"io/ioutil"
)

type QueryResponse struct {
	HistoryID int           `json:"history_id"`
	UUID      string        `json:"uuid"`
	HasNext   bool          `json:"has_next"`
	Records   []interface{} `json:"records"`
}

type QueryRequest struct {
	Query string `json:"query"`
}

func (client *Client) Query(query string) (*QueryResponse, error) {
	res, err := client.post("/api/v1/query", QueryRequest{Query: query})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		return nil, client.createError(res, "query failed")
	}

	var queryResponse *QueryResponse
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &queryResponse)
	if err != nil {
		return nil, err
	}

	return queryResponse, nil
}
