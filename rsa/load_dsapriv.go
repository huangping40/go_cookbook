package main

import (
	"crypto/dsa"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"
)

type dsaPrivateKey struct {
	Version       int
	P, Q, G, Y, X *big.Int
}

var pemBytes = []byte(`-----BEGIN DSA PRIVATE KEY-----
MIIBugIBAAKBgQDCMR4dk2tDY+BBm2fMfzG4WGV/ILCKNMuhsl9yNbsvbet+HG4v
C+gSXC4zp7nxYnO8KljmwkuO8NEfu1yEOEL3+LHn8hfUhUj6lSFszzuxtQKTBYYm
vvlt8982UIaXTqWiRHCVkhRCbGYYYmSNJyd9HYUHppJhy0t642/Bwg4RVwIVALtw
MOntygZZjDtPDr2Iag0inpXzAoGAAjoQITU4sbGwK0fPpgGBQPkws6xm0SOuVPNW
Wzk82i9REBC+pyevT3ZZkx7efTRlriEcvsctRISAiD5KMifw0og1GhukX9L9Jgfz
6GX1ILaVPdtjn3jX5YCFrum9tXnWu9gSpLqbr75LIPtYMBixwiy7Vx1j7+zEXuVz
AQYGLNkCgYBNyslF5i0/ZhugcEVk3BoA1ZapJHNgyMbsZqrpTF3zhJHSMPAvHZX0
ZBhFGmpPGprt/Rcq1CpqlWY4GxnqxbRoFzL/2i5jsOzZKvMu3zcN9/Z+hmBTttjq
SPNGn/PVmdsyxPxeopgjPSKzIGrTwIhUTN3GFsfF4kxc7afNe+MiHwIUdyOjmbsb
7CXCsUNw8/TcGfJFamc=
-----END DSA PRI234523452345234VATE KEY-----
`)

func main() {
	block, _ := pem.Decode(pemBytes)

	var rawPriv dsaPrivateKey
	rest, err := asn1.Unmarshal(block.Bytes, &rawPriv)
	if len(rest) != 0 {
		panic("trailing garbage")
	}
	if err != nil {
		panic(err)
	}

	priv := &dsa.PrivateKey{
		PublicKey: dsa.PublicKey{
			Parameters: dsa.Parameters{
				P: rawPriv.P,
				Q: rawPriv.Q,
				G: rawPriv.G,
			},
			Y: rawPriv.Y,
		},
		X: rawPriv.X,
	}

	fmt.Printf("%#v\n", priv)

	checkY := new(big.Int).Exp(priv.G, priv.X, priv.P)
	if checkY.Cmp(priv.Y) != 0 {
		panic("invalid key")
	}
}
