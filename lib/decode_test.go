package lib

import (
	"io/ioutil"
	"testing"
)

type args struct {
	tokenToBeDecoded string
	hmacSecret       string
	publicKeyFile    string
}

func TestParseToken(t *testing.T) {

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"HMAC successful decode - no signature",
			args {
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMiwiZXhwIjoxNjE2MjM5MDIyfQ.NLixjdr52QBIppYv4IM7CLXbdSJex1LDbhWGDYmD-YM",
				"",
				"",
			},
			false,
		},
		{
			"HMAC invalid signature",
			args {
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMiwiZXhwIjoxNjE2MjM5MDIyfQ.SGKQueLyr4DDqmxb4B9AOz-pLDPMMRMclzX41_ao3AE",
				"",
				"",
			},
			true,
		},
		{
			"HMAC successful decode - valid signature",
			args {
				"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMiwiZXhwIjoxNjE2MjM5MDIyfQ.SGKQueLyr4DDqmxb4B9AOz-pLDPMMRMclzX41_ao3AE",
				"foo",
				"",
			},
			false,
		},
		{
			"RSA failed decode - public key file not provided",
			args {
				"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.WeRGFpXdRflCn3wo5BpIWMgAPHK67DTbEgJ1621nxURE2RUnbvUdK_nrX9R7p_RvifD3vC66q6MxS_qDzttef5vdaiSEmP5lJFeYsr_ieVHL88jL-VmEIEeMbqJbqa3D7NOw7HwuJRSPOMdWefEEs3fm1TGABbcASX1v2q45YjU",
				"",
				"",
			},
			true,
		},
		{
			"RSA failed decode - public key file does not exist",
			args {
				"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.WeRGFpXdRflCn3wo5BpIWMgAPHK67DTbEgJ1621nxURE2RUnbvUdK_nrX9R7p_RvifD3vC66q6MxS_qDzttef5vdaiSEmP5lJFeYsr_ieVHL88jL-VmEIEeMbqJbqa3D7NOw7HwuJRSPOMdWefEEs3fm1TGABbcASX1v2q45YjU",
				"",
				"random_dir/random.key",
			},
			true,
		},
		{
			"RSA successful decode",
			args {
				"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.WeRGFpXdRflCn3wo5BpIWMgAPHK67DTbEgJ1621nxURE2RUnbvUdK_nrX9R7p_RvifD3vC66q6MxS_qDzttef5vdaiSEmP5lJFeYsr_ieVHL88jL-VmEIEeMbqJbqa3D7NOw7HwuJRSPOMdWefEEs3fm1TGABbcASX1v2q45YjU",
				"",
				"../test_data/key.pub",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ParseToken(tt.args.tokenToBeDecoded, tt.args.hmacSecret, tt.args.publicKeyFile, ioutil.Discard); (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
