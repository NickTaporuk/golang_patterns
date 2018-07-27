/*
 * Copyright (c) 2018. Created by Nikolay Kuropatkin
 */

package prototype

import "testing"

func TestClone(t *testing.T) {

	shirtCache := GetShirtsCloner()

	if shirtCache == nil {
		t.Fatal("Received cache was nil")
	}

	_, err := shirtCache.GetClone(White)

	if err != nil {
		t.Error(err)
	}
}
