package neoism

type GraphDB interface {
  CreateNode(p Props) (*Node, error)
  NodesByLabel(label string) ([]*Node, error)
}
