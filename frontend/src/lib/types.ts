import type { components } from "./api";

export type VendorSlug =
  | "arbiters_of_hexis"
  | "cephalon_suda"
  | "new_loka"
  | "red_veil"
  | "steel_meridian"
  | "the_perrin_sequence";

export enum ItemType {
  Mod,
  ArchPart,
  Weapon,
}

export type Item = components["schemas"]["Item"] & {
  type: ItemType;
};

export type Vendor = components["schemas"]["Vendor"] & {
  slug: VendorSlug;
};

export type BasicVendor = components["schemas"]["BasicVendor"] & {
  slug: VendorSlug;
};

export type BasicVendorsResponse = components["schemas"]["BasicVendorsResponseBody"] & {
  updatedAt: Date;
  vendors: BasicVendor[];
};

export type VendorResponse = components["schemas"]["VendorsResponseBody"] & {
  updatedAt: Date;
  vendor: Vendor;
};