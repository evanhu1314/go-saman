// Copyright 2018 The go-saman Authors
// This file is part of the go-saman library.
//
// The go-saman library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-saman library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-saman library. If not, see <http://www.gnu.org/licenses/>.

package mem

import (
	"testing"

	"github.com/saman-org/go-saman/swarm/storage/mock/test"
)

// TestGlobalStore is running test for a GlobalStore
// using test.MockStore function.
func TestGlobalStore(t *testing.T) {
	test.MockStore(t, NewGlobalStore(), 100)
}

// TestImportExport is running tests for importing and
// exporting data between two GlobalStores
// using test.ImportExport function.
func TestImportExport(t *testing.T) {
	test.ImportExport(t, NewGlobalStore(), NewGlobalStore(), 100)
}
