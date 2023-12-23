import * as util from "./util";

const FilePath: string = "../input";
// const FilePath: string = "../test";

function PartOne(rows: string[]): number {
  const values: number[] = [];
  rows.forEach((row) => {
    let first: string = "";
    let second: string = "";

    row.split("").forEach((char) => {
      if (util.isDigit(char)) {
        second = char;
        if (!first) {
          first = char;
        }
      }
    });

    values.push(parseInt(first.concat(second)));
  });

  return values.reduce((sum, val) => sum + val, 0);
}

const inputRows: string[] = util.readLines(FilePath);

console.log(`Puzzle One: ${PartOne(inputRows)}`);
