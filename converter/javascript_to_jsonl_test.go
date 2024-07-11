package converter

import (
	"bufio"
	"bytes"
	"testing"
)

func TestJavascriptToJSONL(t *testing.T) {
	i := bytes.NewBufferString(`window.YTD.tweet.part0 = [
		{
			"tweet": {
				"id_str": "1234567",
				"entities": {
					"hashtags": [ ]
				}
			}
		},
		{
			"tweet": {
				"id_str": "78901234",
				"entities": {
					"hashtags": [ ]
				}
			}
		}
]`)

	s := bufio.NewScanner(i)
	b := bytes.NewBuffer([]byte(""))
	w := bufio.NewWriter(b)

	JavascriptToJSONL(s, w)

	r := string(b.Bytes())
	if r != `{"tweet": {"id_str": "1234567","entities": {"hashtags": [ ]}}}
{"tweet": {"id_str": "78901234","entities": {"hashtags": [ ]}}}` {
		t.Errorf("Unexpected result:\n%s\n", r)
	}
}

func TestJavascriptToJSONLCurlyBracket(t *testing.T) {
	i := bytes.NewBufferString(`window.YTD.tweet.part0 = [
		{
			"tweet": {
				"id_str": "1234{567",
				"entities": {
					"hashtags": [ ]
				}
			}
		},
		{
			"tweet": {
				"id_str": "78901234",
				"entities": {
					"ha{{sht}ags}": [ ]
				}
			}
		}
]`)

	s := bufio.NewScanner(i)
	b := bytes.NewBuffer([]byte(""))
	w := bufio.NewWriter(b)

	JavascriptToJSONL(s, w)

	r := string(b.Bytes())
	if r != `{"tweet": {"id_str": "1234{567","entities": {"hashtags": [ ]}}}
{"tweet": {"id_str": "78901234","entities": {"ha{{sht}ags}": [ ]}}}` {
		t.Errorf("Unexpected result:\n%s\n", r)
	}
}
