// Copyright 2018 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tplimpl

import (
	"github.com/gohugoio/hugo/common/herrors"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

type templateInfo struct {
	template string

	// Used to create some error context in error situations
	fs afero.Fs

	// The filename relative to the fs above.
	filename string

	// The real filename (if possible). Used for logging.
	realFilename string
}

func (info templateInfo) errWithFileContext(what string, err error) error {
	err = errors.Wrapf(err, "file %q: %s:", info.realFilename, what)

	err, _ = herrors.WithFileContextForFile(
		err,
		info.filename,
		info.fs,
		"go-html-template",
		herrors.SimpleLineMatcher)

	return err
}
