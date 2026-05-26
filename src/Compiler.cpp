/*
 * Compiler.cpp
 *
 *  Created on: Dec 5, 2010
 *      Author: dilma
 */

#include "Compiler.h"
#include "MainDataBase.h"

CompilerCollection *compilerSet;

Compiler::Compiler() {

}

CompilerCollection::CompilerCollection() {
	load();
}

void CompilerCollection::load() {
	// Stub - no database available
}

Compiler *CompilerCollection::getById(int id) {
	return NULL;
}
