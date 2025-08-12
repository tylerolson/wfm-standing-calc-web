export enum ItemType {
  Mod,
  ArchPart,
  Weapon,
}

export interface Vendor {
  name: string;
  items: Item[];
}

export interface Item {
  name: string;
  type: ItemType;
  standing: number;
  weightedPrice: number;
  volume: number;
}
