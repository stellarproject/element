package runtime

// Spec is the runtime specification
type Spec struct {
	ID        string            `json:"id,omitempty"`
	Namespace string            `json:"namespace"`
	Image     string            `json:"image"`
	Runtime   string            `json:"runtime"`
	Labels    map[string]string `json:"labels,omitempty"`
}
