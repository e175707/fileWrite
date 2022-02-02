package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("sakamotoyuuya")
    want := "Hi, sakamotoyuuya. Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}