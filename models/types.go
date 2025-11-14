package models

type IndexDocument struct {
	Id    int64
	Words []string
}

func (indexDocument *IndexDocument) Equals(otherDocument IndexDocument) bool {
	return indexDocument.Id == otherDocument.Id
}

type SortOrder string

const (
	ASC  SortOrder = "ASC"
	DESC SortOrder = "DESC"
)

type Query struct {
	Should []string
	Must   []string
	Order  SortOrder
}
