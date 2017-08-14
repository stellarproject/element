package runtime

// Spec is the runtime specification
type Spec struct {
	Namespace string            `json:"namespace"`
	Image     string            `json:"image"`
	Protocol  string            `json:"protocol"`
	Port      int               `json:"port"`
	Runtime   string            `json:"runtime"`
	Labels    map[string]string `json:"labels,omitempty"`
}
