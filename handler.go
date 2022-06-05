package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Represents the expected JSON structure of the datasource. Customise to suit the data coming from the source API.
type PriceDataSource struct {
	Raw struct {
		Eth struct {
			Usd struct {
				Price float64 `json:"PRICE"`
			}
		} `json:"ETH"`
	} `json:"RAW"`
}

// Represents the JSON structure of the EA's data output. Customise the contents of this struct.
type DataOutput struct {
	Price float64 `json:"price"`
}

// Represents the JSON structure of the overall response. Should NOT need customisation.
type ExternalAdapterOutput struct {
	Data  *DataOutput `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// Main request handler of this external adapter. Customise to your heart's content.
func mainHandler(c *gin.Context) {
	defer requestsProcessed.Inc() // increases the metrics counter at the end of the request

	// Fetch the data from the source URL
	res, err := http.Get("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=ETH&tsyms=USD")
	if err != nil {
		c.JSON(http.StatusBadGateway, ExternalAdapterOutput{Error: errors.Wrap(err, "Unable to fetch data from source").Error()})
	} else {
		defer res.Body.Close() // tidies up the open input stream at the end of the request

		// Parse the data retrieved into the PriceDataSource struct
		source := PriceDataSource{}
		err := json.NewDecoder(res.Body).Decode(&source)
		if err != nil {
			c.JSON(http.StatusBadGateway, ExternalAdapterOutput{Error: errors.Wrap(err, "Unable to parse data received from source").Error()})
		} else {
			returnValue := ExternalAdapterOutput{}
			// assign the price to the output data struct. Good place to make any modification.
			returnValue.Data = &DataOutput{Price: source.Raw.Eth.Usd.Price}
			c.JSON(http.StatusOK, returnValue)
		}
	}

}
