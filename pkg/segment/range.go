// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package segment

// T is a required type parameter that must be an integral type.
type T uint64

// A Range represents a contiguous range of T.
type Range struct {
	// Start is the inclusive start of the range.
	Start T

	// End is the exclusive end of the range.
	End T
}

// WellFormed returns true if r.Start <= r.End. All other methods on a Range
// require that the Range is well-formed.
func (r Range) WellFormed() bool {
	return r.Start <= r.End
}

// Length returns the length of the range.
func (r Range) Length() T {
	return r.End - r.Start
}

// Contains returns true if r contains x.
func (r Range) Contains(x T) bool {
	return r.Start <= x && x < r.End
}

// Overlaps returns true if r and r2 overlap.
func (r Range) Overlaps(r2 Range) bool {
	return r.Start < r2.End && r2.Start < r.End
}

// IsSupersetOf returns true if r is a superset of r2; that is, the range r2 is
// contained within r.
func (r Range) IsSupersetOf(r2 Range) bool {
	return r.Start <= r2.Start && r.End >= r2.End
}

// Intersect returns a range consisting of the intersection between r and r2.
// If r and r2 do not overlap, Intersect returns a range with unspecified
// bounds, but for which Length() == 0.
func (r Range) Intersect(r2 Range) Range {
	if r.Start < r2.Start {
		r.Start = r2.Start
	}
	if r.End > r2.End {
		r.End = r2.End
	}
	if r.End < r.Start {
		r.End = r.Start
	}
	return r
}

// CanSplitAt returns true if it is legal to split a segment spanning the range
// r at x; that is, splitting at x would produce two ranges, both of which have
// non-zero length.
func (r Range) CanSplitAt(x T) bool {
	return r.Contains(x) && r.Start < x
}
