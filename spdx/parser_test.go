package spdx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Run("invalid operator", func(t *testing.T) {
		_, err := Parse("0BSD INV AAL")
		assert.Error(t, err)
		assert.IsType(t, InvalidOperatorError{}, err)
		assert.Equal(t, 5, err.(InvalidOperatorError).index)
	})

	t.Run("invalid license", func(t *testing.T) {
		_, err := Parse("0BSD AND inv")
		assert.Error(t, err)
		assert.IsType(t, InvalidLicenseError{}, err)
		assert.Equal(t, 9, err.(InvalidLicenseError).index)
	})

	t.Run("single-license", func(t *testing.T) {
		exp, err := Parse("0BSD")
		assert.NoError(t, err)
		assert.Equal(t, LicenseID("0BSD"), exp)
	})

	t.Run("single-license-and-later", func(t *testing.T) {
		exp, err := Parse("MPL-2.0+")
		assert.NoError(t, err)
		assert.Equal(t, LicenseID("MPL-2.0+"), exp)
	})

	t.Run("simple-expression", func(t *testing.T) {
		exp, err := Parse("0BSD OR AAL AND Abstyles")
		assert.NoError(t, err)
		require.IsType(t, Or{}, exp)
		require.IsType(t, And{}, exp.(Or).Right)
		assert.Equal(t, LicenseID("0BSD"), exp.(Or).Left)
		assert.Equal(t, LicenseID("AAL"), exp.(Or).Right.(And).Left)
		assert.Equal(t, LicenseID("Abstyles"), exp.(Or).Right.(And).Right)
	})

	t.Run("complex-expression", func(t *testing.T) {
		exp, err := Parse("0BSD OR (AAL OR Abstyles) AND Adobe-2006 OR Adobe-Glyph AND (ADSL OR AFL-1.1 AND (AFL-1.2 OR AFL-2.0))")
		/*
			0BSD OR ((AAL OR Abstyles) AND Adobe-2006) OR (Adobe-Glyph AND (ADSL OR AFL-1.1 AND (AFL-1.2 OR AFL-2.0)))
			                       [OR]
			        +---------------+---------------+
			       [OR]                           [AND]
			  +-----+-----+                 +-------+-----+
			 0BSD       [AND]          Adobe-Glyph       [OR]
			      +-------+-------+                 +-----+-----+
			     [OR]         Adobe-2006           ADSL       [AND]
			  +---+---+                                    +----+-----+
			 AAL   Abstyles                             AFL-1.1      [OR]
			                                                     +----+----+
			                                                  AFL-1.2   AFL-2.0
		*/
		assert.NoError(t, err)
		require.IsType(t, Or{}, exp)
		require.IsType(t, Or{}, exp.(Or).Left)
		require.IsType(t, And{}, exp.(Or).Left.(Or).Right)
		require.IsType(t, Or{}, exp.(Or).Left.(Or).Right.(And).Left)
		require.IsType(t, And{}, exp.(Or).Right)
		require.IsType(t, Or{}, exp.(Or).Right.(And).Right)
		require.IsType(t, And{}, exp.(Or).Right.(And).Right.(Or).Right)
		require.IsType(t, Or{}, exp.(Or).Right.(And).Right.(Or).Right.(And).Right)

		assert.Equal(t, LicenseID("0BSD"), exp.(Or).Left.(Or).Left)
		assert.Equal(t, LicenseID("AAL"), exp.(Or).Left.(Or).Right.(And).Left.(Or).Left)
		assert.Equal(t, LicenseID("Abstyles"), exp.(Or).Left.(Or).Right.(And).Left.(Or).Right)
		assert.Equal(t, LicenseID("Adobe-2006"), exp.(Or).Left.(Or).Right.(And).Right)
		assert.Equal(t, LicenseID("Adobe-Glyph"), exp.(Or).Right.(And).Left)
		assert.Equal(t, LicenseID("ADSL"), exp.(Or).Right.(And).Right.(Or).Left)
		assert.Equal(t, LicenseID("AFL-1.1"), exp.(Or).Right.(And).Right.(Or).Right.(And).Left)
		assert.Equal(t, LicenseID("AFL-1.2"), exp.(Or).Right.(And).Right.(Or).Right.(And).Right.(Or).Left)
		assert.Equal(t, LicenseID("AFL-2.0"), exp.(Or).Right.(And).Right.(Or).Right.(And).Right.(Or).Right)
	})
}
