package slackchannelrouter


import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


// New returns an initialized SlackChannelRouter.
//
// 'token' is your secret authorization token that you get from Slack
// to be able to make API requests. If you don't have one already, you
// should be able to get it from here: https://api.slack.com/web
//
// 'channel' is the Slack channel you want to post to. For example:
// "#general".
//
// 'username' is the username you want messages posted to a Slack
// channel by this to appear from. (Note this is NOT the user name
// of your personal account.) For example: "Gobot"
//
// 'iconEmoji' is the avatar image you want messages posted to a
// Slack channel by this to have. For example: ":space_invader:".
func New(token string, channel string, username string, iconEmoji string) *SlackChannelRouter {
	router := SlackChannelRouter{
		token:token,
		channel:channel,
		username:username,
		iconEmoji:iconEmoji,
	}

	return &router
}


// SlackChannelRouter is a Router that posts a chat message to a Slack channel
// when asked to route a message.
type SlackChannelRouter struct {
	token     string
	channel   string
	username  string
	iconEmoji string
}


func (router *SlackChannelRouter) Route(message string, context map[string]interface{}) error {
	err := router.chatPostMessage(message, context)
	return err
}


func (router *SlackChannelRouter) chatPostMessage(message string, context map[string]interface{}) error {

	const href = "https://slack.com/api/chat.postMessage"


	// Construct "attachments" portion of the API request.
	//
	// NOTE that this portion is a JSON string BUT the stuff
	// "above" it cannot be JSON. (It won't work if it is.)
	// The stuff "above" it has to be URL query style.
	attachmentsString := ""
	if lenContext := len(context); nil != context && 0 < lenContext {

		fields := make([]map[string]string, lenContext)

		i := 0
		for key, valueInterface := range context {
			value := fmt.Sprintf("%v", valueInterface)

			fields[i] = map[string]string{
				"title":key,
				"value":value,
			}

			i++
		}


		attachments := []map[string]interface{}{
			{
				"fields": fields,
			},
		}


		jsonBytes, err := json.Marshal(attachments)
		if nil != err {
			return err
		}

		attachmentsString = string(jsonBytes)
	}


	// Construct the body of the HTTP POST request, that will be
	// sent to the Slack chat post message API end point.
	postBody := url.Values{
		"parse":      {"full"},
		"token":      {router.token},
		"channel":    {router.channel},
		"text":       {message},
	}
	if "" != router.username {
		postBody["username"] = []string{router.username}
	}
	if "" != router.iconEmoji {
		postBody["icon_emoji"] = []string{router.iconEmoji}
	}
	if "" != attachmentsString {
		postBody["attachments"] = []string{attachmentsString}
	}

	// Make the HTTP POST request.
	httpResponse, err := http.PostForm(href, postBody)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	// Check that the APi told us the request
	// went through OK.
	body, err := ioutil.ReadAll(httpResponse.Body)
	if nil != err {
		return err
	}

	responseInfo := struct{
		OK    bool   `json:"ok"`
		Error string `json:"error"`
	}{
		OK:false,
		Error:"",
	}

	if err := json.Unmarshal(body, &responseInfo); nil != err {
		return err
	}

	if !responseInfo.OK {
		switch responseInfo.Error {
		case "channel_not_found":
			return singletonChannelNotFoundComplainer
		case "not_in_channel":
			return singletonNotInChannelComplainer
		case "is_archived":
			return singletonIsArchivedComplainer
		case "msg_too_long":
			return singletonMsgTooLongComplainer
		case "no_text":
			return singletonNoTextComplainer
		case "rate_limited":
			return singletonRateLimitedComplainer
		case "not_authed":
			return singletonNotAuthedComplainer
		case "invalid_auth":
			return singletonInvalidAuthComplainer
		case "account_inactive":
			return singletonAccountInactiveComplainer
		default:
			return errors.New(responseInfo.Error)
		}
	}


	// Return.
	return nil
}
