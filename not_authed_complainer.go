package slackchannelrouter


var (
	singletonNotAuthedComplainer = new(internalNotAuthedComplainer)
)


type NotAuthedComplainer interface {
	error
	NotAuthedComplainer()
}


type internalNotAuthedComplainer struct{}


func (complainer *internalNotAuthedComplainer) Error() string {
	return "Value passed for channel was invalid."
}


func (complainer *internalNotAuthedComplainer) NotAuthedComplainer() {
	// Nothing here.
}
