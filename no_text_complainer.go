package slackchannelrouter


var (
	singletonNoTextComplainer = new(internalNoTextComplainer)
)


type NoTextComplainer interface {
	error
	NoTextComplainer()
}


type internalNoTextComplainer struct{}


func (complainer *internalNoTextComplainer) Error() string {
	return "Value passed for channel was invalid."
}


func (complainer *internalNoTextComplainer) NoTextComplainer() {
	// Nothing here.
}
