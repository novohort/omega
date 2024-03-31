function codegen(node) {
  switch (node.type) {
    case "Program":
      return node.body.map(codegen).join("\n");
    case "FunctionDeclaration":
      const bodyCode = node.body.map(codegen).join("\n  ");
      return `fn ${node.name}() {\n  ${bodyCode}\n}`;
    case "CallExpression":
      if (node.name === "out") {
        const paramsCode = node.params.map(codegen).join(", ");
        return `println!(${paramsCode});`;
      }
      return `${node.name}(${node.params.map(codegen).join(", ")});`;
    case "StringLiteral":
      return node.value;
    default:
      throw new TypeError(`Unrecognized node type: ${node.type}`);
  }
}

module.exports = codegen;
