# Interpreter in go &middot;

Learning go by writing an interpreter.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Start](#start)
- [How interpreter works](#how-interpreter-works)
  - [Lexical Analysis](#lexical-analysis)
  - [Parsing](#parsing)
  - [Evaluation](#evaluation)

[//]: # 'AUTO INSERT HEADER PREPUBLISH

## Prerequisites

- [Go](https://golang.org/doc/install), version defined in [go.mod](go.mod) file

## Start

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

### Evaluation

The interpreter walks through the parse tree and executes each operation as it
goes. This could involve doing calculations, running functions, or manipulating
data.
