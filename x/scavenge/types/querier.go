package types


// Query endpoints supported by the scavenge querier
const (
	// TODO: Describe query parameters, update <action> with your query
	// Query<Action>    = "<action>"
	QueryListScavenges = "list"
	QueryGetScavenge   = "get"
	QueryCommit        = "commit"
)


Below you will be able how to set your own queries:
type QueryResScavenges []string
type QueryResList []string
// QueryResList Queries Result Payload for a query



// implement fmt.Stringer
func (n QueryResList) String() string {
	return strings.Join(n[:], "\n")
}

func (n QueryResScavenges) String() string {
	return strings.Join(n[:], "\n")

*/
