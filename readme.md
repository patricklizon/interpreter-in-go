# Interpreter in go &middot;

Learning go by writing an interpreter.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Start](#start)
- [How interpreter works](#how-interpreter-works)
  - [Lexical Analysis](#lexical-analysis)
  - [Parsing](#parsing)
  - [Evaluation](#evaluation)

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
