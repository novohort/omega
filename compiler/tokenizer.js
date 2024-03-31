const fs = require("fs");

function tokenize(code) {
  const tokens = [];
  const tokenSpecs = [
    { type: "string", pattern: /"([^"\\]|\\.)*"/y },
    { type: "keyword", pattern: /\b(fn|void)\b/y },
    { type: "identifier", pattern: /\b[a-zA-Z_][a-zA-Z0-9_]*\b/y },
    { type: "symbol", pattern: /[{}();:]/y },
    { type: "whitespace", pattern: /\s+/y },
  ];

  let cursor = 0;
  while (cursor < code.length) {
    let matched = false;

    for (const { type, pattern } of tokenSpecs) {
      pattern.lastIndex = cursor;
      const match = pattern.exec(code);
      if (match) {
        matched = true;
        tokens.push({ type, value: match[0] });
        cursor = pattern.lastIndex;
        break;
      }
    }

    if (!matched) {
      throw new Error(
        `Unrecognized token starting at position ${cursor}: '${code.slice(
          cursor,
          cursor + 20
        )}'...`
      );
    }
  }

  return tokens.filter((token) => token.type !== "whitespace");
}

module.exports = tokenize;
