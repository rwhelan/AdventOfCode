const std = @import("std");

const InitErrors = error{NoFilePath};

fn getFilePath() InitErrors![*:0]const u8 {
    if (std.os.argv.len < 2) {
        return InitErrors.NoFilePath;
    }

    return std.os.argv[1];
}

pub fn main() !void {
    // var args = std.process.args();
    // while (true) {
    //     const arg = args.next() orelse break;
    //     std.debug.print("{s}\n", .{arg});
    // }

    // const filePath = getFilePath() catch |err|
    //     switch (err) {
    //     InitErrors.NoFilePath => {
    //         std.debug.print("No File Path Given\n", .{});
    //         std.process.exit(1);
    //     },
    // };

    const filePath = try getFilePath();

    std.debug.print("file {s}\n", .{filePath});
}
