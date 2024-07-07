package servers

import (
	"context"

	"github.com/Lornzo/gobinance/commons"
)

func NewServerTimestamp(baseURL string) *ServerTimestamp {
	var serverTimestamp *ServerTimestamp = &ServerTimestamp{
		Base: commons.NewBinanceRestfulBase(),
	}
	serverTimestamp.Base.SetBaseURL(baseURL)
	return serverTimestamp
}

type ServerTimestamp struct {
	Base commons.RestfulBase
}

func (s *ServerTimestamp) GetAPIPathes() []string {
	return []string{"fapi", "v1", "time"}
}

func (s *ServerTimestamp) GetWeights() int {
	return 1
}

func (s *ServerTimestamp) DoRequest(ctx context.Context) (int64, error) {

	type Response struct {
		ServerTime int64 `json:"serverTime"`
	}

	var (
		resp Response
		err  error
	)

	s.initApiUrl()

	if _, err = s.Base.GET(ctx, &resp); err != nil {
		return 0, err
	}

	return resp.ServerTime, nil

}

func (s *ServerTimestamp) initApiUrl() {
	s.Base.SetPathes(s.GetAPIPathes()...)
}
