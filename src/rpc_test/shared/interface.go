package shared

type Owner struct {
	Name string
}

type Task struct {
	ID          int
	Description string
	Owner       Owner
	Status      int
}

type Void struct {
}
