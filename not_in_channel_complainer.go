package slackchannelrouter


var (
	singletonNotInChannelComplainer = new(internalNotInChannelComplainer)
)


type NotInChannelComplainer interface {
	error
	NotInChannelComplainer()
}


type internalNotInChannelComplainer struct{}


func (complainer *internalNotInChannelComplainer) Error() string {
	return "Cannot post user messages to a channel they are not in."
}


func (complainer *internalNotInChannelComplainer) NotInChannelComplainer() {
	// Nothing here.
}
