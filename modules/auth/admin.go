// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/Unknwon/macaron"

	"github.com/macaron-contrib/binding"
)

type AdminCrateUserForm struct {
	LoginType string `binding:"Required"`
	LoginName string
	UserName  string `binding:"Required;AlphaDashDot;MaxSize(35)"`
	Email     string `binding:"Required;Email;MaxSize(254)"`
	Password  string `binding:"MaxSize(255)"`
}

func (f *AdminCrateUserForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}

type AdminEditUserForm struct {
	LoginType    string `binding:"Required"`
	LoginName    string
	FullName     string `binding:"MaxSize(100)"`
	Email        string `binding:"Required;Email;MaxSize(254)"`
	Password     string `binding:"MaxSize(255)"`
	Website      string `binding:"MaxSize(50)"`
	Location     string `binding:"MaxSize(50)"`
	Active       bool
	Admin        bool
	AllowGitHook bool
}

func (f *AdminEditUserForm) Validate(ctx *macaron.Context, errs binding.Errors) binding.Errors {
	return validate(errs, ctx.Data, f, ctx.Locale)
}
