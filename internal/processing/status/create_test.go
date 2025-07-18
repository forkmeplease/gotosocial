// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package status_test

import (
	"net/http"
	"testing"

	apimodel "code.superseriousbusiness.org/gotosocial/internal/api/model"
	"code.superseriousbusiness.org/gotosocial/internal/config"
	"code.superseriousbusiness.org/gotosocial/internal/db"
	"code.superseriousbusiness.org/gotosocial/internal/gtsmodel"
	"code.superseriousbusiness.org/gotosocial/internal/util"
	"github.com/stretchr/testify/suite"
)

type StatusCreateTestSuite struct {
	StatusStandardTestSuite
}

func (suite *StatusCreateTestSuite) TestProcessContentWarningWithQuotationMarks() {
	ctx := suite.T().Context()

	creatingAccount := suite.testAccounts["local_account_1"]
	creatingApplication := suite.testApplications["application_1"]

	statusCreateForm := &apimodel.StatusCreateRequest{
		Status:      "poopoo peepee",
		MediaIDs:    []string{},
		Poll:        nil,
		InReplyToID: "",
		Sensitive:   false,
		SpoilerText: "\"test\"", // these should not be html-escaped when the final text is rendered
		Visibility:  apimodel.VisibilityPublic,
		LocalOnly:   util.Ptr(false),
		ScheduledAt: nil,
		Language:    "en",
		ContentType: apimodel.StatusContentTypePlain,
	}

	apiStatus, err := suite.status.Create(ctx, creatingAccount, creatingApplication, statusCreateForm)
	suite.NoError(err)
	suite.NotNil(apiStatus)

	suite.Equal("\"test\"", apiStatus.SpoilerText)
}

func (suite *StatusCreateTestSuite) TestProcessStatusMarkdownWithUnderscoreEmoji() {
	ctx := suite.T().Context()

	// update the shortcode of the rainbow emoji to surround it in underscores
	if err := suite.db.UpdateWhere(ctx, []db.Where{{Key: "shortcode", Value: "rainbow"}}, "shortcode", "_rainbow_", &gtsmodel.Emoji{}); err != nil {
		suite.FailNow(err.Error())
	}

	creatingAccount := suite.testAccounts["local_account_1"]
	creatingApplication := suite.testApplications["application_1"]

	statusCreateForm := &apimodel.StatusCreateRequest{
		Status:      "poopoo peepee :_rainbow_:",
		MediaIDs:    []string{},
		Poll:        nil,
		InReplyToID: "",
		Sensitive:   false,
		Visibility:  apimodel.VisibilityPublic,
		LocalOnly:   util.Ptr(false),
		ScheduledAt: nil,
		Language:    "en",
		ContentType: apimodel.StatusContentTypeMarkdown,
	}

	apiStatus, err := suite.status.Create(ctx, creatingAccount, creatingApplication, statusCreateForm)
	suite.NoError(err)
	suite.NotNil(apiStatus)

	suite.Equal("<p>poopoo peepee :_rainbow_:</p>", apiStatus.Content)
	suite.NotEmpty(apiStatus.Emojis)
}

func (suite *StatusCreateTestSuite) TestProcessStatusMarkdownWithSpoilerTextEmoji() {
	ctx := suite.T().Context()
	creatingAccount := suite.testAccounts["local_account_1"]
	creatingApplication := suite.testApplications["application_1"]

	statusCreateForm := &apimodel.StatusCreateRequest{
		Status:      "poopoo peepee",
		SpoilerText: "testing something :rainbow:",
		MediaIDs:    []string{},
		Poll:        nil,
		InReplyToID: "",
		Sensitive:   false,
		Visibility:  apimodel.VisibilityPublic,
		LocalOnly:   util.Ptr(false),
		ScheduledAt: nil,
		Language:    "en",
		ContentType: apimodel.StatusContentTypeMarkdown,
	}

	apiStatus, err := suite.status.Create(ctx, creatingAccount, creatingApplication, statusCreateForm)
	suite.NoError(err)
	suite.NotNil(apiStatus)

	suite.Equal("<p>poopoo peepee</p>", apiStatus.Content)
	suite.Equal("testing something :rainbow:", apiStatus.SpoilerText)
	suite.NotEmpty(apiStatus.Emojis)
}

func (suite *StatusCreateTestSuite) TestProcessMediaDescriptionTooShort() {
	ctx := suite.T().Context()

	config.SetMediaDescriptionMinChars(100)

	creatingAccount := suite.testAccounts["local_account_1"]
	creatingApplication := suite.testApplications["application_1"]

	statusCreateForm := &apimodel.StatusCreateRequest{
		Status:      "poopoo peepee",
		MediaIDs:    []string{suite.testAttachments["local_account_1_unattached_1"].ID},
		Poll:        nil,
		InReplyToID: "",
		Sensitive:   false,
		SpoilerText: "",
		Visibility:  apimodel.VisibilityPublic,
		LocalOnly:   util.Ptr(false),
		ScheduledAt: nil,
		Language:    "en",
		ContentType: apimodel.StatusContentTypePlain,
	}

	apiStatus, err := suite.status.Create(ctx, creatingAccount, creatingApplication, statusCreateForm)
	suite.EqualError(err, "media description less than min chars (100)")
	suite.Nil(apiStatus)
}

