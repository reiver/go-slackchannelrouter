package slackchannelrouter


var (
	singletonIsArchivedComplainer = new(internalIsArchivedComplainer)
)


type IsArchivedComplainer interface {
	error
	IsArchivedComplainer()
}


type internalIsArchivedComplainer struct{}


func (complainer *internalIsArchivedComplainer) Error() string {
	return "Value passed for channel was invalid."
}


func (complainer *internalIsArchivedComplainer) IsArchivedComplainer() {
	// Nothing here.
}
