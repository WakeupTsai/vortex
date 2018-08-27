package entity

// ContainerResourceMetrics is the structure for Container Resource Metrics
type ContainerResourceMetrics struct {
	CPUUsagePercentage []SamplePair `json:"cpuUsagePercentage"`
	MemoryUsageBytes   []SamplePair `json:"memoryUsageBytes"`
}

// ContainerStatusMetrics is the structure for Container Status Metrics
type ContainerStatusMetrics struct {
	Status           string `json:"status"`
	WaitingReason    string `json:"waitingReason"`
	TerminatedReason string `json:"terminatedReason"`
	RestartTime      int    `json:"restartTime"`
}

// ContainerDetailMetrics is the structure  for Container Detail Metrics
type ContainerDetailMetrics struct {
	ContainerName string   `json:"containerName"`
	CreatedAt     int      `json:"createAt"`
	Pod           string   `json:"pod"`
	Namespace     string   `json:"namespace"`
	Node          string   `json:"node"`
	Image         string   `json:"image"`
	Command       []string `json:"command"`
}

// ContainerMetrics is the structure for Container Metrics
type ContainerMetrics struct {
	Detail   ContainerDetailMetrics   `json:"detail"`
	Status   ContainerStatusMetrics   `json:"status"`
	Resource ContainerResourceMetrics `json:"resource"`
}
