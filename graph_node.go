package neoism

type GraphNode interface {
	Relate(relType string, destId int, p Props) (*Relationship, error)
}
