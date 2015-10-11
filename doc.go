/*
Package slackchannelrouter provides a router to be used with the flog package ( https://github.com/reiver/go-flog ).


Recommended Usage

The Slack channel router is blocking. And since it is making an API request to
Slack's API servers (across the Internet) it can block for some time.

To deal with this, it is recommended you wrap the Slack channel router in a flog.NonBlockingRouter.
For example:

	var slackChannelRouter flog.Router = slackchannelrouter.New(token, channel, username, iconEmoji)
	
	slackChannelRouter = flog.NewNonBlockingRouter(slackChannelRouter)

Basic Usage

Putting this altogether, basic usage look something like this:

	var slackChannelRouter flog.Router = slackchannelrouter.New(token, channel, username, iconEmoji)
	
	// Note that 'filterMessageForSlackFunc' is some func that decides which messages are
	// allowed to get posted to Slack.
	slackChannelRouter = flog.NewFilteringRouter(slackChannelRouter, filterMessageForSlackFunc)
	
	slackChannelRouter = flog.NewNonBlockingRouter(slackChannelRouter)
	
	
	flogger := flog.New(slackChannelRouter)

(Although likely you would be using other routers too, besides just the Slack channel router.)

Advange Usage

You may want to modify the log before sending it to Slack.

For example, if you logged the following message:

	flogger.Printf("Received error: %v", err)

And again for example, let's say that message gets rendered as:

	Received error: Could not connect to queue server.

Then we might want to make this change the message into the following
before posting it to Slack:

	@group: Received error: Could not connect to queue server. :feelsgood:

To do that we would use the flog.MappingRouter.

So for example using that, plus also the flog.NonBlockingRouter mentioned already:

	var slackChannelRouter flog.Router = slackchannelrouter.New(token, channel, username, iconEmoji)

	slackChannelRouter = NewMappingRouter(slackChannelRouter, func(message string, context map[string]interface{})(string, map[string]interface{}){
		newMessage := fmt.Sprintf("@group: %s :feelsgood:", message)
		return newMessage, context
	})

	slackChannelRouter = flog.NewNonBlockingRouter(slackChannelRouter)

Flogger

For more information on the flogger, see the flog package: https://github.com/reiver/go-flog

*/
package slackchannelrouter
