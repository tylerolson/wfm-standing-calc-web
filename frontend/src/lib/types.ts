export type VendorSlug =
  | "arbiters_of_hexis"
  | "cephalon_suda"
  | "new_loka"
  | "red_veil"
  | "steel_meridian"
  | "the_perrin_sequence";

export interface VendorResponse {
  updatedAt: Date;
  updating: boolean;
  vendor: Vendor;
}

export interface BasicVendorsResponse {
  updatedAt: Date;
  updating: boolean;
  vendors: BasicVendor[];
}

export interface BasicVendor {
  slug: VendorSlug;
  name: string;
  mostProfitable: Item;
  mostVolume: Item;
  mostEfficient: Item;
}

export enum ItemType {
  Mod,
  ArchPart,
  Weapon,
}

export interface Vendor {
  slug: VendorSlug;
  name: string;
  items: Item[];
}

export interface Item {
  name: string;
  type: ItemType;
  standing: number;
  volume: number;
  price: number;
}
