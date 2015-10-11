package slackchannelrouter


var (
	singletonChannelNotFoundComplainer = new(internalChannelNotFoundComplainer)
)


type ChannelNotFoundComplainer interface {
	error
	ChannelNotFoundComplainer()
}


type internalChannelNotFoundComplainer struct{}


func (complainer *internalChannelNotFoundComplainer) Error() string {
	return "Value passed for channel was invalid."
}


func (complainer *internalChannelNotFoundComplainer) ChannelNotFoundComplainer() {
	// Nothing here.
}
