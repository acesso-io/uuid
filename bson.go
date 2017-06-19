// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import "gopkg.in/mgo.v2/bson"



//GetBSON return a bson field from a raw data
func (u UUID) GetBSON() (interface{}, error) {
	return bson.Binary{Kind: 0x04, Data: u[:]}, nil
}

//SetBSON set a bson field from a raw data
func (u *UUID) SetBSON(raw bson.Raw) error {
	doc := bson.Binary{}
	err := raw.Unmarshal(&doc)

	if err != nil {
		return err
	}

	copy(u[:], doc.Data[:])

	return nil
}

//SetBSONBynary set a bson field from a binary data
func (u *UUID) SetBSONBynary(bin bson.Binary) error {
	data := bin.Data
	copy(u[:], data[:])
	return nil
}