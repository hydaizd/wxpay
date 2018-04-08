package wxpay

import (
	"net/http"
	"io/ioutil"
	"errors"
	"encoding/xml"
)

// GetTradeNotification https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_7&index=3
func (this *WXPay) GetTradeNotification(req *http.Request) (*TradeNotification, error) {
	key, err := this.getKey()
	if err != nil {
		return nil, err
	}
	return GetTradeNotification(req, key)
}

func GetTradeNotification(req *http.Request, key string) (noti *TradeNotification, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}

	var data, _ = ioutil.ReadAll(req.Body)

	if _, err := verifyResponseData(data, key); err != nil {
		return nil, err
	}

	if err = xml.Unmarshal(data, &noti); err != nil {
		return nil, err
	}
	return noti, err
}