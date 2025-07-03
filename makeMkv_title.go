package gomakemkv

type MakeMkvTitle struct {
	Properties map[string]MakeMkvValue         `json:"properties"`
	Streams    map[int]map[string]MakeMkvValue `json:"streams"`
}
