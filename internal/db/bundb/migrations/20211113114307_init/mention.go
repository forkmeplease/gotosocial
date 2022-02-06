/*
   GoToSocial
   Copyright (C) 2021-2022 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package gtsmodel

import "time"

// Mention refers to the 'tagging' or 'mention' of a user within a status.
type Mention struct {
	ID               string    `validate:"required,ulid" bun:"type:CHAR(26),pk,nullzero,notnull,unique"`        // id of this item in the database
	CreatedAt        time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"` // when was item created
	UpdatedAt        time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"` // when was item last updated
	StatusID         string    `validate:"required,ulid" bun:"type:CHAR(26),nullzero,notnull"`                  // ID of the status this mention originates from
	Status           *Status   `validate:"-" bun:"rel:belongs-to"`                                              // status referred to by statusID
	OriginAccountID  string    `validate:"required,ulid" bun:"type:CHAR(26),nullzero,notnull"`                  // ID of the mention creator account
	OriginAccountURI string    `validate:"url" bun:",nullzero,notnull"`                                         // ActivityPub URI of the originator/creator of the mention
	OriginAccount    *Account  `validate:"-" bun:"rel:belongs-to"`                                              // account referred to by originAccountID
	TargetAccountID  string    `validate:"required,ulid" bun:"type:CHAR(26),nullzero,notnull"`                  // Mention target/receiver account ID
	TargetAccount    *Account  `validate:"-" bun:"rel:belongs-to"`                                              // account referred to by targetAccountID
	Silent           bool      `validate:"-" bun:",notnull,default:false"`                                      // Prevent this mention from generating a notification?

	/*
		NON-DATABASE CONVENIENCE FIELDS
		These fields are just for convenience while passing the mention
		around internally, to make fewer database calls and whatnot. They're
		not meant to be put in the database!
	*/

	// NameString is for putting in the namestring of the mentioned user
	// before the mention is dereferenced. Should be in a form along the lines of:
	// @whatever_username@example.org
	//
	// This will not be put in the database, it's just for convenience.
	NameString string `validate:"-" bun:"-"`
	// TargetAccountURI is the AP ID (uri) of the user mentioned.
	//
	// This will not be put in the database, it's just for convenience.
	TargetAccountURI string `validate:"-" bun:"-"`
	// TargetAccountURL is the web url of the user mentioned.
	//
	// This will not be put in the database, it's just for convenience.
	TargetAccountURL string `validate:"-" bun:"-"`
	// A pointer to the gtsmodel account of the mentioned account.
}