package main

import "fmt"

type folder struct {
	children []inode
	name string
}

func (f *folder) print(indentation string)  {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		child.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, child := range f.children {
		copy := child.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
