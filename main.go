package main

import (
	"fmt"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func main() {
	repository, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}

	fromHash := plumbing.NewHash("3ad2882a848f4e5cced1088e5d94f9aa16a0efbb")
	fromCommitObj, err := repository.CommitObject(fromHash)
	if err != nil {
		panic(err)
	}

	toHash := plumbing.NewHash("e19b1fc9bddaa78986d1f09e4f24bb5f6e2d2444")
	toCommitObj, err := repository.CommitObject(toHash)
	if err != nil {
		panic(err)
	}

	originPatch, err := fromCommitObj.Patch(toCommitObj)
	if err != nil {
		panic(err)
	}

	// diff.DefaultContextLines = 0
	//_ = originPatch.String()
	fmt.Printf("%s", originPatch.String())
}
