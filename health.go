package element

// Health returns the relative health score of the node (lower is better)
func (a *Agent) Health() int {
	return a.members.GetHealthScore()
}
