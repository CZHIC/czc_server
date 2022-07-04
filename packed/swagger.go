package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GDIE7oaxIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwOj4ffzCT3BjvmXDST2vg93aXNWD/U2y/vJKbOTif3xjpnLr/blhGc42YYYHFl2sMuMRdtgxY0H15paEkqcfhtvnPyXO2X564y5JW9139va76+7/3my6vnPNgpyhy2OybTIRURzOux14Mw1EKhqPabUEcg50STh9qK3G00ai2J9smuPSH5V3nNT6NrPWVp+BoqyWxIsuqM4pHf/5ssWNfA042yaZPAqn8OtUbk6orxWted5xavw5Vn/Tyf5Lbqa47Qv9dhSNtu/36bb3T636/zm9c8r0/Vle6+8Lv/1+PLpxXoalmc2ekiIRkz5V60dyd3kdYlLSUdAvVOygXNh+b6nf/e++eMgX6+xs6jIZ0HpnZbitWvfmR3nfVx9ur+gOSDpneC/0A9TdzddrH7MX54Qvr9M8/Giyye+Xi85+I/bKHbzpxC/u0/CkgIOyOUmqKS3pWhNT9DzPBrEkDTxtWVi+rbCyoop/O03n/utFWv55Rr+tm5V9eqFVj8n/rdZW80f5VjAvr5N+OanoqPn1BZyTr1h419epvq3a8GauI851/48Cthrfy1/3+vpzVFV14J3/32kuMJR/+/mxdvql9fzc9269vpmb+Rx5jiNrOAj3OpXf999kzIl7ktpekUwm6LlwkLPtp+h0Vn/uRkY/v8P8GbnaOgtZ7/MyMDAz8rAAEsdDAwFaKmDHZE6wAnC9Pv5BJBuZDUB3oxMIsyI1IVsMih1wcC2RhCJN60hjMLuFAgQYPjvKMfEgMVhrGwgeSYGJoZOBgaGGCYQDxAAAP//NHQHPfsCAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
