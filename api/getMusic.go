package api

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const Url = "https://autumnfish.cn/"

// PlayListClass 歌单分类
type PlayListClass struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type Musics struct {
	ID int32 `json:"id"`
}

// GetPlayListClass 获取歌单分类
func (p *PlayListClass) GetPlayListClass() {
	url := fmt.Sprintf(Url + "playlist/ho")
	req, _ := http.NewRequest("GET", url, nil)
	_, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Info(err)
	}
}
