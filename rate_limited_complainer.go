package slackchannelrouter


var (
	singletonRateLimitedComplainer = new(internalRateLimitedComplainer)
)


type RateLimitedComplainer interface {
	error
	RateLimitedComplainer()
}


type internalRateLimitedComplainer struct{}


func (complainer *internalRateLimitedComplainer) Error() string {
	return "Value passed for channel was invalid."
}


func (complainer *internalRateLimitedComplainer) RateLimitedComplainer() {
	// Nothing here.
}
