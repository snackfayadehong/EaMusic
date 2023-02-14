package api

import (
	logs "EaMusic/log"
	"EaMusic/utils"
	"encoding/json"
	"fmt"
)

const Url = "https://autumnfish.cn/"

// PlayListClass 歌单分类
type PlayListClass struct {
	Tags []struct {
		PlaylistTag struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"playlistTag"`
	} `json:"tags"`
	Code int `json:"code"`
}

type PlayList struct {
}

// GetPlayListClass 获取歌单分类
func (pc *PlayListClass) GetPlayListClass() (playlistClass *PlayListClass) {
	url := fmt.Sprintf(Url + "playlist/hot")
	res, err := HttpsRequest(url, "GET", nil)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(res, &playlistClass)
	if playlistClass.Code != 200 {
		logs.MyLog.Info(`请求失败, Code: `, playlistClass.Code, err)
		return nil
	} else {
		logs.MyLog.Info("请求成功: ", playlistClass.Tags[0])
	}
	return playlistClass
}

// GetPlayListByClass 根据歌单分类获取热门歌单
func (pl *PlayList) GetPlayListByClass() {
	playListCls := &PlayListClass{}
	playList := playListCls.GetPlayListClass()
	cat := (*playList).Tags[utils.RandNum()].PlaylistTag.Name
	fmt.Println(cat)
	url := fmt.Sprintf(Url + "/top/playlist?cat=" + cat + "&limit=10")
	res, err := HttpsRequest(url, "GET", nil)
	if err != nil {
		return
	}
	fmt.Println(res)
}
