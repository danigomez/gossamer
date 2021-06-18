// Copyright 2021 ChainSafe Systems (ON) Corp.
// This file is part of gossamer.
//
// The gossamer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gossamer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gossamer library. If not, see <http://www.gnu.org/licenses/>.

package modules

import "net/http"

// AccountModule is an RPC module for aliasing System Module
type AccountModule struct {
	sysModule *SystemModule
}

// NewAccountModule creates a new API instance
func NewAccountModule(sm interface{}) *AccountModule {
	return &AccountModule{
		sysModule: sm.(*SystemModule),
	}
}

// NextIndex alias for AccountNextIndex of SystemModule
func (am *AccountModule) NextIndex(r *http.Request, req *StringRequest, res *U64Response) error {
	return am.sysModule.AccountNextIndex(r, req, res)
}