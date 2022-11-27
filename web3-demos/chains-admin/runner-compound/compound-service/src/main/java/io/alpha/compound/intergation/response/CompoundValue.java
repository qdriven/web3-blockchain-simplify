package io.alpha.compound.intergation.response;

import lombok.Builder;
import lombok.Data;

import java.math.BigDecimal;

@Data
@Builder
public class CompoundValue implements Comparable<CompoundValue> {
    private String value;
    public static CompoundValue fromValue(String value){
        return new CompoundValue(value);
    }

    @Override
    public int compareTo(CompoundValue o) {
        return new BigDecimal(value).compareTo(new BigDecimal(o.getValue()));
    }
}
