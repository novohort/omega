function parse(tokens) {
  let current = 0;

  function walk() {
    let token = tokens[current];

    if (token.type === "string") {
      current++;
      return {
        type: "StringLiteral",
        value: token.value,
      };
    }

    if (
      token.type === "identifier" &&
      tokens[current + 1] &&
      tokens[current + 1].value === "("
    ) {
      let node = {
        type: "CallExpression",
        name: token.value,
        params: [],
      };

      token = tokens[++current];
      token = tokens[++current];

      if (tokens[current].type === "string") {
        node.params.push(walk());
        current++;
      }

      current++;
      return node;
    }

    if (token.type === "keyword" && token.value === "fn") {
      current++;
      const name = tokens[current++].value;

      while (
        tokens[current].type !== "symbol" ||
        tokens[current].value !== "{"
      ) {
        current++;
      }

      const body = [];
      if (tokens[current].type === "symbol" && tokens[current].value === "{") {
        current++;
        while (
          !(tokens[current].type === "symbol" && tokens[current].value === "}")
        ) {
          body.push(walk());
        }
        current++;
      }

      return {
        type: "FunctionDeclaration",
        name,
        params: [],
        body,
      };
    }

    if (token.type === "keyword" && (token.value === "let" || token.value === "const")) {
      current++;
      const nameToken = tokens[current++];
      if (nameToken.type !== "identifier") {
        throw new Error(`Expected variable name after '${token.value}', found '${nameToken.type}'`);
      }
      const name = nameToken.value;

      const colonToken = tokens[current++];
      if (colonToken.value !== ":") {
        throw new Error(`Expected ':' after variable name '${name}', found '${colonToken.value}'`);
      }

      const typeToken = tokens[current++];
      if (typeToken.type !== "keyword" || (typeToken.value !== "int" && typeToken.value !== "uint")) {
        throw new Error(`Expected 'int' or 'uint' for variable type, found '${typeToken.value}'`);
      }
      const datatype = typeToken.value;

      const openBraceToken = tokens[current++];
      if (openBraceToken.value !== "{") {
        throw new Error(`Expected '{' before variable initialization, found '${openBraceToken.value}'`);
      }

      const valueToken = tokens[current++];
      if (valueToken.type !== "number") {
        throw new Error(`Expected number for variable initialization, found '${valueToken.type}'`);
      }
      const value = valueToken.value;

      const closeBraceToken = tokens[current++];
      if (closeBraceToken.value !== "}") {
        throw new Error(`Expected '}' after variable initialization, found '${closeBraceToken.value}'`);
      }

      const semicolonToken = tokens[current++];
      if (semicolonToken.value !== ";") {
        throw new Error(`Expected ';' after variable declaration, found '${semicolonToken.value}'`);
      }

      const isMutable = token.value === "let";

      return {
        type: "VariableDeclaration",
        name: name,
        datatype: datatype,
        value: value,
        isMutable: isMutable
      };
    }

    throw new TypeError(`I don't know what this token is: ${token.type} | ${token.value}`);
  }

  let ast = {
    type: "Program",
    body: [],
  };

  while (current < tokens.length) {
    ast.body.push(walk());
  }

  return ast;
}

module.exports = parse;
