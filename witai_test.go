// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import "testing"

func TestNewClient(t *testing.T) {
	c := NewClient("token")
	if c == nil {
		t.Fatalf("client is nil")
	}
	if c.Version != DefaultVersion {
		t.Fatalf("client default version is not set")
	}
}

func TestNewClientWithVersion(t *testing.T) {
	c := NewClientWithVersion("token", "v2")
	if c == nil {
		t.Fatalf("client is nil")
	}
	if c.Version != "v2" {
		t.Fatalf("client v2 version is not set")
	}
}
