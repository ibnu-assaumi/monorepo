package seaotter

import (
	"net/http"
	"time"

	"github.com/Bhinneka/candi/candiutils"
)

type seaotterRESTImpl struct {
	host    string
	authKey string
	httpReq candiutils.HTTPRequest
}

// NewSeaotterServiceREST constructor
func NewSeaotterServiceREST(host string, authKey string) Seaotter {

	return &seaotterRESTImpl{
		host:    host,
		authKey: authKey,
		httpReq: candiutils.NewHTTPRequest(
			candiutils.HTTPRequestSetRetries(5),
			candiutils.HTTPRequestSetSleepBetweenRetry(500*time.Millisecond),
			candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
			candiutils.HTTPRequestSetBreakerName("seaotter"),
		),
	}
}
