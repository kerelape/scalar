/*
 * MIT License
 *
 * Copyright (c) 2023 goctus
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package scalar_test

import (
	"context"
	"testing"

	scalar "github.com/goctus/scalar/pkg"
	"github.com/stretchr/testify/assert"
)

func TestConstant(t *testing.T) {
	t.Run("returns the same value", func(t *testing.T) {
		value, err := scalar.NewConstant("Hello, World!").Value(context.TODO())
		if err != nil {
			t.Errorf("Failed to get the value: %v", err)
		}
		assert.Equal(
			t,
			"Hello, World!",
			value,
			"Value of a Constant must be equal to the value passed into it.",
		)
	})
}
