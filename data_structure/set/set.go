package main

import "fmt"

type customSet struct {
	container map[string]struct{}
}

func makeSet() *customSet {
	return &customSet{
		container: make(map[string]struct{}),
	}
}

func (c *customSet) Exists(key string) bool {
	_, ok := c.container[key]
	return ok
}

func (c *customSet) Add(key string)  {
	c.container[key] = struct{}{}
}

func (c *customSet) Remove(key string) error{
	_, ok := c.container[key]
	if !ok {
		return fmt.Errorf("Remove Error: Item dosn't exist in set")
	}
	delete(c.container, key)
	return  nil
}

func (c *customSet) Size() int {
	return len(c.container)
}

func main() {
	customSet := makeSet()
	customSet.Add("A")
	customSet.Add("B")
	fmt.Printf("Size: %d\n", customSet.Size())
	fmt.Printf("A Exists?: %t\n", customSet.Exists("A"))
	fmt.Printf("B Exists?: %t\n", customSet.Exists("B"))
	customSet.Add("B")
	fmt.Println(customSet.container["A"])
}
