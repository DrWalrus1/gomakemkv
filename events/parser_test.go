package events_test

import (
	"testing"

	"github.com/DrWalrus1/gomakemkv/events"
	"github.com/go-playground/assert/v2"
)

func TestEmptyStringFailsParsing(t *testing.T) {
	actual, err := events.Parse("")

	assert.Equal(t, events.ErrEmptyInput, err)
	assert.Equal(t, nil, actual)
}

func TestStringSanitisation(t *testing.T) {
	actual, err := events.Parse(" SINFO:0,6,31,6121,<b>Track information</b><br>")
	assert.Equal(t, nil, err)
	assert.Equal(t, actual.(*events.StreamInformation).Value, "Track information")
}

func TestCurrentProgressTitleOutputParser(t *testing.T) {
	t.Run("Successful Parse", func(t *testing.T) {
		expected := events.CurrentProgressTitleOutput{
			Code: "1",
			ID:   "1",
			Name: "Test",
		}

		input := "PRGC:1,1,Test"

		actual, err := events.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails when missing prefix", func(t *testing.T) {
		input := "1,1,Test"

		actual, err := events.Parse(input)

		assert.Equal(t, events.ErrPrefixNotFound, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Fails when missing values", func(t *testing.T) {
		input := "PRGC:1,1"

		actual, err := events.Parse(input)

		assert.Equal(t, events.ErrNotEnoughValues, err)
		assert.Equal(t, nil, actual)
	})
}

func TestParseDiscInformationOutput(t *testing.T) {
	t.Run("Successful Parse", func(t *testing.T) {
		expected := events.DiscInformationOutputMessage{
			TitleCount: 1,
		}

		input := "TCOUNT:1"

		actual, err := events.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails when missing prefix", func(t *testing.T) {
		input := "1"

		actual, err := events.Parse(input)

		assert.Equal(t, events.ErrPrefixNotFound, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Fails when missing values", func(t *testing.T) {
		input := "TCOUT:"
		actual, err := events.Parse(input)

		assert.NotEqual(t, nil, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Fails when invalid value is provided", func(t *testing.T) {
		input := "TCOUT:THISISNOTANUMBER"

		actual, err := events.Parse(input)

		assert.NotEqual(t, nil, err)
		assert.Equal(t, nil, actual)
	})
}

func TestParseDiscInfo(t *testing.T) {
	t.Run("Successful Parse", func(t *testing.T) {
		expected := events.DiscInformation{
			ID:            1,
			MessageCodeId: 1,
			Value:         "Value",
		}

		input := "CINFO:1,1,Value"

		actual, err := events.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails when missing prefix", func(t *testing.T) {
		input := "1,CODE,Value"

		actual, err := events.Parse(input)

		assert.Equal(t, events.ErrPrefixNotFound, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Fails when missing values", func(t *testing.T) {
		input := "CINFO:1,CODE"

		actual, err := events.Parse(input)

		assert.Equal(t, events.ErrNotEnoughValues, err)
		assert.Equal(t, nil, actual)
	})
}

func TestDriveScanMessageParser(t *testing.T) {
	t.Run("Parse Successfully", func(t *testing.T) {
		expected := events.DriveScanMessage{
			DriveIndex: "1",
			Visible:    "Drive is empty and open",
			Enabled:    "Drive is empty and open",
			Flags:      "Flags",
			DriveName:  "Drive1",
			DiscName:   "Disc1",
			DeviceName: "/dev/sr0",
		}

		input := "DRV:1,1,1,Flags,Drive1,Disc1,/dev/sr0"

		actual, err := events.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails when missing prefix", func(t *testing.T) {
		input := "1,true,true,Flags,Drive1,Disc1,/dev/sr0"

		actual, err := events.Parse(input)

		assert.Equal(t, events.ErrPrefixNotFound, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Fails when there are not enough values", func(t *testing.T) {
		input := "DRV:1,true,true,Flags,Drive1,/dev/sr0"

		actual, err := events.Parse(input)

		assert.Equal(t, events.ErrNotEnoughValues, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Parses when drives are not visible or enabled", func(t *testing.T) {
		input := "DRV:1,255,999,Flags,Drive1,Disc1,/dev/sr0"

		expected := events.DriveScanMessage{
			DriveIndex: "1",
			Visible:    "Unknown",
			Enabled:    "Unknown",
			Flags:      "Flags",
			DriveName:  "Drive1",
			DiscName:   "Disc1",
			DeviceName: "/dev/sr0",
		}
		actual, err := events.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)

	})
}

func TestParseMessageOutput(t *testing.T) {
	t.Run("Successfully parse", func(t *testing.T) {
		t.Run("No params", func(t *testing.T) {
			expected := events.MessageOutput{
				Code:           "1",
				Flags:          "test",
				ParameterCount: 0,
				RawMessage:     "Test Message",
				FormatMessage:  "Test Message",
			}

			input := "MSG:1,test,0,Test Message,Test Message"

			actual, err := events.Parse(input)

			assert.Equal(t, nil, err)
			assert.Equal(t, expected, actual)

			expected2 := events.MessageOutput{
				Code:           "1011",
				Flags:          "0",
				ParameterCount: 1,
				RawMessage:     "Using LibreDrive mode (v06.3 id=0FA242DD4D0B)",
				FormatMessage:  "%1",
				MessageParams:  []string{"Using LibreDrive mode (v06.3 id=0FA242DD4D0B)"},
			}
			input2 := `MSG:1011,0,1,"Using LibreDrive mode (v06.3 id=0FA242DD4D0B)","%1","Using LibreDrive mode (v06.3 id=0FA242DD4D0B)"`

			actual2, err2 := events.Parse(input2)

			assert.Equal(t, nil, err2)
			assert.Equal(t, expected2, actual2)
		})

		t.Run("One param", func(t *testing.T) {
			expected := events.MessageOutput{
				Code:           "1",
				Flags:          "test",
				ParameterCount: 1,
				RawMessage:     "Test Message",
				FormatMessage:  "Test Message",
				MessageParams:  []string{"Test"},
			}

			input := "MSG:1,test,1,Test Message,Test Message,Test"

			actual, err := events.Parse(input)

			assert.Equal(t, nil, err)
			assert.Equal(t, expected, actual)
		})

		t.Run("Three params", func(t *testing.T) {
			expected := events.MessageOutput{
				Code:           "1",
				Flags:          "test",
				ParameterCount: 3,
				RawMessage:     "Test Message",
				FormatMessage:  "Test Message",
				MessageParams:  []string{"Test1", "Test2", "Test3"},
			}

			input := "MSG:1,test,3,Test Message,Test Message,Test1,Test2,Test3"

			actual, err := events.Parse(input)

			assert.Equal(t, nil, err)
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("Fails when missing prefix", func(t *testing.T) {
		input := "1,test,3,Test Message,Test Message,Test1,Test2,Test3"

		actual, err := events.Parse(input)

		assert.Equal(t, events.ErrPrefixNotFound, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Fails when parameter count mismatch", func(t *testing.T) {
		t.Run("Param count is less than actual", func(t *testing.T) {
			expected := events.MessageOutput{
				Code:           "1",
				Flags:          "test",
				ParameterCount: 1,
				RawMessage:     "Test Message",
				FormatMessage:  "Test Message",
				MessageParams:  []string{"Test1"},
			}

			input := "MSG:1,test,1,Test Message,Test Message,Test1,Test2,Test3"

			actual, err := events.Parse(input)

			assert.Equal(t, nil, err)
			assert.Equal(t, expected, actual)
		})

		t.Run("Param count is greater than actual", func(t *testing.T) {
			expected := events.MessageOutput{
				Code:           "1",
				Flags:          "test",
				ParameterCount: 3,
				RawMessage:     "Test Message",
				FormatMessage:  "Test Message",
			}

			input := "MSG:1,test,3,Test Message,Test Message,Test1"

			actual, err := events.Parse(input)

			assert.Equal(t, nil, err)
			assert.Equal(t, expected, actual)
		})
	})
}

func TestParseProgressBarOutput(t *testing.T) {
	t.Run("Successful parse", func(t *testing.T) {
		expected := events.ProgressBarOutput{
			CurrentProgress: "1",
			TotalProgress:   "100",
			MaxProgress:     "200",
		}
		input := "PRGV:1,100,200"

		actual, err := events.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails to parse", func(t *testing.T) {
		t.Run("Missing prefix", func(t *testing.T) {
			input := "1,100,200"

			actual, err := events.Parse(input)

			assert.Equal(t, events.ErrPrefixNotFound, err)
			assert.Equal(t, nil, actual)
		})

		t.Run("Progress bar missing values", func(t *testing.T) {
			input := "PRGV:1,100"

			actual, err := events.Parse(input)

			assert.Equal(t, events.ErrNotEnoughValues, err)
			assert.Equal(t, nil, actual)
		})
	})
}

func TestParseStreamInfo(t *testing.T) {
	t.Run("Successful parse", func(t *testing.T) {
		expected := events.StreamInformation{
			TitleIndex:    1,
			StreamIndex:   1,
			AttributeId:   1,
			MessageCodeId: 1,
			Value:         "Value",
		}

		input := "SINFO:1,1,1,1,Value"

		actual, err := events.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails to parse", func(t *testing.T) {
		t.Run("Missing prefix", func(t *testing.T) {
			input := "1,CODE,Value"

			actual, err := events.Parse(input)

			assert.Equal(t, events.ErrPrefixNotFound, err)
			assert.Equal(t, nil, actual)
		})

		t.Run("Missing values", func(t *testing.T) {
			input := "SINFO:1,CODE"

			actual, err := events.Parse(input)

			assert.Equal(t, events.ErrNotEnoughValues, err)
			assert.Equal(t, nil, actual)
		})
	})
}

func TestParseTitleInfo(t *testing.T) {
	t.Run("Successful parse", func(t *testing.T) {
		expected := events.TitleInformation{
			TitleIndex:    1,
			AttributeId:   1,
			MessageCodeId: 1,
			Value:         "Value",
		}

		input := "TINFO:1,1,1,Value"

		actual, err := events.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails to parse", func(t *testing.T) {
		input := "1,CODE,Value"

		actual, err := events.Parse(input)

		assert.NotEqual(t, nil, err)
		assert.Equal(t, nil, actual)
	})

	t.Run("Missing values", func(t *testing.T) {
		input := "TINFO:1,CODE"

		actual, err := events.Parse(input)

		assert.Equal(t, events.ErrNotEnoughValues, err)
		assert.Equal(t, nil, actual)
	})
}

func TestParseTotalTitleOutput(t *testing.T) {
	t.Run("Successful parse", func(t *testing.T) {
		expected := events.TotalProgressTitleOutput{
			Code: "1",
			ID:   "1",
			Name: "Test",
		}

		input := "PRGT:1,1,Test"

		actual, err := events.Parse(input)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Fails to parse", func(t *testing.T) {
		t.Run("Missing prefix", func(t *testing.T) {
			input := "1,1,Test"

			actual, err := events.Parse(input)

			assert.Equal(t, events.ErrPrefixNotFound, err)
			assert.Equal(t, nil, actual)
		})

		t.Run("Missing values", func(t *testing.T) {
			input := "PRGT:1,1"

			actual, err := events.Parse(input)

			assert.Equal(t, events.ErrNotEnoughValues, err)
			assert.Equal(t, nil, actual)
		})
	})
}
