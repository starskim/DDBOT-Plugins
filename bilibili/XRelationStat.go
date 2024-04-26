package bilibili

import (
	"github.com/starskim/DDBOT/proxy_pool"
	"github.com/starskim/DDBOT/requests"
	"github.com/starskim/DDBOT/utils"
	"time"
)

const (
	PathXRelationStat = "/x/relation/stat"
)

type XRelationStatRequest struct {
	Mid int64 `json:"vmid"`
}

func XRelationStat(mid int64) (*XRelationStatResponse, error) {
	st := time.Now()
	defer func() {
		ed := time.Now()
		logger.WithField("FuncName", utils.FuncName()).Tracef("cost %v", ed.Sub(st))
	}()
	url := BPath(PathXRelationStat)
	params, err := utils.ToParams(&XRelationStatRequest{
		Mid: mid,
	})
	if err != nil {
		return nil, err
	}
	var opts = []requests.Option{
		requests.ProxyOption(proxy_pool.PreferNone),
		requests.TimeoutOption(time.Second * 15),
		AddUAOption(),
		delete412ProxyOption,
	}
	opts = append(opts, GetVerifyOption()...)
	xrsr := new(XRelationStatResponse)
	err = requests.Get(url, params, xrsr, opts...)
	if err != nil {
		return nil, err
	}
	return xrsr, nil
}
