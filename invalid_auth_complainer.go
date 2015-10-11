package slackchannelrouter


var (
	singletonInvalidAuthComplainer = new(internalInvalidAuthComplainer)
)


type InvalidAuthComplainer interface {
	error
	InvalidAuthComplainer()
}


type internalInvalidAuthComplainer struct{}


func (complainer *internalInvalidAuthComplainer) Error() string {
	return "Value passed for channel was invalid."
}


func (complainer *internalInvalidAuthComplainer) InvalidAuthComplainer() {
	// Nothing here.
}
