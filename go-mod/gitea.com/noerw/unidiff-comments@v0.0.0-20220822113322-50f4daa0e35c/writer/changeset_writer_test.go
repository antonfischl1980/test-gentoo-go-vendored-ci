package writer

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"gitea.com/noerw/unidiff-comments"
)

func TestWillOmitDanglingWhitespaceOnRender(t *testing.T) {
	changeset := unidiff.Changeset{
		Diffs: []*unidiff.Diff{
			{
				FileComments: unidiff.CommentsTree{
					{
						Text: "\n\nevery dangling whitespace    \n\n" +
							"   should be removed\n\n\n",
					},
				},
			},
		},
	}

	buf := bytes.Buffer{}
	err := WriteChangeset(changeset, &buf)
	if err != nil {
		t.Fatal(err)
	}

	output, err := ioutil.ReadFile(
		"../_test/no_dangling_whitespaces.diff")

	if string(output) != buf.String() {
		t.Log("all dangling whitespaces should be trimmed")
		t.Fatal("\n" + makeDiff(buf.String(), string(output)))
	}
}

func makeDiff(actual, expected string) string {
	a, _ := ioutil.TempFile(os.TempDir(), "actual")
	defer func() {
		os.Remove(a.Name())
	}()
	b, _ := ioutil.TempFile(os.TempDir(), "expected")
	defer func() {
		os.Remove(b.Name())
	}()

	a.WriteString(actual)
	b.WriteString(expected)
	cmd := exec.Command("diff", "-u", b.Name(), a.Name())
	buf := bytes.NewBuffer([]byte{})
	cmd.Stdout = buf
	cmd.Run()

	return buf.String()
}
