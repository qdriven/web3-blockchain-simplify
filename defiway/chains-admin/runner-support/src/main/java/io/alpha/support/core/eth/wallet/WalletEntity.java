package io.alpha.support.core.eth.wallet;

import com.alibaba.excel.annotation.ExcelProperty;
import lombok.Data;

@Data
public class WalletEntity {
    @ExcelProperty("priKey")
    private String priKey;
    @ExcelProperty("password")
    private String password;
    @ExcelProperty("jsonFile")
    private String jsonFile;
    @ExcelProperty("usage")
    private String usage;
}
