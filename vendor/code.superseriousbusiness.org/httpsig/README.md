# httpsig

**THIS IS A FORK OF https://github.com/go-fed/httpsig, WHICH WAS NO LONGER MAINTAINED. THANK YOU TO [cjslep](https://github.com/cjslep) FOR ALL YOUR HARD WORK!**

> HTTP Signatures made simple

`go get code.superseriousbusiness.org/httpsig`

Implementation of [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures).

Supports many different combinations of MAC, HMAC signing of hash, or RSA
signing of hash schemes. Its goals are:

* Have a very simple interface for signing and validating
* Support a variety of signing algorithms and combinations
* Support setting either headers (`Authorization` or `Signature`)
* Remaining flexible with headers included in the signing string
* Support both HTTP requests and responses
* Explicitly not support known-cryptographically weak algorithms
* Support automatic signing and validating Digest headers

## How to use

`import "code.superseriousbusiness.org/httpsig"`

### Signing

Signing a request or response requires creating a new `Signer` and using it:

```go
func sign(privateKey crypto.PrivateKey, pubKeyId string, r *http.Request) error {
	prefs := []httpsig.Algorithm{httpsig.RSA_SHA512, httpsig.RSA_SHA256}
	digestAlgorithm := DigestSha256
	// The "Date" and "Digest" headers must already be set on r, as well as r.URL.
	headersToSign := []string{httpsig.RequestTarget, "date", "digest"}
	signer, chosenAlgo, err := httpsig.NewSigner(prefs, digestAlgorithm, headersToSign, httpsig.Signature)
	if err != nil {
		return err
	}
	// To sign the digest, we need to give the signer a copy of the body...
	// ...but it is optional, no digest will be signed if given "nil"
	body := ...
	// If r were a http.ResponseWriter, call SignResponse instead.
	return signer.SignRequest(privateKey, pubKeyId, r, body)
}
```

`Signer`s are not safe for concurrent use by goroutines, so be sure to guard
access:

```go
type server struct {
	signer httpsig.Signer
	mu *sync.Mutex
}

func (s *server) handlerFunc(w http.ResponseWriter, r *http.Request) {
	privateKey := ...
	pubKeyId := ...
	// Set headers and such on w
	s.mu.Lock()
	defer s.mu.Unlock()
	// To sign the digest, we need to give the signer a copy of the response body...
	// ...but it is optional, no digest will be signed if given "nil"
	body := ...
	err := s.signer.SignResponse(privateKey, pubKeyId, w, body)
	if err != nil {
		...
	}
	...
}
```

The `pubKeyId` will be used at verification time.

### Verifying

Verifying requires an application to use the `pubKeyId` to both retrieve the key
needed for verification as well as determine the algorithm to use. Use a
`Verifier`:

```go
func verify(r *http.Request) error {
	verifier, err := httpsig.NewVerifier(r)
	if err != nil {
		return err
	}
	pubKeyId := verifier.KeyId()
	var algo httpsig.Algorithm = ...
	var pubKey crypto.PublicKey = ...
	// The verifier will verify the Digest in addition to the HTTP signature
	return verifier.Verify(pubKey, algo)
}
```

`Verifier`s are not safe for concurrent use by goroutines, but since they are
constructed on a per-request or per-response basis it should not be a common
restriction.

[License-Image]: https://img.shields.io/github/license/go-fed/httpsig?color=blue
[License-Url]: https://opensource.org/licenses/BSD-3-Clause
