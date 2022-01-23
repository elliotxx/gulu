//go:build ignore
// +build ignore

package main

import (
	"fmt"

	. "github.com/elliotxx/gulu/gitutil"
)

func main() {
	latestTag, err := GetLatestTag()
	fmt.Println("GetLatestTag:", latestTag, err)

	tags, err := GetTagList()
	fmt.Println("GetTagList:", tags, err)

	head, err := GetHeadHash()
	fmt.Println("GetHeadHash:", head, err)

	for i, tag := range tags {
		fmt.Println("\t" + tag)
		sha, err := GetTagCommitSha(tag)
		fmt.Println("\tGetTagCommitSha:", i, sha, err)
		isHead, err := IsHeadAtTag(tag)
		fmt.Println("\tIsHeadAtTag:", i, isHead, err)
	}

	dirty, err := IsDirty()
	fmt.Println("IsDirty:", dirty, err)
}
