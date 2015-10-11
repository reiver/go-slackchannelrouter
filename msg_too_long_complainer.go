package slackchannelrouter


var (
	singletonMsgTooLongComplainer = new(internalMsgTooLongComplainer)
)


type MsgTooLongComplainer interface {
	error
	MsgTooLongComplainer()
}


type internalMsgTooLongComplainer struct{}


func (complainer *internalMsgTooLongComplainer) Error() string {
	return "Value passed for channel was invalid."
}


func (complainer *internalMsgTooLongComplainer) MsgTooLongComplainer() {
	// Nothing here.
}
