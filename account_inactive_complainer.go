package slackchannelrouter


var (
	singletonAccountInactiveComplainer = new(internalAccountInactiveComplainer)
)


type AccountInactiveComplainer interface {
	error
	AccountInactiveComplainer()
}


type internalAccountInactiveComplainer struct{}


func (complainer *internalAccountInactiveComplainer) Error() string {
	return "Value passed for channel was invalid."
}


func (complainer *internalAccountInactiveComplainer) AccountInactiveComplainer() {
	// Nothing here.
}
