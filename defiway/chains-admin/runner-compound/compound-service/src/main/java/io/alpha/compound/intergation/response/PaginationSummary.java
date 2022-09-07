package io.alpha.compound.intergation.response;

import lombok.Data;

/**
 * {
 *   "page_number": 1,
 *   "page_size": 100,
 *   "total_entries": 83,
 *   "total_pages": 1,
 * }
 */
@Data
public class PaginationSummary {
    private int page_number;
    private int page_size;
    private int total_entries;
    private int total_pages;
}
