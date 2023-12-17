# Interpreter in go &middot; ![test workflow](https://github.com/patricklizon/interpreter-in-go/actions/workflows/go.yml/badge.svg?event=push)

Learning go by writing an interpreter.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Start](#start)
- [How interpreter works](#how-interpreter-works)
  - [Lexical Analysis](#lexical-analysis)
  - [Parsing](#parsing)
    - [Pratt parsing](#pratt-parsing)
    - [Parse tree](#parse-tree)
  - [Evaluation](#evaluation)
- [Terminology](#terminology)

## Prerequisites

- [Go](https://golang.org/doc/install), version defined in [go.mod](go.mod) file

## Start

To initialize repl run:

```sh
go run main.go
```

## How interpreter works

 Interpreter reads source, then breaks it down into tokens, organizing those
 tokens into a parse tree, and then walking through the tree to execute the
 program.

Source code -> Tokens -> Parse tree -> Execute

### Lexical Analysis

The interpreter reads the program code line by line and breaks it down into
tokens.
Tokens are the smallest units of a program, like variables, operators, brackets,
etc.

**Input:**

```js
    let ten = 5 + 5;
```

**Output:**

```js
    // Collection of tokens that represent the program.
    // in real-life, token would be an object with more properties like
    // file name, line number, etc.
    [
        { type: 'LET', literal: 'let' },
        { type: 'IDENTIFIER', literal: 'ten' },
        { type: 'EQUAL_SIGN', literal: '=' },
        { type: 'INT', literal: '5' },
        { type: 'PLUS_SIGN', literal: '+' },
        { type: 'INT', literal: '5' },
        { type: 'SEMICOLON', literal: ';' },
    ];
```

### Parsing

#### Pratt parsing

Approach for parsing described in this book is called [Pratt parsing](https://en.wikipedia.org/wiki/Pratt_parser). It's a top-down operator precedence parser. It's a way of parsing expressions without using a grammar.

**More can be read here:**

- [Top Down Operator Precedence](https://crockford.com/javascript/tdop/tdop.html)
- [Pratt Parsers: Expression Parsing Made Easy](https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/)

#### Parse tree

The interpreter takes the tokens and organizes them into a data structure called
a parse tree (or syntax tree). This step checks the tokens for syntax errors and
organizes them in a way that reflects the structure of the program.

**Input:**

```js
    let ten = 5 + 5;
```

**Output:**

```js
    // The parse tree is a nested collection of nodes that represent the
    // structure of the program.
    {
        type: 'LET_STATEMENT',
        token: { type: 'LET', literal: 'let' },
        name: {
            type: 'IDENTIFIER',
            token: { type: 'IDENTIFIER', literal: 'ten' },
            value: 'ten',
        },
        value: {
            type: 'INFIX_EXPRESSION',
            token: { type: 'PLUS_SIGN', literal: '+' },
            left: {
                type: 'INT',
                token: { type: 'INT', literal: '5' },
                value: 5,
            },
            operator: '+',
            right: {
                type: 'INT',
                token: { type: 'INT', literal: '5' },
                value: 5,
            },
        },
    };
```

### Evaluation

The interpreter walks through the parse tree and executes each operation as it
goes. This could involve doing calculations, running functions, or manipulating
data.

## Terminology

**Token** - is a single, meaningful unit of information within a program. Tokens
can be thought of as the indivisible "words" in the language.

**Lexical analysis** - is the process of converting a sequence of characters into
a sequence of tokens.

**Parsing** - is the process of analyzing a string of symbols, either in natural
language, computer languages or data structures, conforming to the rules of a
formal grammar.

**Expression** - is a combination of one or more explicit values, constants,
variables, operators, and functions that the programming language interprets
and computes to produce another value.

**Abstract syntax tree (AST)** - is a tree representation of the abstract syntactic
structure of source code written in a programming language. Each node of the tree
denotes a construct occurring in the source code.

**prefix operator** - is an operator that is written before its operand. Prefix
operators include unary logical `not` and arithmetic `minus` operators.

**infix operator** - is an operator that is written between its operands. Infix
operators include arithmetic operators, such as addition, and logical operators,
such as `and` and `or`.

**postfix operator** - is an operator that is written after its operand. Postfix
operators include the `increment` and `decrement` operators.

**precedence** - is a collection of rules that reflect conventions about which
procedures to perform first in order to evaluate a given mathematical expression.

**associativity** - is a property of some binary operations. It describes in which
order operators of the same precedence are evaluated.

**binary operator** - is an operator that operates on two operands and manipulates
them to return a result.

**binary expression** - is an expression that has two operands and one operator.
