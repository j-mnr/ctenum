# ctenum

Compile-time safe string enums for Go!

The idiomatic way of creating an enumeration in Go using `iota` has a
fundamental language design flaw (or feature, depending on what you're doing).
The ability to subtype primitive values and the ability to automatically cast
primitive values to that subtype. This is great when you want to declare some
primitive like: `const myFloat = 4.20`, but terrible when you have a function
signature like: `type ColorEnum uint8; SearchFor(query string, color ColorEnum)
[]Image` because you can do `SearchFor("panthers", 69)`, instead of
`SearchFor("panthers", ColorPink)`. The compiler will not complain about `69`,
it will simply cast it for you, but if your enum only has 8 values, the caller
has given an incorrect value, this happens for all primitives and worse for
`string` there is no `iota` equivalent.

Consider using this tool when you have a known set of `string` values, that
would aid the caller in understanding and using your API/Library.

⚠️ NOTE ⚠️ If you're creating an Enum that you plan to be used by callers of your
API, do **NOT** use [this well intentioned but poorly executed
implementation](#what-not-to-do) that I've seen passed around to _"fix"_ Go
enums.

## What not to do

This _only_ applies to users who are **NOT** using an `internal/` enum or if the
created enum is unexported, i.e. `type roleEnum struct { ... }`.

---

### Problem

Given a `string` enum, we have the problem that a caller can pass in any value
because it will be cast to the underlying type do to `untyped` primitive values.

We don't want the user to pass in the wrong value, so what should we do?
Complicate things.

### (Not so) Solution

```go
type Role struct {
	slug string
}

func (r Role) String() string {
	return r.slug
}

var (
	Unknown   = Role{""}
	Guest     = Role{"guest"}
	Member    = Role{"member"}
	Moderator = Role{"moderator"}
	Admin     = Role{"admin"}
)
```

Now if you have `ShowPage(r Role)`, it would be impossible for the caller to
pass in an incorrect value, correct? We've solved it! ... Or did we? Sure a
caller now has to put in `ShowPage(user.Role)`, but who declares those Roles?
Your library? As **_Global Mutatable State_** 🤔 That doesn't sound safe at all.
😵

I follow a pretty simple rule (for general programming), don't use global state.
If you're looking to have a secure program, it's best to avoid things like
[`http.DefaultClient`](https://pkg.go.dev/net/http@go1.22.0#pkg-variables) in
production because if you accidentally pull in a malicious module they can very
easily do

```go
func init() {
	http.DefaultClient = myMaliciousTrackerClient
}
```

And it doesn't need to be that obvious, a package could easily hide it away in
some private function that doesn't change the values right away until called.

```go
func malicious() {
	myPkg.Unknown   = myPkg.Admin
	myPkg.Guest     = myPkg.Admin
	myPkg.Member    = myPkg.Admin
	myPkg.Moderator = myPkg.Admin
	myPkg.Admin     = myPkg.Admin
}

func HelpfulLibraryFunc( ... ) {
	malicious()
	// Do what is required
}
```

It's very easy, at this point, to say

> This will never happen; this has never happened; this is solely the user's
> fault for not reading all their dependencies; this isn't a big deal; this is
> the least of your worries if you've fallen for this.

etc., etc., but the most malicious CVEs come from what _seem_ like helpful, safe
code. When something like a default value goes unchecked for nearly a decade,
you can go [from logging, all the way to full Remote Code
Execution](https://en.wikipedia.org/wiki/Log4Shell). Stay safe out there,
Gophers.

## Non Goals

This tool is not looking to take over
[`Float`](https://pkg.go.dev/golang.org/x/exp/constraints#Float) or
[`Integer`](https://pkg.go.dev/golang.org/x/exp/constraints#Integer)
type enumerations in Go. It is an anti-pattern to declare an enum of some
number type, just to use it as a `string`. If you need it as a `string` for
debugging, the battle-tested tool
[`stringer`](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) is great for
this. If you're looking for something like `stringer` that _extends_ the
functionality [`enumer`](https://pkg.go.dev/github.com/dmarkham/enumer) could be
useful for you.

## Examples

JSON, YAML, and simple Go files with a `//go:generate` directive are all
supported.

See ['examples/'](./examples) for plenty of ways to declare your enums.
