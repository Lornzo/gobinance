package servers

import (
	"context"

	"github.com/Lornzo/gobinance/commons"
)

func NewPingServer(baseURL string) *PingServer {
	var pingServer *PingServer = &PingServer{
		Base: commons.NewBinanceRestfulBase(),
	}
	pingServer.Base.SetBaseURL(baseURL)
	return pingServer

}

type PingServer struct {
	Base commons.RestfulBase
}

func (p *PingServer) GetAPIPathes() []string {
	return []string{"fapi", "v1", "ping"}
}

func (p *PingServer) GetWeights() int {
	return 1
}

func (p *PingServer) DoRequest(ctx context.Context) error {

	var err error

	p.initApiUrl()

	if _, err = p.Base.GET(ctx, nil); err != nil {
		return err
	}

	return nil
}

func (p *PingServer) initApiUrl() {
	p.Base.SetPathes(p.GetAPIPathes()...)
}
