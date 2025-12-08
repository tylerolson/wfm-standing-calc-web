export interface VendorsResponse {
  updatedAt: Date;
  updating: boolean;
  vendors: Vendor[];
}

export interface BasicVendorsResponse {
  updatedAt: Date;
  updating: boolean;
  vendors: BasicVendor[];
}

export interface BasicVendor {
  slug: string;
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
  slug: string;
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
