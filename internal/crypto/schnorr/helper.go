package schnorr

import (
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
)


var (
	_P = btcec.S256().Params().P // Prime modulus of secp256k1
)


type JacobianPoint struct {
	X, Y, Z *btcec.FieldVal
}
// Point represents a point in affine coordinates on the secp256k1 curve
type Point struct {
	x, y *big.Int // Coordinates in affine form
}

func mulmod(a, b, modulus *big.Int) *big.Int {
	result := new(big.Int).Mul(a, b)
	return new(big.Int).Mod(result, modulus)
}

// Helper function to perform modular addition with big.Int.
func addmod(a, b, modulus *big.Int) *big.Int {
	result := new(big.Int).Add(a, b)
	return new(big.Int).Mod(result, modulus)
}

func invMod(x *big.Int) *big.Int {
	t := new(big.Int)
	q := new(big.Int)
	newT := big.NewInt(1)
	r := new(big.Int).Set(_P)

	for x.Sign() != 0 {
		q.Div(r, x)

		tmp := new(big.Int).Set(t)
		t.Set(newT)
		newT = addmod(tmp, new(big.Int).Sub(_P, mulmod(q, newT, _P)), _P) // addmod(tmp, sub(_P, mulmod(q, newT, _P)), _P)

		tmp.Set(r)
		r.Set(x)
		x.Sub(tmp, new(big.Int).Mul(q, x))
	}

	return t
}
func (jp *JacobianPoint) AddAffinePoint(affinePoint *btcec.PublicKey) {
	// x := &btcec.FieldVal{}
	// y := &btcec.FieldVal{}
	// z := &btcec.FieldVal{}
	// x.SetByteSlice((affinePoint.X().Bytes())[:])
	// y.SetByteSlice((affinePoint.Y().Bytes())[:])
	// z.SetByteSlice((big.NewInt(1).Bytes())[:])
	x1 := new(big.Int).SetBytes((*jp.X.Bytes())[:])
	y1 := new(big.Int).SetBytes((*jp.Y.Bytes())[:])
	z1 := new(big.Int).SetBytes((*jp.Z.Bytes())[:])

	p := Point{}
	p.x = affinePoint.X()
	p.y = affinePoint.Y()

	z1_2 := mulmod(z1, z1, _P)
	h := addmod(mulmod(p.x, z1_2, _P), new(big.Int).Sub(_P, x1), _P)
	h_2 := mulmod(h, h, _P)
	i := mulmod(big.NewInt(4), h_2, _P)

	
	// left := new(big.Int).Add(z1, h) // mulmod(addmod(z1, h, _P), addmod(z1, h, _P), _P);
	left := mulmod(addmod(z1, h, _P), addmod(z1, h, _P), _P)

	mid := new(big.Int).Sub(_P, z1_2)

	right := new(big.Int).Sub(_P, h_2)

	
	// fmt.Printf("mulmod %s, %s\n", h_2, i);

	z := addmod(left, addmod(mid, right, _P), _P)
	
	// Compute v = x1 * i (mod P)
	v := mulmod(x1, i, _P)

	// Compute j = h * i (mod P)
	j := mulmod(h, i, _P)

	// Compute r = 2 * (s - y1) (mod P)
	//r := mulmod(big.NewInt(2), new(big.Int).Sub(_P, y1), _P)
	r := mulmod(big.NewInt(2), addmod(mulmod(p.y, mulmod(z1_2, z1, _P), _P), new(big.Int).Sub(_P, y1), _P),_P);
	// Compute x = r² - j - (2 * v) (mod P)
	r_2 := mulmod(r, r, _P)
	mid = new(big.Int).Sub(_P, j)
	right = new(big.Int).Sub(_P, mulmod(big.NewInt(2), v, _P)) // uint right = _P - mulmod(2, v, _P);

	x := addmod(r_2, addmod(mid, right, _P), _P) // addmod(r_2, addmod(mid, right, _P), _P);
	
	// Compute y = (r * (v - x)) - (2 * y1 * j) (mod P)
	left = mulmod(r, addmod(v, new(big.Int).Sub(_P, x), _P), _P) // mulmod(r, addmod(v, _P - self.x, _P), _P);
	right = new(big.Int).Sub(_P, mulmod(big.NewInt(2), mulmod(y1, j, _P), _P)) // _P - mulmod(2, mulmod(y1, j, _P), _P);

	y := addmod(left, right, _P)
	fmt.Printf("mulmod %s\n", y);
	jp.X.SetByteSlice(x.Bytes())
	jp.Y.SetByteSlice(y.Bytes())
	jp.Z.SetByteSlice(z.Bytes())
}
// Function to convert a Jacobian point to affine coordinates.
func (jp *JacobianPoint) toAffine() *btcec.PublicKey {
	_z := new(big.Int).SetBytes((*jp.Z.Bytes())[:])
	_x :=  new(big.Int).SetBytes((*jp.X.Bytes())[:])
	_y :=  new(big.Int).SetBytes((*jp.Y.Bytes())[:])
	zInv := invMod(_z)
	zInv_2 := mulmod(zInv, zInv, _P)
	x := mulmod(_x, zInv_2, _P) // uint zInv_2 = mulmod(zInv, zInv, _P);
	y := mulmod(_y, mulmod(zInv, zInv_2, _P), _P) // mulmod(self.y, mulmod(zInv, zInv_2, _P), _P);

	rX := (&btcec.FieldVal{})
	rX.SetByteSlice(x.Bytes())
	rY := (&btcec.FieldVal{})
	rY.SetByteSlice(y.Bytes())
	return btcec.NewPublicKey(rX, rY)
}
// addAffinePoint performs point addition in Jacobian coordinates on the secp256k1 curve
func AddAffinePoint(self *JacobianPoint, p *Point) *JacobianPoint {
	// Cache self's coordinates from JacobianPoint
	x1 := new(big.Int).SetBytes(self.X.Bytes()[:])
	y1 := new(big.Int).SetBytes(self.Y.Bytes()[:])
	z1 := new(big.Int).SetBytes(self.Z.Bytes()[:])

	// Compute z1_2 = z1²     (mod P)
	z1_2 := new(big.Int).Mul(z1, z1)
	z1_2.Mod(z1_2, _P)

	// Compute h = u        - x1       (mod P)
	h := new(big.Int).Mul(p.x, z1_2)
	h.Mod(h, _P)
	h.Sub(h, x1)
	if h.Sign() < 0 {
		h.Add(h, _P)
	}

	// Compute h_2 = h²    (mod P)
	h_2 := new(big.Int).Mul(h, h)
	h_2.Mod(h_2, _P)

	// Compute i = 4 * h² (mod P)
	i := new(big.Int).Mul(big.NewInt(4), h_2)
	i.Mod(i, _P)

	// Compute z = (z1 + h)² - z1²       - h²       (mod P)
	left := new(big.Int).Add(z1, h)
	left.Mul(left, left)
	left.Mod(left, _P)

	mid := new(big.Int).Sub(_P, z1_2)
	right := new(big.Int).Sub(_P, h_2)

	iZ := new(big.Int).Add(left, mid)
	iZ.Add(iZ, right)
	iZ.Mod(iZ, _P)
	self.Z.SetByteSlice(iZ.Bytes())

	// Compute v = x1 * i (mod P)
	v := new(big.Int).Mul(x1, i)
	v.Mod(v, _P)

	// Compute j = h * i (mod P)
	j := new(big.Int).Mul(h, i)
	j.Mod(j, _P)

	// Compute r = 2 * (s               - y1)       (mod P)
	s := new(big.Int).Mul(p.y, new(big.Int).Exp(z1, big.NewInt(3), _P))
	s.Mod(s, _P)

	r := new(big.Int).Sub(_P, y1)
	r.Add(r, s)
	r.Mod(r, _P)
	r.Mul(r, big.NewInt(2))
	r.Mod(r, _P)

	// Compute x = r² - j - (2 * v)             (mod P)
	r_2 := new(big.Int).Mul(r, r)
	r_2.Mod(r_2, _P)

	mid = new(big.Int).Sub(_P, j)
	right = new(big.Int).Mul(big.NewInt(2), v)
	right.Mod(right, _P)

	iX := new(big.Int).Add(r_2, mid)
	iX.Add(iX, right)
	iX.Mod(iX, _P)
	self.X.SetByteSlice(iX.Bytes())

	// Compute y = (r * (v - x))       - (2 * y1 * j)       (mod P)
	left = new(big.Int).Sub(v, iX)
	left.Mod(left, _P)
	left.Mul(r, left)
	left.Mod(left, _P)

	right = new(big.Int).Mul(big.NewInt(2), y1)
	right.Mul(right, j)
	right.Mod(right, _P)

	iY := new(big.Int).Sub(left, right)
	iY.Mod(iY, _P)
	self.Y.SetByteSlice(iY.Bytes())
	return self
}