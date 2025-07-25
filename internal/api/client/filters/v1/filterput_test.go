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

package v1_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"

	filtersV1 "code.superseriousbusiness.org/gotosocial/internal/api/client/filters/v1"
	apimodel "code.superseriousbusiness.org/gotosocial/internal/api/model"
	"code.superseriousbusiness.org/gotosocial/internal/config"
	"code.superseriousbusiness.org/gotosocial/internal/gtserror"
	"code.superseriousbusiness.org/gotosocial/internal/oauth"
	"code.superseriousbusiness.org/gotosocial/internal/stream"
	"code.superseriousbusiness.org/gotosocial/testrig"
)

func (suite *FiltersTestSuite) putFilter(
	filterKeywordID string,
	phrase *string,
	context *[]string,
	irreversible *bool,
	wholeWord *bool,
	expiresIn *int,
	expiresInStr *string,
	requestJson *string,
	expectedHTTPStatus int,
	expectedBody string,
) (*apimodel.FilterV1, error) {
	// instantiate recorder + test context
	recorder := httptest.NewRecorder()
	ctx, _ := testrig.CreateGinTestContext(recorder, nil)
	ctx.Set(oauth.SessionAuthorizedAccount, suite.testAccounts["local_account_1"])
	ctx.Set(oauth.SessionAuthorizedToken, oauth.DBTokenToToken(suite.testTokens["local_account_1"]))
	ctx.Set(oauth.SessionAuthorizedApplication, suite.testApplications["application_1"])
	ctx.Set(oauth.SessionAuthorizedUser, suite.testUsers["local_account_1"])

	// create the request
	ctx.Request = httptest.NewRequest(http.MethodPut, config.GetProtocol()+"://"+config.GetHost()+"/api/"+filtersV1.BasePath+"/"+filterKeywordID, nil)
	ctx.Request.Header.Set("accept", "application/json")
	if requestJson != nil {
		ctx.Request.Header.Set("content-type", "application/json")
		ctx.Request.Body = io.NopCloser(strings.NewReader(*requestJson))
	} else {
		ctx.Request.Form = make(url.Values)
		if phrase != nil {
			ctx.Request.Form["phrase"] = []string{*phrase}
		}
		if context != nil {
			ctx.Request.Form["context[]"] = *context
		}
		if irreversible != nil {
			ctx.Request.Form["irreversible"] = []string{strconv.FormatBool(*irreversible)}
		}
		if wholeWord != nil {
			ctx.Request.Form["whole_word"] = []string{strconv.FormatBool(*wholeWord)}
		}
		if expiresIn != nil {
			ctx.Request.Form["expires_in"] = []string{strconv.Itoa(*expiresIn)}
		} else if expiresInStr != nil {
			ctx.Request.Form["expires_in"] = []string{*expiresInStr}
		}
	}

	ctx.AddParam("id", filterKeywordID)

	// trigger the handler
	suite.filtersModule.FilterPUTHandler(ctx)

	// read the response
	result := recorder.Result()
	defer result.Body.Close()

	b, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	errs := gtserror.NewMultiError(2)

	// check code + body
	if resultCode := recorder.Code; expectedHTTPStatus != resultCode {
		errs.Appendf("expected %d got %d", expectedHTTPStatus, resultCode)
		if expectedBody == "" {
			return nil, errs.Combine()
		}
	}

	// if we got an expected body, return early
	if expectedBody != "" {
		if string(b) != expectedBody {
			errs.Appendf("expected %s got %s", expectedBody, string(b))
		}
		return nil, errs.Combine()
	}

	resp := &apimodel.FilterV1{}
	if err := json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (suite *FiltersTestSuite) TestPutFilterFull() {
	homeStream := suite.openHomeStream(suite.testAccounts["local_account_1"])

	id := suite.testFilterKeywords["local_account_1_filter_1_keyword_1"].ID
	phrase := "GNU/Linux"
	context := []string{"home", "public"}
	irreversible := false
	wholeWord := true
	expiresIn := 86400
	filter, err := suite.putFilter(id, &phrase, &context, &irreversible, &wholeWord, &expiresIn, nil, nil, http.StatusOK, "")
	if err != nil {
		suite.FailNow(err.Error())
	}

	suite.Equal(phrase, filter.Phrase)
	filterContext := make([]string, 0, len(filter.Context))
	for _, c := range filter.Context {
		filterContext = append(filterContext, string(c))
	}
	suite.ElementsMatch(context, filterContext)
	suite.Equal(irreversible, filter.Irreversible)
	suite.Equal(wholeWord, filter.WholeWord)
	if suite.NotNil(filter.ExpiresAt) {
		suite.NotEmpty(*filter.ExpiresAt)
	}

	suite.checkStreamed(homeStream, true, "", stream.EventTypeFiltersChanged)
}

func (suite *FiltersTestSuite) TestPutFilterFullJSON() {
	homeStream := suite.openHomeStream(suite.testAccounts["local_account_1"])

	id := suite.testFilterKeywords["local_account_1_filter_1_keyword_1"].ID
	// Use a numeric literal with a fractional part to test the JSON-specific handling for non-integer "expires_in".
	requestJson := `{
		"phrase":"GNU/Linux",
		"context": ["home", "public"],
		"irreversible": false,
		"whole_word": true,
		"expires_in": 86400.1
	}`
	filter, err := suite.putFilter(id, nil, nil, nil, nil, nil, nil, &requestJson, http.StatusOK, "")
	if err != nil {
		suite.FailNow(err.Error())
	}

	suite.Equal("GNU/Linux", filter.Phrase)
	suite.ElementsMatch(
		[]apimodel.FilterContext{
			apimodel.FilterContextHome,
			apimodel.FilterContextPublic,
		},
		filter.Context,
	)
	suite.Equal(false, filter.Irreversible)
	suite.Equal(true, filter.WholeWord)
	if suite.NotNil(filter.ExpiresAt) {
		suite.NotEmpty(*filter.ExpiresAt)
	}

	suite.checkStreamed(homeStream, true, "", stream.EventTypeFiltersChanged)
}

func (suite *FiltersTestSuite) TestPutFilterMinimal() {
	homeStream := suite.openHomeStream(suite.testAccounts["local_account_1"])

	id := suite.testFilterKeywords["local_account_1_filter_1_keyword_1"].ID
	phrase := "GNU/Linux"
	context := []string{"home"}
	filter, err := suite.putFilter(id, &phrase, &context, nil, nil, nil, nil, nil, http.StatusOK, "")
	if err != nil {
		suite.FailNow(err.Error())
	}

	suite.Equal(phrase, filter.Phrase)
	filterContext := make([]string, 0, len(filter.Context))
	for _, c := range filter.Context {
		filterContext = append(filterContext, string(c))
	}
	suite.ElementsMatch(context, filterContext)
	suite.False(filter.Irreversible)
	suite.False(filter.WholeWord)
	suite.Nil(filter.ExpiresAt)

	suite.checkStreamed(homeStream, true, "", stream.EventTypeFiltersChanged)
}

func (suite *FiltersTestSuite) TestPutFilterEmptyPhrase() {
	id := suite.testFilterKeywords["local_account_1_filter_1_keyword_1"].ID
	phrase := ""
	context := []string{"home"}
	_, err := suite.putFilter(id, &phrase, &context, nil, nil, nil, nil, nil, http.StatusUnprocessableEntity, "")
	if err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *FiltersTestSuite) TestPutFilterMissingPhrase() {
	id := suite.testFilterKeywords["local_account_1_filter_1_keyword_1"].ID
	context := []string{"home"}
	_, err := suite.putFilter(id, nil, &context, nil, nil, nil, nil, nil, http.StatusUnprocessableEntity, "")
	if err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *FiltersTestSuite) TestPutFilterEmptyContext() {
	id := suite.testFilterKeywords["local_account_1_filter_1_keyword_1"].ID
	phrase := "GNU/Linux"
	context := []string{}
	_, err := suite.putFilter(id, &phrase, &context, nil, nil, nil, nil, nil, http.StatusUnprocessableEntity, "")
	if err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *FiltersTestSuite) TestPutFilterMissingContext() {
	id := suite.testFilterKeywords["local_account_1_filter_1_keyword_1"].ID
	phrase := "GNU/Linux"
	_, err := suite.putFilter(id, &phrase, nil, nil, nil, nil, nil, nil, http.StatusUnprocessableEntity, "")
	if err != nil {
		suite.FailNow(err.Error())
	}
}

// There should be a filter with this phrase as its title in our test fixtures. Changing ours to that title should fail.
func (suite *FiltersTestSuite) TestPutFilterTitleConflict() {
	id := suite.testFilterKeywords["local_account_1_filter_1_keyword_1"].ID
	phrase := "metasyntactic variables"
	_, err := suite.putFilter(id, &phrase, nil, nil, nil, nil, nil, nil, http.StatusUnprocessableEntity, "")
	if err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *FiltersTestSuite) TestPutAnotherAccountsFilter() {
	id := suite.testFilterKeywords["local_account_2_filter_1_keyword_1"].ID
	phrase := "GNU/Linux"
	context := []string{"home"}
	_, err := suite.putFilter(id, &phrase, &context, nil, nil, nil, nil, nil, http.StatusNotFound, `{"error":"Not Found: filter not found"}`)
	if err != nil {
		suite.FailNow(err.Error())
	}
}

func (suite *FiltersTestSuite) TestPutNonexistentFilter() {
	id := "not_even_a_real_ULID"
	phrase := "GNU/Linux"
	context := []string{"home"}
	_, err := suite.putFilter(id, &phrase, &context, nil, nil, nil, nil, nil, http.StatusNotFound, `{"error":"Not Found: filter keyword not found"}`)
	if err != nil {
		suite.FailNow(err.Error())
	}
}

// setFilterExpiration sets filter expiration.
func (suite *FiltersTestSuite) setFilterExpiration(id string, phrase *string, expiresIn *int, expiresInStr *string, requestJson *string) *apimodel.FilterV1 {
	context := []string{"home"}
	filter, err := suite.putFilter(id, phrase, &context, nil, nil, expiresIn, expiresInStr, requestJson, http.StatusOK, "")
	if err != nil {
		suite.FailNow(err.Error())
	}
	return filter
}

// Regression test for https://codeberg.org/superseriousbusiness/gotosocial/issues/3497
func (suite *FiltersTestSuite) TestPutFilterUnsetExpirationDateEmptyString() {
	filterKeyword := suite.testFilterKeywords["local_account_1_filter_1_keyword_1"]
	id := filterKeyword.ID
	phrase := filterKeyword.Keyword

	// Setup: set an expiration date for the filter.
	expiresIn := 86400
	filter := suite.setFilterExpiration(id, &phrase, &expiresIn, nil, nil)
	if !suite.NotNil(filter.ExpiresAt) {
		suite.FailNow("Test precondition failed")
	}

	// Unset the filter's expiration date by setting it to an empty string.
	expiresInStr := ""
	filter = suite.setFilterExpiration(id, &phrase, nil, &expiresInStr, nil)
	suite.Nil(filter.ExpiresAt)
}

// Regression test related to https://codeberg.org/superseriousbusiness/gotosocial/issues/3497
func (suite *FiltersTestSuite) TestPutFilterUnsetExpirationDateNullJSON() {
	filterKeyword := suite.testFilterKeywords["local_account_1_filter_1_keyword_1"]
	id := filterKeyword.ID
	phrase := filterKeyword.Keyword

	// Setup: set an expiration date for the filter.
	expiresIn := 86400
	filter := suite.setFilterExpiration(id, &phrase, &expiresIn, nil, nil)
	if !suite.NotNil(filter.ExpiresAt) {
		suite.FailNow("Test precondition failed")
	}

	// Unset the filter's expiration date by setting it to a null literal.
	requestJson := `{
		"phrase": "fnord",
		"context": ["home"],
		"expires_in": null
	}`
	filter = suite.setFilterExpiration(id, nil, nil, nil, &requestJson)
	suite.Nil(filter.ExpiresAt)
}
