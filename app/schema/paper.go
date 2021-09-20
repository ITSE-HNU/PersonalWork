package schema

type Topic struct {
	ID    int
	Title string
}

type Paper struct {
	Name  string
	Topic []Topic
}
