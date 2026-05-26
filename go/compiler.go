package main

// Compiler module - Go translation of Compiler.h/cpp
// Original C++ code from HiAsm project, translated to Go

// Compiler represents a compiler configuration
type Compiler struct {
	ID   int
	Name string
	Cmd  string
	Path string
	Ext  string
}

// CompilerCollection manages a collection of compilers
type CompilerCollection struct {
	Compilers []*Compiler
}

// Global compiler collection instance
var CompilerSet *CompilerCollection

// NewCompiler creates a new Compiler instance
func NewCompiler() *Compiler {
	return &Compiler{}
}

// NewCompilerCollection creates and initializes a new CompilerCollection
func NewCompilerCollection() *CompilerCollection {
	cc := &CompilerCollection{
		Compilers: make([]*Compiler, 0),
	}
	cc.Load()
	return cc
}

// Load loads compiler configurations (stub - no database available)
func (cc *CompilerCollection) Load() {
	// Stub - original C++ version loaded from SQLite database
	// No database available, so this remains empty
}

// GetByID returns a compiler by its ID
func (cc *CompilerCollection) GetByID(id int) *Compiler {
	for _, compiler := range cc.Compilers {
		if compiler.ID == id {
			return compiler
		}
	}
	return nil
}

// Add adds a compiler to the collection
func (cc *CompilerCollection) Add(compiler *Compiler) {
	cc.Compilers = append(cc.Compilers, compiler)
}

// Count returns the number of compilers in the collection
func (cc *CompilerCollection) Count() int {
	return len(cc.Compilers)
}

// init initializes the global compiler collection
func init() {
	CompilerSet = NewCompilerCollection()
}