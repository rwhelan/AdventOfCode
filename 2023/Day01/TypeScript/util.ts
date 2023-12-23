import * as fs from "fs";

export function readFile(path: string): string {
  return fs.readFileSync(path, "utf8");
}

export function readLines(path: string): string[] {
  return readFile(path)
    .split("\n")
    .filter((row) => {
      return row.length > 0;
    });
}

export function isDigit(char: string): boolean {
  return char >= "0" && char <= "9";
}
