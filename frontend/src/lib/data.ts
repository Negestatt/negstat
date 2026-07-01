// Display helpers. Product/brand data lives in src/lib/mock/*.json (see api.ts).

/** Formats a whole-Birr amount as an Ethiopian Birr price, e.g. 1200 -> "ETB 1,200". */
export function formatBirr(birr: number): string {
  return `ETB ${birr.toLocaleString("en-US")}`;
}
