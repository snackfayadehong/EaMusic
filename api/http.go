package api

import (
	logs "EaMusic/log"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func HttpsRequest(url string, method string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logs.MyLog.WithFields(logrus.Fields{
			"url": fmt.Sprintf("请求接口地址:%v", url),
		}).Info("请求接口失败,", err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logs.MyLog.WithFields(logrus.Fields{
			"url": fmt.Sprintf("请求接口地址:%v", url),
		}).Info("send https request and get response failure")
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logs.MyLog.Info("关闭resp失败,", err)
		}
	}(resp.Body)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.MyLog.Info("读取resp.body err,", err)
		return nil, err
	}
	return b, err
}
