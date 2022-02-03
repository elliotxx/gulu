package gitutil

import (
	"errors"
	"fmt"
	"os/exec"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

var (
	errMockCombinedOutputErr           = errors.New("mock CombinedOutput error")
	errMockGetRemoteURLErr             = errors.New("mock GetRemoteUrl error")
	errMockGetHeadHashErr              = errors.New("mock GetHeadHash error")
	errMockGetTagCommitShaErr          = errors.New("mock GetTagCommitSha error")
	errMockGetTagCommitShaFromLocalErr = errors.New("mock GetTagCommitShaFromLocal error")
)

func TestGetRemoteUrl(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		url, err := GetRemoteURL()
		assert.Nil(t, err)
		fmt.Println(url)
	})

	t.Run("fail", func(t *testing.T) {
		mockCombinedOutput(nil, errMockCombinedOutputErr)
		defer monkey.UnpatchAll()
		_, err := GetRemoteURL()
		assert.NotNil(t, err)
	})
}

func TestGetLatestTagFromRemote(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockGetRemoteURL("", errMockGetRemoteURLErr)
		defer monkey.UnpatchAll()
		_, err := GetLatestTagFromRemote()
		assert.NotNil(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockGetRemoteURL("", nil)
		mockCombinedOutput(nil, errMockCombinedOutputErr)
		defer monkey.UnpatchAll()
		_, err := GetLatestTagFromRemote()
		assert.NotNil(t, err)
	})
}

func TestGetLatestTagFromLocal(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockGetTagList([]string{"tag1", "tag2"}, nil)
		defer monkey.UnpatchAll()
		_, err := GetLatestTagFromLocal()
		assert.Nil(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockCombinedOutput(nil, ErrEmptyGitTag)
		defer monkey.UnpatchAll()
		_, err := GetLatestTagFromLocal()
		assert.NotNil(t, err)
	})
}

func TestGetHeadHash(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		_, err := GetHeadHash()
		assert.Nil(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockCombinedOutput(nil, errMockCombinedOutputErr)
		defer monkey.UnpatchAll()
		_, err := GetHeadHash()
		assert.NotNil(t, err)
	})
}

func TestGetHeadHashShort(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockGetHeadHash("", errMockGetHeadHashErr)
		defer monkey.UnpatchAll()
		_, err := GetHeadHashShort()
		assert.NotNil(t, err)
	})
}

func TestGetTagCommitSha(t *testing.T) {
	t.Run("wrong tag", func(t *testing.T) {
		_, err := GetTagCommitSha("")
		assert.NotNil(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockGetTagCommitShaFromLocal("", errMockGetTagCommitShaFromLocalErr)
		mockGetTagCommitShaFromRemote("remote sha", nil)
		defer monkey.UnpatchAll()
		_, err := GetTagCommitSha("tag")
		assert.Nil(t, err)
	})
}

func TestGetTagCommitShaFromLocal(t *testing.T) {
	t.Run("fail", func(t *testing.T) {
		mockCombinedOutput(nil, errMockCombinedOutputErr)
		defer monkey.UnpatchAll()
		_, err := GetTagCommitShaFromLocal("")
		assert.NotNil(t, err)
	})
}

func TestGetTagCommitShaFromRemote(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockGetRemoteURL("", errMockGetRemoteURLErr)
		defer monkey.UnpatchAll()
		_, err := GetTagCommitShaFromRemote("")
		assert.NotNil(t, err)
	})
}

func TestIsHeadAtTag(t *testing.T) {
	t.Run("empty tag", func(t *testing.T) {
		_, err := IsHeadAtTag("")
		assert.NotNil(t, err)
	})

	t.Run("fail-1", func(t *testing.T) {
		mockGetTagCommitSha("", errMockGetTagCommitShaErr)
		defer monkey.UnpatchAll()
		_, err := IsHeadAtTag("tag")
		assert.NotNil(t, err)
	})

	t.Run("fail-2", func(t *testing.T) {
		mockGetTagCommitSha("", nil)
		mockGetHeadHash("", errMockGetHeadHashErr)
		defer monkey.UnpatchAll()
		_, err := IsHeadAtTag("tag")
		assert.NotNil(t, err)
	})
}

func TestIsDirty(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		_, err := IsDirty()
		assert.Nil(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockCombinedOutput(nil, errMockCombinedOutputErr)
		defer monkey.UnpatchAll()
		_, err := IsDirty()
		assert.NotNil(t, err)
	})
}

func mockCombinedOutput(output []byte, err error) {
	monkey.Patch((*exec.Cmd).CombinedOutput, func(*exec.Cmd) ([]byte, error) {
		return output, err
	})
}

func mockGetRemoteURL(url string, err error) {
	monkey.Patch(GetRemoteURL, func() (string, error) {
		return url, err
	})
}

func mockGetTagList(tags []string, err error) {
	monkey.Patch(GetTagList, func() ([]string, error) {
		return tags, err
	})
}

func mockGetHeadHash(sha string, err error) {
	monkey.Patch(GetHeadHash, func() (string, error) {
		return sha, err
	})
}

func mockGetTagCommitShaFromLocal(sha string, err error) {
	monkey.Patch(GetTagCommitShaFromLocal, func(tag string) (string, error) {
		return sha, err
	})
}

func mockGetTagCommitShaFromRemote(sha string, err error) {
	monkey.Patch(GetTagCommitShaFromRemote, func(tag string) (string, error) {
		return sha, err
	})
}

func mockGetTagCommitSha(sha string, err error) {
	monkey.Patch(GetTagCommitSha, func(tag string) (string, error) {
		return sha, err
	})
}
