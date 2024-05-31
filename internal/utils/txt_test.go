//nolint:lll
package utils_test

import (
	"strings"
	"testing"

	"github.com/germanbrew/terraform-provider-hetznerdns/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestPlainToTXTRecordValue(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "empty",
			input:  "",
			output: "",
		},
		{
			name:   "small string",
			input:  "test",
			output: "test",
		},
		{
			name:   "small string with quotes",
			input:  `te"st`,
			output: `te"st`,
		},
		{
			name:   "small string with spaces",
			input:  `v=STSv1; id=20230523103000Z`,
			output: `v=STSv1; id=20230523103000Z`,
		},
		{
			name:   "large string",
			input:  strings.Repeat("test", 100),
			output: `"testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttes" "ttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest" `,
		},
		{
			name:   "large string with quotes",
			input:  strings.Repeat(`te"st`, 100),
			output: `"te\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"st" "te\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"st" `,
		},
		{
			name:   "large string with quotes",
			input:  strings.Repeat(`t e s t`, 100),
			output: `"t e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e" " s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s " "tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s t" `,
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.output, utils.PlainToTXTRecordValue(tc.input))
		})
	}
}

func TestTXTRecordToPlainValue(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "empty",
			input:  "",
			output: "",
		},
		{
			name:   "small string",
			input:  "test",
			output: "test",
		},
		{
			name:   "small string with quotes",
			input:  `te"st`,
			output: `te"st`,
		},
		{
			name:   "small string with spaces",
			input:  `v=STSv1; id=20230523103000Z`,
			output: `v=STSv1; id=20230523103000Z`,
		},
		{
			name:   "large string",
			input:  `"testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttes" "ttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest" `,
			output: strings.Repeat("test", 100),
		},
		{
			name:   "large string with quotes",
			input:  `"te\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"st" "te\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"stte\"st" `,
			output: strings.Repeat(`te"st`, 100),
		},
		{
			name:   "large string with spaces",
			input:  `"t e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s t" "t e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s tt e s t" `,
			output: strings.Repeat(`t e s t`, 100),
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.output, utils.TXTRecordToPlainValue(tc.input))
		})
	}
}

func TestPlainToTXTRecordToPlainValue(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name  string
		value string
	}{
		{
			name:  "empty",
			value: "",
		},
		{
			name:  "small string",
			value: "test",
		},
		{
			name:  "small string with quotes",
			value: `te"st`,
		},
		{
			name:  "small string with spaces",
			value: `v=STSv1; id=20230523103000Z`,
		},
		{
			name:  "large string",
			value: strings.Repeat("test", 100),
		},
		{
			name:  "large string with quotes",
			value: strings.Repeat(`te"st`, 100),
		},
		{
			name:  "large string with spaces",
			value: strings.Repeat("t e s t", 100),
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.value, utils.TXTRecordToPlainValue(utils.TXTRecordToPlainValue(tc.value)))
		})
	}
}
