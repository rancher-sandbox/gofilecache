// Copyright 2017 The Go Authors, SUSE LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gofilecache

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestHash(t *testing.T) {
	oldSalt := hashSalt
	hashSalt = nil
	defer func() {
		hashSalt = oldSalt
	}()

	h := NewHash("alice")
	h.Write([]byte("hello world"))
	sum := fmt.Sprintf("%x", h.Sum())
	want := "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f"
	if sum != want {
		t.Errorf("hash(hello world) = %v, want %v", sum, want)
	}
}

func TestHashFile(t *testing.T) {
	f, err := ioutil.TempFile("", "cmd-go-test-")
	if err != nil {
		t.Fatal(err)
	}
	name := f.Name()
	fmt.Fprintf(f, "hello world")
	defer os.Remove(name)
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	var h ActionID // make sure hash result is assignable to ActionID
	h, err = FileHash(name)
	if err != nil {
		t.Fatal(err)
	}
	sum := fmt.Sprintf("%x", h)
	want := "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f"
	if sum != want {
		t.Errorf("hash(hello world) = %v, want %v", sum, want)
	}
}
