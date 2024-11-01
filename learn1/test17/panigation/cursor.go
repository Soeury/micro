package panigation

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

// 基于游标的分页
type Token string

type Page struct {
	NextID        string `json:"next_id"`
	NextTimeAtUTC int64  `json:"next_time_at_utc"`
	PageSize      int64  `json:"page_size"`
}

func (p Page) InValid() bool {
	return p.NextID == "" || p.NextTimeAtUTC == 0 || p.NextTimeAtUTC > time.Now().Unix() || p.PageSize <= 0
}

// Encode 把分页信息进行编码
func (p Page) Encode() Token {

	b, err := json.Marshal(p)
	if err != nil {
		return Token("")
	}

	return Token(base64.StdEncoding.EncodeToString(b))
}

// Decode 基于已有的 token 解析分页信息  // shijiesahng laidaozheli mianduizhe
func (t Token) Decode() Page {

	var result Page
	if len(t) == 0 {
		return result
	}

	bs, err := base64.StdEncoding.DecodeString(string(t))
	if err != nil {
		return result
	}

	err = json.Unmarshal(bs, &result)
	if err != nil {
		return result
	}
	return result
}
