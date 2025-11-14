package models

type IndexDocument struct {
	Id    int64
	Words []string
}

func (indexDocument *IndexDocument) Equals(otherDocument IndexDocument) bool {
	return indexDocument.Id == otherDocument.Id
}
