// (C) 2022 Mike Spreitzer.  All rights reserved.

package boolean

func Eq(vbl, lit string) Atom {
	return Atom{Variable: vbl, CompareTo: lit}
}

func NEq(vbl, lit string) Atom {
	return Atom{Variable: vbl, CompareTo: lit, Negate: true}
}
