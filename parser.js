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

    throw new TypeError("I don't know what this token is: " + token.type);
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
