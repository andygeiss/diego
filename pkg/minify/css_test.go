package minify_test

import (
	"github.com/andygeiss/diego/pkg/assert"
	"github.com/andygeiss/diego/pkg/assert/matchers"
	"github.com/andygeiss/diego/pkg/minify"
	"testing"
)

var (
	string4k  = string(make([]byte, 4096))
	string8k  = string(make([]byte, 4096<<1))
	string16k = string(make([]byte, 4096<<2))
	string32k = string(make([]byte, 4096<<3))
)

func BenchmarkCSS_4k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minify.CSS(string4k)
	}
}

func BenchmarkCSS_8k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minify.CSS(string8k)
	}
}

func BenchmarkCSS_16k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minify.CSS(string16k)
	}
}

func BenchmarkCSS_32k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minify.CSS(string32k)
	}
}

func TestCSS_Newline(t *testing.T) {
	assert.That(t, minify.CSS("*{\r\n}\r\n"), matchers.IsEqual("*{}"))
}

func TestCSS_Tab(t *testing.T) {
	assert.That(t, minify.CSS(".class\t.subclass\t{   }"), matchers.IsEqual(".class .subclass{}"))
}

func TestCSS_Whitespace(t *testing.T) {
	assert.That(t, minify.CSS("  A,   E  >  F  {   }"), matchers.IsEqual("A,E>F{}"))
}

func TestCSS_Class(t *testing.T) {
	assert.That(t, minify.CSS(".class {   }"), matchers.IsEqual(".class{}"))
}

func TestCSS_ClassDecendant(t *testing.T) {
	assert.That(t, minify.CSS(".class .decendant {   }"), matchers.IsEqual(".class .decendant{}"))
}

func TestCSS_Comment(t *testing.T) {
	assert.That(t, minify.CSS("/* Comment */\n.class .decendant {   }"), matchers.IsEqual(".class .decendant{}"))
}

func TestCSS_Attributes(t *testing.T) {
	assert.That(t, minify.CSS(".class + .subclass{ a: x; b: y; }"), matchers.IsEqual(".class+.subclass{a:x;b:y;}"))
}

func TestCSS_VariableDeclaration(t *testing.T) {
	assert.That(t, minify.CSS("    --col-one:		rgb(255,255,255);"), matchers.IsEqual("--col-one:rgb(255,255,255);"))
}
