package shortcuts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAndResolve(t *testing.T) {
	Add("adam", "drew")
	assert.Equal(t, "drew", Resolve("adam"))
}

func TestAddAndDelete(t *testing.T) {
	Add("adam", "drew")
	Delete("adam")
	assert.Equal(t, "adam", Resolve("adam"))
}

func TestGetAll(t *testing.T) {
	Add("adam", "drew")
	Add("ziggy", "dog")
	Add("jonesy", "cat")
	want := map[string]string{
		"adam":   "drew",
		"ziggy":  "dog",
		"jonesy": "cat",
	}
	got := GetAll()
	assert.True(t, fmt.Sprint(want) == fmt.Sprint(got))
}
