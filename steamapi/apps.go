package steamapi

import (
	"net/url"
	"strconv"
)

// AppNews holds a Steam app's news items.
type AppNews []struct {
	Author        string `json:"author"`
	Contents      string `json:"contents"`
	Date          int    `json:"date"`
	FeedLabel     string `json:"feedlabel"`
	FeedName      string `json:"feedname"`
	GID           string `json:"gid"`
	IsExternalURL bool   `json:"is_external_url"`
	Title         string `json:"title"`
	URL           string `json:"url"`
}

// GetNewsForApp returns the latest news of a game specified by its AppID.
func (k Key) GetNewsForApp(appID, count, maxLen int) (AppNews, error) {
	var params = url.Values{}
	params.Add("appid", strconv.Itoa(appID))
	params.Add("count", strconv.Itoa(count))
	params.Add("maxlength", strconv.Itoa(maxLen))

	var respData struct {
		Appnews struct {
			Newsitems AppNews `json:"newsitems"`
		} `json:"appnews"`
	}
	err := requestAPI("ISteamNews", "GetNewsForApp", 2, params, respData)
	if err != nil {
		return AppNews{}, err
	}

	return respData.Appnews.Newsitems, nil
}
