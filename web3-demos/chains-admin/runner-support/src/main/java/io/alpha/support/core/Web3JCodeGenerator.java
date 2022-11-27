package io.alpha.support.core;

import cn.hutool.core.io.FileUtil;
import cn.hutool.core.util.StrUtil;
import org.web3j.codegen.SolidityFunctionWrapperGenerator;

import java.io.File;
import java.nio.file.Paths;

public class Web3JCodeGenerator {

    public static void generateCodes(String[] args) {

        String currentPath = Paths.get("").toAbsolutePath().toString();

        File[] files = FileUtil.ls(currentPath + "/contracts/compound");
        for (File file : files) {
            if (file.getName().contains("usdt")) {
                String contractName = file.getName().replace("mainnet_", "").replaceAll(".json", "");
                String[] inputArgs = new String[]{
                        "-a", currentPath + "/contracts/compound/" + file.getName(),
                        "-o", currentPath + "/contracts/",
                        "-p", "io.alpha.defirunner.protocols.compound.contracts", "-c", StrUtil.upperFirst(contractName) + "Contract"
                };
                SolidityFunctionWrapperGenerator.main(inputArgs);
            }

        }

    }
}