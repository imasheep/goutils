// aliyun.go
package aliyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	g "goutils"
	"goutils/web"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"
)

type PreSortedFormat struct {
	StrSlice []string
}

type AliPubArgs struct {
	Format           string
	Version          string
	AccessKeyId      string
	SignatureMethod  string
	Timestamp        string
	SignatureVersion string
	SignatureNonce   string
	SignatureKey     string
}

func DoAliRequest(aliDomain string, pubArgs AliPubArgs, partArgs interface{}) (res string) {

	url := CreateRequestUrl(aliDomain, pubArgs, partArgs)
	_, res, _ = web.SimGet(url)
	return

}

func CreateRequestUrl(aliDomain string, pubArgs AliPubArgs, partArgs interface{}) string {

	var preSortSlice PreSortedFormat

	preSortSlice.Append(partArgs)
	preSortSlice.Append(pubArgs)

	sort.Strings(preSortSlice.StrSlice)

	var sortstr string
	for i := 0; i < len(preSortSlice.StrSlice); i++ {
		if sortstr == "" {
			sortstr = preSortSlice.StrSlice[i]
		} else {
			sortstr = sortstr + "&" + preSortSlice.StrSlice[i]
		}
	}

	sortstr = strings.Replace(sortstr, "+", "%20", 65535)
	signature := CreateSignature("GET", sortstr, pubArgs.SignatureKey)
	requestUrl := "https://" + aliDomain + "/?" + sortstr + "&Signature=" + signature

	return requestUrl

}

func CreateSignature(method string, stringToSign string, signaturekey string) string {

	seckey := []byte(signaturekey)
	mac := hmac.New(sha1.New, seckey)
	mac.Write([]byte(method + "&%2F&" + url.QueryEscape(stringToSign)))
	signature := url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	return signature
}

func (this *AliPubArgs) Init(apiVersion string,
	accessKeyId string,
	signatureMethod string,
	signatureVersion string,
	signatureKey string) {

	this.Format = "json"
	this.Version = apiVersion
	this.AccessKeyId = accessKeyId
	this.SignatureMethod = signatureMethod
	this.Timestamp = CreateAliTimeStampUtc()
	this.SignatureVersion = signatureVersion
	this.SignatureNonce = g.CreateRandom("string", 32)
	this.SignatureKey = signatureKey

}

func (Prestr *PreSortedFormat) Append(argStruct interface{}) {

	switch reflect.TypeOf(argStruct).Kind() {
	case reflect.Struct:
		k := reflect.TypeOf(argStruct)
		v := reflect.ValueOf(argStruct)
		for i := 0; i < k.NumField(); i++ {
			v_r := fmt.Sprintf("%v", v.Field(i).Interface())
			Prestr.StrSlice = append(Prestr.StrSlice, fmt.Sprintf("%s=%v", k.Field(i).Name, url.QueryEscape(v_r)))
		}
	}

}
func CreateAliTimeStampUtc() string {

	return time.Now().UTC().Format("2006-01-02T15:04:05Z")

}
