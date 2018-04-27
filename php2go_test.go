package php2go

import (
	"testing"
	"reflect"
	"bytes"
	"unicode/utf8"
	"os"
)

func TestTime(t *testing.T) {
	tdate := Date("02/01/2006 15:04:05 PM", 1524799394)
	equal(t, "27/04/2018 11:23:14 AM", tdate)

	tstrtotime, _ := Strtotime("02/01/2006 15:04:05", "02/01/2016 15:04:05")
	equal(t, int64(1451747045), tstrtotime)
}

func TestString(t *testing.T) {
	taddslashes := Addslashes("'wo'简体\"chousha")
	equal(t, `\'wo\'简体\"chousha`, taddslashes)

	ttimestamp := Time()
	gt(t, float64(ttimestamp), 0)

	tmd5 := Md5("123456")
	equal(t, "e10adc3949ba59abbe56e057f20f883e", tmd5)

	tsha1 := Sha1("123456")
	equal(t, "7c4a8d09ca3762af61e59520943dc26494f8941b", tsha1)

	tcrc32 := Crc32("123456")
	equal(t, uint32(158520161), tcrc32)

	tstrrepeat := StrRepeat("简体", 3)
	equal(t, "简体简体简体", tstrrepeat)

	tsubstr := Substr("abcdef", 0, 2)
	equal(t, "ab", tsubstr)

	tstrstr := Strstr("xxx@gmail.com", "@")
	equal(t, "@gmail.com", tstrstr)

	tucfirst := Ucfirst("kello world")
	equal(t, "Kello world", tucfirst)

	tstrpos := Strpos("hello world", "w", 0)
	equal(t, 6, tstrpos)

	timplode := Implode(",", []string{"a", "b", "c"})
	equal(t, "a,b,c", timplode)

	tAddslashes := Addslashes("f'oo b\"ar")
	equal(t, `f\'oo b\"ar`, tAddslashes)

	tStripslashes := Stripslashes("f\\'oo b\\\"ar\\\\a\\\\\\\\\\\\")
	equal(t, `f'oo b"ar\a\\\`, tStripslashes)

	tLevenshtein := Levenshtein("golang", "google", 1, 1, 1)
	equal(t, 4, tLevenshtein)

	var percent float64
	tSimilarText := SimilarText("golang", "google", &percent)
	equal(t, 3, tSimilarText)
	equal(t, float64(50), percent)

	tSoundex := Soundex("Heilbronn")
	equal(t, "H416", tSoundex)

	tuniqid := Uniqid("x")
	equal(t, 14, len(tuniqid))
	equal(t, true, bytes.HasPrefix([]byte(tuniqid), []byte("x")))

	tstrshuffle := StrShuffle("简˚abc")
	equal(t, 5, utf8.RuneCountInString(tstrshuffle))

	tord := Ord("简体")
	equal(t, 31616, tord)

	tchr := Chr(122)
	equal(t, "z", tchr)

	tmbstrlen := MbStrlen("简体 a")
	equal(t, 4, tmbstrlen)

	tnl2br := Nl2br("<a>\n\rxxx\nyy\r简体\r\nn\r\nx", true)
	equal(t, "<a><br />xxx<br />yy<br />简体<br />n<br />x", tnl2br)

	tstrrev := Strrev("abc \t nic %简体10.5()---")
	equal(t, "---)(5.01体简% cin 	 cba", tstrrev)

	tchunksplit := ChunkSplit("abc \t nic %简体10.5()---", 3, "e")
	equal(t, "abce 	 enice %简e体10e.5(e)--e-e", tchunksplit)

	tquotemeta := Quotemeta(".+?[$](*)^简体")
	equal(t, `\.\+\?\[\$\]\(\*\)\^简体`, tquotemeta)

	tHtmlentities := Htmlentities("<html>hello world </html>")
	equal(t, `&lt;html&gt;hello world &lt;/html&gt;`, tHtmlentities)

	tHtmlEntityDecode := HtmlEntityDecode("&lt;html&gt;hello world &lt;/html&gt;")
	equal(t, "<html>hello world </html>", tHtmlEntityDecode)

	tWordwrap := Wordwrap("abc hello world xxx", 5, "\n")
	equal(t, "abc\nhello\nworld\nxxx", tWordwrap)

	tStrWordCount := StrWordCount("a b c")
	equal(t, []string{"a", "b", "c"}, tStrWordCount)
}

func TestArray(t *testing.T) {
	var s1 = make([]interface{}, 3)
	s1[0] = "a"
	s1[1] = "b"
	s1[2] = "c"
	tArrayChunk := ArrayChunk(s1, 2)
	equal(t, [][]interface{}{{"a", "b"}, {"c"}}, tArrayChunk)

	var m1 = make(map[interface{}]interface{}, 3)
	m1[1] = "a"
	m1["a"] = "b"
	m1[2.5] = 1
	tArrayKeyExists := ArrayKeyExists("a", m1)
	equal(t, true, tArrayKeyExists)

	tArrayUnshift := ArrayUnshift(&s1, "x", "y")
	equal(t, 5, tArrayUnshift)
	equal(t, []interface{}{"x", "y", "a", "b", "c"}, s1)

	tArrayPush := ArrayPush(&s1, "u", "v")
	equal(t, 7, tArrayPush)
	equal(t, []interface{}{"x", "y", "a", "b", "c", "u", "v"}, s1)

	tArrayPop := ArrayPop(&s1)
	equal(t, "v", tArrayPop)
	equal(t, []interface{}{"x", "y", "a", "b", "c", "u"}, s1)

	tArrayShift := ArrayShift(&s1)
	equal(t, "x", tArrayShift)
	equal(t, []interface{}{"y", "a", "b", "c", "u"}, s1)

	tarrayfill := ArrayFill(-3, 6, "aaa")
	equal(t, map[int]interface{}{-1: "aaa", 0: "aaa", 1: "aaa", 2: "aaa", -3: "aaa", -2: "aaa"}, tarrayfill)

	tarrayrand := ArrayRand([]interface{}{"a", "b", "c"})
	equal(t, 3, len(tarrayrand))

	var s2 = make([]interface{}, 3)
	s2[0] = "a"
	s2[1] = "b"
	s2[2] = "c"
	tarraypad := ArrayPad(s2, -5, "d")
	equal(t, []interface{}{"d", "d", "a", "b", "c"}, tarraypad)

	var s3 = make([]interface{}, 3, 3)
	s3[0] = "x"
	s3[1] = "y"
	s3[2] = "z"
	tarraycombine := ArrayCombine(s2, s3)
	equal(t, map[interface{}]interface{}{"a": "x", "b": "y", "c": "z"}, tarraycombine)
}

func TestUrl(t *testing.T) {
	tParseUrl, _ := ParseUrl("http://username:password@hostname:9090/path?arg=value#anchor", -1)
	equal(t, map[string]string{"pass": "password", "path": "/path", "query": "arg=value", "fragment": "anchor", "scheme": "http", "host": "hostname", "port": "9090", "user": "username"}, tParseUrl)

	tUrlEncode := UrlEncode("http://golang.org?x y")
	equal(t, "http%3A%2F%2Fgolang.org%3Fx+y", tUrlEncode)

	tUrlDecode, _ := UrlDecode("http%3A%2F%2Fgolang.org%3Fx+y")
	equal(t, "http://golang.org?x y", tUrlDecode)

	tRawurlencode := Rawurlencode("http://golang.org?x y")
	equal(t, "http%3A%2F%2Fgolang.org%3Fx%20y", tRawurlencode)

	tRawurldecode, _ := Rawurldecode("http%3A%2F%2Fgolang.org%3Fx%20y")
	equal(t, "http://golang.org?x y", tRawurldecode)

	tBase64Encode := Base64Encode("This is an encoded string")
	equal(t, "VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw==", tBase64Encode)

	tBase64Decode, _ := Base64Decode("VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw==")
	equal(t, "This is an encoded string", tBase64Decode)
}

func TestMath(t *testing.T) {
	tMax := Max(2, 3.7, 5, 1.1)
	equal(t, float64(5), tMax)

	tMin := Min(2, 3.7, 5, 1.1)
	equal(t, 1.1, tMin)

	tRand := Rand(2, 5)
	rangeValue(t, float64(2), float64(5), float64(tRand))

	tDecbin := Decbin(100)
	equal(t, "1100100", tDecbin)

	tBindec, _ := Bindec(tDecbin)
	equal(t, "100", tBindec)

	tBin2hex, _ := Bin2hex(tDecbin)
	equal(t, "64", tBin2hex)

	tHexdec, _ := Hexdec(tBin2hex)
	equal(t, int64(100), tHexdec)

	tHex2bin, _ := Hex2bin(tBin2hex)
	equal(t, "1100100", tHex2bin)

	tDecoct := Decoct(tHexdec)
	equal(t, "144", tDecoct)

	tOctdec, _ := Octdec(tDecoct)
	equal(t, int64(100), tOctdec)

	tDechex := Dechex(tHexdec)
	equal(t, "64", tDechex)

	tBaseConvert, _ := BaseConvert("64", 16, 2)
	equal(t, "1100100", tBaseConvert)

	tNumberFormat := NumberFormat(1234567890.777, 2, ".", ",")
	equal(t, "1,234,567,890.78", tNumberFormat)

	tIsNumeric := IsNumeric("-0xaF")
	equal(t, true, tIsNumeric)
}

func TestFile(t *testing.T) {
	trealpath1, _ := Realpath("/home/go/../go/test/../")
	equal(t, "/home/go", trealpath1)

	tbasename := Basename("/home/go/src/pkg/php2go.go")
	equal(t, "php2go.go", tbasename)

	tPathinfo := Pathinfo("/home/go/php2go.go.go", -1)
	equal(t, map[string]string{"dirname": "/home/go", "basename": "php2go.go.go", "extension": "go", "filename": "php2go.go"}, tPathinfo)

	tDiskFreeSpace, _ := DiskFreeSpace("/")
	gt(t, float64(tDiskFreeSpace), 0)

	tDiskTotalSpace, _ := DiskTotalSpace("/")
	gte(t, float64(tDiskTotalSpace), 0)

	wd, _ := os.Getwd()
	tfilesize, _ := FileSize(wd)
	gt(t, float64(tfilesize), 0)
}

func TestOther(t *testing.T) {
	tVersionCompare := VersionCompare("1.3-beta", "1.4Rc1", "<")
	equal(t, true, tVersionCompare)

	tMemoryGetUsage := MemoryGetUsage(true)
	gt(t, float64(tMemoryGetUsage), 0)
}

// Expected to be equal.
func equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
	}
}

// Expected to be unequal.
func unequal(t *testing.T, expected, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
	}
}

// Expect a greater than b.
func gt(t *testing.T, a, b float64) {
	if a <= b {
		t.Errorf("Expected %v (type %v) > Got %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

// Expect a greater than or equal to b.
func gte(t *testing.T, a, b float64) {
	if a < b {
		t.Errorf("Expected %v (type %v) > Got %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

// Expected value needs to be within range.
func rangeValue(t *testing.T, min, max, actual float64) {
	if actual < min || actual > max {
		t.Errorf("Expected range of %v-%v (type %v) > Got %v (type %v)", min, max, reflect.TypeOf(min), actual, reflect.TypeOf(actual))
	}
}
