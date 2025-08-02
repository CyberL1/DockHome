package types

type NeededDataType struct {
	Domain     string          `json:"domain"`
	Containers []ContainerData `json:"containers"`
}