func (suite *StatusCreateTestSuite) TestProcessLanguageWithScriptPart() {
	ctx := suite.T().Context()

	creatingAccount := suite.testAccounts["local_account_1"]
	creatingApplication := suite.testApplications["application_1"]

	statusCreateForm := &apimodel.StatusCreateRequest{
		Status:      "你好世界", // hello world
		MediaIDs:    []string{},
		Poll:        nil,
		InReplyToID: "",
		Sensitive:   false,
		SpoilerText: "",
		Visibility:  apimodel.VisibilityPublic,
		LocalOnly:   util.Ptr(false),
		ScheduledAt: nil,
		Language:    "zh-Hans",
		ContentType: apimodel.StatusContentTypePlain,
	}

	apiStatus, err := suite.status.Create(ctx, creatingAccount, creatingApplication, statusCreateForm)
	suite.NoError(err)
	suite.NotNil(apiStatus)

	suite.Equal("zh-Hans", *apiStatus.Language)
}

func (suite *StatusCreateTestSuite) TestProcessReplyToUnthreadedRemoteStatus() {
	ctx := suite.T().Context()

	creatingAccount := suite.testAccounts["local_account_1"]
	creatingApplication := suite.testApplications["application_1"]
	inReplyTo := suite.testStatuses["remote_account_1_status_1"]

	// Reply to a remote status that
	// doesn't have a threadID set on it.
	statusCreateForm := &apimodel.StatusCreateRequest{
		Status:      "boobies",
		MediaIDs:    []string{},
		Poll:        nil,
		InReplyToID: inReplyTo.ID,
		Sensitive:   false,
		SpoilerText: "this is a reply",
		Visibility:  apimodel.VisibilityPublic,
		LocalOnly:   util.Ptr(false),
		ScheduledAt: nil,
		Language:    "en",
		ContentType: apimodel.StatusContentTypePlain,
	}

	apiStatus, err := suite.status.Create(ctx, creatingAccount, creatingApplication, statusCreateForm)
	suite.NoError(err)
	suite.NotNil(apiStatus)

	// ThreadID should be set on the status,
	// even though the replied-to status does
	// not have a threadID.
	dbStatus, dbErr := suite.state.DB.GetStatusByID(ctx, apiStatus.ID)
	if dbErr != nil {
		suite.FailNow(err.Error())
	}
	suite.NotEmpty(dbStatus.ThreadID)
}

func (suite *StatusCreateTestSuite) TestProcessNoContentTypeUsesDefault() {
	ctx := suite.T().Context()
	creatingAccount := suite.testAccounts["local_account_1"]
	creatingApplication := suite.testApplications["application_1"]

	statusCreateForm := &apimodel.StatusCreateRequest{
		Status:      "poopoo peepee",
		SpoilerText: "",
		MediaIDs:    []string{},
		Poll:        nil,
		InReplyToID: "",
		Sensitive:   false,
		Visibility:  apimodel.VisibilityPublic,
		LocalOnly:   util.Ptr(false),
		ScheduledAt: nil,
		Language:    "en",
		ContentType: "",
	}

	apiStatus, errWithCode := suite.status.Create(ctx, creatingAccount, creatingApplication, statusCreateForm)
	suite.NoError(errWithCode)
	suite.NotNil(apiStatus)

	suite.Equal("<p>poopoo peepee</p>", apiStatus.Content)

	// the test accounts don't have settings, so we're comparing to
	// the global default value instead of the requester's default
	suite.Equal(apimodel.StatusContentTypeDefault, apiStatus.ContentType)
}

func (suite *StatusCreateTestSuite) TestProcessInvalidVisibility() {
	ctx := suite.T().Context()
	creatingAccount := suite.testAccounts["local_account_1"]
	creatingApplication := suite.testApplications["application_1"]

	statusCreateForm := &apimodel.StatusCreateRequest{
		Status:      "my tests content is boring",
		SpoilerText: "",
		MediaIDs:    []string{},
		Poll:        nil,
		InReplyToID: "",
		Sensitive:   false,
		Visibility:  "local",
		LocalOnly:   util.Ptr(false),
		ScheduledAt: nil,
		Language:    "en",
		ContentType: apimodel.StatusContentTypePlain,
	}

	apiStatus, errWithCode := suite.status.Create(ctx, creatingAccount, creatingApplication, statusCreateForm)
	suite.Nil(apiStatus)
	suite.Equal(http.StatusUnprocessableEntity, errWithCode.Code())
	suite.Equal("Unprocessable Entity: processVisibility: invalid visibility", errWithCode.Safe())
}

func TestStatusCreateTestSuite(t *testing.T) {
	suite.Run(t, new(StatusCreateTestSuite))
}
