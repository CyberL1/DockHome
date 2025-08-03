package types

type NeededData struct {
	Domain     string          `json:"domain"`
	Containers []ContainerData `json:"containers"`
}
