package io.alpha.support.paths;

import java.nio.file.Paths;

public class PathUtil {

    public static String getJavaWorkspacePath() {
        return Paths.get("").toAbsolutePath().toString();
    }
}
